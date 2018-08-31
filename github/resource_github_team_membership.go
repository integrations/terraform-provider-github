package github

import (
	"context"
	"log"
	"net/http"
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

	log.Printf("[DEBUG] Creating team membership: %s/%s (%s)", teamIdString, username, role)
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

	log.Printf("[DEBUG] Reading team membership: %s/%s", teamIdString, username)
	membership, _, err := client.Organizations.GetTeamMembership(context.TODO(),
		teamId, username)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing team membership %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
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
	username := d.Get("username").(string)

	log.Printf("[DEBUG] Deleting team membership: %s/%s", teamIdString, username)
	_, err = client.Organizations.RemoveTeamMembership(context.TODO(), teamId, username)

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
