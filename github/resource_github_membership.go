package github

import (
	"log"

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
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
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

	username := d.Get("username").(string)
	roleName := d.Get("role").(string)

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Creating membership: %s/%s", orgName, username)
	membership, _, err := client.Organizations.EditOrgMembership(ctx,
		username,
		orgName,
		&github.Membership{
			Role: github.String(roleName),
		},
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(*membership.Organization.Login, *membership.User.Login))

	return resourceGithubMembershipRead(d, meta)
}

func resourceGithubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	_, username, err := parseTwoPartID(d.Id(), "organization", "username")
	if err != nil {
		return err
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
		d.Set("username", membership.User.Login)
		d.Set("role", membership.Role)

		return nil
	}
}

func resourceGithubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Deleting membership: %s", d.Id())
	_, err = client.Organizations.RemoveOrgMembership(ctx,
		d.Get("username").(string), orgName)

	return err
}
