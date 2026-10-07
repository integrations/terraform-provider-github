package github

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"net/http"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubMembership() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubMembershipCreateOrUpdate,
		ReadContext:   resourceGithubMembershipRead,
		UpdateContext: resourceGithubMembershipCreateOrUpdate,
		DeleteContext: resourceGithubMembershipDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: customdiff.All(
			diffETag,
			// ConflictsWith cannot express this because downgrade_on_destroy
			// is a bool with a default; an explicit 'false' would already
			// trigger it.
			func(_ context.Context, d *schema.ResourceDiff, _ any) error {
				email, _ := d.Get("email").(string)
				downgradeOnDestroy, _ := d.Get("downgrade_on_destroy").(bool)
				if email != "" && downgradeOnDestroy {
					return errors.New("downgrade_on_destroy cannot be used together with email")
				}
				return nil
			},
		),

		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ExactlyOneOf:     []string{"username", "email"},
				DiffSuppressFunc: caseInsensitive(),
				Description:      "The user to add to the organization. Exactly one of 'username' or 'email' must be set.",
			},
			"email": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"username", "email"},
				ValidateDiagFunc: func(v any, k cty.Path) diag.Diagnostics {
					value, ok := v.(string)
					if !ok || !strings.Contains(value, "@") {
						return diag.Errorf("%q is not a valid email address", v)
					}
					return nil
				},
				DiffSuppressFunc: caseInsensitive(),
				Description:      "The email address of the person to invite to the organization. The invitee does not need an existing GitHub account. Exactly one of 'username' or 'email' must be set.",
			},
			"role": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateValueFunc([]string{"member", "admin"}),
				Default:          "member",
				Description:      "The role of the user within the organization. Must be one of 'member' or 'admin'.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the membership.",
			},
			"downgrade_on_destroy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Instead of removing the member from the org, you can choose to downgrade their membership to 'member' when this resource is destroyed. This is useful when wanting to downgrade admins while keeping them in the organization",
			},
		},
	}
}

func resourceGithubMembershipCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("expected meta to be *Owner, got %T", meta)
	}
	client := owner.v3client
	orgName := owner.name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	if email, ok := d.Get("email").(string); ok && email != "" {
		return resourceGithubMembershipInvite(ctx, d, owner, email)
	}

	// ExactlyOneOf skips validation when values are unknown at plan time, so
	// an empty username can still reach apply.
	username, ok := d.Get("username").(string)
	if !ok || username == "" {
		return diag.Errorf("exactly one of 'username' or 'email' must be set")
	}
	roleName, ok := d.Get("role").(string)
	if !ok {
		return diag.Errorf("expected role to be a string")
	}

	_, resp, err := client.Organizations.EditOrgMembership(ctx,
		username,
		orgName,
		&github.Membership{
			Role: new(roleName),
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(orgName, username))

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// resourceGithubMembershipInvite handles memberships configured with an email
// address. There is no username to address the membership by (the invitee may
// not even have a GitHub account yet), so the organization invitation API is
// used instead of EditOrgMembership.
func resourceGithubMembershipInvite(ctx context.Context, d *schema.ResourceData, owner *Owner, email string) diag.Diagnostics {
	orgName := owner.name
	roleName, ok := d.Get("role").(string)
	if !ok {
		return diag.Errorf("expected role to be a string")
	}
	// Inviting an email that already has a pending invitation fails, so adopt
	// an existing pending invitation when it matches the configured role.
	invitation, err := findPendingOrgInvitationByEmail(ctx, owner, email)
	if err != nil {
		return diag.FromErr(err)
	}

	if invitation != nil && invitationRoleToMembershipRole(invitation.GetRole()) != roleName {
		if _, err := owner.v3client.Organizations.CancelInvite(ctx, orgName, invitation.GetID()); err != nil {
			return diag.FromErr(err)
		}
		invitation = nil
	}

	if invitation == nil {
		// A lingering failed invitation for the same address would also make
		// the new invitation fail, so remove it first.
		failed, err := findFailedOrgInvitationByEmail(ctx, owner, email)
		if err != nil {
			return diag.FromErr(err)
		}
		if failed != nil {
			if _, err := owner.v3client.Organizations.CancelInvite(ctx, orgName, failed.GetID()); err != nil {
				return diag.FromErr(err)
			}
		}

		_, _, err = owner.v3client.Organizations.CreateOrgInvitation(ctx, orgName, &github.CreateOrgInvitationOptions{
			Email: new(email),
			Role:  new(membershipRoleToInvitationRole(roleName)),
		})
		if err != nil {
			return diag.Errorf("inviting %s to %s: %s (if this person already accepted an earlier invitation, manage their membership with 'username' instead of 'email')", email, orgName, err)
		}
	}

	d.SetId(buildTwoPartID(orgName, email))

	return nil
}

func resourceGithubMembershipRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("expected meta to be *Owner, got %T", meta)
	}
	client := owner.v3client

	orgName := owner.name
	_, usernameOrEmail, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	// GitHub usernames cannot contain "@", so an "@" in the ID means this
	// membership is an email-based invitation. The ID is the only reliable
	// indicator here because imported resources start with an empty state.
	if strings.Contains(usernameOrEmail, "@") {
		return resourceGithubMembershipReadInvitation(ctx, d, owner, usernameOrEmail)
	}

	username := usernameOrEmail
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	membership, resp, err := client.Organizations.GetOrgMembership(ctx,
		username, orgName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, fmt.Sprintf("Removing membership %s from state because it no longer exists in GitHub", d.Id()), map[string]any{
					"membership_id": d.Id(),
				})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("role", membership.GetRole()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubMembershipReadInvitation(ctx context.Context, d *schema.ResourceData, owner *Owner, email string) diag.Diagnostics {
	if err := d.Set("email", email); err != nil {
		return diag.FromErr(err)
	}

	invitation, err := findPendingOrgInvitationByEmail(ctx, owner, email)
	if err != nil {
		return diag.FromErr(err)
	}

	if invitation == nil {
		failed, err := findFailedOrgInvitationByEmail(ctx, owner, email)
		if err != nil {
			return diag.FromErr(err)
		}
		if failed != nil {
			tflog.Info(ctx, fmt.Sprintf("Removing membership %s from state because the invitation failed", d.Id()), map[string]any{
				"membership_id": d.Id(),
			})
			d.SetId("")
			return nil
		}

		// Neither pending nor failed means the invitation was accepted or it
		// was cancelled outside Terraform. The API cannot tell these apart by
		// email, so keep the resource in state rather than re-inviting an
		// existing member on every apply.
		tflog.Info(ctx, fmt.Sprintf("Invitation for membership %s is no longer pending; assuming it was accepted", d.Id()), map[string]any{
			"membership_id": d.Id(),
		})
		return nil
	}

	roleName := invitationRoleToMembershipRole(invitation.GetRole())
	if roleName != "member" && roleName != "admin" {
		return diag.Errorf("invitation for %s has role %q, which is not supported by github_membership", email, invitation.GetRole())
	}
	if err = d.Set("role", roleName); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubMembershipDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("expected meta to be *Owner, got %T", meta)
	}
	client := owner.v3client
	orgName := owner.name
	ctx = context.WithValue(ctx, ctxId, d.Id())

	if email, ok := d.Get("email").(string); ok && email != "" {
		return resourceGithubMembershipCancelInvitation(ctx, owner, email)
	}

	username := d.Get("username").(string)
	downgradeOnDestroy := d.Get("downgrade_on_destroy").(bool)
	downgradeTo := "member"

	if downgradeOnDestroy {
		tflog.Info(ctx, fmt.Sprintf("Downgrading '%s' membership for '%s' to '%s'", orgName, username, downgradeTo), map[string]any{
			"org_name": orgName,
			"username": username,
			"role":     downgradeTo,
		})

		// Check to make sure this member still has access to the organization before downgrading.
		// If we don't do this, the member would just be re-added to the organization.
		var membership *github.Membership
		membership, _, err = client.Organizations.GetOrgMembership(ctx, username, orgName)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
				if ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, fmt.Sprintf("Not downgrading '%s' membership for '%s' because they are not a member of the org anymore", orgName, username), map[string]any{
						"org_name": orgName,
						"username": username,
					})
					return nil
				}
			}

			return diag.FromErr(err)
		}

		if *membership.Role == downgradeTo {
			tflog.Info(ctx, fmt.Sprintf("Not downgrading '%s' membership for '%s' because they are already '%s'", orgName, username, downgradeTo), map[string]any{
				"org_name": orgName,
				"username": username,
				"role":     downgradeTo,
			})
			return nil
		}

		_, _, err = client.Organizations.EditOrgMembership(ctx, username, orgName, &github.Membership{
			Role: new(downgradeTo),
		})
	} else {
		tflog.Info(ctx, fmt.Sprintf("Revoking '%s' membership for '%s'", orgName, username), map[string]any{
			"org_name": orgName,
			"username": username,
		})
		_, err = client.Organizations.RemoveOrgMembership(ctx, username, orgName)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
				if ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, fmt.Sprintf("Not removing '%s' membership for '%s' because they are not a member of the org anymore", orgName, username), map[string]any{
						"org_name": orgName,
						"username": username,
					})
					return nil
				}
			}

			return diag.FromErr(err)
		}
	}

	return diag.FromErr(err)
}

