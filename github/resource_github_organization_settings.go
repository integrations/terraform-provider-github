package github

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationSettingsCreate,
		ReadContext:   resourceGithubOrganizationSettingsRead,
		UpdateContext: resourceGithubOrganizationSettingsUpdate,
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

func resourceGithubOrganizationSettingsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider meta type %T", meta)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	orgInfo, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
	if err != nil {
		return diag.FromErr(err)
	}
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"

	// An attribute is sent on create when the user configured it, including a
	// boolean configured as false (#3493), or when its schema default is a
	// non-zero value. Everything else stays out of the request, which keeps it
	// narrow enough to avoid the validation errors reported in #2305.
	settings := &github.Organization{}

	if _, ok := d.GetOk("billing_email"); ok || isConfigured(d, "billing_email") {
		v, _ := d.Get("billing_email").(string)
		settings.BillingEmail = new(v)
	}
	if _, ok := d.GetOk("company"); ok || isConfigured(d, "company") {
		v, _ := d.Get("company").(string)
		settings.Company = new(v)
	}
	if _, ok := d.GetOk("email"); ok || isConfigured(d, "email") {
		v, _ := d.Get("email").(string)
		settings.Email = new(v)
	}
	if _, ok := d.GetOk("twitter_username"); ok || isConfigured(d, "twitter_username") {
		v, _ := d.Get("twitter_username").(string)
		settings.TwitterUsername = new(v)
	}
	if _, ok := d.GetOk("location"); ok || isConfigured(d, "location") {
		v, _ := d.Get("location").(string)
		settings.Location = new(v)
	}
	if _, ok := d.GetOk("name"); ok || isConfigured(d, "name") {
		v, _ := d.Get("name").(string)
		settings.Name = new(v)
	}
	if _, ok := d.GetOk("description"); ok || isConfigured(d, "description") {
		v, _ := d.Get("description").(string)
		settings.Description = new(v)
	}
	if _, ok := d.GetOk("blog"); ok || isConfigured(d, "blog") {
		v, _ := d.Get("blog").(string)
		settings.Blog = new(v)
	}
	if _, ok := d.GetOk("has_organization_projects"); ok || isConfigured(d, "has_organization_projects") {
		v, _ := d.Get("has_organization_projects").(bool)
		settings.HasOrganizationProjects = new(v)
	}
	if _, ok := d.GetOk("has_repository_projects"); ok || isConfigured(d, "has_repository_projects") {
		v, _ := d.Get("has_repository_projects").(bool)
		settings.HasRepositoryProjects = new(v)
	}
	if _, ok := d.GetOk("default_repository_permission"); ok || isConfigured(d, "default_repository_permission") {
		v, _ := d.Get("default_repository_permission").(string)
		settings.DefaultRepoPermission = new(v)
	}
	if _, ok := d.GetOk("members_can_create_repositories"); ok || isConfigured(d, "members_can_create_repositories") {
		v, _ := d.Get("members_can_create_repositories").(bool)
		settings.MembersCanCreateRepos = new(v)
	}
	if _, ok := d.GetOk("members_can_create_private_repositories"); ok || isConfigured(d, "members_can_create_private_repositories") {
		v, _ := d.Get("members_can_create_private_repositories").(bool)
		settings.MembersCanCreatePrivateRepos = new(v)
	}
	if _, ok := d.GetOk("members_can_create_public_repositories"); ok || isConfigured(d, "members_can_create_public_repositories") {
		v, _ := d.Get("members_can_create_public_repositories").(bool)
		settings.MembersCanCreatePublicRepos = new(v)
	}
	if _, ok := d.GetOk("members_can_create_pages"); ok || isConfigured(d, "members_can_create_pages") {
		v, _ := d.Get("members_can_create_pages").(bool)
		settings.MembersCanCreatePages = new(v)
	}
	if _, ok := d.GetOk("members_can_create_public_pages"); ok || isConfigured(d, "members_can_create_public_pages") {
		v, _ := d.Get("members_can_create_public_pages").(bool)
		settings.MembersCanCreatePublicPages = new(v)
	}
	if _, ok := d.GetOk("members_can_create_private_pages"); ok || isConfigured(d, "members_can_create_private_pages") {
		v, _ := d.Get("members_can_create_private_pages").(bool)
		settings.MembersCanCreatePrivatePages = new(v)
	}
	if _, ok := d.GetOk("members_can_fork_private_repositories"); ok || isConfigured(d, "members_can_fork_private_repositories") {
		v, _ := d.Get("members_can_fork_private_repositories").(bool)
		settings.MembersCanForkPrivateRepos = new(v)
	}
	if _, ok := d.GetOk("web_commit_signoff_required"); ok || isConfigured(d, "web_commit_signoff_required") {
		v, _ := d.Get("web_commit_signoff_required").(bool)
		settings.WebCommitSignoffRequired = new(v)
	}
	if _, ok := d.GetOk("advanced_security_enabled_for_new_repositories"); ok || isConfigured(d, "advanced_security_enabled_for_new_repositories") {
		v, _ := d.Get("advanced_security_enabled_for_new_repositories").(bool)
		settings.AdvancedSecurityEnabledForNewRepos = new(v)
	}
	if _, ok := d.GetOk("dependabot_alerts_enabled_for_new_repositories"); ok || isConfigured(d, "dependabot_alerts_enabled_for_new_repositories") {
		v, _ := d.Get("dependabot_alerts_enabled_for_new_repositories").(bool)
		settings.DependabotAlertsEnabledForNewRepos = new(v)
	}
	if _, ok := d.GetOk("dependabot_security_updates_enabled_for_new_repositories"); ok || isConfigured(d, "dependabot_security_updates_enabled_for_new_repositories") {
		v, _ := d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool)
		settings.DependabotSecurityUpdatesEnabledForNewRepos = new(v)
	}
	if _, ok := d.GetOk("dependency_graph_enabled_for_new_repositories"); ok || isConfigured(d, "dependency_graph_enabled_for_new_repositories") {
		v, _ := d.Get("dependency_graph_enabled_for_new_repositories").(bool)
		settings.DependencyGraphEnabledForNewRepos = new(v)
	}
	if _, ok := d.GetOk("secret_scanning_enabled_for_new_repositories"); ok || isConfigured(d, "secret_scanning_enabled_for_new_repositories") {
		v, _ := d.Get("secret_scanning_enabled_for_new_repositories").(bool)
		settings.SecretScanningEnabledForNewRepos = new(v)
	}
	if _, ok := d.GetOk("secret_scanning_push_protection_enabled_for_new_repositories"); ok || isConfigured(d, "secret_scanning_push_protection_enabled_for_new_repositories") {
		v, _ := d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool)
		settings.SecretScanningPushProtectionEnabledForNewRepos = new(v)
	}

	if isEnterprise {
		if _, ok := d.GetOk("members_can_create_internal_repositories"); ok || isConfigured(d, "members_can_create_internal_repositories") {
			v, _ := d.Get("members_can_create_internal_repositories").(bool)
			settings.MembersCanCreateInternalRepos = new(v)
		}
	}

	tflog.Debug(ctx, "Creating organization settings", map[string]any{
		"org":        owner.name,
		"enterprise": isEnterprise,
		"settings":   settings.String(),
	})

	orgSettings, _, err := owner.v3client.Organizations.Edit(ctx, owner.name, settings)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			tflog.Debug(ctx, "GitHub API error", map[string]any{
				"status":  ghErr.Response.StatusCode,
				"message": ghErr.Message,
				"errors":  ghErr.Errors,
			})
		}

		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(orgSettings.GetID(), 10))

	return nil
}

func resourceGithubOrganizationSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider meta type %T", meta)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	orgInfo, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
	if err != nil {
		return diag.FromErr(err)
	}
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"

	// Only attributes that changed are sent, which keeps unconfigured attributes
	// out of the request and avoids the validation errors reported in #2305.
	settings := &github.Organization{}

	// The API rejects the request without billing_email, so it is always sent.
	if v, ok := d.GetOk("billing_email"); ok {
		billingEmail, _ := v.(string)
		settings.BillingEmail = new(billingEmail)
	}

	if d.HasChange("company") {
		if v, ok := d.GetOk("company"); ok {
			s, _ := v.(string)
			settings.Company = new(s)
		}
	}
	if d.HasChange("email") {
		if v, ok := d.GetOk("email"); ok {
			s, _ := v.(string)
			settings.Email = new(s)
		}
	}
	if d.HasChange("twitter_username") {
		if v, ok := d.GetOk("twitter_username"); ok {
			s, _ := v.(string)
			settings.TwitterUsername = new(s)
		}
	}
	if d.HasChange("location") {
		if v, ok := d.GetOk("location"); ok {
			s, _ := v.(string)
			settings.Location = new(s)
		}
	}
	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			s, _ := v.(string)
			settings.Name = new(s)
		}
	}
	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			s, _ := v.(string)
			settings.Description = new(s)
		}
	}
	if d.HasChange("blog") {
		if v, ok := d.GetOk("blog"); ok {
			s, _ := v.(string)
			settings.Blog = new(s)
		}
	}
	if d.HasChange("has_organization_projects") {
		v, _ := d.Get("has_organization_projects").(bool)
		settings.HasOrganizationProjects = new(v)
	}
	if d.HasChange("has_repository_projects") {
		v, _ := d.Get("has_repository_projects").(bool)
		settings.HasRepositoryProjects = new(v)
	}
	if d.HasChange("default_repository_permission") {
		if v, ok := d.GetOk("default_repository_permission"); ok {
			s, _ := v.(string)
			settings.DefaultRepoPermission = new(s)
		}
	}
	if d.HasChange("members_can_create_repositories") {
		v, _ := d.Get("members_can_create_repositories").(bool)
		settings.MembersCanCreateRepos = new(v)
	}
	if d.HasChange("members_can_create_private_repositories") {
		v, _ := d.Get("members_can_create_private_repositories").(bool)
		settings.MembersCanCreatePrivateRepos = new(v)
	}
	if d.HasChange("members_can_create_public_repositories") {
		v, _ := d.Get("members_can_create_public_repositories").(bool)
		settings.MembersCanCreatePublicRepos = new(v)
	}
	if d.HasChange("members_can_create_pages") {
		v, _ := d.Get("members_can_create_pages").(bool)
		settings.MembersCanCreatePages = new(v)
	}
	if d.HasChange("members_can_create_public_pages") {
		v, _ := d.Get("members_can_create_public_pages").(bool)
		settings.MembersCanCreatePublicPages = new(v)
	}
	if d.HasChange("members_can_create_private_pages") {
		v, _ := d.Get("members_can_create_private_pages").(bool)
		settings.MembersCanCreatePrivatePages = new(v)
	}
	if d.HasChange("members_can_fork_private_repositories") {
		v, _ := d.Get("members_can_fork_private_repositories").(bool)
		settings.MembersCanForkPrivateRepos = new(v)
	}
	if d.HasChange("web_commit_signoff_required") {
		v, _ := d.Get("web_commit_signoff_required").(bool)
		settings.WebCommitSignoffRequired = new(v)
	}
	if d.HasChange("advanced_security_enabled_for_new_repositories") {
		v, _ := d.Get("advanced_security_enabled_for_new_repositories").(bool)
		settings.AdvancedSecurityEnabledForNewRepos = new(v)
	}
	if d.HasChange("dependabot_alerts_enabled_for_new_repositories") {
		v, _ := d.Get("dependabot_alerts_enabled_for_new_repositories").(bool)
		settings.DependabotAlertsEnabledForNewRepos = new(v)
	}
	if d.HasChange("dependabot_security_updates_enabled_for_new_repositories") {
		v, _ := d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool)
		settings.DependabotSecurityUpdatesEnabledForNewRepos = new(v)
	}
	if d.HasChange("dependency_graph_enabled_for_new_repositories") {
		v, _ := d.Get("dependency_graph_enabled_for_new_repositories").(bool)
		settings.DependencyGraphEnabledForNewRepos = new(v)
	}
	if d.HasChange("secret_scanning_enabled_for_new_repositories") {
		v, _ := d.Get("secret_scanning_enabled_for_new_repositories").(bool)
		settings.SecretScanningEnabledForNewRepos = new(v)
	}
	if d.HasChange("secret_scanning_push_protection_enabled_for_new_repositories") {
		v, _ := d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool)
		settings.SecretScanningPushProtectionEnabledForNewRepos = new(v)
	}

	if isEnterprise {
		if d.HasChange("members_can_create_internal_repositories") {
			v, _ := d.Get("members_can_create_internal_repositories").(bool)
			settings.MembersCanCreateInternalRepos = new(v)
		}
	}

	tflog.Debug(ctx, "Updating organization settings", map[string]any{
		"org":        owner.name,
		"enterprise": isEnterprise,
		"settings":   settings.String(),
	})

	if _, _, err := owner.v3client.Organizations.Edit(ctx, owner.name, settings); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			tflog.Debug(ctx, "GitHub API error", map[string]any{
				"status":  ghErr.Response.StatusCode,
				"message": ghErr.Message,
				"errors":  ghErr.Errors,
			})
		}

		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider meta type %T", meta)
	}

	orgSettings, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
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
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider meta type %T", meta)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	tflog.Debug(ctx, "Reverting organization settings to default values", map[string]any{"org": owner.name})

	orgInfo, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
	if err != nil {
		return diag.FromErr(err)
	}
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"

	// Build minimal settings with only required fields
	defaultSettings := &github.Organization{
		BillingEmail: new("email@example.com"),
	}

	// Only add enterprise-specific fields if it's an enterprise org
	if isEnterprise {
		defaultSettings.MembersCanCreateInternalRepos = new(true)
	}

	if _, _, err := owner.v3client.Organizations.Edit(ctx, owner.name, defaultSettings); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
