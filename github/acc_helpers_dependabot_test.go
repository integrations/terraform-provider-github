package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func mustGetDependabotOrganizationPublicKey(t *testing.T) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Dependabot.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
	if err != nil {
		t.Fatalf("failed to get public key for test organization dependabot: %v", err)
	}

	return publicKey
}

func mustCreateTestDependabotOrganizationSecret(t *testing.T, opts ...secretOrgRequestOptionsFunc) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	secretName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	publicKey, _, err := testAccConf.meta.v3client.Dependabot.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
	if err != nil {
		t.Fatalf("failed to get public key for test dependabot  organization secret: %v", err)
	}

	encryptedBytes, err := encryptPlaintext("test", publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test dependabot  organization secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	req := github.SecretOrgRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
		Visibility:     "all",
	}

	for _, fn := range opts {
		if fn != nil {
			fn(&req)
		}
	}

	if _, err := testAccConf.meta.v3client.Dependabot.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, secretName, req); err != nil {
		t.Fatalf("failed to create test dependabot organization secret: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Dependabot.DeleteOrgSecret(context.Background(), testAccConf.meta.name, secretName); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test dependabot organization secret %s: %v", secretName, err)
		}
	})

	return secretName
}

func mustUpdateDependabotOrganizationSecret(t *testing.T, name, value string) {
	t.Helper()

	publicKey := mustGetDependabotOrganizationPublicKey(t)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test organization dependabot secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Dependabot.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, name, github.SecretOrgRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
		Visibility:     "all",
	}); err != nil {
		t.Fatalf("failed to update test organization dependabot secret: %v", err)
	}
}

func mustGetDependabotRepositorPublicKey(t *testing.T, repo *github.Repository) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Dependabot.GetRepoPublicKey(t.Context(), testAccConf.meta.name, repo.GetName())
	if err != nil {
		t.Fatalf("failed to get public key for test repository dependabot: %v", err)
	}

	return publicKey
}

// func mustCreateTestDependabotRepositorySecret(t *testing.T, repo *github.Repository, name, value string) string {
// 	t.Helper()

// 	publicKey := mustGetRepositoryDependabotPublicKey(t, repo)

// 	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
// 	if err != nil {
// 		t.Fatalf("failed to encrypt plaintext for test repository dependabot secret: %v", err)
// 	}
// 	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

// 	if _, err := testAccConf.meta.v3client.Dependabot.CreateOrUpdateRepoSecret(t.Context(), testAccConf.meta.name, repo.GetName(), &github.DependabotEncryptedSecret{
// 		Name:           name,
// 		KeyID:          publicKey.GetKeyID(),
// 		EncryptedValue: encryptedValue,
// 	}); err != nil {
// 		t.Fatalf("failed to create test repository dependabot secret: %v", err)
// 	}

// 	return name
// }

func mustUpdateDependabotRepositorySecret(t *testing.T, repo *github.Repository, name, value string) {
	t.Helper()

	publicKey := mustGetDependabotRepositorPublicKey(t, repo)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository dependabot secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Dependabot.CreateOrUpdateRepoSecret(t.Context(), testAccConf.meta.name, repo.GetName(), name, github.SecretRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to update test repository dependabot secret: %v", err)
	}
}
