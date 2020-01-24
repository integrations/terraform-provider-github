package github

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOrganizationBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrganizationBlockCreate,
		Read:   resourceOrganizationBlockRead,
		Delete: resourceOrganizationBlockDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOrganizationBlockCreate(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	username := d.Get("username").(string)

	log.Printf("[DEBUG] Creating organization block: %s (%s)", username, orgName)
	_, err = client.Organizations.BlockUser(ctx, orgName, username)
	if err != nil {
		return err
	}
	d.SetId(username)

	return resourceOrganizationBlockRead(d, meta)
}

func resourceOrganizationBlockRead(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	username := d.Id()

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Reading organization block: %s (%s)", d.Id(), orgName)
	blocked, resp, err := client.Organizations.IsBlocked(ctx, orgName, username)
	switch apires, apierr := apiResult(resp, err); apires {
	case APINotModified:
		return nil
	case APINotFound:
		// not sure if this will ever be hit, I imagine just returns false?
		log.Printf("[WARN] Removing organization block %s/%s from state because it no longer exists in GitHub", orgName, d.Id())
		d.SetId("")
		return nil
	case APIError:
		return apierr
	default:
		if !blocked {
			d.SetId("")
			return nil
		}

		d.Set("username", username)
		d.Set("etag", resp.Header.Get("ETag"))

		return nil
	}
}

func resourceOrganizationBlockDelete(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	username := d.Id()

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Deleting organization block: %s (%s)", d.Id(), orgName)
	_, err = client.Organizations.UnblockUser(ctx, orgName, username)
	return err
}
