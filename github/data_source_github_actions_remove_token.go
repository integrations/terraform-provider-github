package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsRemoveToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsRemoveTokenRead,

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

func dataSourceGithubActionsRemoveTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	log.Printf("[DEBUG] Creating a GitHub Actions organization remove token for %s", owner)
	token, _, err := client.Actions.CreateOrganizationRemoveToken(context.TODO(), owner)
	if err != nil {
		return fmt.Errorf("error creating a GitHub Actions organization remove token for %s: %s", owner, err)
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
