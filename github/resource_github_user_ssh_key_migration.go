package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubUserSshKeyV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A descriptive name for the new key.",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The public SSH key to add to your GitHub account.",
				DiffSuppressFunc: func(k, oldV, newV string, d *schema.ResourceData) bool {
					newTrimmed := strings.TrimSpace(newV)
					return oldV == newTrimmed
				},
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the SSH key.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubUserSshKeyStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	if rawState == nil {
		return nil, fmt.Errorf("resource state upgrade failed, state is nil")
	}

	// copy d.Id() into key_id
	if id, ok := rawState["id"].(string); ok {
		keyID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("resource state upgrade failed, invalid SSH key ID format: %w", err)
		}
		rawState["key_id"] = keyID
	}

	return rawState, nil
}
