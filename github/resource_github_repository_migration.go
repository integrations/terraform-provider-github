package github

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryResourceV0() *schema.Resource {
	return &schema.Resource{
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
			"only_protected_branches": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Only consider protected branches when looking for default branch.",
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
				Description: "Set to true to create a private repository.",
			},
			"visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Can be public or private.",
			},
			"has_issues": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true to enable the GitHub Issues features on the repository.",
			},
			"has_projects": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true to enable the GitHub Projects features on the repository.",
			},
			"has_downloads": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true to enable the (deprecated) downloads features on the repository.",
			},
			"has_wiki": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true to enable the GitHub Wiki features on the repository.",
			},
			"allow_merge_commit": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to false to disable merge commits on the repository.",
			},
			"allow_squash_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to false to disable squash merges on the repository.",
			},
			"allow_rebase_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to false to disable rebase merges on the repository.",
			},
			"allow_auto_merge": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true to allow auto-merging pull requests on the repository.",
			},
			"squash_merge_commit_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Can be PR_TITLE or COMMIT_OR_PR_TITLE.",
			},
			"squash_merge_commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Can be PR_BODY, COMMIT_MESSAGES, or BLANK.",
			},
			"merge_commit_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Can be PR_TITLE or MERGE_MESSAGE.",
			},
			"merge_commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Can be PR_BODY, PR_TITLE, or BLANK.",
			},
			"default_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default branch of the repository.",
			},
			"archived": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies if the repository should be archived.",
			},
			"branches": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Deprecated: Use github_branch data source instead. The list of branches.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the branch.",
						},
						"protected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the branch is protected.",
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
							Description: "The source branch and directory for the rendered Pages site.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"branch": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The repository branch used to publish the site's source files.",
									},
									"path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The repository directory from which the site publishes.",
									},
								},
							},
						},
						"cname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The custom domain for the repository.",
						},
						"custom_404": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the rendered site has a custom 404 page.",
						},
						"html_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTML URL for the rendered site.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub Pages site's build status.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for the Pages site.",
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
				Description: "URL that can be provided to svn checkout to check out the repository via SVN.",
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
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GraphQL global node id for use with v4 API.",
			},
			"repo_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "GitHub ID for the repository.",
			},
		},
	}
}

func resourceGithubRepositoryInstanceStateUpgradeV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Repository State before migration: %#v", rawState)

	prefix := "branches."

	for k := range rawState {
		if strings.HasPrefix(k, prefix) {
			delete(rawState, k)
		}
	}

	log.Printf("[DEBUG] GitHub Repository State after migration: %#v", rawState)
	return rawState, nil
}
