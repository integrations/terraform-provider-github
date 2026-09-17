package github

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/v89/github"
)

func TestGetEnterpriseCostCenterPaginatesResources(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/v3/enterprises/example/settings/billing/cost-centers/123" {
			t.Fatalf("unexpected request path %q", req.URL.Path)
		}
		if got := req.URL.Query().Get("per_page"); got != "100" {
			t.Fatalf("expected per_page 100, got %q", got)
		}

		switch page := req.URL.Query().Get("page"); page {
		case "1":
			fmt.Fprint(w, `{
				"id": "123",
				"name": "Engineering",
				"state": "active",
				"resources": [
					{"type": "User", "name": "octocat"},
					{"type": "Org", "name": "octo-org"}
				],
				"has_next_page": true
			}`)
		case "2":
			fmt.Fprint(w, `{
				"id": "123",
				"name": "Engineering",
				"state": "active",
				"resources": [
					{"type": "Repo", "name": "octo-org/example"}
				],
				"has_next_page": false
			}`)
		default:
			t.Fatalf("unexpected page %q", page)
		}
	}))
	t.Cleanup(server.Close)

	client, err := github.NewClient(github.WithEnterpriseURLs(server.URL+"/", server.URL+"/"))
	if err != nil {
		t.Fatalf("configuring GitHub client: %v", err)
	}

	costCenter, err := getEnterpriseCostCenter(context.Background(), client, "example", "123")
	if err != nil {
		t.Fatalf("getting cost center: %v", err)
	}
	if costCenter.GetID() != "123" {
		t.Fatalf("expected cost center ID 123, got %q", costCenter.GetID())
	}
	if len(costCenter.Resources) != 3 {
		t.Fatalf("expected 3 resources, got %d", len(costCenter.Resources))
	}
	if costCenter.Resources[2].GetName() != "octo-org/example" {
		t.Fatalf("expected second-page resource, got %q", costCenter.Resources[2].GetName())
	}
}
