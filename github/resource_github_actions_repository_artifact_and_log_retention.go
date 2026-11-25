package github

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// githubDefaultArtifactAndLogRetentionDays is the retention period GitHub applies when it
// has never been customised. The API has no endpoint for clearing an override, so deletion
// restores this value instead.
const githubDefaultArtifactAndLogRetentionDays = 90

// artifactAndLogRetentionError recovers the reason GitHub rejected a retention period. These
// endpoints return the explanation in an "errors" field holding a plain string rather than the
// usual array, which go-github cannot decode, so it discards the parsed error entirely. It does
// restore the raw body on the response, which is the only place the reason survives.
func artifactAndLogRetentionError(err error) diag.Diagnostics {
	ghErr, ok := errors.AsType[*github.ErrorResponse](err)
	if !ok || ghErr.Message != "" || ghErr.Response == nil || ghErr.Response.Body == nil {
		return diag.FromErr(err)
	}

	body, readErr := io.ReadAll(ghErr.Response.Body)
	if readErr != nil {
		return diag.FromErr(err)
	}

	var payload struct {
		Message string `json:"message"`
		Errors  string `json:"errors"`
	}
	if json.Unmarshal(body, &payload) != nil || payload.Errors == "" {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{{
		Severity: diag.Error,
		Summary:  payload.Message,
		Detail:   payload.Errors,
	}}
}

func resourceGithubActionsRepositoryArtifactAndLogRetention() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRepositoryArtifactAndLogRetentionCreate,
		ReadContext:   resourceGithubActionsRepositoryArtifactAndLogRetentionRead,
		UpdateContext: resourceGithubActionsRepositoryArtifactAndLogRetentionUpdate,
		DeleteContext: resourceGithubActionsRepositoryArtifactAndLogRetentionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsRepositoryArtifactAndLogRetentionImport,
		},

		CustomizeDiff: diffRepository,

		Description: "Resource to manage how long GitHub Actions artifacts and logs are retained for a repository.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"days": {
				Type:             schema.TypeInt,
				Required:         true,
				Description:      "Number of days to retain artifacts and logs. Must not exceed 'maximum_allowed_days', which the organization or enterprise imposes.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 400)),
			},
			"maximum_allowed_days": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Longest retention period the repository is allowed to set, as limited by the organization or enterprise.",
			},
		},
	}
}

func resourceGithubActionsRepositoryArtifactAndLogRetentionCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	days, _ := d.Get("days").(int)

	if _, err := client.Repositories.UpdateArtifactAndLogRetentionPeriod(ctx, owner, repoName, github.ArtifactPeriodOpt{
		Days: new(days),
	}); err != nil {
		return artifactAndLogRetentionError(err)
	}

	d.SetId(repoName)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubActionsRepositoryArtifactAndLogRetentionRead(ctx, d, m)
}

func resourceGithubActionsRepositoryArtifactAndLogRetentionRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName := d.Id()

	retention, _, err := client.Repositories.GetArtifactAndLogRetentionPeriod(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing artifact and log retention from state because the repository no longer exists in GitHub", map[string]any{"repository": repoName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("days", retention.GetDays()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("maximum_allowed_days", retention.GetMaximumAllowedDays()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryArtifactAndLogRetentionUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	days, _ := d.Get("days").(int)

	if _, err := client.Repositories.UpdateArtifactAndLogRetentionPeriod(ctx, owner, repoName, github.ArtifactPeriodOpt{
		Days: new(days),
	}); err != nil {
		return artifactAndLogRetentionError(err)
	}

	return resourceGithubActionsRepositoryArtifactAndLogRetentionRead(ctx, d, m)
}

func resourceGithubActionsRepositoryArtifactAndLogRetentionDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName := d.Id()

	retention, _, err := client.Repositories.GetArtifactAndLogRetentionPeriod(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	// A repository may not exceed the period its organization or enterprise allows, so the
	// reset target is capped by that ceiling rather than always being GitHub's default.
	days := min(githubDefaultArtifactAndLogRetentionDays, retention.GetMaximumAllowedDays())

	if _, err := client.Repositories.UpdateArtifactAndLogRetentionPeriod(ctx, owner, repoName, github.ArtifactPeriodOpt{
		Days: new(days),
	}); err != nil {
		return artifactAndLogRetentionError(err)
	}

	return nil
}

func resourceGithubActionsRepositoryArtifactAndLogRetentionImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName := d.Id()

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
