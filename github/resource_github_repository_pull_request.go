package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryPullRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryPullRequestCreate,
		ReadContext:   resourceGithubRepositoryPullRequestRead,
		UpdateContext: resourceGithubRepositoryPullRequestUpdate,
		DeleteContext: resourceGithubRepositoryPullRequestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
				_, baseRepository, _, err := parsePullRequestID(d)
				if err != nil {
					return nil, err
				}
				if err := d.Set("base_repository", baseRepository); err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Owner of the repository. If not provided, the provider's default owner is used.",
			},
			"base_repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the base repository to retrieve the Pull Requests from.",
			},
			"base_ref": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the branch serving as the base of the Pull Request.",
			},
			"head_ref": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the branch serving as the head of the Pull Request.",
			},
			"title": {
				// Even though the documentation does not explicitly mark the
				// title field as required, attempts to create a PR with an
				// empty title result in a "missing_field" validation error
				// (HTTP 422).
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the Pull Request.",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Body of the Pull Request.",
			},
			"maintainer_can_modify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Controls whether the base repository maintainers can modify the Pull Request. Default: 'false'.",
			},
			"base_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Head commit SHA of the Pull Request base.",
			},
			"draft": {
				// The "draft" field is an interesting corner case because while
				// you can create a draft PR through the API, the documentation
				// does not indicate that you can change this field during
				// update:
				//
				// https://docs.github.com/en/rest/reference/pulls#update-a-pull-request
				//
				// And since you cannot manage the lifecycle of this field to
				// reconcile the actual state with the desired one, this field
				// cannot be managed by Terraform.
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates Whether this Pull Request is a draft.",
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
				Description: "List of names of labels on the PR",
			},
			"number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of the Pull Request within the repository.",
			},
			"opened_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Unix timestamp indicating the Pull Request creation time.",
			},
			"opened_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username of the PR creator",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current Pull Request state - can be 'open', 'closed' or 'merged'.",
			},
			"updated_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The timestamp of the last Pull Request update.",
			},
		},
	}
}

func resourceGithubRepositoryPullRequestCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	// For convenience, by default we expect that the base repository and head
	// repository owners are the same, and both belong to the caller, indicating
	// a "PR within the same repo" scenario. The head will *always* belong to
	// the current caller, the base - not necessarily. The base will belong to
	// another namespace in case of forks, and this resource supports them.
	headOwner := meta.(*Owner).name

	baseOwner := headOwner
	if explicitBaseOwner, ok := d.GetOk("owner"); ok {
		baseOwner = explicitBaseOwner.(string)
	}

	title, _ := d.Get("title").(string)
	baseRepository, _ := d.Get("base_repository").(string)
	head, _ := d.Get("head_ref").(string)
	body, _ := d.Get("body").(string)
	maintainerCanModify, _ := d.Get("maintainer_can_modify").(bool)

	if headOwner != baseOwner {
		head = strings.Join([]string{headOwner, head}, ":")
	}

	base, _ := d.Get("base_ref").(string)

	pullRequest, _, err := client.PullRequests.Create(ctx, baseOwner, baseRepository, github.CreatePullRequest{
		Title:               new(title),
		Head:                head,
		Base:                base,
		Body:                new(body),
		MaintainerCanModify: new(maintainerCanModify),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildThreePartID(baseOwner, baseRepository, strconv.Itoa(pullRequest.GetNumber())))

	return resourceGithubRepositoryPullRequestRead(ctx, d, meta)
}

func resourceGithubRepositoryPullRequestRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner, repository, number, err := parsePullRequestID(d)
	if err != nil {
		return diag.FromErr(err)
	}

	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repository, number)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("number", pullRequest.GetNumber()); err != nil {
		return diag.FromErr(err)
	}

	if head := pullRequest.GetHead(); head != nil {
		if err = d.Set("head_ref", head.GetRef()); err != nil {
			return diag.FromErr(err)
		}

		if err = d.Set("head_sha", head.GetSHA()); err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Totally unexpected condition. Better do that than segfault, I guess?
		tflog.Info(ctx, "Head branch missing", map[string]any{"head_ref": d.Get("head_ref")})
		d.SetId("")
		return nil
	}

	if base := pullRequest.GetBase(); base != nil {
		if err = d.Set("base_ref", base.GetRef()); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("base_sha", base.GetSHA()); err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Seme logic as with the missing head branch.
		tflog.Info(ctx, "Base branch missing", map[string]any{"base_ref": d.Get("base_ref")})
		d.SetId("")
		return nil
	}

	if err = d.Set("body", pullRequest.GetBody()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", pullRequest.GetTitle()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("draft", pullRequest.GetDraft()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("maintainer_can_modify", pullRequest.GetMaintainerCanModify()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("number", pullRequest.GetNumber()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("state", pullRequest.GetState()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("opened_at", pullRequest.GetCreatedAt().Unix()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("updated_at", pullRequest.GetUpdatedAt().Unix()); err != nil {
		return diag.FromErr(err)
	}

	if user := pullRequest.GetUser(); user != nil {
		if err = d.Set("opened_by", user.GetLogin()); err != nil {
			return diag.FromErr(err)
		}
	}

	labels := []string{}
	for _, label := range pullRequest.Labels {
		labels = append(labels, label.GetName())
	}
	if err = d.Set("labels", labels); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryPullRequestUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner, repository, number, err := parsePullRequestID(d)
	if err != nil {
		return diag.FromErr(err)
	}

	update := &github.PullRequest{
		Title:               new(d.Get("title").(string)),
		Body:                new(d.Get("body").(string)),
		MaintainerCanModify: new(d.Get("maintainer_can_modify").(bool)),
	}

	if d.HasChange("base_ref") {
		update.Base = &github.PullRequestBranch{
			Ref: new(d.Get("base_ref").(string)),
		}
	}

	_, _, err = client.PullRequests.Edit(ctx, owner, repository, number, update)
	if err == nil {
		return resourceGithubRepositoryPullRequestRead(ctx, d, meta)
	}

	errs := []string{fmt.Sprintf("could not update the Pull Request: %v", err)}

	if diags := resourceGithubRepositoryPullRequestRead(ctx, d, meta); diags.HasError() {
		errs = append(errs, fmt.Sprintf("could not read the Pull Request after the failed update: %s", diags[0].Summary))
	}

	return diag.FromErr(fmt.Errorf("%s", strings.Join(errs, ", ")))
}

func resourceGithubRepositoryPullRequestDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// It's not entirely clear how to treat PR deletion according to Terraform's
	// CRUD semantics. The approach we're taking here is to close the PR unless
	// it's already closed or merged. Merging it feels intuitively wrong in what
	// effectively is a destructor.
	if d.Get("state").(string) != "open" {
		d.SetId("")
		return nil
	}

	client := meta.(*Owner).v3client

	owner, repository, number, err := parsePullRequestID(d)
	if err != nil {
		return diag.FromErr(err)
	}

	update := &github.PullRequest{State: new("closed")}
	if _, _, err = client.PullRequests.Edit(ctx, owner, repository, number, update); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func parsePullRequestID(d *schema.ResourceData) (owner, repository string, number int, err error) {
	var strNumber string

	if owner, repository, strNumber, err = parseID3(d.Id()); err != nil {
		return owner, repository, number, err
	}

	if number, err = strconv.Atoi(strNumber); err != nil {
		err = fmt.Errorf("invalid PR number %s: %w", strNumber, err)
	}

	return owner, repository, number, err
}
