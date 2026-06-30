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

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

const testRandomIDLength = 5

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

func mustCreateTestTeam(t *testing.T, parentID *int64) *github.Team {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	team, _, err := testAccConf.meta.v3client.Teams.CreateTeam(t.Context(), testAccConf.meta.name, github.NewTeam{Name: name, ParentTeamID: parentID, Privacy: new("closed")})
	if err != nil {
		t.Fatalf("failed to create test team: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Teams.DeleteTeamByID(context.Background(), testAccConf.meta.id, team.GetID()); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test team %s: %v", name, err)
		}
	})

	return team
}

func mustRenameTestTeam(t *testing.T, team *github.Team, newName string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Teams.EditTeamBySlug(t.Context(), testAccConf.meta.name, team.GetSlug(), github.NewTeam{Name: newName}, false)
	if err != nil {
		t.Fatalf("failed to rename test team %s to %s: %v", team.GetName(), newName, err)
	}
}

func mustDeleteTestTeam(t *testing.T, team *github.Team) {
	t.Helper()

	if _, err := testAccConf.meta.v3client.Teams.DeleteTeamBySlug(context.Background(), testAccConf.meta.name, team.GetSlug()); err != nil {
		if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
			return
		}
		t.Fatalf("failed to delete test team %s: %v", team.GetName(), err)
	}
}

func mustAddTeamMember(t *testing.T, team *github.Team, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Teams.AddTeamMembershipBySlug(t.Context(), testAccConf.meta.name, team.GetSlug(), username, &github.TeamAddTeamMembershipOptions{Role: "member"})
	if err != nil {
		t.Fatalf("failed to add member %s to test team %s: %v", username, team.GetName(), err)
	}
}

func mustCreateTestRepository(t *testing.T) *github.Repository {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	req := &github.Repository{
		Name:     &name,
		AutoInit: new(true),
	}

	repo, _, err := testAccConf.meta.v3client.Repositories.Create(t.Context(), testAccConf.meta.name, req)
	if err != nil {
		t.Fatalf("failed to create test repository: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Repositories.Delete(context.Background(), testAccConf.meta.name, name); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test repository %s: %v", name, err)
		}
	})

	return repo
}

func mustAddRepositoryCollaborator(t *testing.T, repo *github.Repository, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Repositories.AddCollaborator(t.Context(), testAccConf.meta.name, repo.GetName(), username, &github.RepositoryAddCollaboratorOptions{Permission: "push"})
	if err != nil {
		t.Fatalf("failed to add collaborator %s to test repository %s: %v", username, repo.GetName(), err)
	}
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

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, &github.EncryptedSecret{
		Name:           secretName,
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

func mustCreateTestOrganizationVariable(t *testing.T, value string) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	if _, err := testAccConf.meta.v3client.Actions.CreateOrgVariable(t.Context(), testAccConf.meta.name, &github.ActionsVariable{
		Name:       varName,
		Visibility: new("all"),
		Value:      value,
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

	return varName
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

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateRepoSecret(ctx, testAccConf.meta.name, repo.GetName(), &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to create test repository secret: %v", err)
	}

	return secretName
}

func mustCreateTestRepositoryVariable(t *testing.T, repo *github.Repository, value string) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	if _, err := testAccConf.meta.v3client.Actions.CreateRepoVariable(t.Context(), testAccConf.meta.name, repo.GetName(), &github.ActionsVariable{
		Name:  varName,
		Value: value,
	}); err != nil {
		t.Fatalf("failed to create test repository variable: %v", err)
	}

	return varName
}

func mustCreateTestRepositoryEnvironment(t *testing.T, repo *github.Repository) *github.Environment {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	env, _, err := testAccConf.meta.v3client.Repositories.CreateUpdateEnvironment(t.Context(), testAccConf.meta.name, repo.GetName(), name, &github.CreateUpdateEnvironment{})
	if err != nil {
		t.Fatalf("failed to create test repository environment: %v", err)
	}

	return env
}

func mustCreateTestRepositoryEnvironmentSecret(t *testing.T, repo *github.Repository, env *github.Environment) string {
	t.Helper()

	ctx := t.Context()

	randomID := acctest.RandString(testRandomIDLength)
	secretName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	publicKey, _, err := testAccConf.meta.v3client.Actions.GetEnvPublicKey(ctx, int(repo.GetID()), url.PathEscape(env.GetName()))
	if err != nil {
		t.Fatalf("failed to get public key for test repository environment secret: %v", err)
	}

	encryptedBytes, err := encryptPlaintext("test", publicKey.GetKey())
	if err != nil {
		t.Fatalf("failed to encrypt plaintext for test repository environment secret: %v", err)
	}
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	if _, err := testAccConf.meta.v3client.Actions.CreateOrUpdateEnvSecret(ctx, int(repo.GetID()), env.GetName(), &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedValue,
	}); err != nil {
		t.Fatalf("failed to create test repository environment secret: %v", err)
	}

	return secretName
}

func mustCreateTestRepositoryEnvironmentVariable(t *testing.T, repo *github.Repository, env *github.Environment, value string) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

	if _, err := testAccConf.meta.v3client.Actions.CreateEnvVariable(t.Context(), testAccConf.meta.name, repo.GetName(), url.PathEscape(env.GetName()), &github.ActionsVariable{
		Name:  varName,
		Value: value,
	}); err != nil {
		t.Fatalf("failed to create test repository environment variable: %v", err)
	}

	return varName
}
