package github

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEMUGroupMapping() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEMUGroupMappingCreateOrUpdate,
		ReadContext:   resourceGithubEMUGroupMappingRead,
		UpdateContext: resourceGithubEMUGroupMappingCreateOrUpdate,
		DeleteContext: resourceGithubEMUGroupMappingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEMUGroupMappingImport,
		},
		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Slug of the GitHub team.",
			},
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Integer corresponding to the external group ID to be linked.",
			},
			"group_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the external group.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubEMUGroupMappingRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Trace(ctx, "Reading EMU group mapping", map[string]any{
		"resource_id": d.Id(),
	})

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	id64 := toInt64(d.Get("group_id"))
	teamSlug := d.Get("team_slug").(string)

	tflog.SetField(ctx, "group_id", id64)
	tflog.SetField(ctx, "team_slug", teamSlug)
	tflog.SetField(ctx, "org_name", orgName)

	tflog.Debug(ctx, "Querying external groups linked to team from GitHub API")

	groupsList, resp, err := client.Teams.ListExternalGroupsForTeamBySlug(ctx, orgName, teamSlug)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusBadRequest {
			tflog.Info(ctx, "Removing EMU group mapping from state because the team has explicit members in GitHub", map[string]any{
				"resource_id": d.Id(),
			})
			d.SetId("")
			return nil
		}
		if resp != nil && (resp.StatusCode == http.StatusNotFound) {
			// If the Group is not found, remove it from state
			tflog.Info(ctx, "Removing EMU group mapping from state because team no longer exists in GitHub", map[string]any{
				"resource_id": d.Id(),
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if len(groupsList.Groups) < 1 {
		tflog.Info(ctx, "Removing EMU group mapping from state because no external groups are linked to the team", map[string]any{
			"resource_id": d.Id(),
		})
		d.SetId("")
		return nil
	}

	// A team can only be linked to one external group
	group := groupsList.Groups[0]

	tflog.Debug(ctx, "Successfully retrieved external group from GitHub API", map[string]any{
		"group_id":   group.GetGroupID(),
		"group_name": group.GetGroupName(),
	})

	if group.GetGroupID() != id64 {
		return diag.Errorf("group id mismatch: %d != %d", group.GetGroupID(), id64)
	}

	etag := resp.Header.Get("ETag")
	if err = d.Set("etag", etag); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("group_id", int(group.GetGroupID())); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("group_name", group.GetGroupName()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEMUGroupMappingCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	resourceID := d.Id()
	tflog.Trace(ctx, "Creating or updating EMU group mapping", map[string]any{
		"resource_id": resourceID,
	})

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	tflog.SetField(ctx, "org_name", orgName)

	teamSlug := d.Get("team_slug").(string)
	tflog.SetField(ctx, "team_slug", teamSlug)

	id64 := toInt64(d.Get("group_id"))
	tflog.SetField(ctx, "group_id", id64)
	eg := &github.ExternalGroup{
		GroupID: github.Ptr(id64),
	}

	tflog.Debug(ctx, "Updating connected external group via GitHub API")

	group, resp, err := client.Teams.UpdateConnectedExternalGroup(ctx, orgName, teamSlug, eg)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Successfully updated connected external group")

	newResourceID := fmt.Sprintf("teams/%s/external-groups", teamSlug)
	tflog.Trace(ctx, "Setting resource ID", map[string]any{
		"resource_id": newResourceID,
	})
	d.SetId(newResourceID)

	etag := resp.Header.Get("ETag")
	tflog.Trace(ctx, "Setting state attribute: etag", map[string]any{
		"etag": etag,
	})
	if err = d.Set("etag", etag); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("group_name", group.GetGroupName()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Resource created or updated successfully", map[string]any{
		"resource_id": d.Id(),
	})

	return nil
}

func resourceGithubEMUGroupMappingDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Trace(ctx, "Deleting EMU group mapping", map[string]any{
		"resource_id": d.Id(),
	})

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	teamSlug, ok := d.GetOk("team_slug")
	if !ok {
		return diag.Errorf("could not parse team slug from provided value")
	}

	teamSlugStr := teamSlug.(string)
	tflog.Debug(ctx, "Removing connected external group from team via GitHub API", map[string]any{
		"org_name":    orgName,
		"team_slug":   teamSlugStr,
		"resource_id": d.Id(),
	})

	_, err = client.Teams.RemoveConnectedExternalGroup(ctx, orgName, teamSlugStr)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Successfully removed connected external group from team", map[string]any{
		"org_name":    orgName,
		"team_slug":   teamSlugStr,
		"resource_id": d.Id(),
	})
	return nil
}

func resourceGithubEMUGroupMappingImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	importID := d.Id()
	tflog.Trace(ctx, "Importing EMU group mapping with two-part ID", map[string]any{
		"import_id": importID,
		"strategy":  "two_part_id",
	})

	groupIDString, teamSlug, err := parseTwoPartID(d.Id(), "group_id", "team_slug")
	if err != nil {
		return nil, err
	}
	groupID, err := strconv.Atoi(groupIDString)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Parsed two-part import ID", map[string]any{
		"import_id": importID,
		"group_id":  groupID,
		"team_slug": teamSlug,
	})

	if err := d.Set("group_id", groupID); err != nil {
		return nil, err
	}

	if err := d.Set("team_slug", teamSlug); err != nil {
		return nil, err
	}

	resourceID := fmt.Sprintf("teams/%s/external-groups", teamSlug)
	tflog.Trace(ctx, "Setting resource ID", map[string]any{
		"resource_id": resourceID,
	})
	d.SetId(resourceID)
	return []*schema.ResourceData{d}, nil
}
