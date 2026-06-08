package github

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationPrivateRegistry() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationPrivateRegistryCreate,
		ReadContext:   resourceGithubOrganizationPrivateRegistryRead,
		UpdateContext: resourceGithubOrganizationPrivateRegistryUpdate,
		DeleteContext: resourceGithubOrganizationPrivateRegistryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The auto-generated name of the private registry (computed by GitHub).",
			},
			"registry_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"maven_repository", "nuget_feed", "goproxy_server", "npm_registry", "rubygems_server", "cargo_registry", "composer_repository", "docker_registry", "git_source", "helm_registry", "pub_repository", "python_index", "terraform_registry"}, false)),
				Description:      "The registry type. Can be `maven_repository`, `nuget_feed`, `goproxy_server`, `npm_registry`, `rubygems_server`, `cargo_registry`, `composer_repository`, `docker_registry`, `git_source`, `helm_registry`, `pub_repository`, `python_index`, or `terraform_registry`.",
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the private registry.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username to use when authenticating with the private registry.",
			},
			"replaces_base": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicates whether this private registry should replace the base registry.",
			},
			"secret": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"secret", "encrypted_value"},
				Description:  "The plaintext secret to be encrypted and sent to GitHub. This is used for a token when auth_type is token, and for a password when auth_type is username_password. Required when auth_type is token or username_password.",
			},
			"encrypted_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"secret", "encrypted_value"},
				Description:  "The encrypted value of the secret using the GitHub public key in Base64 format.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ID of the public key used to encrypt the secret. Required if encrypted_value is set.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "private", "selected"}, false)),
				Description:      "Configures the access that repositories have to the organization private registry. Must be one of `all`, `private`, or `selected`.",
			},
			"selected_repository_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "An array of repository IDs that can access the organization private registry.",
			},
			"auth_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "token",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"token", "username_password", "oidc_azure", "oidc_aws", "oidc_jfrog"}, false)),
				Description:      "The authentication type for the private registry. Can be `token`, `username_password`, `oidc_azure`, `oidc_aws`, or `oidc_jfrog`. Defaults to `token`.",
			},
			"oidc_azure_tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The tenant ID of the Azure AD application. Required when auth_type is oidc_azure.",
			},
			"oidc_azure_client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The client ID of the Azure AD application. Required when auth_type is oidc_azure.",
			},
			"oidc_aws_region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The AWS region. Required when auth_type is oidc_aws.",
			},
			"oidc_aws_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The AWS account ID. Required when auth_type is oidc_aws.",
			},
			"oidc_aws_role_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The AWS IAM role name. Required when auth_type is oidc_aws.",
			},
			"oidc_aws_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CodeArtifact domain. Required when auth_type is oidc_aws.",
			},
			"oidc_aws_domain_owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CodeArtifact domain owner. Required when auth_type is oidc_aws.",
			},
			"oidc_jfrog_provider_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The JFrog OIDC provider name. Required when auth_type is oidc_jfrog.",
			},
			"oidc_audience": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The OIDC audience.",
			},
			"oidc_jfrog_identity_mapping_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The JFrog identity mapping name.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the private registry was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the private registry was last updated.",
			},
		},
	}
}

func resourceGithubOrganizationPrivateRegistryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	org := meta.(*Owner).name

	encryptedValue := d.Get("encrypted_value").(string)
	keyID := d.Get("key_id").(string)

	authType := d.Get("auth_type").(string)
	if authType == "token" || authType == "username_password" {
		if keyID == "" || len(encryptedValue) == 0 {
			ki, pk, err := getOrganizationRegistryPublicKeyDetails(ctx, client, org)
			if err != nil {
				return diag.FromErr(err)
			}
			keyID = ki

			if len(encryptedValue) == 0 {
				plaintextValue := d.Get("secret").(string)
				encryptedBytes, err := encryptPlaintext(plaintextValue, pk)
				if err != nil {
					return diag.FromErr(err)
				}
				encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
			}
		}
	}

	payload := github.CreateOrganizationPrivateRegistry{
		RegistryType: github.PrivateRegistryType(d.Get("registry_type").(string)),
		URL:          d.Get("url").(string),
		Visibility:   github.PrivateRegistryVisibility(d.Get("visibility").(string)),
	}

	if v, ok := d.GetOk("username"); ok {
		payload.Username = new(v.(string))
	}
	if v, ok := d.GetOk("replaces_base"); ok {
		payload.ReplacesBase = new(v.(bool))
	}
	if v, ok := d.GetOk("auth_type"); ok {
		payload.AuthType = new(v.(string))
	}

	if encryptedValue != "" {
		payload.EncryptedValue = new(encryptedValue)
		payload.KeyID = new(keyID)
	}

	if v, ok := d.GetOk("oidc_azure_tenant_id"); ok {
		payload.TenantID = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_azure_client_id"); ok {
		payload.ClientID = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_aws_region"); ok {
		payload.AWSRegion = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_aws_account_id"); ok {
		payload.AccountID = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_aws_role_name"); ok {
		payload.RoleName = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_aws_domain"); ok {
		payload.Domain = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_aws_domain_owner"); ok {
		payload.DomainOwner = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_jfrog_provider_name"); ok {
		payload.JFrogOIDCProviderName = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_audience"); ok {
		payload.Audience = new(v.(string))
	}
	if v, ok := d.GetOk("oidc_jfrog_identity_mapping_name"); ok {
		payload.IdentityMappingName = new(v.(string))
	}

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()
		for _, id := range ids {
			payload.SelectedRepositoryIDs = append(payload.SelectedRepositoryIDs, int64(id.(int)))
		}
	}

	registry, _, err := client.PrivateRegistries.CreateOrganizationPrivateRegistry(ctx, org, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(registry.GetName())

	return resourceGithubOrganizationPrivateRegistryRead(ctx, d, meta)
}

func resourceGithubOrganizationPrivateRegistryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	org := meta.(*Owner).name

	registry, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistry(ctx, org, d.Id())
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	if err := d.Set("name", registry.GetName()); err != nil {
		return diag.FromErr(err)
	}
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

func resourceGithubOrganizationPrivateRegistryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	org := meta.(*Owner).name

	encryptedValue := d.Get("encrypted_value").(string)
	keyID := d.Get("key_id").(string)

	authType := d.Get("auth_type").(string)
	if (d.HasChange("secret") || d.HasChange("encrypted_value")) && (authType == "token" || authType == "username_password") {
		ki, pk, err := getOrganizationRegistryPublicKeyDetails(ctx, client, org)
		if err != nil {
			return diag.FromErr(err)
		}
		keyID = ki

		if len(encryptedValue) == 0 {
			plaintextValue := d.Get("secret").(string)
			encryptedBytes, err := encryptPlaintext(plaintextValue, pk)
			if err != nil {
				return diag.FromErr(err)
			}
			encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
		}
	}

	payload := github.UpdateOrganizationPrivateRegistry{
		RegistryType: (*github.PrivateRegistryType)(new(d.Get("registry_type").(string))),
	}

	if d.HasChange("url") {
		payload.URL = new(d.Get("url").(string))
	}
	if d.HasChange("username") {
		payload.Username = new(d.Get("username").(string))
	}
	if d.HasChange("replaces_base") {
		payload.ReplacesBase = new(d.Get("replaces_base").(bool))
	}
	if d.HasChange("visibility") {
		payload.Visibility = (*github.PrivateRegistryVisibility)(new(d.Get("visibility").(string)))
	}
	if d.HasChange("auth_type") {
		payload.AuthType = new(d.Get("auth_type").(string))
	}

	if encryptedValue != "" {
		payload.EncryptedValue = new(encryptedValue)
		payload.KeyID = new(keyID)
	}

	if d.HasChange("oidc_azure_tenant_id") {
		payload.TenantID = new(d.Get("oidc_azure_tenant_id").(string))
	}
	if d.HasChange("oidc_azure_client_id") {
		payload.ClientID = new(d.Get("oidc_azure_client_id").(string))
	}
	if d.HasChange("oidc_aws_region") {
		payload.AWSRegion = new(d.Get("oidc_aws_region").(string))
	}
	if d.HasChange("oidc_aws_account_id") {
		payload.AccountID = new(d.Get("oidc_aws_account_id").(string))
	}
	if d.HasChange("oidc_aws_role_name") {
		payload.RoleName = new(d.Get("oidc_aws_role_name").(string))
	}
	if d.HasChange("oidc_aws_domain") {
		payload.Domain = new(d.Get("oidc_aws_domain").(string))
	}
	if d.HasChange("oidc_aws_domain_owner") {
		payload.DomainOwner = new(d.Get("oidc_aws_domain_owner").(string))
	}
	if d.HasChange("oidc_jfrog_provider_name") {
		payload.JFrogOIDCProviderName = new(d.Get("oidc_jfrog_provider_name").(string))
	}
	if d.HasChange("oidc_audience") {
		payload.Audience = new(d.Get("oidc_audience").(string))
	}
	if d.HasChange("oidc_jfrog_identity_mapping_name") {
		payload.IdentityMappingName = new(d.Get("oidc_jfrog_identity_mapping_name").(string))
	}
	if d.HasChange("selected_repository_ids") {
		v := d.Get("selected_repository_ids").(*schema.Set).List()
		var ids []int64
		for _, id := range v {
			ids = append(ids, int64(id.(int)))
		}
		payload.SelectedRepositoryIDs = ids
	}

	_, err := client.PrivateRegistries.UpdateOrganizationPrivateRegistry(ctx, org, d.Id(), payload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubOrganizationPrivateRegistryRead(ctx, d, meta)
}

func resourceGithubOrganizationPrivateRegistryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	org := meta.(*Owner).name

	_, err := client.PrivateRegistries.DeleteOrganizationPrivateRegistry(ctx, org, d.Id())
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == 404 {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}

func getOrganizationRegistryPublicKeyDetails(ctx context.Context, client *github.Client, org string) (string, string, error) {
	publicKey, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistriesPublicKey(ctx, org)
	if err != nil {
		return "", "", err
	}
	return publicKey.GetKeyID(), publicKey.GetKey(), nil
}
