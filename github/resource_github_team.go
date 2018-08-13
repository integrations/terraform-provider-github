package github

import (
	"context"
	"log"
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
		},
	}
}

func resourceGithubTeamCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	newTeam := &github.NewTeam{
		Name:        d.Get("name").(string),
		Description: github.String(d.Get("description").(string)),
		Privacy:     github.String(d.Get("privacy").(string)),
	}
	if parentTeamID, ok := d.GetOk("parent_team_id"); ok {
		id := int64(parentTeamID.(int))
		newTeam.ParentTeamID = &id
	}

	githubTeam, _, err := client.Organizations.CreateTeam(context.TODO(),
		meta.(*Organization).name, newTeam)
	if err != nil {
		return err
	}

	if ldapDN := d.Get("ldap_dn").(string); ldapDN != "" {
		mapping := &github.TeamLDAPMapping{
			LDAPDN: github.String(ldapDN),
		}
		_, _, err = client.Admin.UpdateTeamLDAPMapping(context.TODO(), *githubTeam.ID, mapping)
		if err != nil {
			return err
		}
	}

	d.SetId(strconv.FormatInt(*githubTeam.ID, 10))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	team, err := getGithubTeam(d, client)
	if err != nil {
		log.Printf("[WARN] GitHub Team (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	d.Set("description", team.Description)
	d.Set("name", team.Name)
	d.Set("privacy", team.Privacy)
	if parent := team.Parent; parent != nil {
		d.Set("parent_team_id", parent.GetID())
	} else {
		d.Set("parent_team_id", "")
	}
	d.Set("ldap_dn", team.GetLDAPDN())
	return nil
}

func resourceGithubTeamUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	team, err := getGithubTeam(d, client)
	if err != nil {
		d.SetId("")
		return nil
	}

	editedTeam := &github.NewTeam{
		Name:        d.Get("name").(string),
		Description: github.String(d.Get("description").(string)),
		Privacy:     github.String(d.Get("privacy").(string)),
	}
	if parentTeamID, ok := d.GetOk("parent_team_id"); ok {
		id := int64(parentTeamID.(int))
		editedTeam.ParentTeamID = &id
	}

	team, _, err = client.Organizations.EditTeam(context.TODO(), *team.ID, editedTeam)
	if err != nil {
		return err
	}

	if d.HasChange("ldap_dn") {
		ldapDN := d.Get("ldap_dn").(string)
		mapping := &github.TeamLDAPMapping{
			LDAPDN: github.String(ldapDN),
		}
		_, _, err = client.Admin.UpdateTeamLDAPMapping(context.TODO(), *team.ID, mapping)
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

	_, err = client.Organizations.DeleteTeam(context.TODO(), id)
	return err
}

func getGithubTeam(d *schema.ResourceData, github *github.Client) (*github.Team, error) {
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return nil, unconvertibleIdErr(d.Id(), err)
	}

	team, _, err := github.Organizations.GetTeam(context.TODO(), id)
	return team, err
}
