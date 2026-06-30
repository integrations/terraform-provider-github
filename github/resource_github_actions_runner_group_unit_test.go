package github

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-github/v88/github"
)

func TestGetOrganizationRunnerGroup_ReturnsNilOn304(t *testing.T) {
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/orgs/test-org/actions/runner-groups/123",
			ExpectedMethod: "GET",
			ExpectedHeaders: map[string]string{
				"If-None-Match": "etag-abc",
			},
			ResponseBody: `{"message":"Not Modified"}`,
			StatusCode:   http.StatusNotModified,
		},
	})
	defer ts.Close()

	httpClient := &http.Client{Transport: NewEtagTransport(http.DefaultTransport)}
	baseURL := ts.URL + "/"
	client, err := github.NewClient(github.WithHTTPClient(httpClient), github.WithURLs(&baseURL, nil))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), ctxEtag, "etag-abc")
	runnerGroup, resp, err := getOrganizationRunnerGroup(client, ctx, "test-org", 123)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if runnerGroup != nil {
		t.Fatalf("expected nil runner group on 304, got: %+v", runnerGroup)
	}
	if resp == nil || resp.StatusCode != http.StatusNotModified {
		t.Fatalf("expected 304 response, got: %+v", resp)
	}
}

func TestGetOrganizationRunnerGroup_ReturnsRunnerGroup(t *testing.T) {
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/orgs/test-org/actions/runner-groups/42",
			ExpectedMethod: "GET",
			ResponseBody:   `{"id":42,"name":"my-group","network_configuration_id":"nc-456"}`,
			StatusCode:     http.StatusOK,
		},
	})
	defer ts.Close()

	httpClient := http.DefaultClient
	baseURL := ts.URL + "/"
	client, err := github.NewClient(github.WithHTTPClient(httpClient), github.WithURLs(&baseURL, nil))
	if err != nil {
		t.Fatal(err)
	}

	runnerGroup, resp, err := getOrganizationRunnerGroup(client, context.Background(), "test-org", 42)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if runnerGroup == nil {
		t.Fatal("expected non-nil runner group")
	}
	if runnerGroup.GetID() != 42 {
		t.Fatalf("expected ID 42, got %d", runnerGroup.GetID())
	}
	if runnerGroup.GetName() != "my-group" {
		t.Fatalf("expected name 'my-group', got %q", runnerGroup.GetName())
	}
	if runnerGroup.GetNetworkConfigurationID() != "nc-456" {
		t.Fatalf("expected network_configuration_id 'nc-456', got %q", runnerGroup.GetNetworkConfigurationID())
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 response, got: %+v", resp)
	}
}

func TestGetEnterpriseRunnerGroup_ReturnsNilOn304(t *testing.T) {
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/enterprises/test-ent/actions/runner-groups/99",
			ExpectedMethod: "GET",
			ExpectedHeaders: map[string]string{
				"If-None-Match": "etag-xyz",
			},
			ResponseBody: `{"message":"Not Modified"}`,
			StatusCode:   http.StatusNotModified,
		},
	})
	defer ts.Close()

	httpClient := &http.Client{Transport: NewEtagTransport(http.DefaultTransport)}
	baseURL := ts.URL + "/"
	client, err := github.NewClient(github.WithHTTPClient(httpClient), github.WithURLs(&baseURL, nil))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), ctxEtag, "etag-xyz")
	runnerGroup, resp, err := getEnterpriseRunnerGroup(client, ctx, "test-ent", 99)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if runnerGroup != nil {
		t.Fatalf("expected nil runner group on 304, got: %+v", runnerGroup)
	}
	if resp == nil || resp.StatusCode != http.StatusNotModified {
		t.Fatalf("expected 304 response, got: %+v", resp)
	}
}
