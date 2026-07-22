package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
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

		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
				ExactlyOneOf:     []string{"username", "user_id"},
				Description:      "The user (login) to add to the organization. Exactly one of `username` or `user_id` must be set.",
			},
			"user_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"username", "user_id"},
				Description:  "The GitHub numeric user ID to add to the organization. Stable across username changes; recommended over `username` for production usage. Exactly one of `username` or `user_id` must be set.",
			},
			"role": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateValueFunc([]string{"member", "admin"}),
				Default:          "member",
				Description:      "The role of the user within the organization. Must be one of 'member' or 'admin'.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
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

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	username, userID, err := resolveMembershipIdentity(ctx, client, d)
	if err != nil {
		return diag.FromErr(err)
	}

	roleName := d.Get("role").(string)

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

	d.SetId(buildTwoPartID(orgName, strconv.FormatInt(userID, 10)))

	if err = d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("user_id", userID); err != nil {
		return diag.FromErr(err)
	}
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

	orgPart, secondPart, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	username, userID, err := loginAndIDFromIDPart(ctx, client, secondPart)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, fmt.Sprintf("Removing membership %s from state because the user no longer exists in GitHub", d.Id()), map[string]any{
				"membership_id": d.Id(),
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	// Lazily migrate legacy IDs of the form `org:username` to `org:user_id`.
	// New resources are always created with the numeric form (see Create).
	if secondPart != strconv.FormatInt(userID, 10) {
		d.SetId(buildTwoPartID(orgPart, strconv.FormatInt(userID, 10)))
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	membership, resp, err := client.Organizations.GetOrgMembership(ctx, username, orgName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
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
	if err = d.Set("user_id", userID); err != nil {
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

	// Username in state is kept fresh by Read, so it reflects the user's
	// current login even after a rename.
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
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) {
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
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) {
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

// resolveMembershipIdentity returns the (login, numeric_id) pair for the
// configured membership, regardless of whether the user supplied `username`
// or `user_id`. The GitHub org membership endpoints only accept the login, so
// when `user_id` is provided we must resolve it via GET /user/{id} first.
func resolveMembershipIdentity(ctx context.Context, client *github.Client, d *schema.ResourceData) (string, int64, error) {
	if v, ok := d.GetOk("user_id"); ok {
		userID := int64(v.(int))
		user, _, err := client.Users.GetByID(ctx, userID)
		if err != nil {
			return "", 0, err
		}
		return user.GetLogin(), user.GetID(), nil
	}

	username := d.Get("username").(string)
	user, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return "", 0, err
	}
	return user.GetLogin(), user.GetID(), nil
}

// loginAndIDFromIDPart resolves the (login, numeric_id) pair from the second
// segment of a resource ID. New resources use `org:<numeric_id>`; legacy
// resources use `org:<username>` and are migrated on the next Read.
func loginAndIDFromIDPart(ctx context.Context, client *github.Client, idPart string) (string, int64, error) {
	if userID, err := strconv.ParseInt(idPart, 10, 64); err == nil {
		user, _, err := client.Users.GetByID(ctx, userID)
		if err != nil {
			return "", 0, err
		}
		return user.GetLogin(), user.GetID(), nil
	}

	user, _, err := client.Users.Get(ctx, idPart)
	if err != nil {
		return "", 0, err
	}
	return user.GetLogin(), user.GetID(), nil
}
