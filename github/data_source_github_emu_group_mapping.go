package github

import (
	"context"
	"encoding/json"

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
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
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
	opts := &github.ListExternalGroupsOptions{}

	groups, _, err := client.Teams.ListExternalGroups(ctx, orgName, opts)
	if err != nil {
		return err
	}

	// convert to JSON in order to martial to format we can return
	jsonGroups, err := json.Marshal(groups.Groups)
	if err != nil {
		return err
	}

	ourGroups := make([]map[string]interface{}, 0)
	err = json.Unmarshal(jsonGroups, &ourGroups)
	if err != nil {
		return err
	}

	if err := d.Set("groups", ourGroups); err != nil {
		return err
	}

	// TODO: set unique identifier based on hash of data here?
	d.SetId("xxx")

	return nil
}
