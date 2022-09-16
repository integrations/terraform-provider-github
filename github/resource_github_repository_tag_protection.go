package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubRepositoryTagProtection() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryTagProtectionCreateOrUpdate,
		Read:   resourceGithubRepositoryTagProtectionRead,
		Delete: resourceGithubRepositoryTagProtectionDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 {
					return nil, fmt.Errorf("Invalid ID specified. Supplied ID must be written as <repository>/<webhook_id>")
				}
				d.Set("repository", parts[0])
				d.Set("tag_protection_id", parts[1])
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
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

func resourceGithubRepositoryTagProtectionCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
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

	return resourceGithubRepositoryTagProtectionRead(d, meta)
}

func resourceGithubRepositoryTagProtectionRead(d *schema.ResourceData, meta interface{}) error {

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	tag_protection, _, err := client.Repositories.ListTagProtection(ctx, owner, repo)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound && d.IsNewResource() {
				return nil
			}
			return err
		}
		return err
	}
	for _, tag := range tag_protection {
		if tag.GetID() == id {
			d.Set("pattern", tag.GetPattern())
		}
	}

	return nil
}

func resourceGithubRepositoryTagProtectionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	tag_protection_id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting tag protection for %s/%s with id %d", owner, repo, tag_protection_id)
	_, error := client.Repositories.DeleteTagProtection(ctx, owner, repo, tag_protection_id)

	return error
}
