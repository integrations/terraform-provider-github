package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusion() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationOIDCCustomPropertyInclusionCreate,
		Read:   resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead,
		Delete: resourceGithubActionsOrganizationOIDCCustomPropertyInclusionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"custom_property_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the custom property to include in the OIDC token.",
			},
		},
	}
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusionCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	customPropertyName := d.Get("custom_property_name").(string)

	body := map[string]string{
		"custom_property_name": customPropertyName,
	}

	req, err := client.NewRequest("POST", fmt.Sprintf("orgs/%s/actions/oidc/customization/properties/repo", orgName), body)
	if err != nil {
		return fmt.Errorf("error creating request to add OIDC custom property inclusion: %w", err)
	}

	_, err = client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error adding OIDC custom property inclusion %q for organization %q: %w", customPropertyName, orgName, err)
	}

	d.SetId(buildTwoPartID(orgName, customPropertyName))

	return resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead(d, meta)
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	orgName, customPropertyName, err := parseTwoPartID(d.Id(), "organization", "custom_property_name")
	if err != nil {
		return err
	}

	err = checkOrganization(meta)
	if err != nil {
		return err
	}

	inclusions, err := listOrgOIDCCustomPropertyInclusions(ctx, client, orgName)
	if err != nil {
		return fmt.Errorf("error reading OIDC custom property inclusions for organization %q: %w", orgName, err)
	}

	found := false
	for _, inclusion := range inclusions {
		if inclusion.PropertyName == customPropertyName {
			found = true
			break
		}
	}

	if !found {
		d.SetId("")
		return nil
	}

	if err := d.Set("custom_property_name", customPropertyName); err != nil {
		return err
	}

	return nil
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusionDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	orgName, customPropertyName, err := parseTwoPartID(d.Id(), "organization", "custom_property_name")
	if err != nil {
		return err
	}

	err = checkOrganization(meta)
	if err != nil {
		return err
	}

	req, err := client.NewRequest("DELETE", fmt.Sprintf("orgs/%s/actions/oidc/customization/properties/repo/%s", orgName, customPropertyName), nil)
	if err != nil {
		return fmt.Errorf("error creating request to delete OIDC custom property inclusion: %w", err)
	}

	_, err = client.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error deleting OIDC custom property inclusion %q for organization %q: %w", customPropertyName, orgName, err)
	}

	return nil
}

// OIDCCustomPropertyInclusion represents a custom property included in OIDC tokens.
type OIDCCustomPropertyInclusion struct {
	PropertyName string `json:"property_name"`
}

// listOrgOIDCCustomPropertyInclusions lists all custom properties included in OIDC tokens for an organization.
func listOrgOIDCCustomPropertyInclusions(ctx context.Context, client *github.Client, orgName string) ([]*OIDCCustomPropertyInclusion, error) {
	req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/oidc/customization/properties/repo", orgName), nil)
	if err != nil {
		return nil, err
	}

	var inclusions []*OIDCCustomPropertyInclusion
	_, err = client.Do(ctx, req, &inclusions)
	if err != nil {
		return nil, err
	}

	return inclusions, nil
}
