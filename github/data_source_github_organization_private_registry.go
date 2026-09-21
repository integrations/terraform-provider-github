package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/integrations/terraform-provider-github/v6/internal/tfpluginv2util"
)

func dataSourceGithubOrganizationPrivateRegistry() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationPrivateRegistryRead,
		Description: "Use this data source to retrieve information about a specific organization private registry.",
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The auto-generated name of the private registry (computed by GitHub).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"registry_type": {
				Description: "The registry type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Description: "The registry URL.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username": {
				Description: "The registry username.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"replaces_base": {
				Description: "Whether the private registry should replace the public base registry.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"visibility": {
				Description: "Configures the access that repositories have to the organization private registry.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"auth_type": {
				Description: "The authentication type for the private registry.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_azure_tenant_id": {
				Description: "The Azure tenant ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_azure_client_id": {
				Description: "The Azure client ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_aws_region": {
				Description: "The AWS region.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_aws_account_id": {
				Description: "The AWS account ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_aws_role_name": {
				Description: "The AWS role name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_aws_domain": {
				Description: "The AWS domain.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_aws_domain_owner": {
				Description: "The AWS domain owner.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_jfrog_provider_name": {
				Description: "The JFrog provider name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_audience": {
				Description: "The JWT audience.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_jfrog_identity_mapping_name": {
				Description: "The JFrog identity mapping name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"selected_repository_ids": {
				Description: "An array of repository IDs that can access the organization private registry.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"created_at": {
				Description: "The time the registry was created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "The time the registry was updated.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceGithubOrganizationPrivateRegistryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	registryName := tfpluginv2util.Get[string](d, "name")

	registry, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistry(ctx, owner, registryName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(registry.GetName())

	if err := setPrivateRegistryAttributes(d, registry); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
