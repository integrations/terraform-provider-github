package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceGithubActionsRemoveToken is a data source that creates a token that can be
// used to remove a self-hosted runner from a repository.
// https://docs.github.com/en/enterprise-cloud@latest/rest/actions/self-hosted-runners?apiVersion=2022-11-28#create-a-remove-token-for-a-repository
func dataSourceGithubActionsRemoveToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsRemoveTokenRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Repository to remove the self-hosted runner from.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Token used to remove a self-hosted runner from a repository.",
			},
			"expires_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The token expiration date.",
			},
		},
	}
}

func dataSourceGithubActionsRemoveTokenRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	log.Printf("[DEBUG] Creating a GitHub Actions repository remove token for %s/%s", owner, repoName)
	token, _, err := client.Actions.CreateRemoveToken(context.TODO(), owner, repoName)
	if err != nil {
		return fmt.Errorf("error creating a GitHub Actions repository remove token for %s/%s: %w", owner, repoName, err)
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
