package github

import (
	"context"
	"strings"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Docs: https://docs.github.com/en/rest/reference/pulls#list-pull-requests
func dataSourceGithubRepositoryPullRequests() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the list of pull requests in a repository.",
		Read:        dataSourceGithubRepositoryPullRequestsRead,
		Schema: map[string]*schema.Schema{
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Owner of the repository. If not provided, the provider's default owner is used.",
			},
			"base_repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the base repository to retrieve the Pull Requests from.",
			},
			"base_ref": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If set, filters Pull Requests by base branch name.",
			},
			"head_ref": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If set, filters Pull Requests by head user or head organization and branch name.",
			},
			"sort_by": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "created",
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"created", "updated", "popularity", "long-running"}, false), "sort_by"),
				Description:      "Indicates what to sort results by. Can be 'created', 'updated', 'popularity', or 'long-running'.",
			},
			"sort_direction": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "asc",
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"asc", "desc"}, false), "sort_direction"),
				Description:      "The direction of the sort. Can be 'asc' or 'desc'.",
			},
			"state": {
				Type:             schema.TypeString,
				Default:          "open",
				Optional:         true,
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"open", "closed", "all"}, false), "state"),
				Description:      "Filters Pull Requests by state. Can be 'open', 'closed', or 'all'.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of Pull Requests matching the filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of the Pull Request within the repository.",
						},
						"base_ref": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the ref (branch) of the Pull Request base.",
						},
						"base_sha": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Head commit SHA of the Pull Request base.",
						},
						"body": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Body of the Pull Request.",
						},
						"draft": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this Pull Request is a draft.",
						},
						"head_owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Owner of the Pull Request head repository.",
						},
						"head_ref": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value of the Pull Request HEAD reference.",
						},
						"head_repository": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Pull Request head repository.",
						},
						"head_sha": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Head commit SHA of the Pull Request head.",
						},
						"labels": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "List of label names set on the Pull Request.",
						},
						"maintainer_can_modify": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the base repository maintainers can modify the Pull Request.",
						},
						"opened_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Unix timestamp indicating the Pull Request creation time.",
						},
						"opened_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GitHub login of the user who opened the Pull Request.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current Pull Request state - can be 'open', 'closed' or 'merged'.",
						},
						"title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The title of the Pull Request.",
						},
						"updated_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timestamp of the last Pull Request update.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryPullRequestsRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	if explicitOwner, ok := d.GetOk("owner"); ok {
		owner = explicitOwner.(string)
	}

	baseRepository := d.Get("base_repository").(string)
	state := d.Get("state").(string)
	head := d.Get("head_ref").(string)
	base := d.Get("base_ref").(string)
	sort := d.Get("sort_by").(string)
	direction := d.Get("sort_direction").(string)

	options := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
		State:       state,
		Head:        head,
		Base:        base,
		Sort:        sort,
		Direction:   direction,
	}

	results := make([]map[string]any, 0)

	for {
		pullRequests, resp, err := client.PullRequests.List(ctx, owner, baseRepository, options)
		if err != nil {
			return err
		}

		for _, pullRequest := range pullRequests {
			result := map[string]any{
				"number":                pullRequest.GetNumber(),
				"body":                  pullRequest.GetBody(),
				"draft":                 pullRequest.GetDraft(),
				"maintainer_can_modify": pullRequest.GetMaintainerCanModify(),
				"opened_at":             pullRequest.GetCreatedAt().Unix(),
				"state":                 pullRequest.GetState(),
				"title":                 pullRequest.GetTitle(),
				"updated_at":            pullRequest.GetUpdatedAt().Unix(),
			}

			if head := pullRequest.GetHead(); head != nil {
				result["head_ref"] = head.GetRef()
				result["head_sha"] = head.GetSHA()

				if headRepo := head.GetRepo(); headRepo != nil {
					result["head_repository"] = headRepo.GetName()

					if headOwner := headRepo.GetOwner(); headOwner != nil {
						result["head_owner"] = headOwner.GetLogin()
					}
				}
			}

			if base := pullRequest.GetBase(); base != nil {
				result["base_ref"] = base.GetRef()
				result["base_sha"] = base.GetSHA()
			}

			labels := []string{}
			for _, label := range pullRequest.Labels {
				labels = append(labels, label.GetName())
			}
			result["labels"] = labels

			if user := pullRequest.GetUser(); user != nil {
				result["opened_by"] = user.GetLogin()
			}

			results = append(results, result)
		}

		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(strings.Join([]string{
		owner,
		baseRepository,
		state,
		head,
		base,
		sort,
		direction,
	}, "/"))

	if err := d.Set("results", results); err != nil {
		return err
	}

	return nil
}
