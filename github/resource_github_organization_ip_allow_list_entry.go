package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func resourceGithubOrganizationIpAllowListEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "Manage a GitHub Organization IP Allow List Entry.",
		CreateContext: resourceGithubOrganizationIpAllowListEntryCreate,
		ReadContext:   resourceGithubOrganizationIpAllowListEntryRead,
		UpdateContext: resourceGithubOrganizationIpAllowListEntryUpdate,
		DeleteContext: resourceGithubOrganizationIpAllowListEntryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubOrganizationIpAllowListEntryImport,
		},

		Schema: map[string]*schema.Schema{
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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

func resourceGithubOrganizationIpAllowListEntryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	organizationID, err := getOrganizationID(ctx, client, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

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
		OwnerID:        githubv4.ID(organizationID),
		AllowListValue: githubv4.String(d.Get("ip").(string)),
		IsActive:       githubv4.Boolean(d.Get("is_active").(bool)),
	}

	if name != "" {
		v := githubv4.String(name)
		input.Name = &v
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

func resourceGithubOrganizationIpAllowListEntryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

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
					Organization struct {
						Login githubv4.String
					} `graphql:"... on Organization"`
				}
			} `graphql:"... on IpAllowListEntry"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
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
	if err := d.Set("name", entry.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip", entry.AllowListValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_active", entry.IsActive); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", entry.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", entry.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationIpAllowListEntryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

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
		v := githubv4.String(name)
		input.Name = &v
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

func resourceGithubOrganizationIpAllowListEntryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

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
	// GraphQL will return a 200 OK if it couldn't find the global ID
	if err != nil && !strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationIpAllowListEntryImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if err := checkOrganization(meta); err != nil {
		return nil, err
	}

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
					Organization struct {
						Login githubv4.String
					} `graphql:"... on Organization"`
				}
			} `graphql:"... on IpAllowListEntry"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
		"id": githubv4.ID(d.Id()),
	}

	err := client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	entry := query.Node.IpAllowListEntry

	if err := d.Set("ip", string(entry.AllowListValue)); err != nil {
		return nil, err
	}
	if err := d.Set("name", entry.Name); err != nil {
		return nil, err
	}
	if err := d.Set("is_active", entry.IsActive); err != nil {
		return nil, err
	}
	if err := d.Set("created_at", entry.CreatedAt); err != nil {
		return nil, err
	}
	if err := d.Set("updated_at", entry.UpdatedAt); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
