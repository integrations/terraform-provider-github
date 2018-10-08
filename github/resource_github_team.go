package github

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}
}

func resourceGithubTeamCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

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
		_, _, err = client.Admin.UpdateTeamLDAPMapping(ctx, *githubTeam.ID, mapping)
		if err != nil {
			return err
		}
	}

	d.SetId(strconv.FormatInt(*githubTeam.ID, 10))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading team: %s", d.Id())
	team, resp, err := client.Teams.GetTeam(ctx, id)
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
	d.Set("description", team.Description)
	d.Set("name", team.Name)
	d.Set("privacy", team.Privacy)
	if parent := team.Parent; parent != nil {
		d.Set("parent_team_id", parent.GetID())
	} else {
		d.Set("parent_team_id", "")
	}
	d.Set("ldap_dn", team.GetLDAPDN())
	d.Set("slug", team.GetSlug())

	return nil
}

func resourceGithubTeamUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

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
	team, _, err := client.Teams.EditTeam(ctx, teamId, editedTeam)
	if err != nil {
		return err
	}

	if d.HasChange("ldap_dn") {
		ldapDN := d.Get("ldap_dn").(string)
		mapping := &github.TeamLDAPMapping{
			LDAPDN: github.String(ldapDN),
		}
		_, _, err = client.Admin.UpdateTeamLDAPMapping(ctx, *team.ID, mapping)
		if err != nil {
			return err
		}
	}

	d.SetId(strconv.FormatInt(*team.ID, 10))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting team: %s", d.Id())
	_, err = client.Teams.DeleteTeam(ctx, id)
	return err
}
