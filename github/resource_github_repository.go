package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubRepository() *schema.Resource {

	return &schema.Resource{
		Create: resourceGithubRepositoryCreate,
		Read:   resourceGithubRepositoryRead,
		Update: resourceGithubRepositoryUpdate,
		Delete: resourceGithubRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Type:     schema.TypeBool,
				Optional: true,
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
			"auto_init": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
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
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
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
		},
	}
}

func resourceGithubRepositoryObject(d *schema.ResourceData) *github.Repository {
	return &github.Repository{
		Name:              github.String(d.Get("name").(string)),
		Description:       github.String(d.Get("description").(string)),
		Homepage:          github.String(d.Get("homepage_url").(string)),
		Private:           github.Bool(d.Get("private").(bool)),
		HasDownloads:      github.Bool(d.Get("has_downloads").(bool)),
		HasIssues:         github.Bool(d.Get("has_issues").(bool)),
		HasProjects:       github.Bool(d.Get("has_projects").(bool)),
		HasWiki:           github.Bool(d.Get("has_wiki").(bool)),
		AllowMergeCommit:  github.Bool(d.Get("allow_merge_commit").(bool)),
		AllowSquashMerge:  github.Bool(d.Get("allow_squash_merge").(bool)),
		AllowRebaseMerge:  github.Bool(d.Get("allow_rebase_merge").(bool)),
		AutoInit:          github.Bool(d.Get("auto_init").(bool)),
		LicenseTemplate:   github.String(d.Get("license_template").(string)),
		GitignoreTemplate: github.String(d.Get("gitignore_template").(string)),
		Archived:          github.Bool(d.Get("archived").(bool)),
		Topics:            expandStringList(d.Get("topics").(*schema.Set).List()),
	}
}

func resourceGithubRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	ctx := context.TODO()

	if _, ok := d.GetOk("default_branch"); ok {
		return fmt.Errorf("Cannot set the default branch on a new repository.")
	}

	repoReq := resourceGithubRepositoryObject(d)
	log.Printf("[DEBUG] create github repository %s/%s", meta.(*Organization).name, *repoReq.Name)
	repo, _, err := client.Repositories.Create(ctx, meta.(*Organization).name, repoReq)
	if err != nil {
		return err
	}
	d.SetId(*repo.Name)

	topics := repoReq.Topics
	if len(topics) > 0 {
		_, _, err = client.Repositories.ReplaceAllTopics(ctx, meta.(*Organization).name, *repoReq.Name, topics)
		if err != nil {
			return err
		}
	}

	return resourceGithubRepositoryUpdate(d, meta)
}

func resourceGithubRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	repoName := d.Id()

	log.Printf("[DEBUG] read github repository %s/%s", meta.(*Organization).name, repoName)
	repo, resp, err := client.Repositories.Get(context.TODO(), meta.(*Organization).name, repoName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Printf(
				"[WARN] removing %s/%s from state because it no longer exists in github",
				meta.(*Organization).name,
				repoName,
			)
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", repoName)
	d.Set("description", repo.Description)
	d.Set("homepage_url", repo.Homepage)
	d.Set("private", repo.Private)
	d.Set("has_issues", repo.HasIssues)
	d.Set("has_wiki", repo.HasWiki)
	d.Set("allow_merge_commit", repo.AllowMergeCommit)
	d.Set("allow_squash_merge", repo.AllowSquashMerge)
	d.Set("allow_rebase_merge", repo.AllowRebaseMerge)
	d.Set("has_downloads", repo.HasDownloads)
	d.Set("full_name", repo.FullName)
	d.Set("default_branch", repo.DefaultBranch)
	d.Set("html_url", repo.HTMLURL)
	d.Set("ssh_clone_url", repo.SSHURL)
	d.Set("svn_url", repo.SVNURL)
	d.Set("git_clone_url", repo.GitURL)
	d.Set("http_clone_url", repo.CloneURL)
	d.Set("archived", repo.Archived)
	d.Set("topics", flattenStringList(repo.Topics))
	return nil
}

func resourceGithubRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	ctx := context.TODO()
	repoReq := resourceGithubRepositoryObject(d)
	// Can only set `default_branch` on an already created repository with the target branches ref already in-place
	if v, ok := d.GetOk("default_branch"); ok {
		branch := v.(string)
		// If branch is "master", and the repository hasn't been initialized yet, setting this value will fail
		if branch != "master" {
			repoReq.DefaultBranch = &branch
		}
	}

	repoName := d.Id()
	log.Printf("[DEBUG] update github repository %s/%s", meta.(*Organization).name, repoName)
	repo, _, err := client.Repositories.Edit(ctx, meta.(*Organization).name, repoName, repoReq)
	if err != nil {
		return err
	}
	d.SetId(*repo.Name)

	if d.HasChange("topics") {
		topics := repoReq.Topics
		_, _, err = client.Repositories.ReplaceAllTopics(ctx, meta.(*Organization).name, *repo.Name, topics)
		if err != nil {
			return err
		}
	}

	return resourceGithubRepositoryRead(d, meta)
}

func resourceGithubRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	repoName := d.Id()
	log.Printf("[DEBUG] delete github repository %s/%s", meta.(*Organization).name, repoName)
	_, err := client.Repositories.Delete(context.TODO(), meta.(*Organization).name, repoName)
	return err
}
