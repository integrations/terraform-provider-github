package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceGithubActionsOrganizationRemoveToken is a data source that creates a token that can be
// used to remove a self-hosted runner from an organization.
// https://docs.github.com/en/enterprise-cloud@latest/rest/actions/self-hosted-runners?apiVersion=2022-11-28#create-a-remove-token-for-an-organization
func dataSourceGithubActionsOrganizationRemoveToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsOrganizationRemoveTokenRead,

		Schema: map[string]*schema.Schema{
			"token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"expires_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubActionsOrganizationRemoveTokenRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	log.Printf("[DEBUG] Creating a GitHub Actions organization remove token for %s", owner)
	token, _, err := client.Actions.CreateOrganizationRemoveToken(context.TODO(), owner)
	if err != nil {
		return fmt.Errorf("error creating a GitHub Actions organization remove token for %s: %w", owner, err)
	}

	d.SetId(owner)
	err = d.Set("token", token.Token)
	if err != nil {
		return err
	}
	err = d.Set("expires_at", token.ExpiresAt.Unix())
	if err != nil {
		return err
	}

	return nil
}
