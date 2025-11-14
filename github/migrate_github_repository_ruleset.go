package github

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func resourceGithubRepositoryRulesetMigrateState(v int, is *terraform.InstanceState, meta any) (*terraform.InstanceState, error) {
	switch v {
	case 1:
		// NOTE: This migration does not change any attributes in the state.
		// It is provided as a signal for the breaking changes in the underlying
		// go-github v67 to v77 upgrade while maintaining state compatibility.

		log.Printf("[INFO] Found GitHub Repository Ruleset State v1; migrating to v2")
		return migrateGithubRepositoryRulesetStateV1toV2(is)
	default:
		return is, fmt.Errorf("unexpected schema version: %d", v)
	}
}

func migrateGithubRepositoryRulesetStateV1toV2(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Printf("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] GitHub Repository Ruleset Attributes before migration: %#v", is.Attributes)

	// No actual attribute changes are needed for the v1 to v2 migration.
	// The breaking changes are in the go-github library structs (Ruleset to RepositoryRuleset)
	// and API method signatures, but the Terraform schema and state structure remain the same.

	log.Printf("[DEBUG] GitHub Repository Ruleset Attributes after State Migration: %#v", is.Attributes)

	return is, nil
}
