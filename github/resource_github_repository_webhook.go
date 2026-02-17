package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryWebhookCreate,
		ReadContext:   resourceGithubRepositoryWebhookRead,
		UpdateContext: resourceGithubRepositoryWebhookUpdate,
		DeleteContext: resourceGithubRepositoryWebhookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 {
					return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <repository>/<webhook_id>")
				}
				if err := d.Set("repository", parts[0]); err != nil {
					return nil, err
				}
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryWebhookResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryWebhookInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository name of the webhook, not including the organization, which will be inferred.",
			},
			"events": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "A list of events which should trigger the webhook",
			},
			"configuration": webhookConfigurationSchema(),
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configuration block for the webhook",
			},
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicate if the webhook should receive events. Defaults to 'true'.",
			},
			"etag": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "An etag representing the webhook.",
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					return true
				},
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

func resourceGithubRepositoryWebhookObject(d *schema.ResourceData) *github.Hook {
	url := d.Get("url").(string)
	active := d.Get("active").(bool)
	events := []string{}
	eventSet := d.Get("events").(*schema.Set)
	for _, v := range eventSet.List() {
		events = append(events, v.(string))
	}

	hook := &github.Hook{
		URL:    &url,
		Events: events,
		Active: &active,
	}

	config := d.Get("configuration").([]any)[0].(map[string]any)
	if len(config) > 0 {
		hook.Config = webhookConfigFromInterface(config)
	}

	return hook
}

func resourceGithubRepositoryWebhookCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	hk := resourceGithubRepositoryWebhookObject(d)

	hook, _, err := client.Repositories.CreateHook(ctx, owner, repoName, hk)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(hook.GetID(), 10))

	// GitHub returns the secret as a string of 8 astrisks "********"
	// We would prefer to store the real secret in state, so we'll
	// write the configuration secret in state from our request to GitHub
	if hook.Config.Secret != nil {
		hook.Config.Secret = hk.Config.Secret
	}

	if err = d.Set("configuration", interfaceFromWebhookConfig(hook.Config)); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryWebhookRead(ctx, d, meta)
}

func resourceGithubRepositoryWebhookRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	hook, _, err := client.Repositories.GetHook(ctx, owner, repoName, hookID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing repository webhook %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	if err = d.Set("url", hook.GetURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("active", hook.GetActive()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("events", hook.Events); err != nil {
		return diag.FromErr(err)
	}

	// GitHub returns the secret as a string of 8 astrisks "********"
	// We would prefer to store the real secret in state, so we'll
	// write the configuration secret in state from what we get from
	// ResourceData
	if len(d.Get("configuration").([]any)) > 0 {
		currentSecret := d.Get("configuration").([]any)[0].(map[string]any)["secret"]

		if hook.Config.Secret != nil {
			hook.Config.Secret = github.Ptr(currentSecret.(string))
		}
	}

	if err = d.Set("configuration", interfaceFromWebhookConfig(hook.Config)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	hk := resourceGithubRepositoryWebhookObject(d)
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, _, err = client.Repositories.EditHook(ctx, owner, repoName, hookID, hk)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryWebhookRead(ctx, d, meta)
}

func resourceGithubRepositoryWebhookDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err = client.Repositories.DeleteHook(ctx, owner, repoName, hookID)
	return diag.FromErr(handleArchivedRepoDelete(err, "repository webhook", d.Id(), owner, repoName))
}
