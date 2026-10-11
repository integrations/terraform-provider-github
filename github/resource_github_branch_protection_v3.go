package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubBranchProtectionV3() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubBranchProtectionV3Create,
		ReadContext:   resourceGithubBranchProtectionV3Read,
		UpdateContext: resourceGithubBranchProtectionV3Update,
		DeleteContext: resourceGithubBranchProtectionV3Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: diffETag,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository name.",
			},
			"branch": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Git branch to protect.",
			},
			"required_status_checks": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Enforce restrictions for required status checks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_admins": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    false,
							Deprecated: "Use enforce_admins instead",
							DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
								return true
							},
						},
						"strict": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Require branches to be up to date before merging.",
						},
						"contexts": {
							Type:       schema.TypeSet,
							Optional:   true,
							Computed:   true,
							Deprecated: "GitHub is deprecating the use of `contexts`. Use a `checks` array instead.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"checks": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the 'context' and 'app_id' like so 'context:app_id'",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							ConflictsWith: []string{"required_status_checks.0.contexts"},
						},
					},
				},
			},
			"required_pull_request_reviews": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Enforce restrictions for pull request reviews.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// FIXME: Remove this deprecated field
						"include_admins": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    false,
							Deprecated: "Use enforce_admins instead",
							DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
								return true
							},
						},
						"dismiss_stale_reviews": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Dismiss approved reviews automatically when a new commit is pushed.",
						},
						"dismissal_users": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of user logins with dismissal access.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"dismissal_teams": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of team slugs with dismissal access. Always use slug of the team, not its name. Each team already has to have access to the repository.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"dismissal_apps": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of apps slugs with dismissal access. Always use slug of the app, not its name. Each app already has to have access to the repository.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"require_code_owner_reviews": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Require an approved review in pull requests including files with a designated code owner.",
						},
						"required_approving_review_count": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          1,
							Description:      "Require 'x' number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6.",
							ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 6)),
						},
						"require_last_push_approval": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Require that the most recent push must be approved by someone other than the last pusher.",
						},
						"bypass_pull_request_allowances": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"teams": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"apps": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"restrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Enforce restrictions for the users and teams that may push to the branch.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of user logins with push access.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"teams": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of team slugs with push access. Always use slug of the team, not its name. Each team already has to have access to the repository.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"apps": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of app slugs with push access.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"enforce_admins": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' enforces status checks for repository administrators.",
			},
			"require_signed_commits": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' requires all commits to be signed with GPG.",
			},
			"require_conversation_resolution": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' requires all conversations on code must be resolved before a pull request can be merged.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the branch protection object.",
			},
		},
	}
}

func resourceGithubBranchProtectionV3Create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	branch := d.Get("branch").(string)

	protectionRequest, err := buildProtectionRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}

	protection, _, err := client.Repositories.UpdateBranchProtection(ctx,
		orgName,
		repoName,
		branch,
		protectionRequest,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := checkBranchRestrictionsUsers(protection.GetRestrictions(), protectionRequest.GetRestrictions()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(repoName, branch))

	if err = requireSignedCommitsUpdate(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubBranchProtectionV3Read(ctx, d, meta)
}

func resourceGithubBranchProtectionV3Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	repoName, branch, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	orgName := meta.(*Owner).name

	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	githubProtection, resp, err := client.Repositories.GetBranchProtection(ctx,
		orgName, repoName, branch)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				if err := requireSignedCommitsRead(ctx, d, meta); err != nil {
					return diag.FromErr(fmt.Errorf("error setting signed commit restriction: %w", err))
				}
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing branch protection from state because it no longer exists in GitHub", map[string]any{"owner": orgName, "repository": repoName, "branch": branch})
				d.SetId("")
				return nil
			}
		}

		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("branch", branch); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("enforce_admins", githubProtection.GetEnforceAdmins().Enabled); err != nil {
		return diag.FromErr(err)
	}
	if rcr := githubProtection.GetRequiredConversationResolution(); rcr != nil {
		if err = d.Set("require_conversation_resolution", rcr.Enabled); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := flattenAndSetRequiredStatusChecks(d, githubProtection); err != nil {
		return diag.FromErr(fmt.Errorf("error setting required_status_checks: %w", err))
	}

	if err := flattenAndSetRequiredPullRequestReviews(d, githubProtection); err != nil {
		return diag.FromErr(fmt.Errorf("error setting required_pull_request_reviews: %w", err))
	}

	if err := flattenAndSetRestrictions(d, githubProtection); err != nil {
		return diag.FromErr(fmt.Errorf("error setting restrictions: %w", err))
	}

	if err := requireSignedCommitsRead(ctx, d, meta); err != nil {
		return diag.FromErr(fmt.Errorf("error setting signed commit restriction: %w", err))
	}

	return nil
}

func resourceGithubBranchProtectionV3Update(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	repoName, branch, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	protectionRequest, err := buildProtectionRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}

	orgName := meta.(*Owner).name
	ctx = context.WithValue(ctx, ctxId, d.Id())

	protection, _, err := client.Repositories.UpdateBranchProtection(ctx,
		orgName,
		repoName,
		branch,
		protectionRequest,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := checkBranchRestrictionsUsers(protection.GetRestrictions(), protectionRequest.GetRestrictions()); err != nil {
		return diag.FromErr(err)
	}

	if protectionRequest.RequiredPullRequestReviews == nil {
		_, err = client.Repositories.RemovePullRequestReviewEnforcement(ctx,
			orgName,
			repoName,
			branch,
		)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(buildTwoPartID(repoName, branch))

	if err = requireSignedCommitsUpdate(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubBranchProtectionV3Read(ctx, d, meta)
}

func resourceGithubBranchProtectionV3Delete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	repoName, branch, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	orgName := meta.(*Owner).name
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err = client.Repositories.RemoveBranchProtection(ctx,
		orgName, repoName, branch)
	return diag.FromErr(err)
}
