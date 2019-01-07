package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubOrganizationWebhook() *schema.Resource {

	return &schema.Resource{
		Create: resourceGithubOrganizationWebhookCreate,
		Read:   resourceGithubOrganizationWebhookRead,
		Update: resourceGithubOrganizationWebhookUpdate,
		Delete: resourceGithubOrganizationWebhookDelete,

		SchemaVersion: 1,
		MigrateState:  resourceGithubWebhookMigrateState,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGithubOrganizationWebhookName,
			},
			"events": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"configuration": webhookConfigurationSchema(),
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func validateGithubOrganizationWebhookName(v interface{}, k string) (ws []string, errors []error) {
	if v.(string) != "web" {
		errors = append(errors, fmt.Errorf("Github: name can only be web"))
	}
	return
}

func resourceGithubOrganizationWebhookObject(d *schema.ResourceData) *github.Hook {
	events := []string{}
	eventSet := d.Get("events").(*schema.Set)
	for _, v := range eventSet.List() {
		events = append(events, v.(string))
	}

	hook := &github.Hook{
		Name:   github.String(d.Get("name").(string)),
		URL:    github.String(d.Get("url").(string)),
		Events: events,
		Active: github.Bool(d.Get("active").(bool)),
	}

	config := d.Get("configuration").([]interface{})
	if len(config) > 0 {
		hook.Config = config[0].(map[string]interface{})
	}

	return hook
}

func resourceGithubOrganizationWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	webhookObj := resourceGithubOrganizationWebhookObject(d)
	ctx := context.Background()

	log.Printf("[DEBUG] Creating organization webhook: %s (%s)",
		webhookObj.GetName(), orgName)
	hook, _, err := client.Organizations.CreateHook(ctx, orgName, webhookObj)

	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*hook.ID, 10))

	return resourceGithubOrganizationWebhookRead(d, meta)
}

func resourceGithubOrganizationWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading organization webhook: %s (%s)", d.Id(), orgName)
	hook, resp, err := client.Organizations.GetHook(ctx, orgName, hookID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing organization webhook %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("name", hook.Name)
	d.Set("url", hook.URL)
	d.Set("active", hook.Active)
	d.Set("events", hook.Events)
	d.Set("configuration", []interface{}{hook.Config})

	return nil
}

func resourceGithubOrganizationWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	webhookObj := resourceGithubOrganizationWebhookObject(d)
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Updating organization webhook: %s (%s)", d.Id(), orgName)

	_, _, err = client.Organizations.EditHook(ctx,
		orgName, hookID, webhookObj)
	if err != nil {
		return err
	}

	return resourceGithubOrganizationWebhookRead(d, meta)
}

func resourceGithubOrganizationWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting organization webhook: %s (%s)", d.Id(), orgName)
	_, err = client.Organizations.DeleteHook(ctx, orgName, hookID)
	return err
}
