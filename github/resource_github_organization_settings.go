package github

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationSettingsCreateOrUpdate,
		Read:   resourceGithubOrganizationSettingsRead,
		Update: resourceGithubOrganizationSettingsCreateOrUpdate,
		Delete: resourceGithubOrganizationSettingsDelete,
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
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"read", "write", "admin", "none"}, false), "default_repository_permission"),
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
		settings.BillingEmail = github.Ptr(billingEmail.(string))
	}

	// Optional string fields - only set if should be included
	if shouldInclude("company") {
		if company, ok := d.GetOk("company"); ok {
			settings.Company = github.Ptr(company.(string))
		}
	}
	if shouldInclude("email") {
		if email, ok := d.GetOk("email"); ok {
			settings.Email = github.Ptr(email.(string))
		}
	}
	if shouldInclude("twitter_username") {
		if twitterUsername, ok := d.GetOk("twitter_username"); ok {
			settings.TwitterUsername = github.Ptr(twitterUsername.(string))
		}
	}
	if shouldInclude("location") {
		if location, ok := d.GetOk("location"); ok {
			settings.Location = github.Ptr(location.(string))
		}
	}
	if shouldInclude("name") {
		if name, ok := d.GetOk("name"); ok {
			settings.Name = github.Ptr(name.(string))
		}
	}
	if shouldInclude("description") {
		if description, ok := d.GetOk("description"); ok {
			settings.Description = github.Ptr(description.(string))
		}
	}
	if shouldInclude("blog") {
		if blog, ok := d.GetOk("blog"); ok {
			settings.Blog = github.Ptr(blog.(string))
		}
	}

	// Boolean fields - only set if should be included
	// Use d.Get() instead of d.GetOk() when shouldInclude() returns true,
	// because we already know the field should be included, and d.Get() correctly handles false values
	if shouldInclude("has_organization_projects") {
		settings.HasOrganizationProjects = github.Ptr(d.Get("has_organization_projects").(bool))
	}
	if shouldInclude("has_repository_projects") {
		settings.HasRepositoryProjects = github.Ptr(d.Get("has_repository_projects").(bool))
	}
	if shouldInclude("default_repository_permission") {
		if defaultRepoPermission, ok := d.GetOk("default_repository_permission"); ok {
			settings.DefaultRepoPermission = github.Ptr(defaultRepoPermission.(string))
		}
	}
	if shouldInclude("members_can_create_repositories") {
		settings.MembersCanCreateRepos = github.Ptr(d.Get("members_can_create_repositories").(bool))
	}
	if shouldInclude("members_can_create_private_repositories") {
		settings.MembersCanCreatePrivateRepos = github.Ptr(d.Get("members_can_create_private_repositories").(bool))
	}
	if shouldInclude("members_can_create_public_repositories") {
		settings.MembersCanCreatePublicRepos = github.Ptr(d.Get("members_can_create_public_repositories").(bool))
	}
	if shouldInclude("members_can_create_pages") {
		settings.MembersCanCreatePages = github.Ptr(d.Get("members_can_create_pages").(bool))
	}
	if shouldInclude("members_can_create_public_pages") {
		settings.MembersCanCreatePublicPages = github.Ptr(d.Get("members_can_create_public_pages").(bool))
	}
	if shouldInclude("members_can_create_private_pages") {
		settings.MembersCanCreatePrivatePages = github.Ptr(d.Get("members_can_create_private_pages").(bool))
	}
	if shouldInclude("members_can_fork_private_repositories") {
		settings.MembersCanForkPrivateRepos = github.Ptr(d.Get("members_can_fork_private_repositories").(bool))
	}
	if shouldInclude("web_commit_signoff_required") {
		settings.WebCommitSignoffRequired = github.Ptr(d.Get("web_commit_signoff_required").(bool))
	}
	if shouldInclude("advanced_security_enabled_for_new_repositories") {
		settings.AdvancedSecurityEnabledForNewRepos = github.Ptr(d.Get("advanced_security_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("dependabot_alerts_enabled_for_new_repositories") {
		settings.DependabotAlertsEnabledForNewRepos = github.Ptr(d.Get("dependabot_alerts_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("dependabot_security_updates_enabled_for_new_repositories") {
		settings.DependabotSecurityUpdatesEnabledForNewRepos = github.Ptr(d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("dependency_graph_enabled_for_new_repositories") {
		settings.DependencyGraphEnabledForNewRepos = github.Ptr(d.Get("dependency_graph_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("secret_scanning_enabled_for_new_repositories") {
		settings.SecretScanningEnabledForNewRepos = github.Ptr(d.Get("secret_scanning_enabled_for_new_repositories").(bool))
	}
	if shouldInclude("secret_scanning_push_protection_enabled_for_new_repositories") {
		settings.SecretScanningPushProtectionEnabledForNewRepos = github.Ptr(d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool))
	}

	// Enterprise-specific field
	if isEnterprise {
		if shouldInclude("members_can_create_internal_repositories") {
			settings.MembersCanCreateInternalRepos = github.Ptr(d.Get("members_can_create_internal_repositories").(bool))
		}
	}

	return settings
}

func resourceGithubOrganizationSettingsCreateOrUpdate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	orgInfo, _, err := client.Organizations.Get(ctx, org)
	if err != nil {
		return err
	}

	// Build settings using helper function
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"
	settings := buildOrganizationSettings(d, isEnterprise)

	// Debug: Log the settings being sent to the API with detailed field information
	log.Printf("[DEBUG] Built settings for org %s (enterprise: %v)", org, isEnterprise)
	if settings.BillingEmail != nil {
		log.Printf("[DEBUG]   BillingEmail: %s", *settings.BillingEmail)
	}
	if settings.Company != nil {
		log.Printf("[DEBUG]   Company: %s", *settings.Company)
	}
	if settings.Email != nil {
		log.Printf("[DEBUG]   Email: %s", *settings.Email)
	}
	if settings.TwitterUsername != nil {
		log.Printf("[DEBUG]   TwitterUsername: %s", *settings.TwitterUsername)
	}
	if settings.Location != nil {
		log.Printf("[DEBUG]   Location: %s", *settings.Location)
	}
	if settings.Name != nil {
		log.Printf("[DEBUG]   Name: %s", *settings.Name)
	}
	if settings.Description != nil {
		log.Printf("[DEBUG]   Description: %s", *settings.Description)
	}
	if settings.Blog != nil {
		log.Printf("[DEBUG]   Blog: %s", *settings.Blog)
	}
	if settings.HasOrganizationProjects != nil {
		log.Printf("[DEBUG]   HasOrganizationProjects: %v", *settings.HasOrganizationProjects)
	}
	if settings.HasRepositoryProjects != nil {
		log.Printf("[DEBUG]   HasRepositoryProjects: %v", *settings.HasRepositoryProjects)
	}
	if settings.DefaultRepoPermission != nil {
		log.Printf("[DEBUG]   DefaultRepoPermission: %s", *settings.DefaultRepoPermission)
	}
	if settings.MembersCanCreateRepos != nil {
		log.Printf("[DEBUG]   MembersCanCreateRepos: %v", *settings.MembersCanCreateRepos)
	}
	if settings.MembersCanCreatePrivateRepos != nil {
		log.Printf("[DEBUG]   MembersCanCreatePrivateRepos: %v", *settings.MembersCanCreatePrivateRepos)
	}
	if settings.MembersCanCreatePublicRepos != nil {
		log.Printf("[DEBUG]   MembersCanCreatePublicRepos: %v", *settings.MembersCanCreatePublicRepos)
	}
	if settings.MembersCanCreateInternalRepos != nil {
		log.Printf("[DEBUG]   MembersCanCreateInternalRepos: %v", *settings.MembersCanCreateInternalRepos)
	}
	if settings.MembersCanCreatePages != nil {
		log.Printf("[DEBUG]   MembersCanCreatePages: %v", *settings.MembersCanCreatePages)
	}
	if settings.MembersCanCreatePublicPages != nil {
		log.Printf("[DEBUG]   MembersCanCreatePublicPages: %v", *settings.MembersCanCreatePublicPages)
	}
	if settings.MembersCanCreatePrivatePages != nil {
		log.Printf("[DEBUG]   MembersCanCreatePrivatePages: %v", *settings.MembersCanCreatePrivatePages)
	}
	if settings.MembersCanForkPrivateRepos != nil {
		log.Printf("[DEBUG]   MembersCanForkPrivateRepos: %v", *settings.MembersCanForkPrivateRepos)
	}
	if settings.WebCommitSignoffRequired != nil {
		log.Printf("[DEBUG]   WebCommitSignoffRequired: %v", *settings.WebCommitSignoffRequired)
	}
	if settings.AdvancedSecurityEnabledForNewRepos != nil {
		log.Printf("[DEBUG]   AdvancedSecurityEnabledForNewRepos: %v", *settings.AdvancedSecurityEnabledForNewRepos)
	}
	if settings.DependabotAlertsEnabledForNewRepos != nil {
		log.Printf("[DEBUG]   DependabotAlertsEnabledForNewRepos: %v", *settings.DependabotAlertsEnabledForNewRepos)
	}
	if settings.DependabotSecurityUpdatesEnabledForNewRepos != nil {
		log.Printf("[DEBUG]   DependabotSecurityUpdatesEnabledForNewRepos: %v", *settings.DependabotSecurityUpdatesEnabledForNewRepos)
	}
	if settings.DependencyGraphEnabledForNewRepos != nil {
		log.Printf("[DEBUG]   DependencyGraphEnabledForNewRepos: %v", *settings.DependencyGraphEnabledForNewRepos)
	}
	if settings.SecretScanningEnabledForNewRepos != nil {
		log.Printf("[DEBUG]   SecretScanningEnabledForNewRepos: %v", *settings.SecretScanningEnabledForNewRepos)
	}
	if settings.SecretScanningPushProtectionEnabledForNewRepos != nil {
		log.Printf("[DEBUG]   SecretScanningPushProtectionEnabledForNewRepos: %v", *settings.SecretScanningPushProtectionEnabledForNewRepos)
	}

	orgSettings, _, err := client.Organizations.Edit(ctx, org, settings)
	if err != nil {
		// Log detailed error information for debugging
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			log.Printf("[DEBUG] GitHub API Error: Status=%d, Message=%s", ghErr.Response.StatusCode, ghErr.Message)
			if len(ghErr.Errors) > 0 {
				for i, apiErr := range ghErr.Errors {
					log.Printf("[DEBUG]   Error[%d]: Resource=%s, Field=%s, Code=%s, Message=%s",
						i, apiErr.Resource, apiErr.Field, apiErr.Code, apiErr.Message)
				}
			}
		}
		return err
	}
	id := strconv.FormatInt(orgSettings.GetID(), 10)
	d.SetId(id)

	return resourceGithubOrganizationSettingsRead(d, meta)
}

func resourceGithubOrganizationSettingsRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.Background()
	org := meta.(*Owner).name

	orgSettings, _, err := client.Organizations.Get(ctx, org)
	if err != nil {
		return err
	}

	if err = d.Set("billing_email", orgSettings.GetBillingEmail()); err != nil {
		return err
	}
	if err = d.Set("company", orgSettings.GetCompany()); err != nil {
		return err
	}
	if err = d.Set("email", orgSettings.GetEmail()); err != nil {
		return err
	}
	if err = d.Set("twitter_username", orgSettings.GetTwitterUsername()); err != nil {
		return err
	}
	if err = d.Set("location", orgSettings.GetLocation()); err != nil {
		return err
	}
	if err = d.Set("name", orgSettings.GetName()); err != nil {
		return err
	}
	if err = d.Set("description", orgSettings.GetDescription()); err != nil {
		return err
	}
	if err = d.Set("has_organization_projects", orgSettings.GetHasOrganizationProjects()); err != nil {
		return err
	}
	if err = d.Set("has_repository_projects", orgSettings.GetHasRepositoryProjects()); err != nil {
		return err
	}
	if err = d.Set("default_repository_permission", orgSettings.GetDefaultRepoPermission()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_repositories", orgSettings.GetMembersCanCreateRepos()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_internal_repositories", orgSettings.GetMembersCanCreateInternalRepos()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_private_repositories", orgSettings.GetMembersCanCreatePrivateRepos()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_public_repositories", orgSettings.GetMembersCanCreatePublicRepos()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_pages", orgSettings.GetMembersCanCreatePages()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_public_pages", orgSettings.GetMembersCanCreatePublicPages()); err != nil {
		return err
	}
	if err = d.Set("members_can_create_private_pages", orgSettings.GetMembersCanCreatePrivatePages()); err != nil {
		return err
	}
	if err = d.Set("members_can_fork_private_repositories", orgSettings.GetMembersCanForkPrivateRepos()); err != nil {
		return err
	}
	if err = d.Set("web_commit_signoff_required", orgSettings.GetWebCommitSignoffRequired()); err != nil {
		return err
	}
	if err = d.Set("blog", orgSettings.GetBlog()); err != nil {
		return err
	}
	if err = d.Set("advanced_security_enabled_for_new_repositories", orgSettings.GetAdvancedSecurityEnabledForNewRepos()); err != nil {
		return err
	}
	if err = d.Set("dependabot_alerts_enabled_for_new_repositories", orgSettings.GetDependabotAlertsEnabledForNewRepos()); err != nil {
		return err
	}
	if err = d.Set("dependabot_security_updates_enabled_for_new_repositories", orgSettings.GetDependabotSecurityUpdatesEnabledForNewRepos()); err != nil {
		return err
	}
	if err = d.Set("dependency_graph_enabled_for_new_repositories", orgSettings.GetDependencyGraphEnabledForNewRepos()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_enabled_for_new_repositories", orgSettings.GetSecretScanningEnabledForNewRepos()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_push_protection_enabled_for_new_repositories", orgSettings.GetSecretScanningPushProtectionEnabledForNewRepos()); err != nil {
		return err
	}
	return nil
}

func resourceGithubOrganizationSettingsDelete(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	log.Printf("[DEBUG] Reverting Organization Settings to default values: %s", org)

	// Get organization info to determine if it's enterprise
	orgInfo, _, err := client.Organizations.Get(ctx, org)
	if err != nil {
		return err
	}

	// Build minimal settings with only required fields
	isEnterprise := orgInfo.GetPlan().GetName() == "enterprise"
	defaultSettings := &github.Organization{
		BillingEmail: github.Ptr("email@example.com"),
	}

	// Only add enterprise-specific fields if it's an enterprise org
	if isEnterprise {
		defaultSettings.MembersCanCreateInternalRepos = github.Ptr(true)
	}

	_, _, err = client.Organizations.Edit(ctx, org, defaultSettings)
	if err != nil {
		return err
	}

	return nil
}
