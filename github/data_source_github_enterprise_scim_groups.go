package github

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseSCIMGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup SCIM groups provisioned for a GitHub enterprise.",
		ReadContext: dataSourceGithubEnterpriseSCIMGroupsRead,

		Schema: map[string]*schema.Schema{
			"enterprise": {
				Description: "The enterprise slug.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"filter": {
				Description: "Optional SCIM filter. See GitHub SCIM enterprise docs for supported filters.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"excluded_attributes": {
				Description: "Optional SCIM excludedAttributes query parameter.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results_per_page": {
				Description:  "Number of results per request (mapped to SCIM 'count'). Used while auto-fetching all pages.",
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      100,
				ValidateFunc: validation.IntBetween(1, 100),
			},

			"schemas": {
				Description: "SCIM response schemas.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"total_results": {
				Description: "The total number of results returned by the SCIM endpoint.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"start_index": {
				Description: "The startIndex from the first SCIM page.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"items_per_page": {
				Description: "The itemsPerPage from the first SCIM page.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"resources": {
				Description: "All SCIM groups.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: enterpriseSCIMGroupSchema(),
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseSCIMGroupsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise := d.Get("enterprise").(string)
	filter := d.Get("filter").(string)
	excluded := d.Get("excluded_attributes").(string)
	count := d.Get("results_per_page").(int)

	groups, first, err := enterpriseSCIMListAllGroups(ctx, client, enterprise, filter, excluded, count)
	if err != nil {
		return diag.FromErr(err)
	}

	flat := make([]any, 0, len(groups))
	for _, g := range groups {
		flat = append(flat, flattenEnterpriseSCIMGroup(g))
	}

	id := fmt.Sprintf("%s/scim-groups", enterprise)
	if filter != "" {
		id = fmt.Sprintf("%s?filter=%s", id, url.QueryEscape(filter))
	}
	if excluded != "" {
		if filter == "" {
			id = fmt.Sprintf("%s?excluded_attributes=%s", id, url.QueryEscape(excluded))
		} else {
			id = fmt.Sprintf("%s&excluded_attributes=%s", id, url.QueryEscape(excluded))
		}
	}

	d.SetId(id)

	_ = d.Set("schemas", first.Schemas)
	_ = d.Set("total_results", first.TotalResults)
	if first.StartIndex > 0 {
		_ = d.Set("start_index", first.StartIndex)
	} else {
		_ = d.Set("start_index", 1)
	}
	if first.ItemsPerPage > 0 {
		_ = d.Set("items_per_page", first.ItemsPerPage)
	} else {
		_ = d.Set("items_per_page", count)
	}
	if err := d.Set("resources", flat); err != nil {
		return diag.FromErr(fmt.Errorf("error setting resources: %w", err))
	}

	return nil
}

func enterpriseSCIMMetaSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SCIM resource type.",
		},
		"created": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The creation timestamp.",
		},
		"last_modified": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The lastModified timestamp.",
		},
		"location": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The resource location.",
		},
		"version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The resource version.",
		},
		"etag": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The resource eTag.",
		},
		"password_changed_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The password changed at timestamp (if present).",
		},
	}
}

func enterpriseSCIMGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"schemas": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "SCIM schemas for this group.",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SCIM group ID.",
		},
		"external_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The external ID for the group.",
		},
		"display_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SCIM group displayName.",
		},
		"members": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Group members.",
			Elem:        &schema.Resource{Schema: enterpriseSCIMGroupMemberSchema()},
		},
		"meta": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Resource metadata.",
			Elem:        &schema.Resource{Schema: enterpriseSCIMMetaSchema()},
		},
	}
}

func enterpriseSCIMGroupMemberSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Member identifier.",
		},
		"ref": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Member reference URL.",
		},
		"display_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Member display name.",
		},
	}
}
