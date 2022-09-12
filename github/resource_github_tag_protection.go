package github

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubTagProtection() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTagProtectionCreateOrUpdate,
		Read:   resourceGithubTagProtectionRead,
		//Update: resourceGithubTagProtectionCreateOrUpdate,
		Delete: resourceGithubTagProtectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pattern": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tag_protection_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceGithubTagProtectionCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	pattern := d.Get("pattern").(string)
	log.Printf("[DEBUG] Creating tag protection for %s/%s with pattern %s", owner, repo, pattern)
	tagProtection, _, err := client.Repositories.CreateTagProtection(ctx, owner, repo, pattern)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(tagProtection.GetID(), 10))

	return resourceGithubTagProtectionRead(d, meta)
}

func resourceGithubTagProtectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	tag_protection, _, err := client.Repositories.ListTagProtection(ctx, owner, repo)
	if err != nil {
		return err
	}
	for _, tag := range tag_protection {
		if tag.GetPattern() == d.Get("pattern").(string) {
			d.Set("tag_protection_id", tag.GetID())
			d.Set("pattern", tag.GetPattern())
		}
	}

	return nil
}

func resourceGithubTagProtectionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	tag_protection_id, error := strconv.ParseInt(d.Id(), 10, 64)
	if error != nil {
		return error
	}
	log.Printf("[DEBUG] Deleting tag protection for %s/%s with id %d", owner, repo, tag_protection_id)
	_, err := client.Repositories.DeleteTagProtection(ctx, owner, repo, tag_protection_id)

	return err
}
