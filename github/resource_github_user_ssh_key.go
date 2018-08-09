package github

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubUserSshKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserSshKeyCreate,
		Read:   resourceGithubUserSshKeyRead,
		Delete: resourceGithubUserSshKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, oldV, newV string, d *schema.ResourceData) bool {
					newTrimmed := strings.TrimSpace(newV)
					return oldV == newTrimmed
				},
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubUserSshKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	title := d.Get("title").(string)
	key := d.Get("key").(string)

	userKey, _, err := client.Users.CreateKey(context.TODO(), &github.Key{
		Title: &title,
		Key:   &key,
	})
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(*userKey.ID, 10))

	return resourceGithubUserSshKeyRead(d, meta)
}

func resourceGithubUserSshKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	key, _, err := client.Users.GetKey(context.TODO(), id)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("title", key.Title)
	d.Set("key", key.Key)
	d.Set("url", key.URL)

	return nil
}

func resourceGithubUserSshKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	_, err = client.Users.DeleteKey(context.TODO(), id)

	return err
}
