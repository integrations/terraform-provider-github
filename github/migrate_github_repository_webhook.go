package github

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func resourceGithubWebhookMigrateState(v int, is *terraform.InstanceState, meta any) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		log.Printf("[INFO] Found GitHub Webhook State v0; migrating to v1")
		return migrateGithubWebhookStateV0toV1(is)
	default:
		return is, fmt.Errorf("unexpected schema version: %d", v)
	}
}

func migrateGithubWebhookStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Printf("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] GitHub Webhook Attributes before migration: %#v", is.Attributes)

	prefix := "configuration."

	delete(is.Attributes, prefix+"%")

	// Read & delete old keys
	oldKeys := make(map[string]string)
	for k, v := range is.Attributes {
		if strings.HasPrefix(k, prefix) {
			oldKeys[k] = v

			// Delete old keys
			delete(is.Attributes, k)
		}
	}

	// Write new keys
	for k, v := range oldKeys {
		newKey := "configuration.0." + strings.TrimPrefix(k, prefix)
		is.Attributes[newKey] = v
	}

	is.Attributes[prefix+"#"] = "1"
	log.Printf("[DEBUG] GitHub Webhook Attributes after State Migration: %#v", is.Attributes)

	return is, nil
}
