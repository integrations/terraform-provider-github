package github

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/go-cty/cty"
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

// stringAttr returns the string value of the named attribute. The resource
// schema guarantees the type, so a failed assertion can only mean the schema
// and this code disagree, in which case the zero value is the safe result.
func stringAttr(d *schema.ResourceData, name string) string {
	v, _ := d.Get(name).(string)

	return v
}

// boolAttr returns the boolean value of the named attribute, with the same
// type guarantee as stringAttr.
func boolAttr(d *schema.ResourceData, name string) bool {
	v, _ := d.Get(name).(bool)

	return v
}

// isConfigured reports whether the named attribute is explicitly set in the
// configuration, as opposed to merely carrying its schema default.
//
// d.GetOk cannot answer this for booleans: it reports a value configured as
// false exactly like an unset one, which is why an explicit false used to be
// dropped from the create payload (#3493). The raw configuration keeps the
// distinction, so a null there means the user did not write the attribute.
//
// The raw configuration is absent on some code paths, such as import. In that
// case there is nothing to read the user's intent from, so this falls back to
// the previous d.GetOk behaviour rather than inventing one.
func isConfigured(d *schema.ResourceData, name string) bool {
	if d.GetRawConfig().IsNull() {
		_, ok := d.GetOk(name)

		return ok
	}

	v, diags := d.GetRawConfigAt(cty.GetAttrPath(name))
	if diags.HasError() {
		_, ok := d.GetOk(name)

		return ok
	}

	return !v.IsNull()
}

// sendOnCreate reports whether the named attribute belongs in the create
// payload: when the user explicitly configured it, or when d.GetOk reports it
// because its schema default is a non-zero value.
//
// The second half keeps the create payload identical to what #2807 sent.
// Attributes with a non-zero default (has_organization_projects,
// default_repository_permission, ...) have always been sent on create, and
// leaving them out would let the org keep a value that differs from the
// schema default, so the first plan after create is not clean and only the
// second apply converges. Attributes defaulting to false or "" are still
// only sent when explicitly configured, which is what keeps the request
// narrow enough to avoid the validation errors reported in #2305.
func sendOnCreate(d *schema.ResourceData, name string) bool {
	if isConfigured(d, name) {
		return true
	}

	_, ok := d.GetOk(name)

	return ok
}

// organizationSettingsForCreate builds the payload sent when the resource is
// created. An attribute is included when the user explicitly set it or when
// its schema default is a non-zero value; everything else stays out of the
// request and the org keeps whatever the API defaults to. Keeping the payload
// narrow is what avoids the validation errors reported in #2305.
//
// Explicit configuration is read from the raw configuration rather than
// d.GetOk, which cannot tell an explicitly configured false from an unset
// attribute and so used to drop a boolean configured as false from the
// create payload (#3493).
func organizationSettingsForCreate(d *schema.ResourceData, isEnterprise bool) *github.Organization {
	settings := &github.Organization{}

	// The API rejects the request without billing_email, so it is always sent.
	if sendOnCreate(d, "billing_email") {
		settings.BillingEmail = new(stringAttr(d, "billing_email"))
	}

	if sendOnCreate(d, "company") {
		settings.Company = new(stringAttr(d, "company"))
	}
	if sendOnCreate(d, "email") {
		settings.Email = new(stringAttr(d, "email"))
	}
	if sendOnCreate(d, "twitter_username") {
		settings.TwitterUsername = new(stringAttr(d, "twitter_username"))
	}
	if sendOnCreate(d, "location") {
		settings.Location = new(stringAttr(d, "location"))
	}
	if sendOnCreate(d, "name") {
		settings.Name = new(stringAttr(d, "name"))
	}
	if sendOnCreate(d, "description") {
		settings.Description = new(stringAttr(d, "description"))
	}
	if sendOnCreate(d, "blog") {
		settings.Blog = new(stringAttr(d, "blog"))
	}

	if sendOnCreate(d, "has_organization_projects") {
		settings.HasOrganizationProjects = new(boolAttr(d, "has_organization_projects"))
	}
	if sendOnCreate(d, "has_repository_projects") {
		settings.HasRepositoryProjects = new(boolAttr(d, "has_repository_projects"))
	}
	if sendOnCreate(d, "default_repository_permission") {
		settings.DefaultRepoPermission = new(stringAttr(d, "default_repository_permission"))
	}
	if sendOnCreate(d, "members_can_create_repositories") {
		settings.MembersCanCreateRepos = new(boolAttr(d, "members_can_create_repositories"))
	}
	if sendOnCreate(d, "members_can_create_private_repositories") {
		settings.MembersCanCreatePrivateRepos = new(boolAttr(d, "members_can_create_private_repositories"))
	}
	if sendOnCreate(d, "members_can_create_public_repositories") {
		settings.MembersCanCreatePublicRepos = new(boolAttr(d, "members_can_create_public_repositories"))
	}
	if sendOnCreate(d, "members_can_create_pages") {
		settings.MembersCanCreatePages = new(boolAttr(d, "members_can_create_pages"))
	}
	if sendOnCreate(d, "members_can_create_public_pages") {
		settings.MembersCanCreatePublicPages = new(boolAttr(d, "members_can_create_public_pages"))
	}
	if sendOnCreate(d, "members_can_create_private_pages") {
		settings.MembersCanCreatePrivatePages = new(boolAttr(d, "members_can_create_private_pages"))
	}
	if sendOnCreate(d, "members_can_fork_private_repositories") {
		settings.MembersCanForkPrivateRepos = new(boolAttr(d, "members_can_fork_private_repositories"))
	}
	if sendOnCreate(d, "web_commit_signoff_required") {
		settings.WebCommitSignoffRequired = new(boolAttr(d, "web_commit_signoff_required"))
	}
	if sendOnCreate(d, "advanced_security_enabled_for_new_repositories") {
		settings.AdvancedSecurityEnabledForNewRepos = new(boolAttr(d, "advanced_security_enabled_for_new_repositories"))
	}
	if sendOnCreate(d, "dependabot_alerts_enabled_for_new_repositories") {
		settings.DependabotAlertsEnabledForNewRepos = new(boolAttr(d, "dependabot_alerts_enabled_for_new_repositories"))
	}
	if sendOnCreate(d, "dependabot_security_updates_enabled_for_new_repositories") {
		settings.DependabotSecurityUpdatesEnabledForNewRepos = new(boolAttr(d, "dependabot_security_updates_enabled_for_new_repositories"))
	}
	if sendOnCreate(d, "dependency_graph_enabled_for_new_repositories") {
		settings.DependencyGraphEnabledForNewRepos = new(boolAttr(d, "dependency_graph_enabled_for_new_repositories"))
	}
	if sendOnCreate(d, "secret_scanning_enabled_for_new_repositories") {
		settings.SecretScanningEnabledForNewRepos = new(boolAttr(d, "secret_scanning_enabled_for_new_repositories"))
	}
	if sendOnCreate(d, "secret_scanning_push_protection_enabled_for_new_repositories") {
		settings.SecretScanningPushProtectionEnabledForNewRepos = new(boolAttr(d, "secret_scanning_push_protection_enabled_for_new_repositories"))
	}

	if isEnterprise {
		if sendOnCreate(d, "members_can_create_internal_repositories") {
			settings.MembersCanCreateInternalRepos = new(boolAttr(d, "members_can_create_internal_repositories"))
		}
	}

	return settings
}

