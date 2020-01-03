package github

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubTeamMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamMembershipCreateOrUpdate,
		Read:   resourceGithubTeamMembershipRead,
		Update: resourceGithubTeamMembershipCreateOrUpdate,
		Delete: resourceGithubTeamMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubTeamMembershipImport,
		},
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubTeamMembershipV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubTeamMembershipStateUpgradeV0,
				Version: 0,
			},
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateNumericIDFunc,
			},
			"user_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateNumericIDFunc,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "member",
				ValidateFunc: validateValueFunc([]string{"member", "maintainer"}),
			},
			"username": {
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

func resourceGithubTeamMembershipCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	teamIDString := d.Get("team_id").(string)
	userIDString := d.Get("user_id").(string)
	role := d.Get("role").(string)

	log.Printf("[DEBUG] Creating team membership: %s/%s (%s)", teamIDString, userIDString, role)

	teamID, _, username, err := getTeamAndUser(teamIDString, userIDString, meta.(*Organization))
	if err != nil {
		return err
	}

	_, _, err = client.Teams.AddTeamMembership(ctx,
		teamID,
		username,
		&github.TeamAddTeamMembershipOptions{
			Role: role,
		},
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(teamIDString, userIDString))

	return resourceGithubTeamMembershipRead(d, meta)
}

func resourceGithubTeamMembershipRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	teamIDString, userIDString, err := parseTwoPartID(d.Id(), "team_id", "user_id")
	if err != nil {
		return err
	}

	teamID, userID, username, err := getTeamAndUser(teamIDString, userIDString, meta.(*Organization))
	if err != nil {
		return err
	}

	d.Set("team_id", teamIdString)
	d.Set("username", username)

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Reading team membership: %s/%s", teamIDString, userIDString)

	membership, resp, err := client.Teams.GetTeamMembership(ctx, teamID, username)
	switch apires, apierr := apiResult(resp, err); apires {
	case APINotModified:
		break
	case APINotFound:
		log.Printf("[WARN] Removing team membership %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	case APIError:
		return apierr
	default:
		d.Set("etag", resp.Header.Get("ETag"))
		d.Set("role", membership.Role)

		return nil
	}

	return nil
}

func resourceGithubTeamMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	teamIDString, userIDString, err := parseTwoPartID(d.Id(), "team_id", "user_id")
	if err != nil {
		return err
	}

	teamID, _, username, err := getTeamAndUser(teamIDString, userIDString, meta.(*Organization))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting team membership: %s/%s", teamIDString, userIDString)
	_, err = client.Teams.RemoveTeamMembership(ctx, teamID, username)

	return err
}

func resourceGithubTeamMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	orgName, err := getOrganization(meta)
	if err != nil {
		return nil, err
	}

	teamString, userString, err := parseTwoPartID(d.Id(), "team_id_or_name", "user_id_or_name")
	if err != nil {
		return nil, err
	}

	var teamID int64
	var userID int64

	log.Printf("[DEBUG] Reading team: %s", teamString)
	// Attempt to parse the string as a numeric ID
	teamID, err = strconv.ParseInt(teamString, 10, 64)
	if err != nil {
		// It wasn't a numeric ID, try to use it as a slug
		team, _, err := client.Teams.GetTeamBySlug(ctx, orgName, teamString)
		if err != nil {
			return nil, err
		}
		teamID = *team.ID
	}

	log.Printf("[DEBUG] Reading user: %s", userString)
	// Attempt to parse the string as a numeric ID
	userID, err = strconv.ParseInt(userString, 10, 64)
	if err != nil {
		// It wasn't a numeric ID, try to use it as a username
		user, _, err := client.Users.Get(ctx, userString)
		if err != nil {
			return nil, err
		}
		userID = *user.ID
	}

	d.SetId(buildTwoPartID(strconv.FormatInt(teamID, 10), strconv.FormatInt(userID, 10)))

	return []*schema.ResourceData{d}, nil
}

func getTeamAndUser(teamIDString string, userIDString string, org *Organization) (teamID, userID int64, username string, err error) {
	teamID, err = strconv.ParseInt(teamIDString, 10, 64)
	if err != nil {
		err = unconvertibleIdErr(teamIDString, err)
		return
	}

	userID, err = strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		err = unconvertibleIdErr(userIDString, err)
		return
	}

	username, ok := org.UserMap.GetUsername(userID, org.client)
	if !ok {
		log.Printf("[DEBUG] Unable to obtain user %d from cache", userID)
		err = fmt.Errorf("Unable to get GitHub user %d", userID)
		return
	}

	return
}
