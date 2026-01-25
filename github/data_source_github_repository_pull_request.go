package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryPullRequest() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryPullRequestRead,
		Schema: map[string]*schema.Schema{
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Owner of the repository. If not provided, the provider's default owner is used.",
			},
			"base_repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the base repository to retrieve the Pull Request from.",
			},
			"number": {
				Type:        schema.TypeInt,
				Required:    true,
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
	}
}

func dataSourceGithubRepositoryPullRequestRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	if expliclitOwner, ok := d.GetOk("owner"); ok {
		owner = expliclitOwner.(string)
	}

	repository := d.Get("base_repository").(string)
	number := d.Get("number").(int)

	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repository, number)
	if err != nil {
		return err
	}

	if head := pullRequest.GetHead(); head != nil {
		if err = d.Set("head_ref", head.GetRef()); err != nil {
			return err
		}
		if err = d.Set("head_sha", head.GetSHA()); err != nil {
			return err
		}

		if headRepo := head.Repo; headRepo != nil {
			if err = d.Set("head_repository", headRepo.GetName()); err != nil {
				return err
			}
		}

		if headUser := head.User; headUser != nil {
			if err = d.Set("head_owner", headUser.GetLogin()); err != nil {
				return err
			}
		}
	}

	if base := pullRequest.GetBase(); base != nil {
		if err = d.Set("base_ref", base.GetRef()); err != nil {
			return err
		}
		if err = d.Set("base_sha", base.GetSHA()); err != nil {
			return err
		}
	}

	if err = d.Set("body", pullRequest.GetBody()); err != nil {
		return err
	}
	if err = d.Set("draft", pullRequest.GetDraft()); err != nil {
		return err
	}
	if err = d.Set("maintainer_can_modify", pullRequest.GetMaintainerCanModify()); err != nil {
		return err
	}
	if err = d.Set("number", pullRequest.GetNumber()); err != nil {
		return err
	}
	if err = d.Set("opened_at", pullRequest.GetCreatedAt().Unix()); err != nil {
		return err
	}
	if err = d.Set("state", pullRequest.GetState()); err != nil {
		return err
	}
	if err = d.Set("title", pullRequest.GetTitle()); err != nil {
		return err
	}
	if err = d.Set("updated_at", pullRequest.GetUpdatedAt().Unix()); err != nil {
		return err
	}

	if user := pullRequest.GetUser(); user != nil {
		if err = d.Set("opened_by", user.GetLogin()); err != nil {
			return err
		}
	}

	labels := []string{}
	for _, label := range pullRequest.Labels {
		labels = append(labels, label.GetName())
	}
	if err = d.Set("labels", labels); err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repository, strconv.Itoa(number)))

	return nil
}
