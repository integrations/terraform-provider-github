package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateCreateOrUpdate,
		Read:   resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead,
		Update: resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateCreateOrUpdate,
		Delete: resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"include_claim_keys": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "A list of OpenID Connect claims.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateCreateOrUpdate(d *schema.ResourceData, m any) error {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name
	ctx := context.Background()

	err := checkOrganization(m)
	if err != nil {
		return err
	}

	includeClaimKeys, _ := d.Get("include_claim_keys").([]any)

	claimsStr := make([]string, len(includeClaimKeys))

	for i, v := range includeClaimKeys {
		s, _ := v.(string)
		claimsStr[i] = s
	}

	_, err = client.Actions.SetOrgOIDCSubjectClaimCustomTemplate(ctx, orgName, github.OIDCSubjectClaimCustomTemplate{
		IncludeClaimKeys: claimsStr,
	})
	if err != nil {
		return err
	}

	d.SetId(orgName)
	return resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead(d, m)
}

func resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead(d *schema.ResourceData, m any) error {
	meta, _ := m.(*Owner)
	client := meta.v3client
	ctx := context.Background()
	orgName := meta.name

	err := checkOrganization(m)
	if err != nil {
		return err
	}

	template, _, err := client.Actions.GetOrgOIDCSubjectClaimCustomTemplate(ctx, orgName)
	if err != nil {
		return err
	}

	if err = d.Set("include_claim_keys", template.IncludeClaimKeys); err != nil {
		return err
	}

	return nil
}

func resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateDelete(d *schema.ResourceData, m any) error {
	// Sets include_claim_keys back to GitHub's defaults
	// https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#resetting-your-customizations
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name
	ctx := context.Background()

	err := checkOrganization(m)
	if err != nil {
		return err
	}

	_, err = client.Actions.SetOrgOIDCSubjectClaimCustomTemplate(ctx, orgName, github.OIDCSubjectClaimCustomTemplate{
		IncludeClaimKeys: []string{"repo", "context"},
	})
	if err != nil {
		return err
	}

	return nil
}
