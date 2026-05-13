package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationPrivateRegistry() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationPrivateRegistryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the private registry.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"registry_type": {
				Description: "The registry type. Can be `maven_repository`, `nuget_feed`, `goproxy_server`, `npm_registry`, `rubygems_server`, `cargo_registry`, `composer_repository`, `docker_registry`, `git_source`, `helm_registry`, `pub_repository`, `python_index`, or `terraform_registry`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replaces_base": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"visibility": {
				Description: "Configures the access that repositories have to the organization private registry. Must be one of `all`, `private`, or `selected`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"auth_type": {
				Description: "The authentication type for the private registry. Can be `token`, `username_password`, `oidc_azure`, `oidc_aws`, or `oidc_jfrog`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"oidc_azure_tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_azure_client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_aws_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_aws_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_aws_role_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_aws_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_aws_domain_owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_jfrog_provider_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_audience": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oidc_jfrog_identity_mapping_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"selected_repository_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubOrganizationPrivateRegistryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	org := meta.(*Owner).name
	registryName := d.Get("name").(string)

	registry, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistry(ctx, org, registryName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(registry.GetName())

	if registry.RegistryType != nil {
		if err := d.Set("registry_type", string(*registry.RegistryType)); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("url", registry.GetURL()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", registry.GetUsername()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("replaces_base", registry.GetReplacesBase()); err != nil {
		return diag.FromErr(err)
	}
	if registry.Visibility != nil {
		if err := d.Set("visibility", string(*registry.Visibility)); err != nil {
			return diag.FromErr(err)
		}
	}

	if registry.AuthType != nil {
		if err := d.Set("auth_type", string(*registry.AuthType)); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("oidc_azure_tenant_id", registry.GetTenantID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_azure_client_id", registry.GetClientID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_aws_region", registry.GetAWSRegion()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_aws_account_id", registry.GetAccountID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_aws_role_name", registry.GetRoleName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_aws_domain", registry.GetDomain()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_aws_domain_owner", registry.GetDomainOwner()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_jfrog_provider_name", registry.GetJFrogOIDCProviderName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_audience", registry.GetAudience()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("oidc_jfrog_identity_mapping_name", registry.GetIdentityMappingName()); err != nil {
		return diag.FromErr(err)
	}

	var repoIDs []any
	for _, id := range registry.SelectedRepositoryIDs {
		repoIDs = append(repoIDs, int(id))
	}
	if err := d.Set("selected_repository_ids", schema.NewSet(schema.HashInt, repoIDs)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", registry.GetCreatedAt().String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
