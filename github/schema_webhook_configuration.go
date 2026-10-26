package github

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func preferWriteOnlyWebhookSecretValidator() schema.ValidateRawResourceConfigFunc {
	configurationItem := cty.UnknownVal(cty.Number)
	return validation.PreferWriteOnlyAttribute(
		cty.GetAttrPath("configuration").Index(configurationItem).GetAttr("secret"),
		cty.GetAttrPath("configuration").Index(configurationItem).GetAttr("secret_wo"),
	)
}

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
					Type:          schema.TypeString,
					Optional:      true,
					Sensitive:     true,
					ConflictsWith: []string{"configuration.0.secret_wo"},
					Description:   "The shared secret for the webhook",
				},
				"secret_wo": {
					Type:          schema.TypeString,
					Optional:      true,
					Sensitive:     true,
					WriteOnly:     true,
					ConflictsWith: []string{"configuration.0.secret"},
					RequiredWith:  []string{"configuration.0.secret_wo_version"},
					Description:   "Write-only shared secret for the webhook. Requires Terraform 1.11 or later.",
				},
				"secret_wo_version": {
					Type:         schema.TypeInt,
					Optional:     true,
					RequiredWith: []string{"configuration.0.secret_wo"},
					ValidateFunc: validation.IntAtLeast(1),
					Description:  "Version of the write-only webhook secret. Change this value to rotate the secret.",
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
