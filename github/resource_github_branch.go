package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubBranch() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubBranchCreate,
		Read:   resourceGithubBranchRead,
		Delete: resourceGithubBranchDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_branch": {
				Type:     schema.TypeString,
				Default:  "master",
				Optional: true,
				ForceNew: true,
			},
			"source_sha": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubBranchCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	branchName := d.Get("branch").(string)
	sourceBranchName := d.Get("source_branch").(string)
	if _, hasSourceSHA := d.GetOk("source_sha"); !hasSourceSHA {
		log.Printf("[DEBUG] Querying source branch state to derive source_sha")
		ref, _, err := client.Git.GetRef(ctx, orgName, repoName, "refs/heads/"+sourceBranchName)
		if err != nil {
			return err
		}
		d.Set("source_sha", *ref.Object.SHA)
	}
	sourceBranchSHA := d.Get("source_sha").(string)

	log.Printf("[DEBUG] Creating repository branch: %s/%s (%s)",
		orgName, repoName, branchName)
	_, _, err = client.Git.CreateRef(ctx, orgName, repoName, &github.Reference{
		Ref:    github.String("refs/heads/" + branchName),
		Object: &github.GitObject{SHA: github.String(sourceBranchSHA)},
	})
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(&repoName, &branchName))

	return resourceGithubBranchRead(d, meta)
}

func resourceGithubBranchRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	repoName, branchName, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Reading repository branch: %s/%s (%s)",
		orgName, repoName, branchName)
	ref, resp, err := client.Git.GetRef(ctx, orgName, repoName, "refs/heads/"+branchName)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing branch %s/%s from state because it no longer exists in Github",
					orgName, repoName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("repository", repoName)
	d.Set("branch", branchName)
	d.Set("ref", *ref.Ref)
	d.Set("sha", *ref.Object.SHA)

	return nil
}

func resourceGithubBranchDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	repoName, branchName, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting repository branch: %s/%s (%s)",
		orgName, repoName, branchName)
	_, err = client.Git.DeleteRef(ctx, orgName, repoName, "refs/heads/"+branchName)

	return err
}
