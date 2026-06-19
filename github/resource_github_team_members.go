package github

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type MemberChange struct {
	Old, New map[string]any
}

func resourceGithubTeamMembers() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubTeamMembersCreate,
		ReadContext:   resourceGithubTeamMembersRead,
		UpdateContext: resourceGithubTeamMembersUpdate,
		DeleteContext: resourceGithubTeamMembersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubTeamMembersImport,
		},

		CustomizeDiff: customdiff.Sequence(diffLegacyTeam, diffTeam),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubTeamMembersV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubTeamMembersStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to authoritatively manage GitHub team members.",

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				Description:      "Slug of the GitHub team to manage membership for.",
			},
			"team_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
				Description:      "ID or slug of the GitHub team to manage membership for.",
				Deprecated:       "Use `team_slug` instead; this field will be made computed only in a future version of the provider.",
			},
			"members": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of users that should be members of the team.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
							Description:      "User to add to the team.",
						},
						"role": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "member",
							Description:      "Role to grant the user within the team; must be one of `member` or `maintainer`.",
							ValidateDiagFunc: validateValueFunc([]string{"member", "maintainer"}),
						},
					},
				},
			},
		},
	}
}

func resourceGithubTeamMembersCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)
	teamIDString, _ := d.Get("team_id").(string)
	members, _ := d.Get("members").(*schema.Set)

	var teamID int64
	if slug == "" {
		team := newLegacyTeamIdentity(teamIDString)
		if s, ok := team.getSlugOK(); ok {
			slug = s
		} else {
			teamID = team.getID()

			s, err := lookupTeamSlug(ctx, client, meta.id, teamID)
			if err != nil {
				return diag.FromErr(err)
			}
			slug = s
		}
	}

	if teamID == 0 {
		id, err := lookupTeamID(ctx, client, orgName, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		teamID = id
	}

	tflog.Debug(ctx, "Creating team members.", map[string]any{"team_slug": slug, "team_id": teamID})

	for _, mMap := range members.List() {
		memb, _ := mMap.(map[string]any)
		username, _ := memb["username"].(string)
		role, _ := memb["role"].(string)

		tflog.Debug(ctx, "Adding member to team.", map[string]any{"team_slug": slug, "username": username, "role": role})

		_, _, err := client.Teams.AddTeamMembershipBySlug(ctx, orgName, slug, username, &github.TeamAddTeamMembershipOptions{Role: role})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(strconv.FormatInt(teamID, 10))

	if err := d.Set("team_slug", slug); err != nil {
		return diag.FromErr(err)
	}

	if teamIDString == "" {
		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubTeamMembersRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)

	teamIDStr, _ := d.Get("team_id").(string)
	if teamIDStr == "" {
		tflog.Debug(ctx, "Looking up team ID from slug.", map[string]any{"team_slug": slug})

		teamID, err := lookupTeamID(ctx, client, orgName, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Debug(ctx, "Reading team members.", map[string]any{"team_slug": slug})

	var teamMembersAndMaintainers []any
	for _, role := range []string{"member", "maintainer"} {
		opts := &github.TeamListTeamMembersOptions{
			Role:        role,
			ListOptions: github.ListOptions{PerPage: maxPerPage},
		}

		for member, err := range client.Teams.ListTeamMembersBySlugIter(ctx, orgName, slug, opts) {
			if err != nil {
				if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, "Team no longer exists, removing from state.", map[string]any{"team_slug": slug})
					d.SetId("")
					return nil
				}
				return diag.FromErr(err)
			}
			teamMembersAndMaintainers = append(teamMembersAndMaintainers, map[string]any{
				"username": strings.ToLower(member.GetLogin()),
				"role":     role,
			})
		}
	}

	if err := d.Set("members", teamMembersAndMaintainers); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubTeamMembersUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)

	if idStr, _ := d.Get("team_id").(string); idStr == "" {
		tflog.Debug(ctx, "Looking up team ID from slug.", map[string]any{"team_slug": slug})

		teamID, err := lookupTeamID(ctx, client, orgName, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return diag.FromErr(err)
		}
	}

	o, n := d.GetChange("members")
	vals := make(map[string]*MemberChange)
	for _, raw := range o.(*schema.Set).List() {
		obj := raw.(map[string]any)
		k := obj["username"].(string)
		vals[k] = &MemberChange{Old: obj}
	}
	for _, raw := range n.(*schema.Set).List() {
		obj := raw.(map[string]any)
		k := obj["username"].(string)
		if _, ok := vals[k]; !ok {
			vals[k] = &MemberChange{}
		}
		vals[k].New = obj
	}

	for username, change := range vals {
		var create, del bool

		switch {
		case change.Old == nil:
			create = true
		case change.New == nil:
			del = true
		case reflect.DeepEqual(change.Old, change.New):
			continue
		default:
			del = true
			create = true
		}

		if del {
			tflog.Debug(ctx, "Removing member from team.", map[string]any{"team_slug": slug, "username": username})

			_, err := client.Teams.RemoveTeamMembershipBySlug(ctx, orgName, slug, username)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		if create {
			role := change.New["role"].(string)

			tflog.Debug(ctx, "Adding member to team.", map[string]any{"team_slug": slug, "username": username, "role": role})

			opts := &github.TeamAddTeamMembershipOptions{Role: role}
			_, _, err := client.Teams.AddTeamMembershipBySlug(ctx, orgName, slug, username, opts)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubTeamMembersDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)

	tflog.Debug(ctx, "Removing all members from team.", map[string]any{"team_slug": slug})

	members := d.Get("members").(*schema.Set)

	for _, member := range members.List() {
		mem := member.(map[string]any)
		username := mem["username"].(string)

		tflog.Debug(ctx, "Removing member from team.", map[string]any{"team_slug": slug, "username": username})

		_, err := client.Teams.RemoveTeamMembershipBySlug(ctx, orgName, slug, username)
		if err != nil {
			// 404 means the team is gone (API returns 204 for missing memberships).
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Team no longer exists, skipping remaining member deletions", map[string]any{"team_slug": slug})
				return nil
			}
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubTeamMembersImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)

	idString := d.Id()

	tflog.Debug(ctx, "Importing team members.", map[string]any{"id": idString})

	team := newLegacyTeamIdentity(idString)
	slug, slugOK := team.getSlugOK()
	teamID := team.getID()
	if slugOK {
		id, err := lookupTeamID(ctx, meta.v3client, meta.name, slug)
		if err != nil {
			return nil, err
		}
		teamID = id

		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return nil, err
		}
	} else {
		s, err := lookupTeamSlug(ctx, meta.v3client, meta.id, teamID)
		if err != nil {
			return nil, err
		}
		slug = s

		if err := d.Set("team_id", idString); err != nil {
			return nil, err
		}
	}

	d.SetId(strconv.FormatInt(teamID, 10))

	if err := d.Set("team_slug", slug); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