// organizationSettingsForUpdate builds the payload sent when the resource is
// updated. Only attributes that actually changed are included, which keeps
// unconfigured attributes out of the request and avoids the API validation
// errors reported in #2305.
func organizationSettingsForUpdate(d *schema.ResourceData, isEnterprise bool) *github.Organization {
	settings := &github.Organization{}

	// The API rejects the request without billing_email, so it is always sent.
	if _, ok := d.GetOk("billing_email"); ok {
		settings.BillingEmail = new(stringAttr(d, "billing_email"))
	}

	if d.HasChange("company") {
		if _, ok := d.GetOk("company"); ok {
			settings.Company = new(stringAttr(d, "company"))
		}
	}
	if d.HasChange("email") {
		if _, ok := d.GetOk("email"); ok {
			settings.Email = new(stringAttr(d, "email"))
		}
	}
	if d.HasChange("twitter_username") {
		if _, ok := d.GetOk("twitter_username"); ok {
			settings.TwitterUsername = new(stringAttr(d, "twitter_username"))
		}
	}
	if d.HasChange("location") {
		if _, ok := d.GetOk("location"); ok {
			settings.Location = new(stringAttr(d, "location"))
		}
	}
	if d.HasChange("name") {
		if _, ok := d.GetOk("name"); ok {
			settings.Name = new(stringAttr(d, "name"))
		}
	}
	if d.HasChange("description") {
		if _, ok := d.GetOk("description"); ok {
			settings.Description = new(stringAttr(d, "description"))
		}
	}
	if d.HasChange("blog") {
		if _, ok := d.GetOk("blog"); ok {
			settings.Blog = new(stringAttr(d, "blog"))
		}
	}

	if d.HasChange("has_organization_projects") {
		settings.HasOrganizationProjects = new(boolAttr(d, "has_organization_projects"))
	}
	if d.HasChange("has_repository_projects") {
		settings.HasRepositoryProjects = new(boolAttr(d, "has_repository_projects"))
	}
	if d.HasChange("default_repository_permission") {
		if _, ok := d.GetOk("default_repository_permission"); ok {
			settings.DefaultRepoPermission = new(stringAttr(d, "default_repository_permission"))
		}
	}
	if d.HasChange("members_can_create_repositories") {
		settings.MembersCanCreateRepos = new(boolAttr(d, "members_can_create_repositories"))
	}
	if d.HasChange("members_can_create_private_repositories") {
		settings.MembersCanCreatePrivateRepos = new(boolAttr(d, "members_can_create_private_repositories"))
	}
	if d.HasChange("members_can_create_public_repositories") {
		settings.MembersCanCreatePublicRepos = new(boolAttr(d, "members_can_create_public_repositories"))
	}
	if d.HasChange("members_can_create_pages") {
		settings.MembersCanCreatePages = new(boolAttr(d, "members_can_create_pages"))
	}
	if d.HasChange("members_can_create_public_pages") {
		settings.MembersCanCreatePublicPages = new(boolAttr(d, "members_can_create_public_pages"))
	}
	if d.HasChange("members_can_create_private_pages") {
		settings.MembersCanCreatePrivatePages = new(boolAttr(d, "members_can_create_private_pages"))
	}
	if d.HasChange("members_can_fork_private_repositories") {
		settings.MembersCanForkPrivateRepos = new(boolAttr(d, "members_can_fork_private_repositories"))
	}
	if d.HasChange("web_commit_signoff_required") {
		settings.WebCommitSignoffRequired = new(boolAttr(d, "web_commit_signoff_required"))
	}
	if d.HasChange("advanced_security_enabled_for_new_repositories") {
		settings.AdvancedSecurityEnabledForNewRepos = new(boolAttr(d, "advanced_security_enabled_for_new_repositories"))
	}
	if d.HasChange("dependabot_alerts_enabled_for_new_repositories") {
		settings.DependabotAlertsEnabledForNewRepos = new(boolAttr(d, "dependabot_alerts_enabled_for_new_repositories"))
	}
	if d.HasChange("dependabot_security_updates_enabled_for_new_repositories") {
		settings.DependabotSecurityUpdatesEnabledForNewRepos = new(boolAttr(d, "dependabot_security_updates_enabled_for_new_repositories"))
	}
	if d.HasChange("dependency_graph_enabled_for_new_repositories") {
		settings.DependencyGraphEnabledForNewRepos = new(boolAttr(d, "dependency_graph_enabled_for_new_repositories"))
	}
	if d.HasChange("secret_scanning_enabled_for_new_repositories") {
		settings.SecretScanningEnabledForNewRepos = new(boolAttr(d, "secret_scanning_enabled_for_new_repositories"))
	}
	if d.HasChange("secret_scanning_push_protection_enabled_for_new_repositories") {
		settings.SecretScanningPushProtectionEnabledForNewRepos = new(boolAttr(d, "secret_scanning_push_protection_enabled_for_new_repositories"))
	}

	if isEnterprise {
		if d.HasChange("members_can_create_internal_repositories") {
			settings.MembersCanCreateInternalRepos = new(boolAttr(d, "members_can_create_internal_repositories"))
		}
	}

	return settings
}

