package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsRegistrationToken() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve a registration token for a GitHub Actions self-hosted runner in a repository.",
		Read:        dataSourceGithubActionsRegistrationTokenRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The token that has been retrieved.",
			},
			"expires_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The token expiration date as a Unix timestamp.",
			},
		},
	}
}

func dataSourceGithubActionsRegistrationTokenRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	log.Printf("[DEBUG] Creating a GitHub Actions repository registration token for %s/%s", owner, repoName)
	token, _, err := client.Actions.CreateRegistrationToken(ctx, owner, repoName)
	if err != nil {
		return fmt.Errorf("error creating a GitHub Actions repository registration token for %s/%s: %w", owner, repoName, err)
	}

	d.SetId(fmt.Sprintf("%s/%s", owner, repoName))
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
