package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func dataSourceGithubRepositoryPullRequestRead(d *schema.ResourceData, meta interface{}) error {
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
		d.Set("head_ref", head.GetRef())
		d.Set("head_sha", head.GetSHA())

		if headRepo := head.Repo; headRepo != nil {
			d.Set("head_repository", headRepo.GetName())
		}

		if headUser := head.User; headUser != nil {
			d.Set("head_owner", headUser.GetLogin())
		}
	}

	if base := pullRequest.GetBase(); base != nil {
		d.Set("base_ref", base.GetRef())
		d.Set("base_sha", base.GetSHA())
	}

	d.Set("body", pullRequest.GetBody())
	d.Set("draft", pullRequest.GetDraft())
	d.Set("maintainer_can_modify", pullRequest.GetMaintainerCanModify())
	d.Set("number", pullRequest.GetNumber())
	d.Set("opened_at", pullRequest.GetCreatedAt().Unix())
	d.Set("state", pullRequest.GetState())
	d.Set("title", pullRequest.GetTitle())
	d.Set("updated_at", pullRequest.GetUpdatedAt().Unix())

	if user := pullRequest.GetUser(); user != nil {
		d.Set("opened_by", user.GetLogin())
	}

	labels := []string{}
	for _, label := range pullRequest.Labels {
		labels = append(labels, label.GetName())
	}
	d.Set("labels", labels)

	d.SetId(buildThreePartID(owner, repository, strconv.Itoa(number)))

	return nil
}
