package github

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func resourceGithubRepositoryMigrateState(v int, is *terraform.InstanceState, meta any) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		log.Printf("[INFO] Found GitHub Repository State v0; migrating to v1")
		return migrateGithubRepositoryStateV0toV1(is)
	default:
		return is, fmt.Errorf("unexpected schema version: %d", v)
	}
}

func migrateGithubRepositoryStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Printf("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] GitHub Repository Attributes before migration: %#v", is.Attributes)

	prefix := "branches."

	for k := range is.Attributes {
		if strings.HasPrefix(k, prefix) {
			delete(is.Attributes, k)
		}
	}

	log.Printf("[DEBUG] GitHub Repository Attributes after State Migration: %#v", is.Attributes)

	return is, nil
}
