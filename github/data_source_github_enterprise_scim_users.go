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
		Description: "Retrieves SCIM users provisioned for a GitHub enterprise.",
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
			"results_per_page": {
				Description:      "Number of results per request (mapped to SCIM 'count'). Used while auto-fetching all pages.",
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          100,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 100)),
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
	count := d.Get("results_per_page").(int)

	users, first, err := enterpriseSCIMListAllUsers(ctx, client, enterprise, filter, count)
	if err != nil {
		return diag.FromErr(err)
	}

	flat := make([]any, 0, len(users))
	for _, user := range users {
		flat = append(flat, flattenEnterpriseSCIMUser(user))
	}

	id := fmt.Sprintf("%s/scim-users", enterprise)
	if filter != "" {
		id = fmt.Sprintf("%s?filter=%s", id, url.QueryEscape(filter))
	}

	d.SetId(id)

	if err := d.Set("schemas", first.Schemas); err != nil {
		return diag.FromErr(err)
	}
	if first.TotalResults != nil {
		if err := d.Set("total_results", *first.TotalResults); err != nil {
			return diag.FromErr(err)
		}
	}
	startIndex := 1
	if first.StartIndex != nil && *first.StartIndex > 0 {
		startIndex = *first.StartIndex
	}
	if err := d.Set("start_index", startIndex); err != nil {
		return diag.FromErr(err)
	}
	itemsPerPage := count
	if first.ItemsPerPage != nil && *first.ItemsPerPage > 0 {
		itemsPerPage = *first.ItemsPerPage
	}
	if err := d.Set("items_per_page", itemsPerPage); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resources", flat); err != nil {
		return diag.FromErr(err)
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
