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
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"full_name"},
			},
			"only_protected_branches": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"branches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protected": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
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
