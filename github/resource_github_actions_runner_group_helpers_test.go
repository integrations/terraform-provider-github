package github

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetRunnerGroupNetworking(t *testing.T) {
	t.Run("returns networking payload", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/orgs/test/actions/runner-groups/123",
				ExpectedMethod: "GET",
				ResponseBody:   `{"network_configuration_id":"network-123"}`,
				StatusCode:     http.StatusOK,
			},
		})
		defer ts.Close()

		httpClient := http.DefaultClient
		client := github.NewClient(httpClient)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		runnerGroup, resp, err := getRunnerGroupNetworking(client, context.Background(), "orgs/test/actions/runner-groups/123")
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil || resp.StatusCode != http.StatusOK {
			t.Fatalf("expected %d response, got %#v", http.StatusOK, resp)
		}
		if runnerGroup == nil || runnerGroup.NetworkConfigurationID == nil {
			t.Fatalf("expected network configuration payload, got %#v", runnerGroup)
		}
		if got := *runnerGroup.NetworkConfigurationID; got != "network-123" {
			t.Fatalf("expected network configuration id %q, got %q", "network-123", got)
		}
	})

	t.Run("swallows 304 not modified", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/orgs/test/actions/runner-groups/123",
				ExpectedMethod: "GET",
				ExpectedHeaders: map[string]string{
					"If-None-Match": "etag-123",
				},
				ResponseBody: `{"message":"Not modified"}`,
				StatusCode:   http.StatusNotModified,
			},
		})
		defer ts.Close()

		httpClient := http.DefaultClient
		httpClient.Transport = NewEtagTransport(http.DefaultTransport)
		client := github.NewClient(httpClient)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		ctx := context.WithValue(context.Background(), ctxEtag, "etag-123")
		runnerGroup, resp, err := getRunnerGroupNetworking(client, ctx, "orgs/test/actions/runner-groups/123")
		if err != nil {
			t.Fatal(err)
		}
		if runnerGroup != nil {
			t.Fatalf("expected nil runner group on 304, got %#v", runnerGroup)
		}
		if resp == nil || resp.StatusCode != http.StatusNotModified {
			t.Fatalf("expected %d response, got %#v", http.StatusNotModified, resp)
		}
	})
}

func TestUpdateRunnerGroupNetworking(t *testing.T) {
	t.Run("sends network configuration id payload", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/orgs/test/actions/runner-groups/123",
				ExpectedMethod: "PATCH",
				ExpectedBody:   []byte("{\"network_configuration_id\":\"network-123\"}\n"),
				ResponseBody:   `{}`,
				StatusCode:     http.StatusNoContent,
			},
		})
		defer ts.Close()

		httpClient := http.DefaultClient
		client := github.NewClient(httpClient)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		networkConfigurationID := "network-123"
		resp, err := updateRunnerGroupNetworking(client, context.Background(), "orgs/test/actions/runner-groups/123", &networkConfigurationID)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil || resp.StatusCode != http.StatusNoContent {
			t.Fatalf("expected %d response, got %#v", http.StatusNoContent, resp)
		}
	})

	t.Run("sends null payload when removing networking", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/orgs/test/actions/runner-groups/123",
				ExpectedMethod: "PATCH",
				ExpectedBody:   []byte("{\"network_configuration_id\":null}\n"),
				ResponseBody:   `{}`,
				StatusCode:     http.StatusNoContent,
			},
		})
		defer ts.Close()

		httpClient := http.DefaultClient
		client := github.NewClient(httpClient)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		resp, err := updateRunnerGroupNetworking(client, context.Background(), "orgs/test/actions/runner-groups/123", nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil || resp.StatusCode != http.StatusNoContent {
			t.Fatalf("expected %d response, got %#v", http.StatusNoContent, resp)
		}
	})
}

func TestSetRunnerGroupNetworkingState(t *testing.T) {
	t.Run("sets network configuration id", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
			"network_configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}, map[string]any{})

		networkConfigurationID := "network-123"
		if err := setRunnerGroupNetworkingState(d, &runnerGroupNetworking{NetworkConfigurationID: &networkConfigurationID}); err != nil {
			t.Fatal(err)
		}

		got, ok := d.GetOk("network_configuration_id")
		if !ok {
			t.Fatal("expected network_configuration_id to be set")
		}
		if got.(string) != networkConfigurationID {
			t.Fatalf("expected network configuration id %q, got %q", networkConfigurationID, got.(string))
		}
	})

	t.Run("clears network configuration id", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
			"network_configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}, map[string]any{"network_configuration_id": "network-123"})

		if err := setRunnerGroupNetworkingState(d, nil); err != nil {
			t.Fatal(err)
		}

		if _, ok := d.GetOk("network_configuration_id"); ok {
			t.Fatalf("expected network_configuration_id to be cleared, got %q", d.Get("network_configuration_id"))
		}
	})
}
