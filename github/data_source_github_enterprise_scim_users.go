package github

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseSCIMUsers() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup SCIM users provisioned for a GitHub enterprise.",
		ReadContext: dataSourceGithubEnterpriseSCIMUsersRead,

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
				Description: "All SCIM users.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: enterpriseSCIMUserSchema(),
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseSCIMUsersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise := d.Get("enterprise").(string)
	filter := d.Get("filter").(string)
	excluded := d.Get("excluded_attributes").(string)
	count := d.Get("results_per_page").(int)

	users, first, err := enterpriseSCIMListAllUsers(ctx, client, enterprise, filter, excluded, count)
	if err != nil {
		return diag.FromErr(err)
	}

	flat := make([]any, 0, len(users))
	for _, u := range users {
		flat = append(flat, flattenEnterpriseSCIMUser(u))
	}

	id := fmt.Sprintf("%s/scim-users", enterprise)
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

func enterpriseSCIMUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"schemas": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "SCIM schemas for this user.",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SCIM user ID.",
		},
		"external_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The external ID for the user.",
		},
		"user_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SCIM userName.",
		},
		"display_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SCIM displayName.",
		},
		"active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the user is active.",
		},
		"name": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "User name object.",
			Elem:        &schema.Resource{Schema: enterpriseSCIMUserNameSchema()},
		},
		"emails": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "User emails.",
			Elem:        &schema.Resource{Schema: enterpriseSCIMUserEmailSchema()},
		},
		"roles": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "User roles.",
			Elem:        &schema.Resource{Schema: enterpriseSCIMUserRoleSchema()},
		},
		"meta": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Resource metadata.",
			Elem:        &schema.Resource{Schema: enterpriseSCIMMetaSchema()},
		},
	}
}

func enterpriseSCIMUserNameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"formatted": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Formatted name.",
		},
		"family_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Family name.",
		},
		"given_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Given name.",
		},
		"middle_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Middle name.",
		},
	}
}

func enterpriseSCIMUserEmailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Email address.",
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Email type.",
		},
		"primary": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this email is primary.",
		},
	}
}

func enterpriseSCIMUserRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Role value.",
		},
		"display": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Role display.",
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Role type.",
		},
		"primary": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this role is primary.",
		},
	}
}
