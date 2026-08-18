package github

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

const testRandomIDLength = 5

type testCreateOpts struct {
	name *string
}

type testCreateOptsFunc func(t *testing.T, opts *testCreateOpts)

func withTestCreateName(name string) testCreateOptsFunc {
	return func(t *testing.T, opts *testCreateOpts) {
		opts.name = &name
	}
}

func mustGetTestMockResponse(t *testing.T, uri string, statusCode int, body any) *mockResponse {
	resp := &mockResponse{
		ExpectedUri: uri,
		StatusCode:  statusCode,
	}

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("failed to marshal mock response body: %v", err)
		}
		resp.ResponseBody = string(bodyBytes)
	}

	return resp
}

func mustCreateTestGitHubClient(t *testing.T, baseURL string, opts ...github.ClientOptionsFunc) *github.Client {
	client, err := github.NewClient(append([]github.ClientOptionsFunc{github.WithURLs(&baseURL, nil)}, opts...)...)
	if err != nil {
		t.Fatalf("failed to create GitHub client: %s", err)
	}
	return client
}

func mustGetOrganizationRole(t *testing.T, roleID int64) *github.CustomOrgRole {
	t.Helper()

	role, _, err := testAccConf.meta.v3client.Organizations.GetOrgRole(t.Context(), testAccConf.meta.name, roleID)
	if err != nil {
		t.Fatalf("failed to get test organization role: %v", err)
	}
	return role
}

func mustAssignOrganizationRoleToUser(t *testing.T, username string, roleID int64) {
	t.Helper()

	_, err := testAccConf.meta.v3client.Organizations.AssignOrgRoleToUser(t.Context(), testAccConf.meta.name, username, roleID)
	if err != nil {
		t.Fatalf("failed to add user %s to test organization role %d: %v", username, roleID, err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Organizations.RemoveOrgRoleFromUser(context.Background(), testAccConf.meta.name, username, roleID); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to remove user %s from test organization role %d: %v", username, roleID, err)
		}
	})
}

func mustCreateTestOrganizationRepositoryRole(t *testing.T) *github.CustomRepoRoles {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	role, _, err := testAccConf.meta.v3client.Organizations.CreateCustomRepoRole(t.Context(), testAccConf.meta.name, &github.CreateOrUpdateCustomRepoRoleOptions{
		Name:        &name,
		Description: new("Test organization repository role."),
		BaseRole:    new("read"),
		Permissions: []string{"reopen_issue"},
	})
	if err != nil {
		t.Fatalf("failed to create test organization repository role: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Organizations.DeleteCustomRepoRole(context.Background(), testAccConf.meta.name, role.GetID()); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test organization repository role %s: %v", name, err)
		}
	})

	return role
}

func mustCreateTestOrganizationRepositoryCustomProperty(t *testing.T, valType string, allowed []string) *github.CustomProperty {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	req := &github.CustomProperty{
		PropertyName:  &name,
		ValueType:     github.PropertyValueType(valType),
		AllowedValues: allowed,
	}

	prop, _, err := testAccConf.meta.v3client.Organizations.CreateOrUpdateCustomProperty(t.Context(), testAccConf.meta.name, name, req)
	if err != nil {
		t.Fatalf("failed to create test organization repository custom property: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Organizations.RemoveCustomProperty(context.Background(), testAccConf.meta.name, name); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test organization repository custom property %s: %v", name, err)
		}
	})

	return prop
}

func mustGetUser(t *testing.T, username string) *github.User {
	t.Helper()

	user, _, err := testAccConf.meta.v3client.Users.Get(t.Context(), username)
	if err != nil {
		t.Fatalf("failed to get user %s: %v", username, err)
	}

	return user
}

