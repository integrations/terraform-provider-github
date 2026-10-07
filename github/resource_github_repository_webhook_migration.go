package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryWebhookResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"events": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"configuration": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceGithubRepositoryWebhookInstanceStateUpgradeV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	tflog.Debug(ctx, "GitHub Repository Webhook State before migration", map[string]any{"field_count": len(rawState)})

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
	tflog.Debug(ctx, "GitHub Repository Webhook State after migration", map[string]any{"field_count": len(rawState)})

	return rawState, nil
}
