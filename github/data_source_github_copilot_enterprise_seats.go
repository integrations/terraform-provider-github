package github

import (
	"cmp"
	"context"
	"slices"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubCopilotEnterpriseSeats() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubCopilotEnterpriseSeatsRead,
		Description: "Data source to list the GitHub Copilot seats billed to an enterprise, across its organizations and enterprise teams.",
		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Description:      "The slug of the enterprise.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"total_seats": {
				Description: "Number of Copilot seats billed to the enterprise. A user who has access through several organizations or enterprise teams is counted once.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"seats": {
				Description: "Copilot seat assignments, sorted by assignee login. A user who has access through several organizations or enterprise teams has one entry for each of them.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assignee_login": {
							Description: "Login of the user the seat is assigned to.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"assigning_team_slug": {
							Description: "Slug of the team or enterprise team through which the seat is granted. Empty when the seat is assigned directly.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"plan_type": {
							Description: "Copilot plan of the seat, for example `business` or `enterprise`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"created_at": {
							Description: "Time the assignee was last granted access to Copilot, in RFC 3339 format.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_activity_at": {
							Description: "Time of the assignee's last Copilot activity, in RFC 3339 format. Empty when no activity has been recorded.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_activity_editor": {
							Description: "Editor used for the assignee's last Copilot activity. Empty when no activity has been recorded.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"pending_cancellation_date": {
							Description: "Date, in `YYYY-MM-DD` format, on which the seat will be cancelled. Empty when no cancellation is pending.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubCopilotEnterpriseSeatsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	enterpriseSlug, _ := d.Get("enterprise_slug").(string)

	tflog.Debug(ctx, "Reading GitHub Copilot enterprise seats", map[string]any{"enterprise_slug": enterpriseSlug})

	// total_seats is only returned on the response envelope; the *Iter method yields the seats alone.
	firstPage, _, err := client.Copilot.ListCopilotEnterpriseSeats(ctx, enterpriseSlug, &github.ListOptions{PerPage: 1})
	if err != nil {
		return diag.Errorf("error listing Copilot seats for enterprise %s: %v", enterpriseSlug, err)
	}

	seats := make([]*github.CopilotSeatDetails, 0)
	for seat, err := range client.Copilot.ListCopilotEnterpriseSeatsIter(ctx, enterpriseSlug, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			return diag.Errorf("error listing Copilot seats for enterprise %s: %v", enterpriseSlug, err)
		}
		seats = append(seats, seat)
	}

	// The API does not document an order: sort so that the list is stable between reads.
	slices.SortStableFunc(seats, func(a, b *github.CopilotSeatDetails) int {
		return cmp.Or(
			cmp.Compare(copilotSeatAssigneeLogin(a), copilotSeatAssigneeLogin(b)),
			cmp.Compare(a.GetAssigningTeam().GetSlug(), b.GetAssigningTeam().GetSlug()),
		)
	})

	d.SetId(enterpriseSlug)

	if err := d.Set("total_seats", firstPage.TotalSeats); err != nil {
		return diag.Errorf("error setting total_seats: %v", err)
	}
	if err := d.Set("seats", flattenCopilotSeats(seats)); err != nil {
		return diag.Errorf("error setting seats: %v", err)
	}

	return nil
}

func flattenCopilotSeats(seats []*github.CopilotSeatDetails) []map[string]any {
	result := make([]map[string]any, 0, len(seats))
	for _, seat := range seats {
		result = append(result, map[string]any{
			"assignee_login":            copilotSeatAssigneeLogin(seat),
			"assigning_team_slug":       seat.GetAssigningTeam().GetSlug(),
			"plan_type":                 seat.GetPlanType(),
			"created_at":                formatCopilotSeatTime(seat.CreatedAt),
			"last_activity_at":          formatCopilotSeatTime(seat.LastActivityAt),
			"last_activity_editor":      seat.GetLastActivityEditor(),
			"pending_cancellation_date": seat.GetPendingCancellationDate(),
		})
	}
	return result
}

// The REST API documents the assignee as a user; it can be null, for example for a deleted account.
func copilotSeatAssigneeLogin(seat *github.CopilotSeatDetails) string {
	if user, ok := seat.GetUser(); ok {
		return user.GetLogin()
	}
	return ""
}

func formatCopilotSeatTime(t *github.Timestamp) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
