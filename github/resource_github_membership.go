package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	membershipDowngradeToMember              = "member"
	membershipDowngradeToOutsideCollaborator = "outside_collaborator"
	membershipStateActive                    = "active"
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

		CustomizeDiff: diffETag,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
				Description:      "The user to add to the organization.",
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
				Description: "Instead of removing the member from the org, you can choose to downgrade their membership when this resource is destroyed. This is useful when wanting to downgrade admins while keeping them in the organization, or to downgrade members while keeping their access to public repositories intact.",
			},
			"downgrade_to": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateValueFunc([]string{membershipDowngradeToMember, membershipDowngradeToOutsideCollaborator}),
				Default:          membershipDowngradeToMember,
				Description:      "The target membership state when downgrade_on_destroy is true. Must be one of 'member' or 'outside_collaborator'.",
			},
		},
	}
}

func resourceGithubMembershipCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	username := d.Get("username").(string)
	roleName := d.Get("role").(string)
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
		if !d.HasChange("role") {
			return resourceGithubMembershipRead(ctx, d, meta)
		}
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

func resourceGithubMembershipRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	_, username, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
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

func resourceGithubMembershipDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx = context.WithValue(ctx, ctxId, d.Id())

	username := d.Get("username").(string)
	downgradeOnDestroy := d.Get("downgrade_on_destroy").(bool)
	downgradeTo, ok := d.Get("downgrade_to").(string)
	if !ok || downgradeTo == "" {
		downgradeTo = membershipDowngradeToMember
	}

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

		if downgradeTo == membershipDowngradeToMember && membership.GetRole() == downgradeTo {
			tflog.Info(ctx, fmt.Sprintf("Not downgrading '%s' membership for '%s' because they are already '%s'", orgName, username, downgradeTo), map[string]any{
				"org_name": orgName,
				"username": username,
				"role":     downgradeTo,
			})
			return nil
		}

		if downgradeTo == membershipDowngradeToOutsideCollaborator && membership.GetState() != membershipStateActive {
			return diag.Errorf("cannot downgrade %q to outside collaborator because organization membership is %q, not active; the user must accept the organization invitation first", username, membership.GetState())
		}

		switch downgradeTo {
		case membershipDowngradeToMember:
			_, _, err = client.Organizations.EditOrgMembership(ctx, username, orgName, &github.Membership{
				Role: new(downgradeTo),
			})
		case membershipDowngradeToOutsideCollaborator:
			_, err = client.Organizations.ConvertMemberToOutsideCollaborator(ctx, orgName, username)
		default:
			return diag.Errorf("%s is an invalid value for argument downgrade_to", downgradeTo)
		}
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
