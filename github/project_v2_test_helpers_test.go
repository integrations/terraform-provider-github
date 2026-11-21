package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shurcooL/githubv4"
)

type projectV2GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

func newProjectV2TestClient(t *testing.T, response func(request projectV2GraphQLRequest) string) (*githubv4.Client, *[]projectV2GraphQLRequest) {
	t.Helper()
	requests := make([]projectV2GraphQLRequest, 0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var payload projectV2GraphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("decoding GraphQL request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		requests = append(requests, payload)
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, response(payload)); err != nil {
			t.Errorf("writing GraphQL response: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	return githubv4.NewEnterpriseClient(server.URL, server.Client()), &requests
}

func projectV2GraphQLInput(t *testing.T, request projectV2GraphQLRequest) map[string]any {
	t.Helper()
	input, ok := request.Variables["input"].(map[string]any)
	if !ok {
		t.Fatalf("GraphQL request has no input object: %#v", request.Variables)
	}
	return input
}

func assertProjectV2GraphQLInput(t *testing.T, request projectV2GraphQLRequest, expected map[string]any) {
	t.Helper()
	input := projectV2GraphQLInput(t, request)
	for key, value := range expected {
		if got := input[key]; got != value {
			t.Fatalf("unexpected GraphQL input %s: got %#v, want %#v; input=%#v", key, got, value, input)
		}
	}
}
