package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v67/github"
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

// buildOrganizationSettings creates a github.Organization struct with only the fields that are explicitly configured
func buildOrganizationSettings(d *schema.ResourceData, isEnterprise bool) *github.Organization {
	settings := &github.Organization{}

	// Required field
	if billingEmail, ok := d.GetOk("billing_email"); ok {
		settings.BillingEmail = github.String(billingEmail.(string))
	}

	// Optional string fields - only set if explicitly configured
	if company, ok := d.GetOk("company"); ok {
		settings.Company = github.String(company.(string))
	}
	if email, ok := d.GetOk("email"); ok {
		settings.Email = github.String(email.(string))
	}
	if twitterUsername, ok := d.GetOk("twitter_username"); ok {
		settings.TwitterUsername = github.String(twitterUsername.(string))
	}
	if location, ok := d.GetOk("location"); ok {
		settings.Location = github.String(location.(string))
	}
	if name, ok := d.GetOk("name"); ok {
		settings.Name = github.String(name.(string))
	}
	if description, ok := d.GetOk("description"); ok {
		settings.Description = github.String(description.(string))
	}
	if blog, ok := d.GetOk("blog"); ok {
		settings.Blog = github.String(blog.(string))
	}

	// Boolean fields - only set if explicitly configured
	if hasOrgProjects, ok := d.GetOk("has_organization_projects"); ok {
		settings.HasOrganizationProjects = github.Bool(hasOrgProjects.(bool))
	}
	if hasRepoProjects, ok := d.GetOk("has_repository_projects"); ok {
		settings.HasRepositoryProjects = github.Bool(hasRepoProjects.(bool))
	}
	if defaultRepoPermission, ok := d.GetOk("default_repository_permission"); ok {
		settings.DefaultRepoPermission = github.String(defaultRepoPermission.(string))
	}
	if membersCanCreateRepos, ok := d.GetOk("members_can_create_repositories"); ok {
		settings.MembersCanCreateRepos = github.Bool(membersCanCreateRepos.(bool))
	}
	if _, exists := d.GetOkExists("members_can_create_private_repositories"); exists {
		settings.MembersCanCreatePrivateRepos = github.Bool(d.Get("members_can_create_private_repositories").(bool))
	}
	if membersCanCreatePublicRepos, ok := d.GetOk("members_can_create_public_repositories"); ok {
		settings.MembersCanCreatePublicRepos = github.Bool(membersCanCreatePublicRepos.(bool))
	}
	if membersCanCreatePages, ok := d.GetOk("members_can_create_pages"); ok {
		settings.MembersCanCreatePages = github.Bool(membersCanCreatePages.(bool))
	}
	if membersCanCreatePublicPages, ok := d.GetOk("members_can_create_public_pages"); ok {
		settings.MembersCanCreatePublicPages = github.Bool(membersCanCreatePublicPages.(bool))
	}
	if membersCanCreatePrivatePages, ok := d.GetOk("members_can_create_private_pages"); ok {
		settings.MembersCanCreatePrivatePages = github.Bool(membersCanCreatePrivatePages.(bool))
	}
	if membersCanForkPrivateRepos, ok := d.GetOk("members_can_fork_private_repositories"); ok {
		settings.MembersCanForkPrivateRepos = github.Bool(membersCanForkPrivateRepos.(bool))
	}
	if webCommitSignoffRequired, ok := d.GetOk("web_commit_signoff_required"); ok {
		settings.WebCommitSignoffRequired = github.Bool(webCommitSignoffRequired.(bool))
	}
	if advancedSecurityEnabled, ok := d.GetOk("advanced_security_enabled_for_new_repositories"); ok {
		settings.AdvancedSecurityEnabledForNewRepos = github.Bool(advancedSecurityEnabled.(bool))
	}
	if dependabotAlertsEnabled, ok := d.GetOk("dependabot_alerts_enabled_for_new_repositories"); ok {
		settings.DependabotAlertsEnabledForNewRepos = github.Bool(dependabotAlertsEnabled.(bool))
	}
	if dependabotSecurityUpdatesEnabled, ok := d.GetOk("dependabot_security_updates_enabled_for_new_repositories"); ok {
		settings.DependabotSecurityUpdatesEnabledForNewRepos = github.Bool(dependabotSecurityUpdatesEnabled.(bool))
	}
	if dependencyGraphEnabled, ok := d.GetOk("dependency_graph_enabled_for_new_repositories"); ok {
		settings.DependencyGraphEnabledForNewRepos = github.Bool(dependencyGraphEnabled.(bool))
	}
	if secretScanningEnabled, ok := d.GetOk("secret_scanning_enabled_for_new_repositories"); ok {
		settings.SecretScanningEnabledForNewRepos = github.Bool(secretScanningEnabled.(bool))
	}
	if secretScanningPushProtectionEnabled, ok := d.GetOk("secret_scanning_push_protection_enabled_for_new_repositories"); ok {
		settings.SecretScanningPushProtectionEnabledForNewRepos = github.Bool(secretScanningPushProtectionEnabled.(bool))
	}

	// Enterprise-specific field
	if isEnterprise {
		if _, exists := d.GetOkExists("members_can_create_internal_repositories"); exists {
			settings.MembersCanCreateInternalRepos = github.Bool(d.Get("members_can_create_internal_repositories").(bool))
		}
	}

	return settings
}

func resourceGithubOrganizationSettingsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
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

	orgSettings, _, err := client.Organizations.Edit(ctx, org, settings)
	if err != nil {
		return err
	}
	id := strconv.FormatInt(orgSettings.GetID(), 10)
	d.SetId(id)

	return resourceGithubOrganizationSettingsRead(d, meta)
}

func resourceGithubOrganizationSettingsRead(d *schema.ResourceData, meta interface{}) error {
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

func resourceGithubOrganizationSettingsDelete(d *schema.ResourceData, meta interface{}) error {
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
		BillingEmail: github.String("email@example.com"),
	}

	// Only add enterprise-specific fields if it's an enterprise org
	if isEnterprise {
		defaultSettings.MembersCanCreateInternalRepos = github.Bool(true)
	}

	_, _, err = client.Organizations.Edit(ctx, org, defaultSettings)
	if err != nil {
		return err
	}

	return nil
}