// logOrganizationSettings records the payload about to be sent to the API.
func logOrganizationSettings(org string, isEnterprise bool, settings *github.Organization) {
	log.Printf("[DEBUG] Built settings for org %s (enterprise: %v)", org, isEnterprise)
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
}

// organizationPlanIsEnterprise reports whether the organization is on an
// enterprise plan, which gates the enterprise-only attributes.
func organizationPlanIsEnterprise(ctx context.Context, owner *Owner) (bool, error) {
	orgInfo, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
	if err != nil {
		return false, err
	}

	return orgInfo.GetPlan().GetName() == "enterprise", nil
}

// editOrganizationSettings sends the payload and refreshes state from the API.
func editOrganizationSettings(ctx context.Context, d *schema.ResourceData, owner *Owner, settings *github.Organization) diag.Diagnostics {
	orgSettings, _, err := owner.v3client.Organizations.Edit(ctx, owner.name, settings)
	if err != nil {
		// Log detailed error information for debugging
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			log.Printf("[DEBUG] GitHub API Error: Status=%d, Message=%s", ghErr.Response.StatusCode, ghErr.Message)
			if len(ghErr.Errors) > 0 {
				for i, apiErr := range ghErr.Errors {
					log.Printf("[DEBUG]   Error[%d]: Resource=%s, Field=%s, Code=%s, Message=%s",
						i, apiErr.Resource, apiErr.Field, apiErr.Code, apiErr.Message)
				}
			}
		}

		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(orgSettings.GetID(), 10))

	return resourceGithubOrganizationSettingsRead(ctx, d, owner)
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

	isEnterprise, err := organizationPlanIsEnterprise(ctx, owner)
	if err != nil {
		return diag.FromErr(err)
	}

	settings := organizationSettingsForCreate(d, isEnterprise)
	logOrganizationSettings(owner.name, isEnterprise, settings)

	return editOrganizationSettings(ctx, d, owner, settings)
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

	isEnterprise, err := organizationPlanIsEnterprise(ctx, owner)
	if err != nil {
		return diag.FromErr(err)
	}

	settings := organizationSettingsForUpdate(d, isEnterprise)
	logOrganizationSettings(owner.name, isEnterprise, settings)

	return editOrganizationSettings(ctx, d, owner, settings)
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

	log.Printf("[DEBUG] Reverting Organization Settings to default values: %s", owner.name)

	isEnterprise, err := organizationPlanIsEnterprise(ctx, owner)
	if err != nil {
		return diag.FromErr(err)
	}

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
