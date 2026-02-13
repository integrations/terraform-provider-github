package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryRead,

		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
				Description:   "Full name of the repository (in org/name format).",
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"full_name"},
				Description:   "The name of the repository.",
			},
			"description": {
				Type:        schema.TypeString,
				Default:     nil,
				Optional:    true,
				Description: "A description of the repository.",
			},
			"homepage_url": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "URL of a page describing the project.",
			},
			"private": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository is private.",
			},
			"visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether the repository is public, private or internal.",
			},
			"has_issues": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository has GitHub Issues enabled.",
			},
			"has_discussions": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository has GitHub Discussions enabled.",
			},
			"has_projects": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository has the GitHub Projects enabled.",
			},
			"has_downloads": {
				Type:        schema.TypeBool,
				Computed:    true,
				Deprecated:  "This attribute is no longer in use, but it hasn't been removed yet. It will be removed in a future version. See https://github.com/orgs/community/discussions/102145#discussioncomment-8351756",
				Description: "Whether the repository has Downloads feature enabled.",
			},
			"has_wiki": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository has the GitHub Wiki enabled.",
			},
			"is_template": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository is a template repository.",
			},
			"fork": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository is a fork.",
			},
			"allow_merge_commit": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository allows merge commits.",
			},
			"allow_squash_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository allows squash merges.",
			},
			"allow_rebase_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository allows rebase merges.",
			},
			"allow_auto_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository allows auto-merging pull requests.",
			},
			"allow_update_branch": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository allows updating pull request branches.",
			},
			"allow_forking": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository allows private forking.",
			},
			"squash_merge_commit_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default value for a squash merge commit title.",
			},
			"squash_merge_commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default value for a squash merge commit message.",
			},
			"merge_commit_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default value for a merge commit title.",
			},
			"merge_commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default value for a merge commit message.",
			},
			"default_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the default branch of the repository.",
			},
			"primary_language": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The primary language used in the repository.",
			},
			"archived": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository is archived.",
			},
			"repository_license": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An array of GitHub repository licenses.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the license file (e.g., 'LICENSE').",
						},
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The path to the license file within the repository.",
						},
						"license": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The license details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A key representing the license type (e.g., 'apache-2.0').",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the license (e.g., 'Apache License 2.0').",
									},
									"url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to access information about the license on GitHub.",
									},
									"spdx_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SPDX identifier for the license (e.g., 'Apache-2.0').",
									},
									"html_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to view the license details on GitHub.",
									},
									"featured": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates if the license is featured.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A description of the license.",
									},
									"implementation": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Details about the implementation of the license.",
									},
									"permissions": {
										Type:        schema.TypeSet,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Permissions associated with the license.",
									},
									"conditions": {
										Type:        schema.TypeSet,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Conditions associated with the license.",
									},
									"limitations": {
										Type:        schema.TypeSet,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Limitations associated with the license.",
									},
									"body": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The text of the license.",
									},
								},
							},
						},
						"sha": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SHA hash of the license file.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the license file in bytes.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to access information about the license file on GitHub.",
						},
						"html_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to view the license file on GitHub.",
						},
						"git_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to access information about the license file as a Git blob.",
						},
						"download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to download the raw content of the license file.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the content (e.g., 'file').",
						},
						"content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Content of the license file, encoded by encoding scheme mentioned below.",
						},
						"encoding": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encoding used for the content (e.g., 'base64').",
						},
					},
				},
			},
			"pages": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The repository's GitHub Pages configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The source configuration for GitHub Pages.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"branch": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The branch used for GitHub Pages.",
									},
									"path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The path within the branch used for GitHub Pages.",
									},
								},
							},
						},
						"build_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of build for GitHub Pages.",
						},
						"cname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The custom domain for the GitHub Pages site.",
						},
						"custom_404": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the GitHub Pages site has a custom 404 page.",
						},
						"html_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for the GitHub Pages site.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the GitHub Pages site.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API URL for the GitHub Pages configuration.",
						},
					},
				},
			},
			"topics": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "The list of topics of the repository.",
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to the repository on the web.",
			},
			"ssh_clone_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that can be provided to git clone to clone the repository via SSH.",
			},
			"svn_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that can be provided to svn checkout to check out the repository via GitHub's Subversion protocol emulation.",
			},
			"git_clone_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that can be provided to git clone to clone the repository anonymously via the git protocol.",
			},
			"http_clone_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that can be provided to git clone to clone the repository via HTTPS.",
			},
			"template": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The repository source template configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner of the template repository.",
						},
						"repository": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the template repository.",
						},
					},
				},
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GraphQL global node ID for use with v4 API.",
			},
			"repo_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "GitHub ID for the repository.",
			},
			"delete_branch_on_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether head branches are automatically deleted on pull request merges.",
			},
		},
	}
}

func dataSourceGithubRepositoryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	var repoName string

	if fullName, ok := d.GetOk("full_name"); ok {
		var err error
		owner, repoName, err = splitRepoFullName(fullName.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if name, ok := d.GetOk("name"); ok {
		repoName = name.(string)
	}

	if repoName == "" {
		return diag.Errorf("one of %q or %q has to be provided", "full_name", "name")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] Missing GitHub repository %s/%s", owner, repoName)
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	d.SetId(repoName)

	_ = d.Set("name", repo.GetName())
	_ = d.Set("description", repo.GetDescription())
	_ = d.Set("homepage_url", repo.GetHomepage())
	_ = d.Set("private", repo.GetPrivate())
	_ = d.Set("visibility", repo.GetVisibility())
	_ = d.Set("has_issues", repo.GetHasIssues())
	_ = d.Set("has_discussions", repo.GetHasDiscussions())
	_ = d.Set("has_wiki", repo.GetHasWiki())
	_ = d.Set("is_template", repo.GetIsTemplate())
	_ = d.Set("fork", repo.GetFork())
	_ = d.Set("allow_merge_commit", repo.GetAllowMergeCommit())
	_ = d.Set("allow_squash_merge", repo.GetAllowSquashMerge())
	_ = d.Set("allow_rebase_merge", repo.GetAllowRebaseMerge())
	_ = d.Set("allow_auto_merge", repo.GetAllowAutoMerge())
	_ = d.Set("allow_forking", repo.GetAllowForking())
	_ = d.Set("squash_merge_commit_title", repo.GetSquashMergeCommitTitle())
	_ = d.Set("squash_merge_commit_message", repo.GetSquashMergeCommitMessage())
	_ = d.Set("merge_commit_title", repo.GetMergeCommitTitle())
	_ = d.Set("merge_commit_message", repo.GetMergeCommitMessage())
	_ = d.Set("has_downloads", repo.GetHasDownloads())
	_ = d.Set("full_name", repo.GetFullName())
	_ = d.Set("default_branch", repo.GetDefaultBranch())
	_ = d.Set("primary_language", repo.GetLanguage())
	_ = d.Set("html_url", repo.GetHTMLURL())
	_ = d.Set("ssh_clone_url", repo.GetSSHURL())
	_ = d.Set("svn_url", repo.GetSVNURL())
	_ = d.Set("git_clone_url", repo.GetGitURL())
	_ = d.Set("http_clone_url", repo.GetCloneURL())
	_ = d.Set("archived", repo.GetArchived())
	_ = d.Set("node_id", repo.GetNodeID())
	_ = d.Set("repo_id", repo.GetID())
	_ = d.Set("has_projects", repo.GetHasProjects())
	_ = d.Set("delete_branch_on_merge", repo.GetDeleteBranchOnMerge())
	_ = d.Set("allow_update_branch", repo.GetAllowUpdateBranch())

	if repo.GetHasPages() {
		pages, _, err := client.Repositories.GetPagesInfo(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("pages", flattenPages(pages)); err != nil {
			return diag.Errorf("error setting pages: %v", err)
		}
	} else {
		err = d.Set("pages", flattenPages(nil))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if repo.License != nil {
		repository_license, _, err := client.Repositories.License(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("repository_license", flattenRepositoryLicense(repository_license)); err != nil {
			return diag.Errorf("error setting repository_license: %v", err)
		}
	} else {
		_ = d.Set("repository_license", flattenRepositoryLicense(nil))
	}

	if repo.TemplateRepository != nil {
		err = d.Set("template", []any{
			map[string]any{
				"owner":      repo.TemplateRepository.Owner.Login,
				"repository": repo.TemplateRepository.Name,
			},
		})
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		err = d.Set("template", []any{})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err = d.Set("topics", flattenStringList(repo.Topics))
	if err != nil {
		return diag.FromErr(err)
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
