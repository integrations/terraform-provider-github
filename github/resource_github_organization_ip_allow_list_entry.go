package github

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func resourceGithubOrganizationIpAllowListEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationIpAllowListEntryCreate,
		Read:   resourceGithubOrganizationIpAllowListEntryRead,
		Update: resourceGithubOrganizationIpAllowListEntryUpdate,
		Delete: resourceGithubOrganizationIpAllowListEntryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceGithubOrganizationIpAllowListEntryCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	// First, get the organization ID as we need it for the mutation
	var getOrgQuery struct {
		Organization struct {
			ID githubv4.ID
		} `graphql:"organization(login: $login)"`
	}

	variables := map[string]interface{}{
		"login": githubv4.String(orgName),
	}

	err = client.Query(ctx, &getOrgQuery, variables)
	if err != nil {
		return err
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
			}
		} `graphql:"createIpAllowListEntry(input: $input)"`
	}

	name := d.Get("name").(string)
	input := githubv4.CreateIpAllowListEntryInput{
		OwnerID:        getOrgQuery.Organization.ID,
		AllowListValue: githubv4.String(d.Get("ip").(string)),
		IsActive:       githubv4.Boolean(d.Get("is_active").(bool)),
	}

	if name != "" {
		input.Name = githubv4.NewString(githubv4.String(name))
	}

	err = client.Mutate(ctx, &mutation, input, nil)
	if err != nil {
		return err
	}

	d.SetId(string(mutation.CreateIpAllowListEntry.IpAllowListEntry.ID))

	return resourceGithubOrganizationIpAllowListEntryRead(d, meta)
}

func resourceGithubOrganizationIpAllowListEntryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	var query struct {
		Node struct {
			IpAllowListEntry struct {
				ID             githubv4.String
				AllowListValue githubv4.String
				Name           githubv4.String
				IsActive       githubv4.Boolean
				CreatedAt      githubv4.String
			} `graphql:"... on IpAllowListEntry"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": githubv4.ID(d.Id()),
	}

	err := client.Query(ctx, &query, variables)
	if err != nil {
		if strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
			log.Printf("[INFO] Removing IP allow list entry (%s) from state because it no longer exists in GitHub", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	entry := query.Node.IpAllowListEntry

	d.Set("ip", entry.AllowListValue)
	d.Set("name", entry.Name)
	d.Set("is_active", entry.IsActive)

	return nil
}

func resourceGithubOrganizationIpAllowListEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	var mutation struct {
		UpdateIpAllowListEntry struct {
			IpAllowListEntry struct {
				ID             githubv4.String
				AllowListValue githubv4.String
				Name           githubv4.String
				IsActive       githubv4.Boolean
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
		return err
	}

	return resourceGithubOrganizationIpAllowListEntryRead(d, meta)
}

func resourceGithubOrganizationIpAllowListEntryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

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
		return err
	}

	d.SetId("")
	return nil
}
