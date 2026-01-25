package github

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryWebhookResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the webhook.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository.",
			},
			"events": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "A list of events which should trigger the webhook.",
			},
			"configuration": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Webhook configuration options.",
			},
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
		},
	}
}

func resourceGithubRepositoryWebhookInstanceStateUpgradeV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Repository Webhook State before migration: %#v", rawState)

	prefix := "configuration."
	delete(rawState, prefix+"%")

	// Read & delete old keys
	oldKeys := make(map[string]any)
	for k, v := range rawState {
		if strings.HasPrefix(k, prefix) {
			oldKeys[k] = v

			// Delete old keys
			delete(rawState, k)
		}
	}

	// Write new keys
	for k, v := range oldKeys {
		newKey := "configuration.0." + strings.TrimPrefix(k, prefix)
		rawState[newKey] = v
	}

	rawState[prefix+"#"] = "1"

	log.Printf("[DEBUG] GitHub Repository Webhook State after migration: %#v", rawState)

	return rawState, nil
}
