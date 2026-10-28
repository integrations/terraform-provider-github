package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/integrations/terraform-provider-github/v6/internal/tfpluginv2util"
)

func resourceGithubOrganizationPrivateRegistry() *schema.Resource {
	return &schema.Resource{
		Description:   "This resource allows you to create and manage an organization private registry for Dependabot.",
		CreateContext: resourceGithubOrganizationPrivateRegistryCreate,
		ReadContext:   resourceGithubOrganizationPrivateRegistryRead,
		UpdateContext: resourceGithubOrganizationPrivateRegistryUpdate,
		DeleteContext: resourceGithubOrganizationPrivateRegistryDelete,
		CustomizeDiff: customdiff.All(resourceGithubOrganizationPrivateRegistryDiff, diffPrivateRegistrySecret, diffSecretVariableVisibility),
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
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"maven_repository", "nuget_feed", "goproxy_server", "npm_registry", "rubygems_server", "cargo_registry", "composer_repository", "docker_registry", "git_source", "helm_registry", "hex_organization", "hex_repository", "pub_repository", "python_index", "terraform_registry"}, false)),
				Description:      "The registry type. Must be one of `maven_repository`, `nuget_feed`, `goproxy_server`, `npm_registry`, `rubygems_server`, `cargo_registry`, `composer_repository`, `docker_registry`, `git_source`, `helm_registry`, `hex_organization`, `hex_repository`, `pub_repository`, `python_index`, or `terraform_registry`.",
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
			"value": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"value_encrypted"},
				Description:   "The plaintext secret to be encrypted and sent to GitHub. This is used for a token when `auth_type` is `token`, and for a password when `auth_type` is `username_password`. One of `value` or `value_encrypted` is required for those auth types.",
			},
			"value_encrypted": {
				Type:             schema.TypeString,
				Optional:         true,
				Sensitive:        true,
				ConflictsWith:    []string{"value"},
				RequiredWith:     []string{"key_id"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
				Description:      "The secret encrypted with the organization private registries public key, in Base64 format. One of `value` or `value_encrypted` is required when `auth_type` is `token` or `username_password`.",
			},
			"key_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				RequiredWith:  []string{"value_encrypted"},
				ConflictsWith: []string{"value"},
				Description:   "ID of the public key used to encrypt the secret. Required if `value_encrypted` is set.",
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
				Set:      schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "An array of repository IDs that can access the organization private registry. Only valid when `visibility` is `selected`. GitHub does not return this list when reading a registry, so it is kept from the configuration and cannot be recovered by `terraform import`.",
			},
			"auth_type": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          "token",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"token", "username_password", "oidc_azure", "oidc_aws", "oidc_jfrog"}, false)),
				Description:      "The authentication type for the private registry. Can be `token`, `username_password`, `oidc_azure`, `oidc_aws`, or `oidc_jfrog`. Defaults to `token`. Cannot be changed after creation.",
			},
			"oidc_azure_tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The tenant ID of the Azure AD application. Required when `auth_type` is `oidc_azure`.",
			},
			"oidc_azure_client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The client ID of the Azure AD application. Required when `auth_type` is `oidc_azure`.",
			},
			"oidc_aws_region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The AWS region. Required when `auth_type` is `oidc_aws`.",
			},
			"oidc_aws_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The AWS account ID. Required when `auth_type` is `oidc_aws`.",
			},
			"oidc_aws_role_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The AWS IAM role name. Required when `auth_type` is `oidc_aws`.",
			},
			"oidc_aws_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CodeArtifact domain. Required when `auth_type` is `oidc_aws`.",
			},
			"oidc_aws_domain_owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CodeArtifact domain owner. Required when `auth_type` is `oidc_aws`.",
			},
			"oidc_jfrog_provider_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The JFrog OIDC provider name. Required when `auth_type` is `oidc_jfrog`.",
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
				Description: "The timestamp when the private registry was last updated by this resource.",
			},
			"remote_updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the private registry was last updated in GitHub. A difference with `updated_at` means the secret was changed outside of Terraform.",
			},
		},
	}
}

