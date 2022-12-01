package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v48/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationSettingsCreateOrUpdate,
		Read:   resourceGithubOrganizationSettingsRead,
		Update: resourceGithubOrganizationSettingsCreateOrUpdate,
		Delete: resourceGithubOrganizationSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"billing_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"company": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"twitter_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"has_organization_projects": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"has_repository_projects": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"default_repository_permission": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "read",
				ValidateFunc: validation.StringInSlice([]string{"read", "write", "admin", "none"}, false),
			},
			"members_can_create_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"members_can_create_internal_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.",
			},
			"members_can_create_private_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"members_can_create_public_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"members_can_create_pages": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"members_can_create_public_pages": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"members_can_create_private_pages": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"members_can_fork_private_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"web_commit_signoff_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"blog": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"advanced_security_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dependabot_alerts_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dependabot_security_updates_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dependency_graph_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"secret_scanning_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"secret_scanning_push_protection_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceGithubOrganizationSettingsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	settings := github.Organization{
		BillingEmail:                       github.String(d.Get("billing_email").(string)),
		Company:                            github.String(d.Get("company").(string)),
		Email:                              github.String(d.Get("email").(string)),
		TwitterUsername:                    github.String(d.Get("twitter_username").(string)),
		Location:                           github.String(d.Get("location").(string)),
		Name:                               github.String(d.Get("name").(string)),
		Description:                        github.String(d.Get("description").(string)),
		HasOrganizationProjects:            github.Bool(d.Get("has_organization_projects").(bool)),
		HasRepositoryProjects:              github.Bool(d.Get("has_repository_projects").(bool)),
		DefaultRepoPermission:              github.String(d.Get("default_repository_permission").(string)),
		MembersCanCreateRepos:              github.Bool(d.Get("members_can_create_repositories").(bool)),
		MembersCanCreatePrivateRepos:       github.Bool(d.Get("members_can_create_private_repositories").(bool)),
		MembersCanCreatePublicRepos:        github.Bool(d.Get("members_can_create_public_repositories").(bool)),
		MembersCanCreatePages:              github.Bool(d.Get("members_can_create_pages").(bool)),
		MembersCanCreatePublicPages:        github.Bool(d.Get("members_can_create_public_pages").(bool)),
		MembersCanCreatePrivatePages:       github.Bool(d.Get("members_can_create_private_pages").(bool)),
		MembersCanForkPrivateRepos:         github.Bool(d.Get("members_can_fork_private_repositories").(bool)),
		WebCommitSignoffRequired:           github.Bool(d.Get("web_commit_signoff_required").(bool)),
		Blog:                               github.String(d.Get("blog").(string)),
		AdvancedSecurityEnabledForNewRepos: github.Bool(d.Get("advanced_security_enabled_for_new_repositories").(bool)),
		DependabotAlertsEnabledForNewRepos: github.Bool(d.Get("dependabot_alerts_enabled_for_new_repositories").(bool)),
		DependabotSecurityUpdatesEnabledForNewRepos:    github.Bool(d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool)),
		DependencyGraphEnabledForNewRepos:              github.Bool(d.Get("dependency_graph_enabled_for_new_repositories").(bool)),
		SecretScanningEnabledForNewRepos:               github.Bool(d.Get("secret_scanning_enabled_for_new_repositories").(bool)),
		SecretScanningPushProtectionEnabledForNewRepos: github.Bool(d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool)),
	}

	enterpriseSettings := github.Organization{
		BillingEmail:                       github.String(d.Get("billing_email").(string)),
		Company:                            github.String(d.Get("company").(string)),
		Email:                              github.String(d.Get("email").(string)),
		TwitterUsername:                    github.String(d.Get("twitter_username").(string)),
		Location:                           github.String(d.Get("location").(string)),
		Name:                               github.String(d.Get("name").(string)),
		Description:                        github.String(d.Get("description").(string)),
		HasOrganizationProjects:            github.Bool(d.Get("has_organization_projects").(bool)),
		HasRepositoryProjects:              github.Bool(d.Get("has_repository_projects").(bool)),
		DefaultRepoPermission:              github.String(d.Get("default_repository_permission").(string)),
		MembersCanCreateRepos:              github.Bool(d.Get("members_can_create_repositories").(bool)),
		MembersCanCreateInternalRepos:      github.Bool(d.Get("members_can_create_internal_repositories").(bool)),
		MembersCanCreatePrivateRepos:       github.Bool(d.Get("members_can_create_private_repositories").(bool)),
		MembersCanCreatePublicRepos:        github.Bool(d.Get("members_can_create_public_repositories").(bool)),
		MembersCanCreatePages:              github.Bool(d.Get("members_can_create_pages").(bool)),
		MembersCanCreatePublicPages:        github.Bool(d.Get("members_can_create_public_pages").(bool)),
		MembersCanCreatePrivatePages:       github.Bool(d.Get("members_can_create_private_pages").(bool)),
		MembersCanForkPrivateRepos:         github.Bool(d.Get("members_can_fork_private_repositories").(bool)),
		WebCommitSignoffRequired:           github.Bool(d.Get("web_commit_signoff_required").(bool)),
		Blog:                               github.String(d.Get("blog").(string)),
		AdvancedSecurityEnabledForNewRepos: github.Bool(d.Get("advanced_security_enabled_for_new_repositories").(bool)),
		DependabotAlertsEnabledForNewRepos: github.Bool(d.Get("dependabot_alerts_enabled_for_new_repositories").(bool)),
		DependabotSecurityUpdatesEnabledForNewRepos:    github.Bool(d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool)),
		DependencyGraphEnabledForNewRepos:              github.Bool(d.Get("dependency_graph_enabled_for_new_repositories").(bool)),
		SecretScanningEnabledForNewRepos:               github.Bool(d.Get("secret_scanning_enabled_for_new_repositories").(bool)),
		SecretScanningPushProtectionEnabledForNewRepos: github.Bool(d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool)),
	}

	enterpriseSettingsNoFork := github.Organization{
		BillingEmail:                       github.String(d.Get("billing_email").(string)),
		Company:                            github.String(d.Get("company").(string)),
		Email:                              github.String(d.Get("email").(string)),
		TwitterUsername:                    github.String(d.Get("twitter_username").(string)),
		Location:                           github.String(d.Get("location").(string)),
		Name:                               github.String(d.Get("name").(string)),
		Description:                        github.String(d.Get("description").(string)),
		HasOrganizationProjects:            github.Bool(d.Get("has_organization_projects").(bool)),
		HasRepositoryProjects:              github.Bool(d.Get("has_repository_projects").(bool)),
		DefaultRepoPermission:              github.String(d.Get("default_repository_permission").(string)),
		MembersCanCreateRepos:              github.Bool(d.Get("members_can_create_repositories").(bool)),
		MembersCanCreateInternalRepos:      github.Bool(d.Get("members_can_create_internal_repositories").(bool)),
		MembersCanCreatePrivateRepos:       github.Bool(d.Get("members_can_create_private_repositories").(bool)),
		MembersCanCreatePublicRepos:        github.Bool(d.Get("members_can_create_public_repositories").(bool)),
		MembersCanCreatePages:              github.Bool(d.Get("members_can_create_pages").(bool)),
		MembersCanCreatePublicPages:        github.Bool(d.Get("members_can_create_public_pages").(bool)),
		MembersCanCreatePrivatePages:       github.Bool(d.Get("members_can_create_private_pages").(bool)),
		WebCommitSignoffRequired:           github.Bool(d.Get("web_commit_signoff_required").(bool)),
		Blog:                               github.String(d.Get("blog").(string)),
		AdvancedSecurityEnabledForNewRepos: github.Bool(d.Get("advanced_security_enabled_for_new_repositories").(bool)),
		DependabotAlertsEnabledForNewRepos: github.Bool(d.Get("dependabot_alerts_enabled_for_new_repositories").(bool)),
		DependabotSecurityUpdatesEnabledForNewRepos:    github.Bool(d.Get("dependabot_security_updates_enabled_for_new_repositories").(bool)),
		DependencyGraphEnabledForNewRepos:              github.Bool(d.Get("dependency_graph_enabled_for_new_repositories").(bool)),
		SecretScanningEnabledForNewRepos:               github.Bool(d.Get("secret_scanning_enabled_for_new_repositories").(bool)),
		SecretScanningPushProtectionEnabledForNewRepos: github.Bool(d.Get("secret_scanning_push_protection_enabled_for_new_repositories").(bool)),
	}

	orgPlan, _, err := client.Organizations.Edit(ctx, org, nil)
	if err != nil {
		return err
	}

	if orgPlan.GetPlan().GetName() == "enterprise" {
		if _, ok := d.GetOk("members_can_fork_private_repositories"); !ok {
			orgSettings, _, err := client.Organizations.Edit(ctx, org, &enterpriseSettingsNoFork)
			if err != nil {
				return err
			}
			id := strconv.FormatInt(orgSettings.GetID(), 10)
			d.SetId(id)
		} else if _, ok := d.GetOk("members_can_fork_private_repositories"); ok {
			orgSettings, _, err := client.Organizations.Edit(ctx, org, &enterpriseSettings)
			if err != nil {
				return err
			}
			id := strconv.FormatInt(orgSettings.GetID(), 10)
			d.SetId(id)
		}
	} else {
		orgSettings, _, err := client.Organizations.Edit(ctx, org, &settings)
		if err != nil {
			return err
		}
		id := strconv.FormatInt(orgSettings.GetID(), 10)
		d.SetId(id)
	}

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

	d.Set("billing_email", orgSettings.GetBillingEmail())
	d.Set("company", orgSettings.GetCompany())
	d.Set("email", orgSettings.GetEmail())
	d.Set("twitter_username", orgSettings.GetTwitterUsername())
	d.Set("location", orgSettings.GetLocation())
	d.Set("name", orgSettings.GetName())
	d.Set("description", orgSettings.GetDescription())
	d.Set("has_organization_projects", orgSettings.GetHasOrganizationProjects())
	d.Set("has_repository_projects", orgSettings.GetHasRepositoryProjects())
	d.Set("default_repository_permission", orgSettings.GetDefaultRepoPermission())
	d.Set("members_can_create_repositories", orgSettings.GetMembersCanCreateRepos())
	d.Set("members_can_create_internal_repositories", orgSettings.GetMembersCanCreateInternalRepos())
	d.Set("members_can_create_private_repositories", orgSettings.GetMembersCanCreatePrivateRepos())
	d.Set("members_can_create_public_repositories", orgSettings.GetMembersCanCreatePublicRepos())
	d.Set("members_can_create_pages", orgSettings.GetMembersCanCreatePages())
	d.Set("members_can_create_public_pages", orgSettings.GetMembersCanCreatePublicPages())
	d.Set("members_can_create_private_pages", orgSettings.GetMembersCanCreatePrivatePages())
	d.Set("members_can_fork_private_repositories", orgSettings.GetMembersCanForkPrivateRepos())
	d.Set("web_commit_signoff_required", orgSettings.GetWebCommitSignoffRequired())
	d.Set("blog", orgSettings.GetBlog())
	d.Set("advanced_security_enabled_for_new_repositories", orgSettings.GetAdvancedSecurityEnabledForNewRepos())
	d.Set("dependabot_alerts_enabled_for_new_repositories", orgSettings.GetDependabotAlertsEnabledForNewRepos())
	d.Set("dependabot_security_updates_enabled_for_new_repositories", orgSettings.GetDependabotSecurityUpdatesEnabledForNewRepos())
	d.Set("dependency_graph_enabled_for_new_repositories", orgSettings.GetDependencyGraphEnabledForNewRepos())
	d.Set("secret_scanning_enabled_for_new_repositories", orgSettings.GetSecretScanningEnabledForNewRepos())
	d.Set("secret_scanning_push_protection_enabled_for_new_repositories", orgSettings.GetSecretScanningPushProtectionEnabledForNewRepos())

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

	// This will set org settings to default values
	settings := github.Organization{
		BillingEmail:                       github.String("email@example.com"),
		Company:                            github.String(""),
		Email:                              github.String(""),
		TwitterUsername:                    github.String(""),
		Location:                           github.String(""),
		Name:                               github.String(""),
		Description:                        github.String(""),
		HasOrganizationProjects:            github.Bool(true),
		HasRepositoryProjects:              github.Bool(true),
		DefaultRepoPermission:              github.String("read"),
		MembersCanCreateRepos:              github.Bool(true),
		MembersCanCreatePrivateRepos:       github.Bool(true),
		MembersCanCreatePublicRepos:        github.Bool(true),
		MembersCanCreatePages:              github.Bool(false),
		MembersCanCreatePublicPages:        github.Bool(true),
		MembersCanCreatePrivatePages:       github.Bool(true),
		MembersCanForkPrivateRepos:         github.Bool(false),
		WebCommitSignoffRequired:           github.Bool(false),
		Blog:                               github.String(""),
		AdvancedSecurityEnabledForNewRepos: github.Bool(false),
		DependabotAlertsEnabledForNewRepos: github.Bool(false),
		DependabotSecurityUpdatesEnabledForNewRepos:    github.Bool(false),
		DependencyGraphEnabledForNewRepos:              github.Bool(false),
		SecretScanningEnabledForNewRepos:               github.Bool(false),
		SecretScanningPushProtectionEnabledForNewRepos: github.Bool(false),
	}

	enterpriseSettings := github.Organization{
		BillingEmail:                       github.String("email@example.com"),
		Company:                            github.String(""),
		Email:                              github.String(""),
		TwitterUsername:                    github.String(""),
		Location:                           github.String(""),
		Name:                               github.String(""),
		Description:                        github.String(""),
		HasOrganizationProjects:            github.Bool(true),
		HasRepositoryProjects:              github.Bool(true),
		DefaultRepoPermission:              github.String("read"),
		MembersCanCreateRepos:              github.Bool(true),
		MembersCanCreatePrivateRepos:       github.Bool(true),
		MembersCanCreateInternalRepos:      github.Bool(true),
		MembersCanCreatePublicRepos:        github.Bool(true),
		MembersCanCreatePages:              github.Bool(false),
		MembersCanCreatePublicPages:        github.Bool(true),
		MembersCanCreatePrivatePages:       github.Bool(true),
		MembersCanForkPrivateRepos:         github.Bool(false),
		WebCommitSignoffRequired:           github.Bool(false),
		Blog:                               github.String(""),
		AdvancedSecurityEnabledForNewRepos: github.Bool(false),
		DependabotAlertsEnabledForNewRepos: github.Bool(false),
		DependabotSecurityUpdatesEnabledForNewRepos:    github.Bool(false),
		DependencyGraphEnabledForNewRepos:              github.Bool(false),
		SecretScanningEnabledForNewRepos:               github.Bool(false),
		SecretScanningPushProtectionEnabledForNewRepos: github.Bool(false),
	}

	enterpriseSettingsNoFork := github.Organization{
		BillingEmail:                       github.String("email@example.com"),
		Company:                            github.String(""),
		Email:                              github.String(""),
		TwitterUsername:                    github.String(""),
		Location:                           github.String(""),
		Name:                               github.String(""),
		Description:                        github.String(""),
		HasOrganizationProjects:            github.Bool(true),
		HasRepositoryProjects:              github.Bool(true),
		DefaultRepoPermission:              github.String("read"),
		MembersCanCreateRepos:              github.Bool(true),
		MembersCanCreatePrivateRepos:       github.Bool(true),
		MembersCanCreateInternalRepos:      github.Bool(true),
		MembersCanCreatePublicRepos:        github.Bool(true),
		MembersCanCreatePages:              github.Bool(false),
		MembersCanCreatePublicPages:        github.Bool(true),
		MembersCanCreatePrivatePages:       github.Bool(true),
		WebCommitSignoffRequired:           github.Bool(false),
		Blog:                               github.String(""),
		AdvancedSecurityEnabledForNewRepos: github.Bool(false),
		DependabotAlertsEnabledForNewRepos: github.Bool(false),
		DependabotSecurityUpdatesEnabledForNewRepos:    github.Bool(false),
		DependencyGraphEnabledForNewRepos:              github.Bool(false),
		SecretScanningEnabledForNewRepos:               github.Bool(false),
		SecretScanningPushProtectionEnabledForNewRepos: github.Bool(false),
	}

	log.Printf("[DEBUG] Reverting Organization Settings to default values: %s", org)
	orgPlan, _, err := client.Organizations.Edit(ctx, org, nil)
	if err != nil {
		return err
	}
	if orgPlan.GetPlan().GetName() == "enterprise" {
		if _, ok := d.GetOk("members_can_fork_private_repositories"); !ok {
			_, _, err := client.Organizations.Edit(ctx, org, &enterpriseSettingsNoFork)
			if err != nil {
				return err
			}
		} else if _, ok := d.GetOk("members_can_fork_private_repositories"); ok {
			_, _, err := client.Organizations.Edit(ctx, org, &enterpriseSettings)
			if err != nil {
				return err
			}
		}
	} else {
		_, _, err := client.Organizations.Edit(ctx, org, &settings)
		if err != nil {
			return err
		}
	}

	return nil
}
