package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationInvitation() *schema.Resource {
	return &schema.Resource{
		Description: "Invite a user to a GitHub organization by email address or GitHub user ID.",

		CreateContext: resourceGithubOrganizationInvitationCreate,
		ReadContext:   resourceGithubOrganizationInvitationRead,
		DeleteContext: resourceGithubOrganizationInvitationDelete,

		Schema: map[string]*schema.Schema{
			"email": {
				Description:  "The email address of the person to invite to the organization. Exactly one of `email` or `invitee_id` must be set.",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"email", "invitee_id"},
			},
			"invitee_id": {
				Description:  "The GitHub user ID of the person to invite. Exactly one of `email` or `invitee_id` must be set.",
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"email", "invitee_id"},
			},
			"role": {
				Description:      "The role for the new member. Must be one of `admin`, `direct_member`, or `billing_manager`. Defaults to `direct_member`.",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          "direct_member",
				ValidateDiagFunc: validateValueFunc([]string{"admin", "direct_member", "billing_manager"}),
			},
			"invitation_id": {
				Description: "The ID of the invitation that was created.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"login": {
				Description: "The GitHub username of the invited user (if available).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceGithubOrganizationInvitationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	opts := &github.CreateOrgInvitationOptions{
		Role: github.Ptr(d.Get("role").(string)),
	}

	if v, ok := d.GetOk("email"); ok {
		opts.Email = github.Ptr(v.(string))
	}

	if v, ok := d.GetOk("invitee_id"); ok {
		opts.InviteeID = github.Ptr(int64(v.(int)))
	}

	invitation, _, err := client.Organizations.CreateOrgInvitation(ctx, orgName, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	invitationID := invitation.GetID()
	d.SetId(strconv.FormatInt(invitationID, 10))

	if err = d.Set("invitation_id", int(invitationID)); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("login", invitation.GetLogin()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationInvitationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	invitationID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	// There is no single-item GET endpoint for org invitations,
	// so we paginate through the pending invitations list to find ours.
	opts := &github.ListOptions{
		PerPage: maxPerPage,
	}

	var invitation *github.Invitation
	for {
		invitations, resp, err := client.Organizations.ListPendingOrgInvitations(ctx, orgName, opts)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, inv := range invitations {
			if inv.GetID() == invitationID {
				invitation = inv
				break
			}
		}

		if invitation != nil || resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	if invitation == nil {
		// Invitation was accepted, cancelled, or expired — remove from state
		tflog.Info(ctx, fmt.Sprintf("Removing organization invitation %s from state because it is no longer pending", d.Id()), map[string]any{
			"invitation_id": d.Id(),
		})
		d.SetId("")
		return nil
	}

	if err = d.Set("invitation_id", int(invitation.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("role", invitation.GetRole()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("login", invitation.GetLogin()); err != nil {
		return diag.FromErr(err)
	}
	if invitation.GetEmail() != "" {
		if err = d.Set("email", invitation.GetEmail()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubOrganizationInvitationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	invitationID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err = client.Organizations.CancelInvite(ctx, orgName, invitationID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			// Already cancelled or accepted — not an error
			tflog.Info(ctx, fmt.Sprintf("Organization invitation %s was already cancelled or accepted", d.Id()), map[string]any{
				"invitation_id": d.Id(),
			})
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
