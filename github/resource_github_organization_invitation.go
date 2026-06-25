package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationInvitation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationInvitationCreate,
		ReadContext:   resourceGithubOrganizationInvitationRead,
		DeleteContext: resourceGithubOrganizationInvitationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"email", "invitee_id"},
				Description:  "The email address of the person to invite. Exactly one of email or invitee_id must be set.",
			},
			"invitee_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"email", "invitee_id"},
				Description:  "The GitHub user ID of the person to invite. Exactly one of email or invitee_id must be set.",
			},
			"role": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          "direct_member",
				ValidateDiagFunc: validateValueFunc([]string{"admin", "direct_member", "billing_manager"}),
				Description:      "The role for the new member. Must be one of admin, direct_member, or billing_manager. Defaults to direct_member.",
			},
			"login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitHub username of the invited user, if available at invitation time.",
			},
		},
	}
}

func resourceGithubOrganizationInvitationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	opts := &github.CreateOrgInvitationOptions{
		Role: github.Ptr(d.Get("role").(string)),
	}

	if v, ok := d.GetOk("email"); ok {
		opts.Email = github.Ptr(v.(string))
	}
	if v, ok := d.GetOk("invitee_id"); ok {
		opts.InviteeID = github.Ptr(int64(v.(int)))
	}

	invitation, _, err := client.Organizations.CreateOrgInvitation(ctx, org, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(invitation.GetID(), 10))

	if err := d.Set("login", invitation.GetLogin()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationInvitationRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	invitationID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	// There is no single-item GET endpoint for org invitations;
	// scan the pending list to find ours.
	opts := &github.ListOptions{PerPage: maxPerPage}
	var found *github.Invitation
	for {
		invitations, resp, err := client.Organizations.ListPendingOrgInvitations(ctx, org, opts)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, inv := range invitations {
			if inv.GetID() == invitationID {
				found = inv
				break
			}
		}
		if found != nil || resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	if found == nil {
		// Invitation is no longer pending. Check if the invitee accepted and
		// became an org member — if so the resource goal is achieved; keep in
		// state and let Terraform consider it up-to-date. If the invitee is
		// not a member either, the invitation expired or was cancelled and
		// must be removed from state so Terraform can re-invite.
		login := d.Get("login").(string)
		if login != "" {
			isMember, _, err := client.Organizations.IsMember(ctx, org, login)
			if err != nil {
				return diag.FromErr(err)
			}
			if isMember {
				tflog.Info(ctx, "Invitation was accepted; user is an org member — keeping in state", map[string]any{"login": login})
				return nil
			}
		}
		tflog.Info(ctx, "Organization invitation no longer pending and invitee is not a member, removing from state", map[string]any{"invitation_id": d.Id()})
		d.SetId("")
		return nil
	}

	if err := d.Set("role", found.GetRole()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("login", found.GetLogin()); err != nil {
		return diag.FromErr(err)
	}
	// Only set email if the resource was created with email — the API always
	// returns the invitee's email even for invitee_id-based invitations, which
	// would cause a perpetual diff when invitee_id is the configured field.
	if _, usingEmail := d.GetOk("email"); usingEmail {
		if err := d.Set("email", found.GetEmail()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubOrganizationInvitationDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	invitationID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.CancelInvite(ctx, org, invitationID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Organization invitation no longer exists, skipping cancel", map[string]any{"invitation_id": d.Id()})
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
