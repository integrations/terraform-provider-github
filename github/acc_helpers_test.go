package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	meta, err := getTestMeta()
	if err != nil {
		t.Fatalf("failed to get test meta: %v", err)
	}

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	req := &github.CustomProperty{
		PropertyName:  &name,
		ValueType:     github.PropertyValueType(valType),
		AllowedValues: allowed,
	}

	prop, _, err := meta.v3client.Organizations.CreateOrUpdateCustomProperty(t.Context(), meta.name, name, req)
	if err != nil {
		t.Fatalf("failed to create test organization repository custom property: %v", err)
	}

	t.Cleanup(func() {
		if _, err := meta.v3client.Organizations.RemoveCustomProperty(context.Background(), meta.name, name); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test organization repository custom property %s: %v", name, err)
		}
	})

	return prop
}

func mustCreateTestRepository(t *testing.T) *github.Repository {
	t.Helper()

	meta, err := getTestMeta()
	if err != nil {
		t.Fatalf("failed to get test meta: %v", err)
	}

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	req := &github.Repository{
		Name:     &name,
		AutoInit: new(true),
	}

	repo, _, err := meta.v3client.Repositories.Create(t.Context(), meta.name, req)
	if err != nil {
		t.Fatalf("failed to create test repository: %v", err)
	}

	t.Cleanup(func() {
		if _, err := meta.v3client.Repositories.Delete(context.Background(), meta.name, name); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test repository %s: %v", name, err)
		}
	})

	return repo
}
