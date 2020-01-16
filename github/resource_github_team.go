package github

import (
	"context"

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
	client := meta.(*Owner).client
	n := d.Get("name").(string)
	desc := d.Get("description").(string)
	p := d.Get("privacy").(string)

	newTeam := &github.NewTeam{
		Name:        n,
		Description: &desc,
		Privacy:     &p,
	}
	if parentTeamID, ok := d.GetOk("parent_team_id"); ok {
		id := int64(parentTeamID.(int))
		newTeam.ParentTeamID = &id
	}

	githubTeam, _, err := client.Organizations.CreateTeam(context.TODO(), meta.(*Owner).name, newTeam)
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

	d.SetId(fromGithubID(githubTeam.ID))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client

	team, err := getGithubTeam(d, client)
	if err != nil {
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
	client := meta.(*Owner).client
	team, err := getGithubTeam(d, client)

	if err != nil {
		d.SetId("")
		return nil
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	privacy := d.Get("privacy").(string)

	editedTeam := &github.NewTeam{
		Name:        name,
		Description: &description,
		Privacy:     &privacy,
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

	d.SetId(fromGithubID(team.ID))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client
	id := toGithubID(d.Id())
	_, err := client.Organizations.DeleteTeam(context.TODO(), id)
	return err
}

func getGithubTeam(d *schema.ResourceData, github *github.Client) (*github.Team, error) {
	id := toGithubID(d.Id())
	team, _, err := github.Organizations.GetTeam(context.TODO(), id)
	return team, err
}
