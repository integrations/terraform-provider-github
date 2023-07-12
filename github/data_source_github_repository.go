package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	err = d.Set("name", repo.GetName())
	if err != nil {
		return err
	}
	err = d.Set("description", repo.GetDescription())
	if err != nil {
		return err
	}
	err = d.Set("homepage_url", repo.GetHomepage())
	if err != nil {
		return err
	}
	err = d.Set("private", repo.GetPrivate())
	if err != nil {
		return err
	}
	err = d.Set("visibility", repo.GetVisibility())
	if err != nil {
		return err
	}
	err = d.Set("has_issues", repo.GetHasIssues())
	if err != nil {
		return err
	}
	err = d.Set("has_discussions", repo.GetHasDiscussions())
	if err != nil {
		return err
	}
	err = d.Set("has_wiki", repo.GetHasWiki())
	if err != nil {
		return err
	}
	err = d.Set("is_template", repo.GetIsTemplate())
	if err != nil {
		return err
	}
	err = d.Set("fork", repo.GetFork())
	if err != nil {
		return err
	}
	err = d.Set("allow_merge_commit", repo.GetAllowMergeCommit())
	if err != nil {
		return err
	}
	err = d.Set("allow_squash_merge", repo.GetAllowSquashMerge())
	if err != nil {
		return err
	}
	err = d.Set("allow_rebase_merge", repo.GetAllowRebaseMerge())
	if err != nil {
		return err
	}
	err = d.Set("allow_auto_merge", repo.GetAllowAutoMerge())
	if err != nil {
		return err
	}
	err = d.Set("squash_merge_commit_title", repo.GetSquashMergeCommitTitle())
	if err != nil {
		return err
	}
	err = d.Set("squash_merge_commit_message", repo.GetSquashMergeCommitMessage())
	if err != nil {
		return err
	}
	err = d.Set("merge_commit_title", repo.GetMergeCommitTitle())
	if err != nil {
		return err
	}
	err = d.Set("merge_commit_message", repo.GetMergeCommitMessage())
	if err != nil {
		return err
	}
	err = d.Set("has_downloads", repo.GetHasDownloads())
	if err != nil {
		return err
	}
	err = d.Set("full_name", repo.GetFullName())
	if err != nil {
		return err
	}
	err = d.Set("default_branch", repo.GetDefaultBranch())
	if err != nil {
		return err
	}
	err = d.Set("html_url", repo.GetHTMLURL())
	if err != nil {
		return err
	}
	err = d.Set("ssh_clone_url", repo.GetSSHURL())
	if err != nil {
		return err
	}
	err = d.Set("svn_url", repo.GetSVNURL())
	if err != nil {
		return err
	}
	err = d.Set("git_clone_url", repo.GetGitURL())
	if err != nil {
		return err
	}
	err = d.Set("http_clone_url", repo.GetCloneURL())
	if err != nil {
		return err
	}
	err = d.Set("archived", repo.GetArchived())
	if err != nil {
		return err
	}
	err = d.Set("node_id", repo.GetNodeID())
	if err != nil {
		return err
	}
	err = d.Set("repo_id", repo.GetID())
	if err != nil {
		return err
	}
	err = d.Set("has_projects", repo.GetHasProjects())
	if err != nil {
		return err
	}

	if repo.GetHasPages() {
		pages, _, err := client.Repositories.GetPagesInfo(context.TODO(), owner, repoName)
		if err != nil {
			return err
		}
		if err := d.Set("pages", flattenPages(pages)); err != nil {
			return fmt.Errorf("error setting pages: %w", err)
		}
	} else {
		err = d.Set("pages", flattenPages(nil))
		if err != nil {
			return err
		}
	}

	if repo.TemplateRepository != nil {
		err = d.Set("template", []interface{}{
			map[string]interface{}{
				"owner":      repo.TemplateRepository.Owner.Login,
				"repository": repo.TemplateRepository.Name,
			},
		})
		if err != nil {
			return err
		}
	} else {
		err = d.Set("template", []interface{}{})
		if err != nil {
			return err
		}
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
