package github

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v42/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func resourceGithubTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamCreate,
		Read:   resourceGithubTeamRead,
		Update: resourceGithubTeamUpdate,
		Delete: resourceGithubTeamDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: customdiff.Sequence(
			customdiff.ComputedIf("slug", func(d *schema.ResourceDiff, meta interface{}) bool {
				return d.HasChange("name")
			}),
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"privacy": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "secret",
				ValidateFunc: validateValueFunc([]string{"secret", "closed"}),
			},
			"parent_team_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldap_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_default_maintainer": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceGithubTeamCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client

	ownerName := meta.(*Owner).name
	name := d.Get("name").(string)

	newTeam := github.NewTeam{
		Name:        name,
		Description: github.String(d.Get("description").(string)),
		Privacy:     github.String(d.Get("privacy").(string)),
	}

	if parentTeamID, ok := d.GetOk("parent_team_id"); ok {
		id := int64(parentTeamID.(int))
		newTeam.ParentTeamID = &id
	}
	ctx := context.Background()

	githubTeam, _, err := client.Teams.CreateTeam(ctx,
		ownerName, newTeam)
	if err != nil {
		return err
	}

	create_default_maintainer := d.Get("create_default_maintainer").(bool)
	if !create_default_maintainer {
		log.Printf("[DEBUG] Removing default maintainer from team: %s (%s)", name, ownerName)
		if err := removeDefaultMaintainer(*githubTeam.Slug, meta); err != nil {
			return err
		}
	}

	if ldapDN := d.Get("ldap_dn").(string); ldapDN != "" {
		mapping := &github.TeamLDAPMapping{
			LDAPDN: github.String(ldapDN),
		}
		_, _, err = client.Admin.UpdateTeamLDAPMapping(ctx, githubTeam.GetID(), mapping)
		if err != nil {
			return err
		}
	}

	d.SetId(strconv.FormatInt(githubTeam.GetID(), 10))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	team, resp, err := client.Teams.GetTeamByID(ctx, orgId, id)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing team %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("description", team.GetDescription())
	d.Set("name", team.GetName())
	d.Set("privacy", team.GetPrivacy())
	if parent := team.Parent; parent != nil {
		d.Set("parent_team_id", parent.GetID())
	} else {
		d.Set("parent_team_id", "")
	}
	d.Set("ldap_dn", team.GetLDAPDN())
	d.Set("slug", team.GetSlug())
	d.Set("node_id", team.GetNodeID())
	d.Set("members_count", team.GetMembersCount())

	return nil
}

func resourceGithubTeamUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	editedTeam := github.NewTeam{
		Name:        d.Get("name").(string),
		Description: github.String(d.Get("description").(string)),
		Privacy:     github.String(d.Get("privacy").(string)),
	}
	if parentTeamID, ok := d.GetOk("parent_team_id"); ok {
		id := int64(parentTeamID.(int))
		editedTeam.ParentTeamID = &id
	}

	teamId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	team, _, err := client.Teams.EditTeamByID(ctx, orgId, teamId, editedTeam, false)
	if err != nil {
		return err
	}

	if d.HasChange("ldap_dn") {
		ldapDN := d.Get("ldap_dn").(string)
		mapping := &github.TeamLDAPMapping{
			LDAPDN: github.String(ldapDN),
		}
		_, _, err = client.Admin.UpdateTeamLDAPMapping(ctx, team.GetID(), mapping)
		if err != nil {
			return err
		}
	}

	d.SetId(strconv.FormatInt(team.GetID(), 10))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, err = client.Teams.DeleteTeamByID(ctx, orgId, id)
	/*
		When deleting a team and it failed, we need to check if it has already been deleted meanwhile.
		This could be the case when deleting nested teams via Terraform by looping through a module
		or resource and the parent team might have been deleted already. If the parent team had
		been deleted already (via parallel runs), the child team is also already gone (deleted by
		GitHub automatically).
		So we're checking if it still exists and if not, simply remove it from TF state.
	*/
	if err != nil {
		// Fetch the team in order to see if it exists or not (http 404)
		_, _, err = client.Teams.GetTeamByID(ctx, orgId, id)
		if err != nil {
			if ghErr, ok := err.(*github.ErrorResponse); ok {
				if ghErr.Response.StatusCode == http.StatusNotFound {
					// If team we failed to delete does not exist, remove it from TF state.
					log.Printf("[WARN] Removing team: %s from state because it no longer exists",
						d.Id())
					d.SetId("")
					return nil
				}
			}
		}
	}
	return err
}

func removeDefaultMaintainer(teamSlug string, meta interface{}) error {

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	v4client := meta.(*Owner).v4client

	type User struct {
		Login githubv4.String
	}

	var query struct {
		Organization struct {
			Team struct {
				Members struct {
					Nodes []User
				}
			} `graphql:"team(slug:$slug)"`
		} `graphql:"organization(login:$login)"`
	}
	variables := map[string]interface{}{
		"slug":  githubv4.String(teamSlug),
		"login": githubv4.String(orgName),
	}

	err := v4client.Query(meta.(*Owner).StopContext, &query, variables)
	if err != nil {
		return err
	}

	for _, user := range query.Organization.Team.Members.Nodes {
		_, err := client.Teams.RemoveTeamMembershipBySlug(meta.(*Owner).StopContext, orgName, teamSlug, string(user.Login))
		if err != nil {
			return err
		}
	}

	return nil
}
