package github

import (
	"context"
	"encoding/base64"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Agent secrets have their own keys; Actions keys cannot be used here.
func agentSecretValue(ctx context.Context, d *schema.ResourceData, meta *Owner, repository string) (github.SecretRequest, error) {
	if value, ok := d.GetOk("value_encrypted"); ok {
		keyID, _ := d.Get("key_id").(string)
		encryptedValue, _ := value.(string)
		return github.SecretRequest{KeyID: keyID, EncryptedValue: encryptedValue}, nil
	}

	var key *github.PublicKey
	var err error
	if repository == "" {
		key, _, err = meta.v3client.Agents.GetOrgPublicKey(ctx, meta.name)
	} else {
		key, _, err = meta.v3client.Agents.GetRepoPublicKey(ctx, meta.name, repository)
	}
	if err != nil {
		return github.SecretRequest{}, err
	}
	plaintext, _ := d.Get("value").(string)
	value, err := encryptPlaintext(plaintext, key.GetKey())
	if err != nil {
		return github.SecretRequest{}, err
	}
	return github.SecretRequest{KeyID: key.GetKeyID(), EncryptedValue: base64.StdEncoding.EncodeToString(value)}, nil
}

func setAgentSecretTimestamps(d *schema.ResourceData, secret *github.Secret, written bool) error {
	if err := d.Set("created_at", secret.CreatedAt.String()); err != nil {
		return err
	}
	// Preserve the last provider write so diffSecret can detect external updates.
	updatedAt, _ := d.Get("updated_at").(string)
	if written || updatedAt == "" {
		if err := d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return err
		}
	}
	return d.Set("remote_updated_at", secret.UpdatedAt.String())
}
