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

	_, _, err := client.Copilot.GetSeatDetails(ctx, org, username)
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
