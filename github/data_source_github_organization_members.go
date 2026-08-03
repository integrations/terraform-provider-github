package github

import (
	"context"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationMembers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationMembersRead,

		Description: "Data source to list all organization members.",

		Schema: map[string]*schema.Schema{
			"members": {
				Description: "Organization members.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the member.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"node_id": {
							Description: "Node ID of the member.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"login": {
							Description: "Login of the member.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationMembersRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	members := make([]map[string]any, 0)
	for member, err := range meta.v3client.Organizations.ListMembersIter(ctx, meta.name, &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: meta.maxPerPage}}) {
		if err != nil {
			return diag.FromErr(err)
		}

		m := map[string]any{
			"id":      member.GetID(),
			"node_id": member.GetNodeID(),
			"login":   member.GetLogin(),
		}
		members = append(members, m)
	}

	d.SetId(meta.name)

	if err := d.Set("members", members); err != nil {
		return diag.Errorf("error setting members: %v", err)
	}

	return nil
}
