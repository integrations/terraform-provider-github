package github

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationSettingsCreateOrUpdate,
		ReadContext:   resourceGithubOrganizationSettingsRead,
		UpdateContext: resourceGithubOrganizationSettingsCreateOrUpdate,
		DeleteContext: resourceGithubOrganizationSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"billing_email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The billing email address for the organization.",
			},
			"company": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The company name for the organization.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The email address for the organization.",
			},
			"twitter_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Twitter username for the organization.",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The location for the organization.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name for the organization.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the organization.",
			},
			"has_organization_projects": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization projects are enabled for the organization.",
			},
			"has_repository_projects": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not repository projects are enabled for the organization.",
			},
			"default_repository_permission": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "read",
				Description:      "The default permission for organization members to create new repositories. Can be one of 'read', 'write', 'admin' or 'none'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"read", "write", "admin", "none"}, false)),
			},
			"members_can_create_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization members can create new repositories.",
			},
			"members_can_create_internal_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not organization members can create new internal repositories. For Enterprise Organizations only.",
			},
			"members_can_create_private_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization members can create new private repositories.",
			},
			"members_can_create_public_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization members can create new public repositories.",
			},
			"members_can_create_pages": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization members can create new pages.",
			},
			"members_can_create_public_pages": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization members can create new public pages.",
			},
			"members_can_create_private_pages": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not organization members can create new private pages.",
			},
			"members_can_fork_private_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not organization members can fork private repositories.",
			},
			"web_commit_signoff_required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not commit signatures are required for commits to the organization.",
			},
			"blog": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The blog URL for the organization.",
			},
			"advanced_security_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: " Whether or not advanced security is enabled for new repositories.",
			},
			"dependabot_alerts_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not dependabot alerts are enabled for new repositories.",
			},
			"dependabot_security_updates_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: " Whether or not dependabot security updates are enabled for new repositories.",
			},
			"dependency_graph_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not dependency graph is enabled for new repositories.",
			},
			"secret_scanning_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not secret scanning is enabled for new repositories.",
			},
			"secret_scanning_push_protection_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not secret scanning push protection is enabled for new repositories.",
			},
		},
	}
}

