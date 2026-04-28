package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryCodeScanningDefaultSetup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryCodeScanningDefaultSetupCreateOrUpdate,
		ReadContext:   resourceGithubRepositoryCodeScanningDefaultSetupRead,
		UpdateContext: resourceGithubRepositoryCodeScanningDefaultSetupCreateOrUpdate,
		DeleteContext: resourceGithubRepositoryCodeScanningDefaultSetupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryCodeScanningDefaultSetupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		CustomizeDiff: diffRepository,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub repository name.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub repository.",
			},
			"state": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired state of code scanning default setup. Must be `configured` or `not-configured`.",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"configured", "not-configured"}, false),
				),
			},
			"query_suite": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The query suite to use. Must be `default` or `extended`.",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"default", "extended"}, false),
				),
			},
			"languages": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "The languages to enable for code scanning. Supported values include `actions`, `c-cpp`, `csharp`, `go`, `java-kotlin`, `javascript-typescript`, `python`, `ruby`, `swift`.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceGithubRepositoryCodeScanningDefaultSetupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	state := d.Get("state").(string)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.Errorf("error reading repository %s/%s: %s", owner, repoName, err)
	}
	if repo.GetArchived() {
		return diag.Errorf("repository %s/%s is archived", owner, repoName)
	}

	options := &github.UpdateDefaultSetupConfigurationOptions{
		State: state,
	}

	if v, ok := d.GetOk("query_suite"); ok {
		qs := v.(string)
		options.QuerySuite = &qs
	}

	if v, ok := d.GetOk("languages"); ok {
		options.Languages = expandStringList(v.(*schema.Set).List())
	}

	_, _, err = client.CodeScanning.UpdateDefaultSetupConfiguration(ctx, owner, repoName, options)
	if err != nil {
		// 202 Accepted is expected — go-github surfaces it as AcceptedError
		var acceptedErr *github.AcceptedError
		if !errors.As(err, &acceptedErr) {
			return diag.Errorf("error updating code scanning default setup for %s/%s: %s", owner, repoName, err)
		}
	}

	d.SetId(repoName)

	var timeout time.Duration
	if d.IsNewResource() {
		timeout = d.Timeout(schema.TimeoutCreate)
	} else {
		timeout = d.Timeout(schema.TimeoutUpdate)
	}

	config, err := waitForCodeScanningState(ctx, client, owner, repoName, state, timeout)
	if err != nil {
		return diag.Errorf("error waiting for code scanning default setup state for %s/%s: %s", owner, repoName, err)
	}

	return setCodeScanningDefaultSetupState(d, config)
}

func resourceGithubRepositoryCodeScanningDefaultSetupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Repository not found, removing from state", map[string]any{
				"owner":      owner,
				"repository": repoName,
			})
			d.SetId("")
			return nil
		}
		return diag.Errorf("error reading repository %s/%s: %s", owner, repoName, err)
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.Errorf("error setting repository_id: %s", err)
	}

	config, _, err := client.CodeScanning.GetDefaultSetupConfiguration(ctx, owner, repoName)
	if err != nil {
		return diag.Errorf("error reading code scanning default setup for %s/%s: %s", owner, repoName, err)
	}

	return setCodeScanningDefaultSetupState(d, config)
}

func resourceGithubRepositoryCodeScanningDefaultSetupDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	options := &github.UpdateDefaultSetupConfigurationOptions{
		State: "not-configured",
	}

	_, _, err := client.CodeScanning.UpdateDefaultSetupConfiguration(ctx, owner, repoName, options)
	if err != nil {
		var acceptedErr *github.AcceptedError
		var ghErr *github.ErrorResponse
		switch {
		case errors.As(err, &acceptedErr):
			// 202 Accepted is expected
		case errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound:
			// repository already gone
			return nil
		default:
			return diag.Errorf("error disabling code scanning default setup for %s/%s: %s", owner, repoName, err)
		}
	}

	tflog.Info(ctx, "Code scanning default setup disabled", map[string]any{
		"owner":      owner,
		"repository": repoName,
	})
	return nil
}

func resourceGithubRepositoryCodeScanningDefaultSetupImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	repoName := d.Id()
	if repoName == "" {
		return nil, fmt.Errorf("repository name must not be empty")
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func setCodeScanningDefaultSetupState(d *schema.ResourceData, config *github.DefaultSetupConfiguration) diag.Diagnostics {
	if err := d.Set("state", config.GetState()); err != nil {
		return diag.Errorf("error setting state: %s", err)
	}
	if err := d.Set("query_suite", config.GetQuerySuite()); err != nil {
		return diag.Errorf("error setting query_suite: %s", err)
	}
	if err := d.Set("languages", config.Languages); err != nil {
		return diag.Errorf("error setting languages: %s", err)
	}
	return nil
}

func waitForCodeScanningState(ctx context.Context, client *github.Client, owner, repo, targetState string, timeout time.Duration) (*github.DefaultSetupConfiguration, error) {
	conf := &retry.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{targetState},
		Timeout:    timeout,
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Refresh: func() (any, string, error) {
			config, _, err := client.CodeScanning.GetDefaultSetupConfiguration(ctx, owner, repo)
			if err != nil {
				return nil, "", err
			}
			state := config.GetState()
			if state == targetState {
				return config, state, nil
			}
			return config, "pending", nil
		},
	}

	result, err := conf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("code scanning default setup returned nil result for %s/%s", owner, repo)
	}

	return result.(*github.DefaultSetupConfiguration), nil
}
