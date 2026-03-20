package github

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubTeamExternalGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Data source to retrieve external groups for a specific GitHub team.",
		ReadContext: dataSourceGithubTeamExternalGroupsRead,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the GitHub team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"external_groups": {
				Description: "List of external groups connected to the team.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The internal ID of the group.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"group_name": {
							Description: "The display name of the group.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_at": {
							Description: "The time of the last update for this group.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubTeamExternalGroupsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	orgName := meta.name
	slug, ok := d.Get("slug").(string)
	if !ok {
		return diag.Errorf("expected type of %s to be string", d.Get("slug"))
	}

	externalGroups, _, err := meta.v3client.Teams.ListExternalGroupsForTeamBySlug(ctx, orgName, slug)
	if err != nil {
		// Teams with explicit members cannot be externally managed, so they have no external groups.
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil &&
			ghErr.Response.StatusCode == http.StatusBadRequest &&
			strings.Contains(ghErr.Message, "has explicit members") {
			externalGroups = &github.ExternalGroupList{}
		} else {
			return diag.FromErr(err)
		}
	}

	groups := make([]map[string]any, 0)
	for _, group := range externalGroups.Groups {
		g := map[string]any{
			"group_id":   group.GetGroupID(),
			"group_name": group.GetGroupName(),
			"updated_at": group.GetUpdatedAt().Format(time.RFC3339),
		}

		groups = append(groups, g)
	}

	id, err := buildID(orgName, slug)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("external_groups", groups); err != nil {
		return diag.Errorf("error setting external_groups: %v", err)
	}

	return nil
}
