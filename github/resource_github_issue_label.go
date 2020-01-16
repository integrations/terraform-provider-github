package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubIssueLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubIssueLabelCreateOrUpdate,
		Read:   resourceGithubIssueLabelRead,
		Update: resourceGithubIssueLabelCreateOrUpdate,
		Delete: resourceGithubIssueLabelDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// resourceGithubIssueLabelCreateOrUpdate idempotently creates or updates an
// issue label. Issue labels are keyed off of their "name", so pre-existing
// issue labels result in a 422 HTTP error if they exist outside of Terraform.
// Normally this would not be an issue, except new repositories are created with
// a "default" set of labels, and those labels easily conflict with custom ones.
//
// This function will first check if the label exists, and then issue an update,
// otherwise it will create. This is also advantageous in that we get to use the
// same function for two schema funcs.

func resourceGithubIssueLabelCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	name := d.Get("name").(string)
	color := d.Get("color").(string)

	label := &github.Label{
		Name:  github.String(name),
		Color: github.String(color),
	}

	log.Printf("[DEBUG] Querying label existence: %s (%s/%s)",
		name, orgName, repoName)
	existing, resp, err := client.Issues.GetLabel(context.TODO(),
		orgName, repoName, name)
	if err != nil && resp.StatusCode != http.StatusNotFound {
		return err
	}

	if existing != nil {
		label.Description = github.String(d.Get("description").(string))

		log.Printf("[DEBUG] Updating label: %s:%s (%s/%s)",
			name, color, orgName, repoName)

		// Pull out the original name. If we already have a resource, this is the
		// parsed ID. If not, it's the value given to the resource.
		var originalName string
		if d.Id() == "" {
			originalName = name
		} else {
			var err error
			_, originalName, err = parseTwoPartID(d.Id())
			if err != nil {
				return err
			}
		}

		_, _, err := client.Issues.EditLabel(context.TODO(),
			orgName, repoName, originalName, label)
		if err != nil {
			return err
		}
	} else {
		if v, ok := d.GetOk("description"); ok {
			label.Description = github.String(v.(string))
		}

		log.Printf("[DEBUG] Creating label: %s:%s (%s/%s)",
			name, color, orgName, repoName)
		_, resp, err := client.Issues.CreateLabel(context.TODO(),
			orgName, repoName, label)
		if resp != nil {
			log.Printf("[DEBUG] Response from creating label: %#v", *resp)
		}
		if err != nil {
			return err
		}
	}

	d.SetId(buildTwoPartID(&repoName, &name))

	return resourceGithubIssueLabelRead(d, meta)
}

func resourceGithubIssueLabelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client
	repoName, name, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	orgName := meta.(*Organization).name
	log.Printf("[DEBUG] Reading label: %s (%s/%s)", name, orgName, repoName)
	githubLabel, _, err := client.Issues.GetLabel(context.TODO(),
		orgName, repoName, name)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing label %s (%s/%s) from state because it no longer exists in GitHub",
					name, orgName, repoName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("repository", repoName)
	d.Set("name", name)
	d.Set("color", githubLabel.Color)
	d.Set("description", githubLabel.Description)
	d.Set("url", githubLabel.URL)

	return nil
}

func resourceGithubIssueLabelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client

	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Deleting label: %s (%s/%s)", name, orgName, repoName)
	_, err := client.Issues.DeleteLabel(context.TODO(),
		orgName, repoName, name)

	return err
}
