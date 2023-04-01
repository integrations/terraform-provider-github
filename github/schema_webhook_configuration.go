package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func webhookConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Configuration for the webhook.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					Description: "The URL of the webhook.",
				},
				"content_type": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The content type for the payload. Valid values are either 'form' or 'json'.",
				},
				"secret": {
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					Description: "The shared secret for the webhook",
				},
				"insecure_ssl": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Insecure SSL boolean toggle. Defaults to 'false'.",
				},
			},
		},
	}
}
