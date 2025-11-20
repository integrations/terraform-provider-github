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
				Type:     schema.TypeString,
				Optional: true,
			},
			"base_repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"number": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"base_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"base_sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"body": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"draft": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"head_owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"head_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"head_repository": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"head_sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "List of names of labels on the PR",
			},
			"maintainer_can_modify": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"opened_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"opened_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username of the PR creator",
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRepositoryPullRequestRead(d *schema.ResourceData, meta any) error {
	ctx := context.TODO()
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
