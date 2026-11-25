package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-github/v92/github"
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

	role, _, err := testAccConf.meta.v3client.Organizations.CreateCustomRepoRole(t.Context(), testAccConf.meta.name, github.CreateCustomRepoRoleRequest{
		Name:        name,
		Description: new("Test organization repository role."),
		BaseRole:    "read",
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

func mustAddRepositoryCollaborator(t *testing.T, repo *github.Repository, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Repositories.AddCollaborator(t.Context(), testAccConf.meta.name, repo.GetName(), username, &github.RepositoryAddCollaboratorOptions{Permission: "push"})
	if err != nil {
		t.Fatalf("failed to add collaborator %s to test repository %s: %v", username, repo.GetName(), err)
	}
}
