package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationWebhookCreate,
		ReadContext:   resourceGithubOrganizationWebhookRead,
		UpdateContext: resourceGithubOrganizationWebhookUpdate,
		DeleteContext: resourceGithubOrganizationWebhookDelete,
		ValidateRawResourceConfigFuncs: []schema.ValidateRawResourceConfigFunc{
			preferWriteOnlyWebhookSecretValidator(),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: diffETag,

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubOrganizationWebhookResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubOrganizationWebhookInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"events": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list of events which should trigger the webhook.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
			"configuration": webhookConfigurationSchema(),
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the webhook.",
			},
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicate if the webhook should receive events.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the organization webhook.",
			},
		},
	}
}

func resourceGithubOrganizationWebhookObject(d *schema.ResourceData) *github.Hook {
	events := []string{}
	eventSet := d.Get("events").(*schema.Set)
	for _, v := range eventSet.List() {
		events = append(events, v.(string))
	}

	hook := &github.Hook{
		URL:    new(d.Get("url").(string)),
		Events: events,
		Active: new(d.Get("active").(bool)),
	}

	config := d.Get("configuration").([]any)
	if len(config) > 0 {
		hook.Config = webhookConfigFromInterface(config[0].(map[string]any))
		if secret, configured := d.GetOk("configuration.0.secret"); configured {
			hook.Config.Secret = new(secret.(string))
		}
	}

	return hook
}

func resourceGithubOrganizationWebhookCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	webhookObj := resourceGithubOrganizationWebhookObject(d)
	if _, configured := d.GetOk("configuration.0.secret_wo_version"); configured {
		secret, diags := readRawWriteOnlyString(d, webhookSecretWriteOnlyPath)
		if diags.HasError() {
			return diags
		}
		if webhookObj.Config == nil {
			webhookObj.Config = &github.HookConfig{}
		}
		webhookObj.Config.Secret = new(secret)
	}

	hook, _, err := client.Organizations.CreateHook(ctx, orgName, webhookObj)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(hook.GetID(), 10))

	return resourceGithubOrganizationWebhookRead(ctx, d, meta)
}

func resourceGithubOrganizationWebhookRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	hook, resp, err := client.Organizations.GetHook(ctx, orgName, hookID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing organization webhook %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
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

	if err = d.Set("configuration", interfaceFromWebhookConfigPreservingState(hook.Config, d)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	webhookObj := resourceGithubOrganizationWebhookObject(d)
	if d.HasChange("configuration.0.secret") {
		if _, configured := d.GetOk("configuration.0.secret"); !configured && webhookObj.Config != nil {
			webhookObj.Config.Secret = new("")
		}
	}
	if _, configured := d.GetOk("configuration.0.secret_wo_version"); d.HasChange("configuration.0.secret_wo_version") && configured {
		secret, diags := readRawWriteOnlyString(d, webhookSecretWriteOnlyPath)
		if diags.HasError() {
			return diags
		}
		if webhookObj.Config == nil {
			webhookObj.Config = &github.HookConfig{}
		}
		webhookObj.Config.Secret = new(secret)
	}
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, _, err = client.Organizations.EditHook(ctx,
		orgName, hookID, webhookObj)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubOrganizationWebhookRead(ctx, d, meta)
}

func resourceGithubOrganizationWebhookDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err = client.Organizations.DeleteHook(ctx, orgName, hookID)
	return diag.FromErr(err)
}

func webhookConfigFromInterface(config map[string]any) *github.HookConfig {
	hookConfig := &github.HookConfig{}
	if config["url"] != nil {
		hookConfig.URL = new(config["url"].(string))
	}
	if config["content_type"] != nil {
		hookConfig.ContentType = new(config["content_type"].(string))
	}
	if config["insecure_ssl"] != nil {
		if insecureSsl, ok := config["insecure_ssl"].(bool); ok {
			if insecureSsl {
				hookConfig.InsecureSSL = new("1")
			} else {
				hookConfig.InsecureSSL = new("0")
			}
		} else {
			if config["insecure_ssl"] == "1" || config["insecure_ssl"] == "true" {
				hookConfig.InsecureSSL = new("1")
			} else {
				hookConfig.InsecureSSL = new("0")
			}
		}
	}
	return hookConfig
}

func interfaceFromWebhookConfig(config *github.HookConfig) []any {
	cfg := map[string]any{}
	if config == nil {
		return []any{cfg}
	}
	if config.URL != nil {
		cfg["url"] = *config.URL
	}
	if config.ContentType != nil {
		cfg["content_type"] = *config.ContentType
	}
	if config.InsecureSSL != nil {
		cfg["insecure_ssl"] = *config.InsecureSSL == "1"
	}
	if config.Secret != nil {
		cfg["secret"] = *config.Secret
	}
	return []any{cfg}
}

func interfaceFromWebhookConfigPreservingState(config *github.HookConfig, d *schema.ResourceData) []any {
	flattened := interfaceFromWebhookConfig(config)
	cfg := flattened[0].(map[string]any)

	if secret, configured := d.GetOk("configuration.0.secret"); configured {
		cfg["secret"] = secret.(string)
	} else {
		delete(cfg, "secret")
	}
	if version, configured := d.GetOk("configuration.0.secret_wo_version"); configured {
		cfg["secret_wo_version"] = version.(int)
	}

	return flattened
}
