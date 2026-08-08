package github

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubTeamExternalGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve external groups for a specific GitHub team.",
		ReadContext: dataSourceGithubTeamExternalGroupsRead,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the GitHub team.",
			},
			"external_groups": {
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

func dataSourceGithubTeamExternalGroupsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("expected type of %s to be *Owner", meta)
	}
	err := checkOrganization(owner)
	if err != nil {
		return diag.FromErr(err)
	}
	client := owner.v3client
	orgName := owner.name
	slug, ok := d.Get("slug").(string)
	if !ok {
		return diag.Errorf("expected type of %s to be string", d.Get("slug"))
	}

	externalGroups, _, err := client.Teams.ListExternalGroupsForTeamBySlug(ctx, orgName, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	// convert to JSON in order to marshal to format we can return
	jsonGroups, err := json.Marshal(externalGroups.Groups)
	if err != nil {
		return diag.FromErr(err)
	}

	groupsState := make([]map[string]any, 0)
	err = json.Unmarshal(jsonGroups, &groupsState)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("external_groups", groupsState); err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(orgName, slug)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}