func mustCreateTestBranch(t *testing.T, repo *github.Repository) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	// Get the SHA of the default branch
	defaultBranch, _, err := testAccConf.meta.v3client.Repositories.GetBranch(t.Context(), testAccConf.meta.name, repo.GetName(), repo.GetDefaultBranch(), 0)
	if err != nil {
		t.Fatalf("failed to get default branch for test repository %s: %v", repo.GetName(), err)
	}

	req := github.CreateRef{
		Ref: fmt.Sprintf("refs/heads/%s", name),
		SHA: defaultBranch.GetCommit().GetSHA(),
	}
	if _, _, err := testAccConf.meta.v3client.Git.CreateRef(t.Context(), testAccConf.meta.name, repo.GetName(), req); err != nil {
		t.Fatalf("failed to create test branch %s for repository %s: %v", name, repo.GetName(), err)
	}

	return name
}

func mustAddRepositoryCollaborator(t *testing.T, repo *github.Repository, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Repositories.AddCollaborator(t.Context(), testAccConf.meta.name, repo.GetName(), username, &github.RepositoryAddCollaboratorOptions{Permission: "push"})
	if err != nil {
		t.Fatalf("failed to add collaborator %s to test repository %s: %v", username, repo.GetName(), err)
	}
}

func mustGetOrganizationPublicKey(t *testing.T) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
	if err != nil {
		t.Fatalf("failed to get public key for test organization: %v", err)
	}

	return publicKey
}

func mustCreateTestOrganizationSecret(t *testing.T) string {
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

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, secretName, github.OrgSecretRequest{
		Visibility:     "all",
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
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

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, name, github.OrgSecretRequest{
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

	if _, err := testAccConf.meta.v3client.Actions.CreateOrgVariable(t.Context(), testAccConf.meta.name, github.OrgActionsVariableCreateRequest{
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

	if _, err := testAccConf.meta.v3client.Actions.CreateRepoVariable(t.Context(), testAccConf.meta.name, repo.GetName(), github.ActionsVariableCreateRequest{
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

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetEnvPublicKey(t.Context(), testAccConf.meta.name, repo.GetName(), url.PathEscape(env.GetName()))
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

	if _, err := testAccConf.meta.v3client.Actions.CreateEnvVariable(t.Context(), testAccConf.meta.name, repo.GetName(), url.PathEscape(env.GetName()), github.ActionsVariableCreateRequest{
		Name:  varName,
		Value: varValue,
	}); err != nil {
		t.Fatalf("failed to create test repository environment variable: %v", err)
	}

	return varName
}

func mustGetOrganizationDependabotPublicKey(t *testing.T) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Dependabot.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
	if err != nil {
		t.Fatalf("failed to get public key for test organization dependabot: %v", err)
	}

	return publicKey
}

func mustUpdateOrganizationDependabotSecret(t *testing.T, name, value string) {
	t.Helper()

	publicKey := mustGetOrganizationDependabotPublicKey(t)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test organization dependabot secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Dependabot.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, &github.DependabotEncryptedSecret{
		Name:           name,
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
		Visibility:     "all",
	}); err != nil {
		t.Fatalf("failed to update test organization dependabot secret: %v", err)
	}
}

func mustGetRepositoryDependabotPublicKey(t *testing.T, repo *github.Repository) *github.PublicKey {
	t.Helper()

	publicKey, _, err := testAccConf.meta.v3client.Dependabot.GetRepoPublicKey(t.Context(), testAccConf.meta.name, repo.GetName())
	if err != nil {
		t.Fatalf("failed to get public key for test repository dependabot: %v", err)
	}

	return publicKey
}

// func mustCreateTestRepositoryDependabotSecret(t *testing.T, repo *github.Repository, name, value string) string {
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

func mustUpdateRepositoryDependabotSecret(t *testing.T, repo *github.Repository, name, value string) {
	t.Helper()

	publicKey := mustGetRepositoryDependabotPublicKey(t, repo)

	encryptedBytes, err := encryptPlaintext(value, publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository dependabot secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Dependabot.CreateOrUpdateRepoSecret(t.Context(), testAccConf.meta.name, repo.GetName(), &github.DependabotEncryptedSecret{
		Name:           name,
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to update test repository dependabot secret: %v", err)
	}
}
