package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func diffProjectV2Repository(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
	return diffRepositoryWithOwner(ctx, diff, meta, "repository_owner")
}

func diffProjectV2Team(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
	return diffTeamWithOrganization(ctx, diff, meta, "organization")
}

func diffProjectV2Owner(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
	if diff.Id() == "" || !diff.HasChanges("owner", "owner_type") {
		return nil
	}

	owner := projectV2OwnerLoginFromDiff(diff, meta)
	account, _, err := projectV2OwnerMetadata(meta).v3client.Users.Get(ctx, owner)
	if err != nil {
		return fmt.Errorf("querying Projects V2 owner %q: %w", owner, err)
	}

	requestedKind := projectV2OwnerKindFromDiff(diff, meta)
	actualKind := projectV2OwnerUser
	if account.GetType() == "Organization" {
		actualKind = projectV2OwnerOrganization
	}
	if requestedKind != actualKind {
		return fmt.Errorf("projects V2 owner %q is a %s, not a %s", owner, actualKind, requestedKind)
	}

	ownerID, ok := diff.GetOk("owner_id")
	if !ok || int64(ownerID.(int)) != account.GetID() {
		return forceNewChangedFields(diff, "owner", "owner_type")
	}
	return nil
}

func projectV2OwnerLoginFromDiff(diff *schema.ResourceDiff, meta any) string {
	if owner, ok := diff.Get("owner").(string); ok && owner != "" {
		return owner
	}
	return projectV2OwnerMetadata(meta).name
}

func projectV2OwnerKindFromDiff(diff *schema.ResourceDiff, meta any) string {
	if kind, ok := diff.Get("owner_type").(string); ok && kind != "" {
		return kind
	}
	return projectV2DefaultOwnerKind(meta)
}

func projectV2DefaultOwnerKind(meta any) string {
	if projectV2OwnerMetadata(meta).IsOrganization {
		return projectV2OwnerOrganization
	}
	return projectV2OwnerUser
}