func registryAuthTypeUsesSecret(authType string) bool {
	return authType == "token" || authType == "username_password"
}

// resourceGithubOrganizationPrivateRegistryDiff enforces that token and username_password registries carry a secret,
// and that OIDC registries do not, since the API silently ignores a secret sent for them.
func resourceGithubOrganizationPrivateRegistryDiff(_ context.Context, d *schema.ResourceDiff, _ any) error {
	if _, ok := d.GetOk("selected_repository_ids"); ok && tfpluginv2util.Get[string](d, "visibility") != "selected" {
		return fmt.Errorf("cannot use selected_repository_ids without visibility being set to selected")
	}

	authType := tfpluginv2util.Get[string](d, "auth_type")
	// A value that is unknown at plan time (for example from random_password) counts as present.
	_, hasValue := d.GetOk("value")
	hasValue = hasValue || !d.NewValueKnown("value")
	_, hasEncrypted := d.GetOk("value_encrypted")
	hasEncrypted = hasEncrypted || !d.NewValueKnown("value_encrypted")

	if registryAuthTypeUsesSecret(authType) {
		if !hasValue && !hasEncrypted {
			return fmt.Errorf("one of `value,value_encrypted` must be specified when auth_type is %q", authType)
		}
		return nil
	}

	// key_id is Optional+Computed, so it carries a stale value across an auth_type replacement and is not checked here.
	if hasValue || hasEncrypted {
		return fmt.Errorf("`value` and `value_encrypted` cannot be set when auth_type is %q", authType)
	}
	return nil
}

// diffPrivateRegistrySecret runs the shared secret drift detection only for auth types that store a secret.
func diffPrivateRegistrySecret(ctx context.Context, d *schema.ResourceDiff, m any) error {
	if !registryAuthTypeUsesSecret(tfpluginv2util.Get[string](d, "auth_type")) {
		return nil
	}
	return diffSecret(ctx, d, m)
}

