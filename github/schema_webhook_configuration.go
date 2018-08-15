package github

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func webhookConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:     schema.TypeString,
					Required: true,
				},
				"content_type": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"secret": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
					DiffSuppressFunc: func(k, oldV, newV string, d *schema.ResourceData) bool {
						// Undocumented GitHub feature where API returns 8 asterisks in place of the secret
						maskedSecret := "********"
						if oldV == maskedSecret {
							return true
						}

						return oldV == newV
					},
				},
				"insecure_ssl": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}
