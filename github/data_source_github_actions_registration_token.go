package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubActionsRegistrationToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsRegistrationTokenRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

func dataSourceGithubActionsRegistrationTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	log.Printf("[DEBUG] Creating a GitHub Actions repository registration token for %s/%s", owner, repoName)
	token, _, err := client.Actions.CreateRegistrationToken(context.TODO(), owner, repoName)
	if err != nil {
		return fmt.Errorf("error creating a GitHub Actions repository registration token for %s/%s: %s", owner, repoName, err)
	}

	d.SetId(fmt.Sprintf("%s/%s", owner, repoName))
	d.Set("token", token.Token)
	d.Set("expires_at", token.ExpiresAt.Unix())

	return nil
}
