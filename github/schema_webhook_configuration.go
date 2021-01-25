package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func webhookConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:      schema.TypeString,
					Required:  true,
					Sensitive: true,
				},
				"content_type": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"secret": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
				},
				"insecure_ssl": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}
}
