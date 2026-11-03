package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	codeQualityStateConfigured    = "configured"
	codeQualityStateNotConfigured = "not-configured"
)

func resourceGithubRepositoryCodeQuality() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryCodeQualityCreate,
		ReadContext:   resourceGithubRepositoryCodeQualityRead,
		UpdateContext: resourceGithubRepositoryCodeQualityUpdate,
		DeleteContext: resourceGithubRepositoryCodeQualityDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryCodeQualityImport,
		},

		Description: "Manages the GitHub Code Quality setup for a repository.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository to configure Code Quality for.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the repository to configure Code Quality for.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether Code Quality is enabled for the repository. Defaults to `true`.",
			},
			"languages": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateValueFunc([]string{"csharp", "go", "java-kotlin", "javascript-typescript", "python", "ruby"}),
				},
				Description: "The languages to analyze. Can be `csharp`, `go`, `java-kotlin`, `javascript-typescript`, `python` or `ruby`. If not set, GitHub selects the languages detected in the repository.",
			},
			"runner_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validateValueFunc([]string{"standard", "labeled"}),
				Description:      "The runner type used for the analysis. Can be `standard` or `labeled`.",
			},
			"runner_label": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The runner label used for the analysis when `runner_type` is `labeled`.",
			},
		},

		CustomizeDiff: diffRepository,
	}
}

func resourceGithubRepositoryCodeQualityCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Creating repository code quality setup", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	if repo.GetArchived() {
		return diag.Errorf("cannot configure code quality on archived repository %s/%s", owner, repoName)
	}

	if err := updateCodeQualitySetup(ctx, client, owner, repoName, expandCodeQualitySetup(d)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(repo.GetID(), 10))

	if err = d.Set("repository_id", repo.GetID()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCodeQualityRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Reading repository code quality setup", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	setup, _, err := client.CodeQuality.GetSetup(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing repository code quality setup from state because it no longer exists", map[string]any{"owner": owner, "repo_name": repoName})
			d.SetId("")
			return nil
		}
		return diag.Errorf("error reading repository code quality setup: %s", err.Error())
	}

	enabled := setup.GetState() == codeQualityStateConfigured
	if err = d.Set("enabled", enabled); err != nil {
		return diag.FromErr(err)
	}

	// Languages and runner settings only apply while Code Quality is enabled, so keep the
	// configured values instead of reporting drift for a disabled setup.
	if !enabled {
		return nil
	}

	if err = d.Set("languages", flattenStringList(setup.Languages)); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("runner_type", setup.GetRunnerType()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("runner_label", setup.GetRunnerLabel()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCodeQualityUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Updating repository code quality setup", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	// A repository rename alone needs no API call; re-sending the setup would start a new analysis run.
	if !d.HasChanges("enabled", "languages", "runner_type", "runner_label") {
		return nil
	}

	if err := updateCodeQualitySetup(ctx, client, owner, repoName, expandCodeQualitySetup(d)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCodeQualityDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Deleting repository code quality setup", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	err := updateCodeQualitySetup(ctx, client, owner, repoName, github.CodeQualityUpdateSetupRequest{
		State: new(codeQualityStateNotConfigured),
	})
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Repository code quality setup no longer exists", map[string]any{"owner": owner, "repo_name": repoName})
			return nil
		}
		return diag.FromErr(handleArchivedRepoDelete(err, "repository code quality setup", d.Id(), owner, repoName))
	}

	return nil
}

func resourceGithubRepositoryCodeQualityImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "Importing repository code quality setup", map[string]any{"id": d.Id()})
	repoName := d.Id()
	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	meta, _ := m.(*Owner)
	owner := meta.name
	client := meta.v3client

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(repo.GetID(), 10))

	if err = d.Set("repository_id", repo.GetID()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func expandCodeQualitySetup(d *schema.ResourceData) github.CodeQualityUpdateSetupRequest {
	if enabled, _ := d.Get("enabled").(bool); !enabled {
		return github.CodeQualityUpdateSetupRequest{State: new(codeQualityStateNotConfigured)}
	}

	req := github.CodeQualityUpdateSetupRequest{State: new(codeQualityStateConfigured)}
	if languages, ok := d.Get("languages").(*schema.Set); ok && languages.Len() > 0 {
		req.Languages = expandStringList(languages.List())
	}
	if runnerType, _ := d.Get("runner_type").(string); runnerType != "" {
		req.RunnerType = new(runnerType)
	}
	if runnerLabel, _ := d.Get("runner_label").(string); runnerLabel != "" {
		req.RunnerLabel = new(runnerLabel)
	}
	return req
}

// updateCodeQualitySetup applies the setup, treating GitHub's 202 (update scheduled in the
// background) as success.
func updateCodeQualitySetup(ctx context.Context, client *github.Client, owner, repoName string, req github.CodeQualityUpdateSetupRequest) error {
	_, _, err := client.CodeQuality.UpdateSetup(ctx, owner, repoName, req)
	if _, ok := errors.AsType[*github.AcceptedError](err); ok {
		return nil
	}
	return err
}
