package github

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEMUGroupMapping() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEMUGroupMappingCreate,
		ReadContext:   resourceGithubEMUGroupMappingRead,
		UpdateContext: resourceGithubEMUGroupMappingUpdate,
		DeleteContext: resourceGithubEMUGroupMappingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEMUGroupMappingImport,
		},
		CustomizeDiff: diffTeam,
		Description:   "Manages the mapping of an external group to a GitHub team.",
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the GitHub team.",
			},
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Slug of the GitHub team.",
			},
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
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
		SchemaVersion: 2,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubEMUGroupMappingV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubEMUGroupMappingStateUpgradeV0,
				Version: 0,
			},
			{
				Type:    resourceGithubEMUGroupMappingV1().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubEMUGroupMappingStateUpgradeV1,
				Version: 1,
			},
		},
	}
}

func resourceGithubEMUGroupMappingCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Trace(ctx, "Creating EMU group mapping")

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)

	groupID := toInt64(d.Get("group_id"))
	eg := &github.ExternalGroup{
		GroupID: github.Ptr(groupID),
	}

	tflog.Debug(ctx, "Connecting external group to team via GitHub API", map[string]any{
		"org_name":  orgName,
		"team_slug": teamSlug,
		"group_id":  groupID,
	})

	group, resp, err := client.Teams.UpdateConnectedExternalGroup(ctx, orgName, teamSlug, eg)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Successfully updated connected external group")

	teamID, err := lookupTeamID(ctx, meta.(*Owner), teamSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	newResourceID, err := buildID(strconv.FormatInt(groupID, 10), strconv.FormatInt(teamID, 10))
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Setting resource ID", map[string]any{
		"resource_id": newResourceID,
	})
	d.SetId(newResourceID)

	tflog.Trace(ctx, "Setting team_id", map[string]any{
		"team_id": teamID,
	})
	if err := d.Set("team_id", int(teamID)); err != nil {
		return diag.FromErr(err)
	}

	etag := resp.Header.Get("ETag")
	tflog.Trace(ctx, "Setting state attribute: etag", map[string]any{
		"etag": etag,
	})
	if err := d.Set("etag", etag); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Setting group_name", map[string]any{
		"group_name": group.GetGroupName(),
	})
	if err := d.Set("group_name", group.GetGroupName()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Resource created successfully", map[string]any{
		"resource_id": d.Id(),
	})

	return nil
}

func resourceGithubEMUGroupMappingRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	tflog.Trace(ctx, "Reading EMU group mapping")

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	groupID := toInt64(d.Get("group_id"))
	teamSlug := d.Get("team_slug").(string)

	tflog.Debug(ctx, "Querying external groups linked to team from GitHub API", map[string]any{
		"org_name":  orgName,
		"team_slug": teamSlug,
	})

	groupsList, resp, err := client.Teams.ListExternalGroupsForTeamBySlug(ctx, orgName, teamSlug)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusBadRequest {
			tflog.Info(ctx, "Removing EMU group mapping from state because the team has explicit members in GitHub")
			d.SetId("")
			return nil
		}
		if resp != nil && (resp.StatusCode == http.StatusNotFound) {
			// If the Group is not found, remove it from state
			tflog.Info(ctx, "Removing EMU group mapping from state because team no longer exists in GitHub")
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if len(groupsList.Groups) < 1 {
		tflog.Info(ctx, "Removing EMU group mapping from state because no external groups are linked to the team")
		d.SetId("")
		return nil
	}

	// A team can only be linked to one external group
	group := groupsList.Groups[0]

	tflog.Debug(ctx, "Successfully retrieved external group from GitHub API", map[string]any{
		"configured_group_id": groupID,
		"upstream_group_id":   group.GetGroupID(),
		"group_name":          group.GetGroupName(),
	})

	if group.GetGroupID() != groupID {
		return diag.Errorf("group id mismatch: %d != %d", group.GetGroupID(), groupID)
	}

	etag := resp.Header.Get("ETag")
	if err := d.Set("etag", etag); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("group_id", int(group.GetGroupID())); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("group_name", group.GetGroupName()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEMUGroupMappingUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	tflog.Trace(ctx, "Updating EMU group mapping")

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)

	groupID := toInt64(d.Get("group_id"))
	eg := &github.ExternalGroup{
		GroupID: github.Ptr(groupID),
	}

	if d.HasChange("team_slug") {

		tflog.Debug(ctx, "Updating connected external group via GitHub API", map[string]any{
			"org_name":  orgName,
			"team_slug": teamSlug,
			"group_id":  groupID,
		})

		group, resp, err := client.Teams.UpdateConnectedExternalGroup(ctx, orgName, teamSlug, eg)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Debug(ctx, "Successfully updated connected external group")

		etag := resp.Header.Get("ETag")
		tflog.Trace(ctx, "Setting state attribute: etag", map[string]any{
			"etag": etag,
		})
		if err := d.Set("etag", etag); err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("group_name", group.GetGroupName()); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Trace(ctx, "Updated successfully")

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

	// <group-id>:<team-slug>
	groupIDString, teamSlug, err := parseID2(d.Id())
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

	teamID, err := lookupTeamID(ctx, meta.(*Owner), teamSlug)
	if err != nil {
		return nil, err
	}

	if err := d.Set("team_id", int(teamID)); err != nil {
		return nil, err
	}

	if err := d.Set("group_id", groupID); err != nil {
		return nil, err
	}

	if err := d.Set("team_slug", teamSlug); err != nil {
		return nil, err
	}

	resourceID, err := buildID(groupIDString, strconv.FormatInt(teamID, 10))
	if err != nil {
		return nil, err
	}

	tflog.Trace(ctx, "Setting resource ID", map[string]any{
		"resource_id": resourceID,
	})
	d.SetId(resourceID)

	return []*schema.ResourceData{d}, nil
}
