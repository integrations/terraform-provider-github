package github

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func resourceGithubActionsSecretMigrateState(v int, is *terraform.InstanceState, meta any) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		log.Printf("[INFO] Found GitHub Actions Secret State v0; migrating to v1")
		return migrateGithubActionsSecretStateV0toV1(is)
	default:
		return is, fmt.Errorf("unexpected schema version: %d", v)
	}
}

func migrateGithubActionsSecretStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Printf("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] GitHub Actions Secret Attributes before migration: %#v", is.Attributes)

	// Add the destroy_on_drift field with default value true if it doesn't exist
	if _, ok := is.Attributes["destroy_on_drift"]; !ok {
		is.Attributes["destroy_on_drift"] = "true"
	}

	log.Printf("[DEBUG] GitHub Actions Secret Attributes after State Migration: %#v", is.Attributes)

	return is, nil
}
