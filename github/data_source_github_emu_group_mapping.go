package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v44/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubEMUGroupMapping() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEMUGroupMappingRead,
		Schema: map[string]*schema.Schema{
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"team_slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEMUGroupMappingRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	opts := &github.ListExternalGroupsOptions{
		// could provide a DisplayName here
	}

	groups, resp, err := client.Teams.ListExternalGroups(ctx, orgName, opts)
	if err != nil {
		return err
	}

	fmt.Printf("response: %v", resp)

	// need to flatten/format first
	d.Set("groups", groups)
	return nil
}
