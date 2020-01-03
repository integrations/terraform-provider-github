package github

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubMembershipCreateOrUpdate,
		Read:   resourceGithubMembershipRead,
		Update: resourceGithubMembershipCreateOrUpdate,
		Delete: resourceGithubMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubMembershipImport,
		},
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubMembershipV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubMembershipStateUpgradeV0,
				Version: 0,
			},
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateNumericIDFunc,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateValueFunc([]string{"member", "admin"}),
				Default:      "member",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubMembershipCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	userIDString := d.Get("user_id").(string)
	role := d.Get("role").(string)

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Creating membership: %s/%s (%s)", orgName, userIDString, role)

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(userIDString, err)
	}

	username, ok := meta.(*Organization).UserMap.GetUsername(userID, client)
	if !ok {
		return fmt.Errorf("Unable to get GitHub user %d", userID)
	}

	membership, _, err := client.Organizations.EditOrgMembership(ctx,
		username,
		orgName,
		&github.Membership{
			Role: &role,
		},
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(*membership.Organization.Login, userIDString))

	return resourceGithubMembershipRead(d, meta)
}

func resourceGithubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName, userIDString, err := parseTwoPartID(d.Id(), "organization", "user_id")
	if err != nil {
		return err
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(userIDString, err)
	}

	username, ok := meta.(*Organization).UserMap.GetUsername(userID, client)
	if !ok {
		return fmt.Errorf("Unable to get GitHub user %d", userID)
	}

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Reading membership: %s", d.Id())

	membership, resp, err := client.Organizations.GetOrgMembership(ctx, username, orgName)
	switch apires, apierr := apiResult(resp, err); apires {
	case APINotModified:
		return nil
	case APINotFound:
		log.Printf("[WARN] Removing membership %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	case APIError:
		return apierr
	default:
		d.Set("etag", resp.Header.Get("ETag"))
		d.Set("role", membership.Role)

		return nil
	}
}

func resourceGithubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName, userIDString, err := parseTwoPartID(d.Id(), "organization", "user_id")
	if err != nil {
		return err
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(userIDString, err)
	}

	username, ok := meta.(*Organization).UserMap.GetUsername(userID, client)
	if !ok {
		return fmt.Errorf("Unable to get GitHub user %d", userID)
	}

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Deleting membership: %s", d.Id())
	_, err = client.Organizations.RemoveOrgMembership(ctx, username, orgName)

	return err
}

func resourceGithubMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	orgName, userString, err := parseTwoPartID(d.Id(), "organization", "user_id_or_name")
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Reading user: %s", userString)
	// Attempt to parse the string as a numeric ID
	userID, err := strconv.ParseInt(userString, 10, 64)
	if err != nil {
		// It wasn't a numeric ID, try to use it as a username
		user, _, err := client.Users.Get(ctx, userString)
		if err != nil {
			return nil, err
		}
		userID = *user.ID
	}

	d.SetId(buildTwoPartID(orgName, strconv.FormatInt(userID, 10)))

	return []*schema.ResourceData{d}, nil
}
