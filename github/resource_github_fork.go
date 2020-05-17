package github

import (
	"context"
	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"net/http"
	"regexp"

	"strings"
)

func resourceGithubFork() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGithubForkRead,
		Create: resourceGithubForkCreate,
		Update: resourceGithubForkUpdate,
		Delete: resourceGithubForkDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fork_from_owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fork_from_repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fork_into_organization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"homepage_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Can only be set after initial fork creation, and only if repository is also private",
			},
			"has_issues": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"has_projects": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"has_downloads": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"has_wiki": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_template": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_merge_commit": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allow_squash_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allow_rebase_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"delete_branch_on_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"default_branch": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Can only be set after initial repository creation, and only if the target branch exists",
			},
			"license_template": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"gitignore_template": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"archived": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"topics": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`), "must include only lowercase alphanumeric characters or hyphens and cannot start with a hyphen"),
				},
			},
			"full_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"html_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_clone_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"svn_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"git_clone_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_clone_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubForkObject(d *schema.ResourceData) *github.Repository {
	return &github.Repository{
		Description:         github.String(d.Get("description").(string)),
		Homepage:            github.String(d.Get("homepage_url").(string)),
		Private:             github.Bool(d.Get("private").(bool)),
		HasDownloads:        github.Bool(d.Get("has_downloads").(bool)),
		HasIssues:           github.Bool(d.Get("has_issues").(bool)),
		HasProjects:         github.Bool(d.Get("has_projects").(bool)),
		HasWiki:             github.Bool(d.Get("has_wiki").(bool)),
		IsTemplate:          github.Bool(d.Get("is_template").(bool)),
		AllowMergeCommit:    github.Bool(d.Get("allow_merge_commit").(bool)),
		AllowSquashMerge:    github.Bool(d.Get("allow_squash_merge").(bool)),
		AllowRebaseMerge:    github.Bool(d.Get("allow_rebase_merge").(bool)),
		DeleteBranchOnMerge: github.Bool(d.Get("delete_branch_on_merge").(bool)),
		LicenseTemplate:     github.String(d.Get("license_template").(string)),
		GitignoreTemplate:   github.String(d.Get("gitignore_template").(string)),
		Archived:            github.Bool(d.Get("archived").(bool)),
		Topics:              expandStringList(d.Get("topics").(*schema.Set).List()),
	}
}

func resourceGithubForkCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client
	ctx := context.Background()
	owner := d.Get("fork_from_owner").(string)
	repoName := d.Get("fork_from_repository").(string)

	log.Printf("[DEBUG] Create a fork from %s/%s", owner, repoName)
	opts := &github.RepositoryCreateForkOptions{}
	if v, ok := d.GetOk("fork_into_organization"); ok {
		opts.Organization = v.(string)
	}
	repo, _, err := client.Repositories.CreateFork(ctx, owner, repoName, opts)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode != http.StatusAccepted {
				return err
			}
		}
	}
	d.SetId(repo.GetFullName())

	return resourceGithubForkUpdate(d, meta)
}

func resourceGithubForkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client

	log.Printf("[DEBUG] Reading fork: %s", d.Id())

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}
	forkParts := strings.Split(d.Id(), "/")
	owner := forkParts[0]
	repoName := forkParts[1]
	repo, resp, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing repository fork %s/%s from state because it no longer exists in GitHub",
					owner, repoName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("name", repo.GetName())
	d.Set("description", repo.GetDescription())
	d.Set("homepage_url", repo.GetHomepage())
	d.Set("private", repo.GetPrivate())
	d.Set("has_issues", repo.GetHasIssues())
	d.Set("has_projects", repo.GetHasProjects())
	d.Set("has_wiki", repo.GetHasWiki())
	d.Set("is_template", repo.GetIsTemplate())
	d.Set("allow_merge_commit", repo.GetAllowMergeCommit())
	d.Set("allow_squash_merge", repo.GetAllowSquashMerge())
	d.Set("allow_rebase_merge", repo.GetAllowRebaseMerge())
	d.Set("delete_branch_on_merge", repo.GetDeleteBranchOnMerge())
	d.Set("has_downloads", repo.GetHasDownloads())
	d.Set("full_name", repo.GetFullName())
	d.Set("default_branch", repo.GetDefaultBranch())
	d.Set("html_url", repo.GetHTMLURL())
	d.Set("ssh_clone_url", repo.GetSSHURL())
	d.Set("svn_url", repo.GetSVNURL())
	d.Set("git_clone_url", repo.GetGitURL())
	d.Set("http_clone_url", repo.GetCloneURL())
	d.Set("archived", repo.GetArchived())
	d.Set("topics", flattenStringList(repo.Topics))
	d.Set("node_id", repo.GetNodeID())

	if repo.TemplateRepository != nil {
		d.Set("template", []interface{}{
			map[string]interface{}{
				"owner":      repo.TemplateRepository.Owner.Login,
				"repository": repo.TemplateRepository.Name,
			},
		})
	} else {
		d.Set("template", []interface{}{})
	}
	return nil
}

func resourceGithubForkUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client

	repoReq := resourceGithubForkObject(d)
	if d.HasChanges("name") {
		_, n := d.GetChange("name")
		if len(n.(string)) > 0 {
			repoReq.Name = github.String(n.(string))
		}
	}
	// Can only set `default_branch` on an already created repository with the target branches ref already in-place
	if v, ok := d.GetOk("default_branch"); ok {
		branch := v.(string)
		// If branch is "master", and the repository hasn't been initialized yet, setting this value will fail
		if branch != "master" {
			repoReq.DefaultBranch = &branch
		}
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	forkName := d.Id()
	forkParts := strings.Split(forkName, "/")
	owner := forkParts[0]
	repoName := forkParts[1]
	if repoReq.Name != nil && repoName == *repoReq.Name {
		repoName = *repoReq.Name
	}
	log.Printf("[DEBUG] Updating fork: %s/%s", owner, repoName)
	repo, _, err := client.Repositories.Edit(ctx, owner, repoName, repoReq)
	if err != nil {
		return err
	}
	d.SetId(repo.GetFullName())

	if d.HasChange("topics") {
		topics := repoReq.Topics
		_, _, err = client.Repositories.ReplaceAllTopics(ctx, owner, repo.GetName(), topics)
		if err != nil {
			return err
		}
	}

	return resourceGithubForkRead(d, meta)
}

func resourceGithubForkDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	forkName := d.Id()
	forkParts := strings.Split(forkName, "/")
	log.Printf("[DEBUG] Deleting fork: %s/%s", forkParts[0], forkParts[1])
	_, err := client.Repositories.Delete(ctx, forkParts[0], forkParts[1])

	return err
}