// registryEncryptedSecret returns the encrypted secret and key ID to send to GitHub,
// encrypting the plaintext value with the organization public key when no encrypted value is configured.
func registryEncryptedSecret(ctx context.Context, client *github.Client, owner string, d *schema.ResourceData) (string, string, error) {
	encryptedValue := tfpluginv2util.Get[string](d, "value_encrypted")
	keyID := tfpluginv2util.Get[string](d, "key_id")

	if keyID != "" && encryptedValue != "" {
		return encryptedValue, keyID, nil
	}

	publicKey, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistriesPublicKey(ctx, owner)
	if err != nil {
		return "", "", err
	}
	keyID = publicKey.GetKeyID()

	if encryptedValue == "" {
		encryptedBytes, err := encryptPlaintext(tfpluginv2util.Get[string](d, "value"), publicKey.GetKey())
		if err != nil {
			return "", "", err
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	return encryptedValue, keyID, nil
}

func registrySelectedRepositoryIDs(d *schema.ResourceData) []int64 {
	var ids []int64
	for _, id := range tfpluginv2util.GetSet[int](d, "selected_repository_ids", true) {
		ids = append(ids, int64(id))
	}
	return ids
}

// optString returns nil for an unset attribute so omitempty drops it from the request body.
func optString(d *schema.ResourceData, key string) *string {
	if v, ok := tfpluginv2util.GetOk[string](d, key); ok {
		return new(v)
	}
	return nil
}

// privateRegistryWriteRequest holds the desired state shared by the create and update request bodies.
type privateRegistryWriteRequest struct {
	url                   string
	username              *string
	replacesBase          bool
	visibility            github.PrivateRegistryVisibility
	selectedRepositoryIDs []int64
	encryptedValue        *string
	keyID                 *string
	tenantID              *string
	clientID              *string
	awsRegion             *string
	accountID             *string
	roleName              *string
	domain                *string
	domainOwner           *string
	jfrogOIDCProviderName *string
	audience              *string
	identityMappingName   *string
}

// privateRegistryDesiredState maps the full configuration to the request body so Create and Update send the same shape.
// The secret is included only when sendSecret is set; the returned key ID is written to state once the API call succeeds.
func privateRegistryDesiredState(ctx context.Context, client *github.Client, owner string, d *schema.ResourceData, sendSecret bool) (*privateRegistryWriteRequest, string, error) {
	req := &privateRegistryWriteRequest{
		url:                   tfpluginv2util.Get[string](d, "url"),
		username:              optString(d, "username"),
		replacesBase:          tfpluginv2util.Get[bool](d, "replaces_base"),
		visibility:            github.PrivateRegistryVisibility(tfpluginv2util.Get[string](d, "visibility")),
		selectedRepositoryIDs: registrySelectedRepositoryIDs(d),
		tenantID:              optString(d, "oidc_azure_tenant_id"),
		clientID:              optString(d, "oidc_azure_client_id"),
		awsRegion:             optString(d, "oidc_aws_region"),
		accountID:             optString(d, "oidc_aws_account_id"),
		roleName:              optString(d, "oidc_aws_role_name"),
		domain:                optString(d, "oidc_aws_domain"),
		domainOwner:           optString(d, "oidc_aws_domain_owner"),
		jfrogOIDCProviderName: optString(d, "oidc_jfrog_provider_name"),
		audience:              optString(d, "oidc_audience"),
		identityMappingName:   optString(d, "oidc_jfrog_identity_mapping_name"),
	}

	if !sendSecret || !registryAuthTypeUsesSecret(tfpluginv2util.Get[string](d, "auth_type")) {
		return req, "", nil
	}

	encryptedValue, keyID, err := registryEncryptedSecret(ctx, client, owner, d)
	if err != nil {
		return nil, "", err
	}
	req.encryptedValue = new(encryptedValue)
	req.keyID = new(keyID)

	return req, keyID, nil
}

func setPrivateRegistryAttributes(d *schema.ResourceData, registry *github.PrivateRegistry) error {
	fields := map[string]any{
		"name":       registry.GetName(),
		"url":        registry.GetURL(),
		"created_at": registry.GetCreatedAt().String(),
	}
	// Optional fields are only written when the API returned them, so a fresh import matches a fresh create.
	optional := map[string]*string{
		"username":                         registry.Username,
		"oidc_azure_tenant_id":             registry.TenantID,
		"oidc_azure_client_id":             registry.ClientID,
		"oidc_aws_region":                  registry.AWSRegion,
		"oidc_aws_account_id":              registry.AccountID,
		"oidc_aws_role_name":               registry.RoleName,
		"oidc_aws_domain":                  registry.Domain,
		"oidc_aws_domain_owner":            registry.DomainOwner,
		"oidc_jfrog_provider_name":         registry.JFrogOIDCProviderName,
		"oidc_audience":                    registry.Audience,
		"oidc_jfrog_identity_mapping_name": registry.IdentityMappingName,
	}
	for key, value := range optional {
		if value != nil {
			fields[key] = *value
		}
	}
	if registry.ReplacesBase != nil {
		fields["replaces_base"] = registry.GetReplacesBase()
	}
	// The GET endpoint omits selected_repository_ids, so the configured list stays in state unless the API returned one.
	if registry.SelectedRepositoryIDs != nil {
		fields["selected_repository_ids"] = registry.SelectedRepositoryIDs
	}
	if registry.RegistryType != nil {
		fields["registry_type"] = string(*registry.RegistryType)
	}
	if registry.Visibility != nil {
		fields["visibility"] = string(*registry.Visibility)
	}
	if registry.AuthType != nil {
		fields["auth_type"] = string(*registry.AuthType)
	}

	for key, value := range fields {
		if err := d.Set(key, value); err != nil {
			return err
		}
	}
	return nil
}

func resourceGithubOrganizationPrivateRegistryCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	desired, keyID, err := privateRegistryDesiredState(ctx, client, owner, d, true)
	if err != nil {
		return diag.FromErr(err)
	}

	authType := tfpluginv2util.Get[string](d, "auth_type")
	payload := github.CreateOrganizationPrivateRegistry{
		RegistryType:          github.PrivateRegistryType(tfpluginv2util.Get[string](d, "registry_type")),
		AuthType:              new(authType),
		URL:                   desired.url,
		Username:              desired.username,
		ReplacesBase:          new(desired.replacesBase),
		Visibility:            desired.visibility,
		SelectedRepositoryIDs: desired.selectedRepositoryIDs,
		EncryptedValue:        desired.encryptedValue,
		KeyID:                 desired.keyID,
		TenantID:              desired.tenantID,
		ClientID:              desired.clientID,
		AWSRegion:             desired.awsRegion,
		AccountID:             desired.accountID,
		RoleName:              desired.roleName,
		Domain:                desired.domain,
		DomainOwner:           desired.domainOwner,
		JFrogOIDCProviderName: desired.jfrogOIDCProviderName,
		Audience:              desired.audience,
		IdentityMappingName:   desired.identityMappingName,
	}

	registry, _, err := client.PrivateRegistries.CreateOrganizationPrivateRegistry(ctx, owner, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(registry.GetName())

	if err := d.Set("name", registry.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if keyID != "" {
		if err := d.Set("key_id", keyID); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("created_at", registry.GetCreatedAt().String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("remote_updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationPrivateRegistryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	registry, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistry(ctx, owner, d.Id())
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing organization private registry from state because it no longer exists in GitHub", map[string]any{"name": d.Id()})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := setPrivateRegistryAttributes(d, registry); err != nil {
		return diag.FromErr(err)
	}

	// updated_at tracks the last write made by this resource and remote_updated_at tracks GitHub,
	// so diffSecret can detect a registry changed outside of Terraform.
	if tfpluginv2util.Get[string](d, "updated_at") == "" {
		if err := d.Set("updated_at", registry.GetUpdatedAt().String()); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("remote_updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationPrivateRegistryUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	// updated_at changes when diffSecret detected an out-of-band change, so the secret is re-sent to restore it.
	sendSecret := d.HasChanges("value", "value_encrypted", "key_id", "updated_at")
	desired, keyID, err := privateRegistryDesiredState(ctx, client, owner, d, sendSecret)
	if err != nil {
		return diag.FromErr(err)
	}

	payload := github.UpdateOrganizationPrivateRegistry{
		URL:                   new(desired.url),
		Username:              desired.username,
		ReplacesBase:          new(desired.replacesBase),
		Visibility:            new(desired.visibility),
		SelectedRepositoryIDs: desired.selectedRepositoryIDs,
		EncryptedValue:        desired.encryptedValue,
		KeyID:                 desired.keyID,
		TenantID:              desired.tenantID,
		ClientID:              desired.clientID,
		AWSRegion:             desired.awsRegion,
		AccountID:             desired.accountID,
		RoleName:              desired.roleName,
		Domain:                desired.domain,
		DomainOwner:           desired.domainOwner,
		JFrogOIDCProviderName: desired.jfrogOIDCProviderName,
		Audience:              desired.audience,
		IdentityMappingName:   desired.identityMappingName,
	}

	// omitempty drops a nil username, so a removed username is sent as an empty string to clear it.
	if d.HasChange("username") && payload.Username == nil {
		payload.Username = new("")
	}

	if _, err := client.PrivateRegistries.UpdateOrganizationPrivateRegistry(ctx, owner, d.Id(), payload); err != nil {
		return diag.FromErr(err)
	}

	if keyID != "" {
		if err := d.Set("key_id", keyID); err != nil {
			return diag.FromErr(err)
		}
	}

	// The update endpoint returns no body, so the registry is fetched once to record the new timestamp.
	registry, _, err := client.PrivateRegistries.GetOrganizationPrivateRegistry(ctx, owner, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("remote_updated_at", registry.GetUpdatedAt().String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationPrivateRegistryDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	_, err := client.PrivateRegistries.DeleteOrganizationPrivateRegistry(ctx, owner, d.Id())
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
