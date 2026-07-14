package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shurcooL/githubv4"
)

func newProjectV2TestClient(t *testing.T, response func(query string) string) (*githubv4.Client, *[]string) {
	t.Helper()
	requests := make([]string, 0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var payload struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("decoding GraphQL request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		requests = append(requests, payload.Query)
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, response(payload.Query)); err != nil {
			t.Errorf("writing GraphQL response: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	return githubv4.NewEnterpriseClient(server.URL, server.Client()), &requests
}
