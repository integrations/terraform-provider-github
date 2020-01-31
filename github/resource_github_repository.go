package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceGithubRepository() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner": {
							Type:     schema.TypeString,
							Required: true,
						},
						"repository": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},

		Create: resourceGithubRepositoryCreate,
		Read:   resourceGithubRepositoryRead,
		Update: resourceGithubRepositoryUpdate,
		Delete: resourceGithubRepositoryDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				d.Set("auto_init", false)
				return []*schema.ResourceData{d}, nil
			},
		},

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryUpgradeV0,
				Version: 0,
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
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	if branchName, hasDefaultBranch := d.GetOk("default_branch"); hasDefaultBranch && (branchName != "master") {
		return fmt.Errorf("Cannot set the default branch on a new repository to something other than 'master'.")
	}

	repoReq := resourceGithubRepositoryObject(d)
	orgName := meta.(*Organization).name
	repoName := repoReq.GetName()
	ctx := context.Background()

	log.Printf("[DEBUG] Creating repository: %s/%s", orgName, repoName)

	if template, ok := d.GetOk("template"); ok {
		templateConfigBlocks := template.([]interface{})

		for _, templateConfigBlock := range templateConfigBlocks {
			templateConfigMap, ok := templateConfigBlock.(map[string]interface{})
			if !ok {
				return errors.New("failed to unpack template configuration block")
			}

			templateRepo := templateConfigMap["repository"].(string)
			templateRepoOwner := templateConfigMap["owner"].(string)
			templateRepoReq := github.TemplateRepoRequest{
				Name:        &repoName,
				Owner:       &orgName,
				Description: github.String(d.Get("description").(string)),
				Private:     github.Bool(d.Get("private").(bool)),
			}

			repo, _, err := client.Repositories.CreateFromTemplate(ctx,
				templateRepoOwner,
				templateRepo,
				&templateRepoReq,
			)

			if err != nil {
				return err
			}

			d.SetId(repo.GetNodeID())
		}
	} else {
		// Create without a repository template
		repo, _, err := client.Repositories.Create(ctx, orgName, repoReq)
		if err != nil {
			return err
		}
		d.SetId(repo.GetNodeID())
	}

	topics := repoReq.Topics
	if len(topics) > 0 {
		_, _, err = client.Repositories.ReplaceAllTopics(ctx, orgName, repoName, topics)
		if err != nil {
			return err
		}
	}

	return resourceGithubRepositoryUpdate(d, meta)
}

func resourceGithubRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	orgName := meta.(*Organization).name
	repoName := d.Get("name").(string)

	log.Printf("[DEBUG] Reading repository: %s/%s", orgName, repoName)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	repo, resp, err := client.Repositories.Get(ctx, orgName, repoName)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing repository %s/%s from state because it no longer exists in GitHub",
					orgName, repoName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("name", repoName)
	d.Set("description", repo.Description)
	d.Set("homepage_url", repo.Homepage)
	d.Set("private", repo.Private)
	d.Set("has_issues", repo.HasIssues)
	d.Set("has_projects", repo.HasProjects)
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

func resourceGithubRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	repoReq := resourceGithubRepositoryObject(d)
	// Can only set `default_branch` on an already created repository with the target branches ref already in-place
	if v, ok := d.GetOk("default_branch"); ok {
		branch := v.(string)
		// If branch is "master", and the repository hasn't been initialized yet, setting this value will fail
		if branch != "master" {
			repoReq.DefaultBranch = &branch
		}
	}

	repoName := d.Get("name").(string)
	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Updating repository: %s/%s", orgName, repoName)
	repo, _, err := client.Repositories.Edit(ctx, orgName, repoName, repoReq)
	if err != nil {
		return err
	}
	d.SetId(repo.GetNodeID())

	if d.HasChange("topics") {
		topics := repoReq.Topics
		_, _, err = client.Repositories.ReplaceAllTopics(ctx, orgName, *repo.Name, topics)
		if err != nil {
			return err
		}
	}

	return resourceGithubRepositoryRead(d, meta)
}

func resourceGithubRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	repoName := d.Get("name").(string)
	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting repository: %s/%s", orgName, repoName)
	_, err = client.Repositories.Delete(ctx, orgName, repoName)

	return err
}
