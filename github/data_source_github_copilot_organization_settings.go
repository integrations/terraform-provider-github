package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubCopilotOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubCopilotOrganizationSettingsRead,

		Schema: map[string]*schema.Schema{
			"seat_management_setting": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How Copilot seats are assigned: assign_selected, all_members, or unconfigured.",
			},
			"public_code_suggestions": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether Copilot can suggest code matching public repositories: allow or block.",
			},
			"seat_breakdown": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Breakdown of Copilot seat usage.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total number of Copilot seats.",
						},
						"added_this_cycle": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Seats added in the current billing cycle.",
						},
						"pending_invitation": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Seats pending invitation acceptance.",
						},
						"pending_cancellation": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Seats pending cancellation.",
						},
						"active_this_cycle": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Seats active in the current billing cycle.",
						},
						"inactive_this_cycle": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Seats inactive in the current billing cycle.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubCopilotOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	billing, _, err := client.Copilot.GetCopilotBilling(ctx, org)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return diag.Errorf("Copilot is not enabled for organization %q", org)
		}
		return diag.FromErr(err)
	}

	d.SetId(org)

	if err := d.Set("seat_management_setting", billing.GetSeatManagementSetting()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("public_code_suggestions", billing.GetPublicCodeSuggestions()); err != nil {
		return diag.FromErr(err)
	}
	if breakdown := billing.GetSeatBreakdown(); breakdown != nil {
		if err := d.Set("seat_breakdown", []any{map[string]any{
			"total":                breakdown.GetTotal(),
			"added_this_cycle":     breakdown.GetAddedThisCycle(),
			"pending_invitation":   breakdown.GetPendingInvitation(),
			"pending_cancellation": breakdown.GetPendingCancellation(),
			"active_this_cycle":    breakdown.GetActiveThisCycle(),
			"inactive_this_cycle":  breakdown.GetInactiveThisCycle(),
		}}); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
