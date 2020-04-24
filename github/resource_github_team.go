package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		},
	}
}

func resourceGithubTeamCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	orgName := meta.(*Organization).name
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

	log.Printf("[DEBUG] Creating team: %s (%s)", name, orgName)
	githubTeam, _, err := client.Teams.CreateTeam(ctx,
		orgName, newTeam)
	if err != nil {
		return err
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

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading team: %s", d.Id())
	var team *github.Team
	var resp *github.Response
	if meta.(*Organization).isEnterprise {
		team, resp, err = GetEnterpriseTeamByID(ctx, client, id)
	} else {
		team, resp, err = client.Teams.GetTeamByID(ctx, orgId, id)
	}

	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing team %s from state because it no longer exists in GitHub",
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

	return nil
}

func resourceGithubTeamUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

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

	log.Printf("[DEBUG] Updating team: %s", d.Id())
	var team *github.Team
	if meta.(*Organization).isEnterprise {
		team, _, err = EditEnterpriseTeamByID(ctx, client, teamId, editedTeam)
	} else {
		team, _, err = client.Teams.EditTeamByID(ctx, orgId, teamId, editedTeam, false)
	}
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

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting team: %s", d.Id())
	if meta.(*Organization).isEnterprise {
		_, err = DeleteEnterpriseTeamByID(ctx, client, id)
	} else {
		_, err = client.Teams.DeleteTeamByID(ctx, orgId, id)
	}
	return err
}

// API functionality below is no longer available in go-github v29.0.3+.
// Naming conventions reflect Enterprise Github Account support.
// Code taken from go-github v29.0.2 as a temporary work-around to [GH-404] and [GH-434].
func GetEnterpriseTeamByID(ctx context.Context, client *github.Client, id int64) (*github.Team, *github.Response, error) {
	u := fmt.Sprintf("teams/%v", id)
	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	t := new(github.Team)
	resp, err := client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

func EditEnterpriseTeamByID(ctx context.Context, client *github.Client, id int64, team github.NewTeam) (*github.Team, *github.Response, error) {
	u := fmt.Sprintf("teams/%v", id)
	req, err := client.NewRequest("PATCH", u, team)

	if err != nil {
		return nil, nil, err
	}

	t := new(github.Team)
	resp, err := client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

func DeleteEnterpriseTeamByID(ctx context.Context, client *github.Client, id int64) (*github.Response, error) {
	u := fmt.Sprintf("teams/%v", id)
	req, err := client.NewRequest("DELETE", u, nil)

	if err != nil {
		return nil, err
	}

	return client.Do(ctx, req, nil)
}