// buildOrganizationSettings creates a github.Organization struct with only the fields that are explicitly configured.
// For updates, it only includes fields that have actually changed to avoid API validation errors.
func buildOrganizationSettings(d *schema.ResourceData, isEnterprise bool) *github.Organization {
	settings := &github.Organization{}

	// Check if this is an update (has ID) or create (no ID)
	isUpdate := d.Id() != ""

	// Helper function to check if field should be included
	shouldInclude := func(fieldName string) bool {
		if !isUpdate {
			// For creates, include if explicitly configured
			_, ok := d.GetOk(fieldName)
			return ok
		}
		// For updates, only include if the field has changed
		return d.HasChange(fieldName)
	}

	// Required field - always include if configured (API requires it even if unchanged)
	if billingEmail, ok := d.GetOk("billing_email"); ok {
		settings.BillingEmail = new(billingEmail.(string))
	}

	// Optional string fields - only set if should be included
	if shouldInclude("company") {
		if company, ok := d.GetOk("company"); ok {
			settings.Company = new(company.(string))
		}
	}
	if shouldInclude("email") {
		if email, ok := d.GetOk("email"); ok {
			settings.Email = new(email.(string))
		}
	}
	if shouldInclude("twitter_username") {
		if twitterUsername, ok := d.GetOk("twitter_username"); ok {
			settings.TwitterUsername = new(twitterUsername.(string))
		}
	}
	if shouldInclude("location") {
		if location, ok := d.GetOk("location"); ok {
			settings.Location = new(location.(string))
		}
	}
	if shouldInclude("name") {
		if name, ok := d.GetOk("name"); ok {
			settings.Name = new(name.(string))
		}
	}
	if shouldInclude("description") {
		if description, ok := d.GetOk("description"); ok {
			settings.Description = new(description.(string))
		}
	}
	if shouldInclude("blog") {
		if blog, ok := d.GetOk("blog"); ok {
			settings.Blog = new(blog.(string))
		}
	}

	// Boolean fields - only set if should be included
	// Use d.Get() instead of d.GetOk() when shouldInclude() returns true,
	// because we already know the field should be included, and d.Get() correctly handles false values
	if shouldInclude("has_organization_projects") {
		settings.HasOrganizationProjects = new(d.Get("has_organization_projects").(bool))
	}
	if shouldInclude("has_repository_projects") {
		settings.HasRepositoryProjects = new(d.Get("has_repository_projects").(bool))
	}
	if shouldInclude("default_repository_permission") {
		if defaultRepoPermission, ok := d.GetOk("default_repository_permission"); ok {
			settings.DefaultRepoPermission = new(defaultRepoPermission.(string))
		}
	}
	if shouldInclude("members_can_create_repositories") {
		settings.MembersCanCreateRepos = new(d.Get("members_can_create_repositories").(bool))
	}
	if shouldInclude("members_can_create_private_repositories") {
		settings.MembersCanCreatePrivateRepos = new(d.Get("members_can_create_private_repositories").(bool))
	}
	if shouldInclude("members_can_create_public_repositories") {
		settings.MembersCanCreatePublicRepos = new(d.Get("members_can_create_public_repositories").(bool))
	}
	if shouldInclude("members_can_create_pages") {
		settings.MembersCanCreatePages = new(d.Get("members_can_create_pages").(bool))
	}
	if shouldInclude("members_can_create_public_pages") {
		settings.MembersCanCreatePublicPages = new(d.Get("members_can_create_public_pages").(bool))
	}
	if shouldInclude("members_can_create_private_pages") {
		settings.MembersCanCreatePrivatePages = new(d.Get("members_can_create_private_pages").(bool))
	}
	if shouldInclude("members_can_fork_private_repositories") {
		settings.MembersCanForkPrivateRepos = new(d.Get("members_can_fork_private_repositories").(bool))
	}
	if shouldInclude("web_commit_signoff_required") {
		settings.WebCommitSignoffRequired = new(d.Get("web_commit_signoff_required").(bool))
	}
	if shouldInclude("advanced_security_enabled_for_new_repositories") {
		settings.AdvancedSecurityEnabledForNewRepos = new(d.Get("advanced_security_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("dependabot_alerts_enabled_for_new_repositories") {
		settings.DependabotAlertsEnabledForNewRepos = new(d.Get("dependabot_alerts_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("dependabot_security_updates_enabled_for_new_repositories") {
		settings.DependabotSecurityUpdatesEnabledForNewRepos = new(d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("dependency_graph_enabled_for_new_repositories") {
		settings.DependencyGraphEnabledForNewRepos = new(d.Get("dependency_graph_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("secret_scanning_enabled_for_new_repositories") {
		settings.SecretScanningEnabledForNewRepos = new(d.Get("secret_scanning_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("secret_scanning_push_protection_enabled_for_new_repositories") {
		settings.SecretScanningPushProtectionEnabledForNewRepos = new(d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool))
	}

	// Enterprise-specific field
	if isEnterprise {
		if shouldInclude("members_can_create_internal_repositories") {
			settings.MembersCanCreateInternalRepos = new(d.Get("members_can_create_internal_repositories").(bool))
		}
	}

	return settings
}

func resourceGithubOrganizationSettingsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	ctx = context.WithValue(ctx, ctxId, d.Id())
	org := meta.(*Owner).name

	orgInfo, _, err := client.Organizations.Get(ctx, org)
	if err != nil {
		return diag.FromErr(err)
	}

	// Build settings using helper function
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"
	settings := buildOrganizationSettings(d, isEnterprise)
	tflog.Debug(ctx, "Built organization settings", map[string]any{"owner": org, "is_enterprise": isEnterprise})
	if settings.HasOrganizationProjects != nil {
		tflog.Debug(ctx, "HasOrganizationProjects", map[string]any{"settings_has_organization_projects": *settings.HasOrganizationProjects})
	}
	if settings.HasRepositoryProjects != nil {
		tflog.Debug(ctx, "HasRepositoryProjects", map[string]any{"settings_has_repository_projects": *settings.HasRepositoryProjects})
	}
	if settings.DefaultRepoPermission != nil {
		tflog.Debug(ctx, "DefaultRepoPermission", map[string]any{"settings_default_repo_permission": *settings.DefaultRepoPermission})
	}
	if settings.MembersCanCreateRepos != nil {
		tflog.Debug(ctx, "MembersCanCreateRepos", map[string]any{"settings_members_can_create_repos": *settings.MembersCanCreateRepos})
	}
	if settings.MembersCanCreatePrivateRepos != nil {
		tflog.Debug(ctx, "MembersCanCreatePrivateRepos", map[string]any{"settings_members_can_create_private_repos": *settings.MembersCanCreatePrivateRepos})
	}
	if settings.MembersCanCreatePublicRepos != nil {
		tflog.Debug(ctx, "MembersCanCreatePublicRepos", map[string]any{"settings_members_can_create_public_repos": *settings.MembersCanCreatePublicRepos})
	}
	if settings.MembersCanCreateInternalRepos != nil {
		tflog.Debug(ctx, "MembersCanCreateInternalRepos", map[string]any{"settings_members_can_create_internal_repos": *settings.MembersCanCreateInternalRepos})
	}
	if settings.MembersCanCreatePages != nil {
		tflog.Debug(ctx, "MembersCanCreatePages", map[string]any{"settings_members_can_create_pages": *settings.MembersCanCreatePages})
	}
	if settings.MembersCanCreatePublicPages != nil {
		tflog.Debug(ctx, "MembersCanCreatePublicPages", map[string]any{"settings_members_can_create_public_pages": *settings.MembersCanCreatePublicPages})
	}
	if settings.MembersCanCreatePrivatePages != nil {
		tflog.Debug(ctx, "MembersCanCreatePrivatePages", map[string]any{"settings_members_can_create_private_pages": *settings.MembersCanCreatePrivatePages})
	}
	if settings.MembersCanForkPrivateRepos != nil {
		tflog.Debug(ctx, "MembersCanForkPrivateRepos", map[string]any{"settings_members_can_fork_private_repos": *settings.MembersCanForkPrivateRepos})
	}
	if settings.WebCommitSignoffRequired != nil {
		tflog.Debug(ctx, "WebCommitSignoffRequired", map[string]any{"settings_web_commit_signoff_required": *settings.WebCommitSignoffRequired})
	}
	if settings.AdvancedSecurityEnabledForNewRepos != nil {
		tflog.Debug(ctx, "AdvancedSecurityEnabledForNewRepos", map[string]any{"settings_advanced_security_enabled_for_new_repos": *settings.AdvancedSecurityEnabledForNewRepos})
	}
	if settings.DependabotAlertsEnabledForNewRepos != nil {
		tflog.Debug(ctx, "DependabotAlertsEnabledForNewRepos", map[string]any{"settings_dependabot_alerts_enabled_for_new_repos": *settings.DependabotAlertsEnabledForNewRepos})
	}
	if settings.DependabotSecurityUpdatesEnabledForNewRepos != nil {
		tflog.Debug(ctx, "DependabotSecurityUpdatesEnabledForNewRepos", map[string]any{"settings_dependabot_security_updates_enabled_for_new_repos": *settings.DependabotSecurityUpdatesEnabledForNewRepos})
	}
	if settings.DependencyGraphEnabledForNewRepos != nil {
		tflog.Debug(ctx, "DependencyGraphEnabledForNewRepos", map[string]any{"settings_dependency_graph_enabled_for_new_repos": *settings.DependencyGraphEnabledForNewRepos})
	}
	if settings.SecretScanningEnabledForNewRepos != nil {
		tflog.Debug(ctx, "SecretScanningEnabledForNewRepos", map[string]any{"settings_secret_scanning_enabled_for_new_repos": *settings.SecretScanningEnabledForNewRepos})
	}
	if settings.SecretScanningPushProtectionEnabledForNewRepos != nil {
		tflog.Debug(ctx, "SecretScanningPushProtectionEnabledForNewRepos", map[string]any{"settings_secret_scanning_push_protection_enabled_for_new_repos": *settings.SecretScanningPushProtectionEnabledForNewRepos})
	}

	orgSettings, _, err := client.Organizations.Edit(ctx, org, settings)
	if err != nil {
		// Log detailed error information for debugging
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			tflog.Debug(ctx, "GitHub API error", map[string]any{"status_code": ghErr.Response.StatusCode, "message": ghErr.Message})
			if len(ghErr.Errors) > 0 {
				for i, apiErr := range ghErr.Errors {
					tflog.Debug(ctx, "GitHub API validation error", map[string]any{"error_index": i, "resource": apiErr.Resource, "field": apiErr.Field, "code": apiErr.Code, "message": apiErr.Message})
				}
			}
		}
		return diag.FromErr(err)
	}
	id := strconv.FormatInt(orgSettings.GetID(), 10)
	d.SetId(id)

	return resourceGithubOrganizationSettingsRead(ctx, d, meta)
}

func resourceGithubOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client

	org := meta.(*Owner).name

	orgSettings, _, err := client.Organizations.Get(ctx, org)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("billing_email", orgSettings.GetBillingEmail()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("company", orgSettings.GetCompany()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("email", orgSettings.GetEmail()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("twitter_username", orgSettings.GetTwitterUsername()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("location", orgSettings.GetLocation()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", orgSettings.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", orgSettings.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("has_organization_projects", orgSettings.GetHasOrganizationProjects()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("has_repository_projects", orgSettings.GetHasRepositoryProjects()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("default_repository_permission", orgSettings.GetDefaultRepoPermission()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_repositories", orgSettings.GetMembersCanCreateRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_internal_repositories", orgSettings.GetMembersCanCreateInternalRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_private_repositories", orgSettings.GetMembersCanCreatePrivateRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_public_repositories", orgSettings.GetMembersCanCreatePublicRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_pages", orgSettings.GetMembersCanCreatePages()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_public_pages", orgSettings.GetMembersCanCreatePublicPages()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_create_private_pages", orgSettings.GetMembersCanCreatePrivatePages()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("members_can_fork_private_repositories", orgSettings.GetMembersCanForkPrivateRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("web_commit_signoff_required", orgSettings.GetWebCommitSignoffRequired()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("blog", orgSettings.GetBlog()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("advanced_security_enabled_for_new_repositories", orgSettings.GetAdvancedSecurityEnabledForNewRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependabot_alerts_enabled_for_new_repositories", orgSettings.GetDependabotAlertsEnabledForNewRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependabot_security_updates_enabled_for_new_repositories", orgSettings.GetDependabotSecurityUpdatesEnabledForNewRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependency_graph_enabled_for_new_repositories", orgSettings.GetDependencyGraphEnabledForNewRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning_enabled_for_new_repositories", orgSettings.GetSecretScanningEnabledForNewRepos()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning_push_protection_enabled_for_new_repositories", orgSettings.GetSecretScanningPushProtectionEnabledForNewRepos()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubOrganizationSettingsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	ctx = context.WithValue(ctx, ctxId, d.Id())
	org := meta.(*Owner).name
	tflog.Debug(ctx, "Reverting Organization Settings to default values", map[string]any{"owner": org})

	// Get organization info to determine if it's enterprise
	orgInfo, _, err := client.Organizations.Get(ctx, org)
	if err != nil {
		return diag.FromErr(err)
	}

	// Build minimal settings with only required fields
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"
	defaultSettings := &github.Organization{
		BillingEmail: new("email@example.com"),
	}

	// Only add enterprise-specific fields if it's an enterprise org
	if isEnterprise {
		defaultSettings.MembersCanCreateInternalRepos = new(true)
	}

	_, _, err = client.Organizations.Edit(ctx, org, defaultSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
