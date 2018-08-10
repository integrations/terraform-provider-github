package github

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubUserGpgKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserGpgKeyCreate,
		Read:   resourceGithubUserGpgKeyRead,
		Delete: resourceGithubUserGpgKeyDelete,

		Schema: map[string]*schema.Schema{
			"armored_public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubUserGpgKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	pubKey := d.Get("armored_public_key").(string)

	key, _, err := client.Users.CreateGPGKey(context.TODO(), pubKey)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(*key.ID, 10))

	return resourceGithubUserGpgKeyRead(d, meta)
}

func resourceGithubUserGpgKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	key, _, err := client.Users.GetGPGKey(context.TODO(), id)
	if err != nil {
		log.Printf("[WARN] GitHub User GPG Key (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("key_id", key.KeyID)

	return nil
}

func resourceGithubUserGpgKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	_, err = client.Users.DeleteGPGKey(context.TODO(), id)

	return err
}
