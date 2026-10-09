package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	username := d.Get("username").(string)
	roleName := d.Get("role").(string)
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
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
				tflog.Info(ctx, "Removing membership from state because it no longer exists in GitHub", map[string]any{"membership_id": d.Id()})

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
	downgradeTo := "member"

	if downgradeOnDestroy {
		tflog.Info(ctx, "Downgrading organization membership", map[string]any{"owner": orgName, "username": username, "role": downgradeTo})

		// Check to make sure this member still has access to the organization before downgrading.
		// If we don't do this, the member would just be re-added to the organization.
		var membership *github.Membership
		membership, _, err = client.Organizations.GetOrgMembership(ctx, username, orgName)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
				if ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, "Not downgrading organization membership because the user is no longer a member", map[string]any{"owner": orgName, "username": username})

					return nil
				}
			}

			return diag.FromErr(err)
		}

		if *membership.Role == downgradeTo {
			tflog.Info(ctx, "Not downgrading organization membership because the member already has the requested role", map[string]any{"owner": orgName, "username": username, "role": downgradeTo})

			return nil
		}

		_, _, err = client.Organizations.EditOrgMembership(ctx, username, orgName, &github.Membership{
			Role: new(downgradeTo),
		})
	} else {
		tflog.Info(ctx, "Revoking organization membership", map[string]any{"owner": orgName, "username": username})

		_, err = client.Organizations.RemoveOrgMembership(ctx, username, orgName)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
				if ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, "Not removing organization membership because the user is no longer a member", map[string]any{"owner": orgName, "username": username})

					return nil
				}
			}

			return diag.FromErr(err)
		}
	}

	return diag.FromErr(err)
}
