package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubCopilotTeamSeatAssignment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubCopilotTeamSeatAssignmentCreate,
		ReadContext:   resourceGithubCopilotTeamSeatAssignmentRead,
		DeleteContext: resourceGithubCopilotTeamSeatAssignmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"team": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
				Description:      "The slug of the team to assign Copilot seats.",
			},
		},
	}
}

func resourceGithubCopilotTeamSeatAssignmentCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	team := d.Get("team").(string)

	_, _, err := client.Copilot.AddCopilotTeams(ctx, org, []string{team})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(team)

	return resourceGithubCopilotTeamSeatAssignmentRead(ctx, d, m)
}

func resourceGithubCopilotTeamSeatAssignmentRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	teamSlug := d.Id()

	// The Copilot API has no single-team seat lookup; scan all seats for a team assignee matching our slug.
	opts := &github.ListOptions{PerPage: 100}
	for {
		resp_data, resp, err := client.Copilot.ListCopilotSeats(ctx, org, opts)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Copilot team seat assignment no longer exists, removing from state", map[string]any{"team": teamSlug})
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}

		for _, seat := range resp_data.Seats {
			if t, ok := seat.GetTeam(); ok && t.GetSlug() == teamSlug {
				if err := d.Set("team", teamSlug); err != nil {
					return diag.FromErr(err)
				}
				return nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	// Team not found in any seat — it was removed outside Terraform.
	tflog.Info(ctx, "Copilot team seat assignment no longer exists, removing from state", map[string]any{"team": teamSlug})
	d.SetId("")
	return nil
}

func resourceGithubCopilotTeamSeatAssignmentDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	teamSlug := d.Id()

	_, _, err := client.Copilot.RemoveCopilotTeams(ctx, org, []string{teamSlug})
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Copilot team seat assignment no longer exists, skipping delete", map[string]any{"team": teamSlug})
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
