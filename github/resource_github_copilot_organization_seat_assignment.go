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

func resourceGithubCopilotOrganizationSeatAssignment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubCopilotOrganizationSeatAssignmentCreate,
		ReadContext:   resourceGithubCopilotOrganizationSeatAssignmentRead,
		DeleteContext: resourceGithubCopilotOrganizationSeatAssignmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
				Description:      "The login of the user to assign a Copilot seat.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the seat was first assigned.",
			},
			"pending_cancellation_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date on which the seat will be cancelled, if pending cancellation.",
			},
			"last_activity_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the user's last Copilot activity.",
			},
			"last_activity_editor": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Editor used in the user's last Copilot activity.",
			},
			"plan_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Copilot plan type for this seat (e.g. business, enterprise).",
			},
		},
	}
}

func resourceGithubCopilotOrganizationSeatAssignmentCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	username := d.Get("username").(string)

	_, _, err := client.Copilot.AddCopilotUsers(ctx, org, []string{username})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(username)

	return resourceGithubCopilotOrganizationSeatAssignmentRead(ctx, d, m)
}

func resourceGithubCopilotOrganizationSeatAssignmentRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	username := d.Id()

	seat, _, err := client.Copilot.GetSeatDetails(ctx, org, username)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Copilot seat assignment no longer exists, removing from state", map[string]any{"username": username})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}

	createdAt := seat.GetCreatedAt()
	if err := d.Set("created_at", createdAt.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("pending_cancellation_date", seat.GetPendingCancellationDate()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_activity_editor", seat.GetLastActivityEditor()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("plan_type", seat.GetPlanType()); err != nil {
		return diag.FromErr(err)
	}

	if lastActivity := seat.GetLastActivityAt(); !lastActivity.IsZero() {
		if err := d.Set("last_activity_at", lastActivity.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubCopilotOrganizationSeatAssignmentDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	username := d.Id()

	_, _, err := client.Copilot.RemoveCopilotUsers(ctx, org, []string{username})
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Copilot seat assignment no longer exists, skipping delete", map[string]any{"username": username})
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
