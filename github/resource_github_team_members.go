package github

import (
	"context"
	"errors"
	"iter"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub team id or slug",
			},
			"members": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of team members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
							Description:      "The user to add to the team.",
						},
						"role": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "member",
							Description:      "The role of the user within the team. Must be one of 'member' or 'maintainer'.",
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
	orgId := meta.id
	orgName := meta.name

	teamIdString := d.Get("team_id").(string)
	team := newLegacyTeamIdentity(teamIdString)

	members := d.Get("members").(*schema.Set)
	for _, mMap := range members.List() {
		memb := mMap.(map[string]any)
		username := memb["username"].(string)
		role := memb["role"].(string)

		log.Printf("[DEBUG] Creating team membership: %s/%s (%s)", teamIdString, username, role)

		opts := &github.TeamAddTeamMembershipOptions{Role: role}
		var err error
		if slug, ok := team.getSlugOK(); ok {
			_, _, err = client.Teams.AddTeamMembershipBySlug(ctx, orgName, slug, username, opts)
		} else {
			_, _, err = client.Teams.AddTeamMembershipByID(ctx, orgId, team.getID(), username, opts)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(teamIdString)

	return nil
}

func resourceGithubTeamMembersUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgId := meta.id
	orgName := meta.name

	teamIdString := d.Get("team_id").(string)
	team := newLegacyTeamIdentity(teamIdString)

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
			log.Printf("[DEBUG] Deleting team membership: %s/%s", teamIdString, username)

			var err error
			if slug, ok := team.getSlugOK(); ok {
				_, err = client.Teams.RemoveTeamMembershipBySlug(ctx, orgName, slug, username)
			} else {
				_, err = client.Teams.RemoveTeamMembershipByID(ctx, orgId, team.getID(), username)
			}
			if err != nil {
				return diag.FromErr(err)
			}
		}

		if create {
			role := change.New["role"].(string)

			log.Printf("[DEBUG] Creating team membership: %s/%s (%s)", teamIdString, username, role)

			opts := &github.TeamAddTeamMembershipOptions{Role: role}
			var err error
			if slug, ok := team.getSlugOK(); ok {
				_, _, err = client.Teams.AddTeamMembershipBySlug(ctx, orgName, slug, username, opts)
			} else {
				_, _, err = client.Teams.AddTeamMembershipByID(ctx, orgId, team.getID(), username, opts)
			}
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(teamIdString)

	return nil
}

func resourceGithubTeamMembersRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgId := meta.id
	orgName := meta.name

	teamIdString := d.Get("team_id").(string)
	if teamIdString == "" && !d.IsNewResource() {
		log.Printf("[DEBUG] Importing team with id %q", d.Id())
		teamIdString = d.Id()
	}

	team := newLegacyTeamIdentity(teamIdString)

	// We intentionally set these early to allow reconciliation
	// from an upstream bug which emptied team_id in state
	// See https://github.com/integrations/terraform-provider-github/issues/323
	if err := d.Set("team_id", teamIdString); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Reading team members: %s", teamIdString)

	var teamMembersAndMaintainers []any

	for _, role := range []string{"member", "maintainer"} {
		opts := &github.TeamListTeamMembersOptions{
			Role:        role,
			ListOptions: github.ListOptions{PerPage: maxPerPage},
		}

		var seq iter.Seq2[*github.User, error]
		if slug, ok := team.getSlugOK(); ok {
			seq = client.Teams.ListTeamMembersBySlugIter(ctx, orgName, slug, opts)
		} else {
			seq = client.Teams.ListTeamMembersByIDIter(ctx, orgId, team.getID(), opts)
		}

		for member, err := range seq {
			if err != nil {
				if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, "Team no longer exists, removing from state", map[string]any{"team_id": teamIdString})
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

func resourceGithubTeamMembersDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgId := meta.id
	orgName := meta.name

	teamIdString := d.Get("team_id").(string)
	team := newLegacyTeamIdentity(teamIdString)

	members := d.Get("members").(*schema.Set)

	for _, member := range members.List() {
		mem := member.(map[string]any)
		username := mem["username"].(string)

		log.Printf("[DEBUG] Deleting team membership: %s/%s", teamIdString, username)

		var err error
		if slug, ok := team.getSlugOK(); ok {
			_, err = client.Teams.RemoveTeamMembershipBySlug(ctx, orgName, slug, username)
		} else {
			_, err = client.Teams.RemoveTeamMembershipByID(ctx, orgId, team.getID(), username)
		}
		if err != nil {
			// 404 means the team is gone (API returns 204 for missing memberships).
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Team no longer exists, skipping remaining member deletions", map[string]any{"team_id": teamIdString})
				return nil
			}
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubTeamMembersImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)

	teamId, err := getTeamID(ctx, meta, d.Id())
	if err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(teamId, 10))

	return []*schema.ResourceData{d}, nil
}