func resourceGithubMembershipCancelInvitation(ctx context.Context, owner *Owner, email string) diag.Diagnostics {
	orgName := owner.name

	invitation, err := findPendingOrgInvitationByEmail(ctx, owner, email)
	if err != nil {
		return diag.FromErr(err)
	}
	if invitation == nil {
		// Failed invitations linger in the organization until removed, so
		// cancel those as well.
		invitation, err = findFailedOrgInvitationByEmail(ctx, owner, email)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if invitation == nil {
		// Surface this as a warning: if the invitation was accepted, the
		// person keeps their organization access after the destroy.
		return diag.Diagnostics{{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("No pending invitation for %s found to cancel", email),
			Detail:   fmt.Sprintf("If the invitation was accepted, %s is still a member of %s. Manage the membership with 'username' to remove or change it.", email, orgName),
		}}
	}

	tflog.Info(ctx, fmt.Sprintf("Cancelling '%s' invitation for '%s'", orgName, email), map[string]any{
		"org_name": orgName,
		"email":    email,
	})

	if _, err := owner.v3client.Organizations.CancelInvite(ctx, orgName, invitation.GetID()); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func findPendingOrgInvitationByEmail(ctx context.Context, owner *Owner, email string) (*github.Invitation, error) {
	opts := &github.ListOptions{PerPage: owner.maxPerPage}
	return findOrgInvitationByEmail(owner.v3client.Organizations.ListPendingOrgInvitationsIter(ctx, owner.name, opts), email)
}

func findFailedOrgInvitationByEmail(ctx context.Context, owner *Owner, email string) (*github.Invitation, error) {
	opts := &github.ListOptions{PerPage: owner.maxPerPage}
	return findOrgInvitationByEmail(owner.v3client.Organizations.ListFailedOrgInvitationsIter(ctx, owner.name, opts), email)
}

// findOrgInvitationByEmail returns the organization invitation addressed to
// email, or nil if there is none. There is no endpoint to fetch an invitation
// directly, so the list has to be scanned.
func findOrgInvitationByEmail(invitations iter.Seq2[*github.Invitation, error], email string) (*github.Invitation, error) {
	for invitation, err := range invitations {
		if err != nil {
			return nil, err
		}
		if strings.EqualFold(invitation.GetEmail(), email) {
			return invitation, nil
		}
	}
	return nil, nil
}

// The invitation API uses a different role vocabulary than the membership
// API: an invited 'direct_member' becomes a 'member' once the invitation is
// accepted.
func membershipRoleToInvitationRole(role string) string {
	if role == "member" {
		return "direct_member"
	}
	return role
}

func invitationRoleToMembershipRole(role string) string {
	if role == "direct_member" {
		return "member"
	}
	return role
}
