package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func resourceGithubEnterprisePrivateRepositoryForkingSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Manages the private repository forking policy for a GitHub Enterprise.",
		Create:      resourceGithubEnterprisePrivateRepositoryForkingSettingCreateOrUpdate,
		Read:        resourceGithubEnterprisePrivateRepositoryForkingSettingRead,
		Update:      resourceGithubEnterprisePrivateRepositoryForkingSettingCreateOrUpdate,
		Delete:      resourceGithubEnterprisePrivateRepositoryForkingSettingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: func(_ context.Context, diff *schema.ResourceDiff, _ any) error {
			settingValue := diff.Get("setting_value").(string)
			policyValue := diff.Get("policy_value").(string)

			if settingValue == "ENABLED" && policyValue == "" {
				return fmt.Errorf("policy_value is required when setting_value is ENABLED")
			}
			if settingValue != "ENABLED" && policyValue != "" {
				return fmt.Errorf("policy_value must not be set when setting_value is %s", settingValue)
			}
			return nil
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"setting_value": {
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"ENABLED",
					"DISABLED",
					"NO_POLICY",
				}, false)),
				Description: "Whether private repository forking is enabled for the enterprise. Must be one of: ENABLED, DISABLED, NO_POLICY.",
			},
			"policy_value": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"ENTERPRISE_ORGANIZATIONS",
					"SAME_ORGANIZATION",
					"SAME_ORGANIZATION_USER_ACCOUNTS",
					"ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS",
					"USER_ACCOUNTS",
					"EVERYWHERE",
				}, false)),
				Description: "Where members can fork private repositories. Required when setting_value is ENABLED. Must be one of: ENTERPRISE_ORGANIZATIONS, SAME_ORGANIZATION, SAME_ORGANIZATION_USER_ACCOUNTS, ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS, USER_ACCOUNTS, EVERYWHERE.",
			},
		},
	}
}

func resourceGithubEnterprisePrivateRepositoryForkingSettingCreateOrUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)

	enterpriseID, err := getEnterpriseID(ctx, client, enterpriseSlug)
	if err != nil {
		return fmt.Errorf("error resolving enterprise ID for slug %q: %w", enterpriseSlug, err)
	}

	settingValue := githubv4.EnterpriseEnabledDisabledSettingValue(d.Get("setting_value").(string))

	input := githubv4.UpdateEnterpriseAllowPrivateRepositoryForkingSettingInput{
		EnterpriseID: enterpriseID,
		SettingValue: settingValue,
	}

	if v, ok := d.GetOk("policy_value"); ok {
		pv := githubv4.EnterpriseAllowPrivateRepositoryForkingPolicyValue(v.(string))
		input.PolicyValue = &pv
	}

	var mutate struct {
		UpdateEnterpriseAllowPrivateRepositoryForkingSetting struct {
			Enterprise struct {
				ID githubv4.ID
			}
			Message githubv4.String
		} `graphql:"updateEnterpriseAllowPrivateRepositoryForkingSetting(input: $input)"`
	}

	log.Printf("[DEBUG] Updating private repository forking setting for enterprise: %s", enterpriseSlug)
	err = client.Mutate(ctx, &mutate, input, nil)
	if err != nil {
		return fmt.Errorf("error updating private repository forking setting for enterprise %q: %w", enterpriseSlug, err)
	}

	d.SetId(enterpriseSlug)

	return resourceGithubEnterprisePrivateRepositoryForkingSettingRead(d, meta)
}

func resourceGithubEnterprisePrivateRepositoryForkingSettingRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	enterpriseSlug := d.Id()

	var query struct {
		Enterprise struct {
			OwnerInfo struct {
				AllowPrivateRepositoryForkingSetting            githubv4.EnterpriseEnabledDisabledSettingValue
				AllowPrivateRepositoryForkingSettingPolicyValue githubv4.EnterpriseAllowPrivateRepositoryForkingPolicyValue
			}
		} `graphql:"enterprise(slug: $slug)"`
	}

	variables := map[string]any{
		"slug": githubv4.String(enterpriseSlug),
	}

	log.Printf("[DEBUG] Reading private repository forking setting for enterprise: %s", enterpriseSlug)
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return fmt.Errorf("error reading private repository forking setting for enterprise %q: %w", enterpriseSlug, err)
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}

	settingValue := string(query.Enterprise.OwnerInfo.AllowPrivateRepositoryForkingSetting)
	if err := d.Set("setting_value", settingValue); err != nil {
		return err
	}

	if settingValue == "ENABLED" {
		if err := d.Set("policy_value", string(query.Enterprise.OwnerInfo.AllowPrivateRepositoryForkingSettingPolicyValue)); err != nil {
			return err
		}
	} else {
		if err := d.Set("policy_value", ""); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubEnterprisePrivateRepositoryForkingSettingDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	enterpriseSlug := d.Id()

	enterpriseID, err := getEnterpriseID(ctx, client, enterpriseSlug)
	if err != nil {
		return fmt.Errorf("error resolving enterprise ID for slug %q: %w", enterpriseSlug, err)
	}

	input := githubv4.UpdateEnterpriseAllowPrivateRepositoryForkingSettingInput{
		EnterpriseID: enterpriseID,
		SettingValue: githubv4.EnterpriseEnabledDisabledSettingValueNoPolicy,
	}

	var mutate struct {
		UpdateEnterpriseAllowPrivateRepositoryForkingSetting struct {
			Enterprise struct {
				ID githubv4.ID
			}
		} `graphql:"updateEnterpriseAllowPrivateRepositoryForkingSetting(input: $input)"`
	}

	log.Printf("[DEBUG] Resetting private repository forking setting to NO_POLICY for enterprise: %s", enterpriseSlug)
	err = client.Mutate(ctx, &mutate, input, nil)
	if err != nil {
		return fmt.Errorf("error resetting private repository forking setting for enterprise %q: %w", enterpriseSlug, err)
	}

	return nil
}
