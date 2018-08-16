package github

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubTeamMembership() *schema.Resource {

	return &schema.Resource{
		Create: resourceGithubTeamMembershipCreateOrUpdate,
		Read:   resourceGithubTeamMembershipRead,
		Update: resourceGithubTeamMembershipCreateOrUpdate,
		Delete: resourceGithubTeamMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "member",
				ValidateFunc: validateValueFunc([]string{"member", "maintainer"}),
			},
		},
	}
}

func resourceGithubTeamMembershipCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}

	username := d.Get("username").(string)
	role := d.Get("role").(string)

	_, _, err = client.Organizations.AddTeamMembership(context.TODO(),
		teamId,
		username,
		&github.OrganizationAddTeamMembershipOptions{
			Role: role,
		},
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(&teamIdString, &username))

	return resourceGithubTeamMembershipRead(d, meta)
}

func resourceGithubTeamMembershipRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	teamIdString, username, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}

	membership, _, err := client.Organizations.GetTeamMembership(context.TODO(),
		teamId, username)
	if err != nil {
		d.SetId("")
		return nil
	}

	team, user := getTeamAndUserFromURL(membership.URL)

	d.Set("username", user)
	d.Set("role", membership.Role)
	d.Set("team_id", team)

	return nil
}

func resourceGithubTeamMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}

	_, err = client.Organizations.RemoveTeamMembership(context.TODO(),
		teamId,
		d.Get("username").(string))

	return err
}

func getTeamAndUserFromURL(url *string) (string, string) {
	var team, user string

	urlSlice := strings.Split(*url, "/")
	for v := range urlSlice {
		if urlSlice[v] == "teams" {
			team = urlSlice[v+1]
		}
		if urlSlice[v] == "memberships" {
			user = urlSlice[v+1]
		}
	}
	return team, user
}
