package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationWebhookCreate,
		ReadContext:   resourceGithubOrganizationWebhookRead,
		UpdateContext: resourceGithubOrganizationWebhookUpdate,
		DeleteContext: resourceGithubOrganizationWebhookDelete,
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
	if eventSet, ok := d.Get("events").(*schema.Set); ok {
		for _, v := range eventSet.List() {
			vStr, _ := v.(string)
			events = append(events, vStr)
		}
	}

	url, _ := d.Get("url").(string)
	active, _ := d.Get("active").(bool)
	hook := &github.Hook{
		URL:    new(url),
		Events: events,
		Active: new(active),
	}

	config, _ := d.Get("configuration").([]any)
	if len(config) > 0 {
		cfgMap, _ := config[0].(map[string]any)
		hook.Config = webhookConfigFromInterface(cfgMap)
	}

	return hook
}

func resourceGithubOrganizationWebhookCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client

	orgName := meta.name
	webhookObj := resourceGithubOrganizationWebhookObject(d)

	hook, _, err := client.Organizations.CreateHook(ctx, orgName, webhookObj)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(hook.GetID(), 10))

	// GitHub returns the secret as a string of 8 astrisks "********"
	// We would prefer to store the real secret in state, so we'll
	// write the configuration secret in state from our request to GitHub
	if hook.Config.Secret != nil {
		hook.Config.Secret = webhookObj.Config.Secret
	}

	if err = d.Set("configuration", interfaceFromWebhookConfig(hook.Config)); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubOrganizationWebhookRead(ctx, d, meta)
}

func resourceGithubOrganizationWebhookRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client

	orgName := meta.name
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		etag, _ := d.Get("etag").(string)
		ctx = context.WithValue(ctx, ctxEtag, etag)
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

	if meta.storeEtag {
		if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
			return diag.FromErr(err)
		}
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
	if cfg, ok := d.Get("configuration").([]any); ok && len(cfg) > 0 {
		cfgMap, _ := cfg[0].(map[string]any)
		currentSecret := cfgMap["secret"]

		if hook.Config.Secret != nil {
			secretStr, _ := currentSecret.(string)
			hook.Config.Secret = new(secretStr)
		}
	}

	if err = d.Set("configuration", interfaceFromWebhookConfig(hook.Config)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationWebhookUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client
	orgName := meta.name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	webhookObj := resourceGithubOrganizationWebhookObject(d)
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

func resourceGithubOrganizationWebhookDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client

	orgName := meta.name
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
		urlStr, _ := config["url"].(string)
		hookConfig.URL = new(urlStr)
	}
	if config["content_type"] != nil {
		contentTypeStr, _ := config["content_type"].(string)
		hookConfig.ContentType = new(contentTypeStr)
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
	if config["secret"] != nil {
		secretStr, _ := config["secret"].(string)
		hookConfig.Secret = new(secretStr)
	}
	return hookConfig
}

func interfaceFromWebhookConfig(config *github.HookConfig) []any {
	cfg := map[string]any{}
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
