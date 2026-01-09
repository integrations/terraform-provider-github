package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func importEMUGroupMappingWithTwoPartID(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
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

func importEMUGroupMappingWithIntegerID(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	importID := d.Id()
	tflog.Trace(ctx, "Importing EMU group mapping with integer ID", map[string]any{
		"import_id": importID,
		"strategy":  "integer_id",
	})

	groupID, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Parsed integer import ID", map[string]any{
		"import_id": importID,
		"group_id":  groupID,
	})

	if err := d.Set("group_id", groupID); err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	tflog.Debug(ctx, "Querying external group from GitHub API for import", map[string]any{
		"org_name":  orgName,
		"group_id":  groupID,
		"import_id": importID,
	})

	group, _, err := client.Teams.GetExternalGroup(ctx, orgName, int64(groupID))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Successfully retrieved external group from GitHub API for import", map[string]any{
		"org_name":   orgName,
		"group_id":   groupID,
		"team_count": len(group.Teams),
	})

	if len(group.Teams) != 1 {
		tflog.Info(ctx, "Multiple teams found for external group during import", map[string]any{
			"org_name":   orgName,
			"group_id":   groupID,
			"team_count": len(group.Teams),
			"import_id":  importID,
		})
		return nil, fmt.Errorf("could not get team_slug from %v number of teams", len(group.Teams))
	}

	teamSlug := *group.Teams[0].TeamName
	tflog.Trace(ctx, "Setting state attribute: team_slug", map[string]any{
		"team_slug": teamSlug,
	})
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
