package github

import (
	"errors"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserCreate,
		Read:   resourceGithubUserRead,
		Update: nil,
		Delete: resourceGithubUserDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubUserImport,
		},

		Schema: map[string]*schema.Schema{
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

func resourceGithubUserCreate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("The github_user resource must be imported, it cannot be created.")
}

func resourceGithubUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	ctx := prepareResourceContext(d)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	log.Printf("[DEBUG] Reading user by ID: %d", id)
	user, resp, err := client.Users.GetByID(ctx, id)
	switch apires, apierr := apiResult(resp, err); apires {
	case APINotModified:
		break
	case APINotFound:
		log.Printf("[WARN] Removing user %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	case APIError:
		return apierr
	default:
		d.Set("etag", resp.Header.Get("ETag"))
		d.Set("username", user.Login)
	}

	username := d.Get("username").(string)
	log.Printf("[DEBUG] Adding user %s/%d to cache", username, id)
	meta.(*Organization).UserMap.AddUser(username, id)

	return nil
}

func resourceGithubUserDelete(d *schema.ResourceData, meta interface{}) error {
	// this operation cannot be performed, but should be silently ignored
	return nil
}

func resourceGithubUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	userString := d.Id()

	log.Printf("[DEBUG] Reading user: %s", userString)
	// Attempt to parse the string as a numeric ID
	userId, err := strconv.ParseInt(userString, 10, 64)
	if err != nil {
		// It wasn't a numeric ID, try to use it as a username
		if user, _, err := client.Users.Get(ctx, userString); err != nil {
			return nil, err
		} else {
			userId = *user.ID
		}
	}

	d.SetId(strconv.FormatInt(userId, 10))

	return []*schema.ResourceData{d}, nil
}
