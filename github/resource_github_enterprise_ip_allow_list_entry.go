package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func resourceGithubEnterpriseIpAllowListEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "Manage a GitHub Enterprise IP Allow List Entry.",
		CreateContext: resourceGithubEnterpriseIpAllowListEntryCreate,
		ReadContext:   resourceGithubEnterpriseIpAllowListEntryRead,
		UpdateContext: resourceGithubEnterpriseIpAllowListEntryUpdate,
		DeleteContext: resourceGithubEnterpriseIpAllowListEntryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise to apply the IP allow list entry to.",
			},
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An IP address or range of IP addresses in CIDR notation.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An optional name for the IP allow list entry.",
			},
			"is_active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the entry is currently active.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the entry was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the entry was last updated.",
			},
		},
	}
}

func resourceGithubEnterpriseIpAllowListEntryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v4client

	// First, get the enterprise ID as we need it for the mutation
	enterpriseSlug := d.Get("enterprise_slug").(string)
	enterpriseID, err := getEnterpriseID(ctx, client, enterpriseSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	// Then create the IP allow list entry
	var mutation struct {
		CreateIpAllowListEntry struct {
			IpAllowListEntry struct {
				ID             githubv4.String
				AllowListValue githubv4.String
				Name           githubv4.String
				IsActive       githubv4.Boolean
				CreatedAt      githubv4.String
				UpdatedAt      githubv4.String
			}
		} `graphql:"createIpAllowListEntry(input: $input)"`
	}

	name := d.Get("name").(string)
	input := githubv4.CreateIpAllowListEntryInput{
		OwnerID:        githubv4.ID(enterpriseID),
		AllowListValue: githubv4.String(d.Get("ip").(string)),
		IsActive:       githubv4.Boolean(d.Get("is_active").(bool)),
	}

	if name != "" {
		input.Name = githubv4.NewString(githubv4.String(name))
	}

	err = client.Mutate(ctx, &mutation, input, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(mutation.CreateIpAllowListEntry.IpAllowListEntry.ID))

	if err := d.Set("created_at", mutation.CreateIpAllowListEntry.IpAllowListEntry.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", mutation.CreateIpAllowListEntry.IpAllowListEntry.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseIpAllowListEntryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v4client

	var query struct {
		Node struct {
			IpAllowListEntry struct {
				ID             githubv4.String
				AllowListValue githubv4.String
				Name           githubv4.String
				IsActive       githubv4.Boolean
				CreatedAt      githubv4.String
				UpdatedAt      githubv4.String
				Owner          struct {
					Enterprise struct {
						Slug githubv4.String
					} `graphql:"... on Enterprise"`
				}
			} `graphql:"... on IpAllowListEntry"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": githubv4.ID(d.Id()),
	}

	err := client.Query(ctx, &query, variables)
	if err != nil {
		if strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
			tflog.Info(ctx, "Removing IP allow list entry from state because it no longer exists in GitHub", map[string]any{
				"id": d.Id(),
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	entry := query.Node.IpAllowListEntry

	d.Set("ip", entry.AllowListValue)
	if err := d.Set("name", entry.Name); err != nil {
		return diag.FromErr(err)
	}
	d.Set("is_active", entry.IsActive)
	d.Set("created_at", entry.CreatedAt)
	d.Set("updated_at", entry.UpdatedAt)
	d.Set("enterprise_slug", entry.Owner.Enterprise.Slug)

	return nil
}

func resourceGithubEnterpriseIpAllowListEntryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v4client

	var mutation struct {
		UpdateIpAllowListEntry struct {
			IpAllowListEntry struct {
				ID             githubv4.String
				AllowListValue githubv4.String
				Name           githubv4.String
				IsActive       githubv4.Boolean
				UpdatedAt      githubv4.String
			}
		} `graphql:"updateIpAllowListEntry(input: $input)"`
	}

	name := d.Get("name").(string)
	input := githubv4.UpdateIpAllowListEntryInput{
		IPAllowListEntryID: githubv4.ID(d.Id()),
		AllowListValue:     githubv4.String(d.Get("ip").(string)),
		IsActive:           githubv4.Boolean(d.Get("is_active").(bool)),
	}

	if name != "" {
		input.Name = githubv4.NewString(githubv4.String(name))
	}

	err := client.Mutate(ctx, &mutation, input, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("updated_at", mutation.UpdateIpAllowListEntry.IpAllowListEntry.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseIpAllowListEntryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v4client

	var mutation struct {
		DeleteIpAllowListEntry struct {
			ClientMutationID githubv4.String
		} `graphql:"deleteIpAllowListEntry(input: $input)"`
	}

	input := githubv4.DeleteIpAllowListEntryInput{
		IPAllowListEntryID: githubv4.ID(d.Id()),
	}

	err := client.Mutate(ctx, &mutation, input, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
