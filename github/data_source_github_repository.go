package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v54/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryRead,

		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"full_name"},
			},
			"description": {
				Type:     schema.TypeString,
				Default:  nil,
				Optional: true,
			},
			"homepage_url": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
			},
			"private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"has_issues": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_discussions": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_projects": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_downloads": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_wiki": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_template": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fork": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_merge_commit": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_squash_merge": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_rebase_merge": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_auto_merge": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"squash_merge_commit_title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"squash_merge_commit_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"merge_commit_title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"merge_commit_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"pages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"branch": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"build_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_404": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"html_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"topics": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"template": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repo_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	var repoName string

	if fullName, ok := d.GetOk("full_name"); ok {
		var err error
		owner, repoName, err = splitRepoFullName(fullName.(string))
		if err != nil {
			return err
		}
	}
	if name, ok := d.GetOk("name"); ok {
		repoName = name.(string)
	}

	if repoName == "" {
		return fmt.Errorf("one of %q or %q has to be provided", "full_name", "name")
	}

	repo, _, err := client.Repositories.Get(context.TODO(), owner, repoName)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] Missing GitHub repository %s/%s", owner, repoName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(repoName)

	d.Set("name", repo.GetName())
	d.Set("description", repo.GetDescription())
	d.Set("homepage_url", repo.GetHomepage())
	d.Set("private", repo.GetPrivate())
	d.Set("visibility", repo.GetVisibility())
	d.Set("has_issues", repo.GetHasIssues())
	d.Set("has_discussions", repo.GetHasDiscussions())
	d.Set("has_wiki", repo.GetHasWiki())
	d.Set("is_template", repo.GetIsTemplate())
	d.Set("fork", repo.GetFork())
	d.Set("allow_merge_commit", repo.GetAllowMergeCommit())
	d.Set("allow_squash_merge", repo.GetAllowSquashMerge())
	d.Set("allow_rebase_merge", repo.GetAllowRebaseMerge())
	d.Set("allow_auto_merge", repo.GetAllowAutoMerge())
	d.Set("squash_merge_commit_title", repo.GetSquashMergeCommitTitle())
	d.Set("squash_merge_commit_message", repo.GetSquashMergeCommitMessage())
	d.Set("merge_commit_title", repo.GetMergeCommitTitle())
	d.Set("merge_commit_message", repo.GetMergeCommitMessage())
	d.Set("has_downloads", repo.GetHasDownloads())
	d.Set("full_name", repo.GetFullName())
	d.Set("default_branch", repo.GetDefaultBranch())
	d.Set("primary_language", repo.GetLanguage())
	d.Set("html_url", repo.GetHTMLURL())
	d.Set("ssh_clone_url", repo.GetSSHURL())
	d.Set("svn_url", repo.GetSVNURL())
	d.Set("git_clone_url", repo.GetGitURL())
	d.Set("http_clone_url", repo.GetCloneURL())
	d.Set("archived", repo.GetArchived())
	d.Set("node_id", repo.GetNodeID())
	d.Set("repo_id", repo.GetID())
	d.Set("has_projects", repo.GetHasProjects())

	if repo.GetHasPages() {
		pages, _, err := client.Repositories.GetPagesInfo(context.TODO(), owner, repoName)
		if err != nil {
			return err
		}
		if err := d.Set("pages", flattenPages(pages)); err != nil {
			return fmt.Errorf("error setting pages: %w", err)
		}
	} else {
		d.Set("pages", flattenPages(nil))
	}

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

	err = d.Set("topics", flattenStringList(repo.Topics))
	if err != nil {
		return err
	}

	return nil
}

func splitRepoFullName(fullName string) (string, string, error) {
	parts := strings.Split(fullName, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected full name format (%q), expected owner/repo_name", fullName)
	}
	return parts[0], parts[1], nil
}
