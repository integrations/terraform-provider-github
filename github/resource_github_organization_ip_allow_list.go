package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func resourceGithubOrganizationIpAllowList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGithubOrganizationIpAllowListRead,
		Create: resourceGithubOrganizationIpAllowListCreate,
		Update: resourceGithubOrganizationIpAllowListUpdate,
		Delete: resourceGithubOrganizationIpAllowListDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGithubOrganizationIpAllowListImport,
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the IP allow list entry.",
			},
			"allow_list_value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A single IP address or range of IP addresses in CIDR notation.",
			},
			"is_active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the entry is currently active. Default is true.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifies the date and time when the object was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifies the date and time when the object was last updated.",
			},
		},
	}
}

func resourceGithubOrganizationIpAllowListImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgName := meta.(*Owner).name

	// Fetch all IP allow list entries for the org.
	ipAllowListEntries, err := getOrganizationIpAllowListEntries(meta)
	if err != nil {
		return nil, err
	}

	// For user convenience, we support importing ip allow list entries
	// by both their graphql node id or by their ip allow list value.
	// For example, "IALE_kwHNBQk8ps2ADPDc" or "192.168.1.1/24"
	valueToImport := d.Id()
	ipAllowListEntryId := ""

	for index := range ipAllowListEntries {
		ipAllowListEntry := ipAllowListEntries[index]
		if string(ipAllowListEntry.ID) == valueToImport || string(ipAllowListEntry.AllowListValue) == valueToImport {
			ipAllowListEntryId = string(ipAllowListEntry.ID)
			break
		}
	}

	if ipAllowListEntryId == "" {
		return nil, fmt.Errorf("Organization %s does not have an IP allow list entry for %s.", orgName, valueToImport)
	}

	d.SetId(ipAllowListEntryId)
	err = resourceGithubOrganizationIpAllowListRead(d, meta)
	if err != nil {
		return nil, err
	}

	// resourceGithubOrganizationIpAllowListRead calls d.SetId("") if the ip entry does not exist
	if d.Id() == "" {
		return nil, fmt.Errorf("Organization %s does not have an IP allow list entry for %s.", orgName, valueToImport)
	}

	return []*schema.ResourceData{d}, nil
}

func resourceGithubOrganizationIpAllowListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	var query struct {
		Node struct {
			IpAllowListEntry struct {
				ID             githubv4.String
				Name           githubv4.String
				AllowListValue githubv4.String
				IsActive       githubv4.Boolean
				CreatedAt      githubv4.String
				UpdatedAt      githubv4.String
			} `graphql:"... on IpAllowListEntry"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": d.Id(),
	}

	err := client.Query(ctx, &query, variables)
	if err != nil {
		if githubv4IsNodeNotFoundError(err) {
			log.Printf("[INFO] Removing ip allow list entry %s from state because it no longer exists in GitHub", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", query.Node.IpAllowListEntry.Name)
	d.Set("allow_list_value", query.Node.IpAllowListEntry.AllowListValue)
	d.Set("is_active", query.Node.IpAllowListEntry.IsActive)
	d.Set("created_at", query.Node.IpAllowListEntry.CreatedAt)
	d.Set("updated_at", query.Node.IpAllowListEntry.UpdatedAt)

	return nil
}

func resourceGithubOrganizationIpAllowListCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	orgId := meta.(*Owner).nodeId
	name := d.Get("name").(string)
	allowListValue := d.Get("allow_list_value").(string)
	isActive := d.Get("is_active").(bool)

	var mutate struct {
		CreateIpAllowListEntry struct {
			IpAllowListEntry struct {
				ID githubv4.String
			} `graphql:"ipAllowListEntry"`
		} `graphql:"createIpAllowListEntry(input: $input)"`
	}

	input := githubv4.CreateIpAllowListEntryInput{
		OwnerID:        githubv4.NewID(orgId),
		Name:           (*githubv4.String)(&name),
		AllowListValue: githubv4.String(allowListValue),
		IsActive:       githubv4.Boolean(isActive),
	}

	err := client.Mutate(ctx, &mutate, input, nil)
	if err != nil {
		return err
	}

	d.SetId(string(mutate.CreateIpAllowListEntry.IpAllowListEntry.ID))

	return resourceGithubOrganizationIpAllowListRead(d, meta)
}

func resourceGithubOrganizationIpAllowListUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	name := d.Get("name").(string)
	allowListValue := d.Get("allow_list_value").(string)
	isActive := d.Get("is_active").(bool)

	var mutate struct {
		UpdateIpAllowListEntry struct {
			IpAllowListEntry struct {
				ID githubv4.String
			} `graphql:"ipAllowListEntry"`
		} `graphql:"updateIpAllowListEntry(input: $input)"`
	}

	input := githubv4.UpdateIpAllowListEntryInput{
		IPAllowListEntryID: d.Id(),
		Name:               (*githubv4.String)(&name),
		AllowListValue:     githubv4.String(allowListValue),
		IsActive:           githubv4.Boolean(isActive),
	}

	err := client.Mutate(ctx, &mutate, input, nil)
	if err != nil {
		return err
	}

	d.SetId(string(mutate.UpdateIpAllowListEntry.IpAllowListEntry.ID))

	return resourceGithubOrganizationIpAllowListRead(d, meta)
}

func resourceGithubOrganizationIpAllowListDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v4client
	ctx := context.Background()

	var mutate struct {
		DeleteIpAllowListEntry struct {
			IpAllowListEntry struct {
				ID githubv4.String
			} `graphql:"ipAllowListEntry"`
		} `graphql:"deleteIpAllowListEntry(input: $input)"`
	}

	input := githubv4.DeleteIpAllowListEntryInput{
		IPAllowListEntryID: d.Id(),
	}

	err := client.Mutate(ctx, &mutate, input, nil)
	return err
}
