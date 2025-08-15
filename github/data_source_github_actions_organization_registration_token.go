package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationRegistrationToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsOrganizationRegistrationTokenRead,

		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubActionsOrganizationRegistrationTokenRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	log.Printf("[DEBUG] Creating a GitHub Actions organization registration token for %s", owner)
	token, _, err := client.Actions.CreateOrganizationRegistrationToken(context.TODO(), owner)
	if err != nil {
		return fmt.Errorf("error creating a GitHub Actions organization registration token for %s: %s", owner, err)
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
