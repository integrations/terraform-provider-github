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

func mustGetOrganizationPublicKey(t *testing.T) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
	if err != nil {
		t.Fatalf("failed to get public key for test organization: %v", err)
	}

	return publicKey
}

type secretOrgRequestOptionsFunc func(*github.SecretOrgRequest)

func mustCreateTestOrganizationSecret(t *testing.T, opts ...secretOrgRequestOptionsFunc) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	secretName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
	if err != nil {
		t.Fatalf("failed to get public key for test organization secret: %v", err)
	}

	encryptedBytes, err := encryptPlaintext("test", publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test organization secret: %v", err)
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

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, secretName, req); err != nil {
		t.Fatalf("failed to create test organization secret: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Actions.DeleteOrgSecret(context.Background(), testAccConf.meta.name, secretName); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test organization secret %s: %v", secretName, err)
		}
	})

	return secretName
}

func mustUpdateOrganizationSecret(t *testing.T, name, value string) {
	t.Helper()

	publicKey := mustGetOrganizationPublicKey(t)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test organization secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, name, github.SecretOrgRequest{
		Visibility:     "all",
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to update test organization secret: %v", err)
	}
}

func mustCreateTestOrganizationVariable(t *testing.T, name, value *string) {
	t.Helper()

	var varName, varValue string

	if name != nil {
		varName = strings.ToUpper(*name)
	} else {
		randomID := acctest.RandString(testRandomIDLength)
		varName = strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))
	}

	if value != nil {
		varValue = *value
	} else {
		varValue = acctest.RandString(16)
	}

	if _, err := testAccConf.meta.v3client.Actions.CreateOrgVariable(t.Context(), testAccConf.meta.name, github.ActionsCreateOrgVariableRequest{
		Name:       varName,
		Visibility: "all",
		Value:      varValue,
	}); err != nil {
		t.Fatalf("failed to create test organization variable: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Actions.DeleteOrgVariable(context.Background(), testAccConf.meta.name, varName); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test organization variable %s: %v", varName, err)
		}
	})

	// return varName
}

func mustCreateTestRepositorySecret(t *testing.T, repo *github.Repository) string {
	t.Helper()

	ctx := t.Context()

	randomID := acctest.RandString(testRandomIDLength)
	secretName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetRepoPublicKey(ctx, testAccConf.meta.name, repo.GetName())
	if err != nil {
		t.Fatalf("failed to get public key for test repository secret: %v", err)
	}

	encryptedBytes, err := encryptPlaintext("test", publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateRepoSecret(ctx, testAccConf.meta.name, repo.GetName(), secretName, github.SecretRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to create test repository secret: %v", err)
	}

	return secretName
}

func mustGetTestRepositoryPublicKey(t *testing.T, repo *github.Repository) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetRepoPublicKey(t.Context(), testAccConf.meta.name, repo.GetName())
	if err != nil {
		t.Fatalf("failed to get public key for test repository: %v", err)
	}

	return publicKey
}

func mustUpdateTestRepositorySecret(t *testing.T, repo *github.Repository, name, value string) {
	t.Helper()

	publicKey := mustGetTestRepositoryPublicKey(t, repo)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateRepoSecret(t.Context(), testAccConf.meta.name, repo.GetName(), name, github.SecretRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to update test repository secret: %v", err)
	}
}

func mustCreateTestRepositoryVariable(t *testing.T, repo *github.Repository, name, value *string) string {
	t.Helper()

	var varName, varValue string

	if name != nil {
		varName = strings.ToUpper(*name)
	} else {
		randomID := acctest.RandString(testRandomIDLength)
		varName = strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))
	}

	if value != nil {
		varValue = *value
	} else {
		varValue = acctest.RandString(16)
	}

	if _, err := testAccConf.meta.v3client.Actions.CreateRepoVariable(t.Context(), testAccConf.meta.name, repo.GetName(), github.ActionsCreateVariableRequest{
		Name:  varName,
		Value: varValue,
	}); err != nil {
		t.Fatalf("failed to create test repository variable: %v", err)
	}

	return varName
}

func mustCreateTestRepositoryEnvironment(t *testing.T, repo *github.Repository, optArgs ...testCreateOptsFunc) *github.Environment {
	t.Helper()

	opts := &testCreateOpts{}
	for _, opt := range optArgs {
		opt(t, opts)
	}

	var name string
	if opts.name != nil {
		name = *opts.name
	} else {
		randomID := acctest.RandString(testRandomIDLength)
		name = fmt.Sprintf("%s%s", testResourcePrefix, randomID)
	}

	env, _, err := testAccConf.meta.v3client.Repositories.CreateUpdateEnvironment(t.Context(), testAccConf.meta.name, repo.GetName(), name, &github.CreateUpdateEnvironment{})
	if err != nil {
		t.Fatalf("failed to create test repository environment: %v", err)
	}

	return env
}

func mustGetTestRepositoryEnvironmentPublicKey(t *testing.T, repo *github.Repository, env *github.Environment) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetEnvPublicKey(t.Context(), testAccConf.meta.name, repo.GetName(), env.GetName())
	if err != nil {
		t.Fatalf("failed to get public key for test repository environment: %v", err)
	}

	return publicKey
}

func mustCreateTestRepositoryEnvironmentSecret(t *testing.T, repo *github.Repository, env *github.Environment, value string) string {
	t.Helper()

	ctx := t.Context()

	randomID := acctest.RandString(testRandomIDLength)
	secretName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	publicKey := mustGetTestRepositoryEnvironmentPublicKey(t, repo, env)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository environment secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateEnvSecret(ctx, testAccConf.meta.name, repo.GetName(), env.GetName(), secretName, github.SecretRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to create test repository environment secret: %v", err)
	}

	return secretName
}

func mustUpdateTestRepositoryEnvironmentSecret(t *testing.T, repo *github.Repository, env *github.Environment, name, value string) {
	t.Helper()

	ctx := t.Context()

	publicKey := mustGetTestRepositoryEnvironmentPublicKey(t, repo, env)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository environment secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateEnvSecret(ctx, testAccConf.meta.name, repo.GetName(), env.GetName(), name, github.SecretRequest{
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to update test repository environment secret: %v", err)
	}
}

func mustCreateTestRepositoryEnvironmentVariable(t *testing.T, repo *github.Repository, env *github.Environment, name, value *string) string {
	t.Helper()

	var varName string
	var varValue string

	if name != nil {
		varName = strings.ToUpper(*name)
	} else {
		randomID := acctest.RandString(testRandomIDLength)
		varName = strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))
	}

	if value != nil {
		varValue = *value
	} else {
		varValue = acctest.RandString(16)
	}

	if _, err := testAccConf.meta.v3client.Actions.CreateEnvVariable(t.Context(), testAccConf.meta.name, repo.GetName(), env.GetName(), github.ActionsCreateVariableRequest{
		Name:  varName,
		Value: varValue,
	}); err != nil {
		t.Fatalf("failed to create test repository environment variable: %v", err)
	}

	return varName
}
