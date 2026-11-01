package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/google/go-github/v92/github"
)

func Test_getFileCommit(t *testing.T) {
	t.Parallel()

	commit := func(sha, message string, files ...*github.CommitFile) *github.RepositoryCommit {
		return &github.RepositoryCommit{SHA: new(sha), Commit: &github.Commit{Message: new(message)}, Files: files}
	}
	changed := &github.CommitFile{Filename: new("main.yaml"), Status: new("modified")}
	removed := &github.CommitFile{Filename: new("main.yaml"), Status: new("removed")}

	for _, d := range []struct {
		testName string
		// The file's history, newest first, one slice per page.
		pages    [][]*github.RepositoryCommit
		wantSHA  string
		wantErr  bool
		wantReqs []string
	}{
		{
			testName: "stops_at_the_first_match",
			pages: [][]*github.RepositoryCommit{
				{commit("m", "Merge branch 'feature'"), commit("a", "Change the file", changed)},
				{commit("b", "Create the file", changed)},
			},
			wantSHA:  "a",
			wantReqs: []string{"list 1", "get a"},
		},
		{
			testName: "reads_the_next_page_without_a_match",
			pages: [][]*github.RepositoryCommit{
				{commit("m", "Merge branch 'feature'"), commit("r", "Remove the file", removed)},
				{commit("b", "Create the file", changed)},
			},
			wantSHA:  "b",
			wantReqs: []string{"list 1", "get r", "list 2", "get b"},
		},
		{
			testName: "fails_without_a_matching_commit",
			pages: [][]*github.RepositoryCommit{
				{commit("m", "Merge branch 'feature'")},
			},
			wantErr:  true,
			wantReqs: []string{"list 1"},
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			details := map[string]*github.RepositoryCommit{}
			for _, page := range d.pages {
				for _, c := range page {
					details[c.GetSHA()] = c
				}
			}

			var mu sync.Mutex
			var requests []string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				defer mu.Unlock()

				if sha, ok := strings.CutPrefix(r.URL.Path, "/repos/owner/repo/commits/"); ok {
					requests = append(requests, "get "+sha)
					_ = json.NewEncoder(w).Encode(details[sha])
					return
				}

				page := 1
				if p := r.URL.Query().Get("page"); p != "" {
					page, _ = strconv.Atoi(p)
				}
				requests = append(requests, fmt.Sprintf("list %d", page))
				if page < len(d.pages) {
					w.Header().Set("Link", fmt.Sprintf(`<http://%s/repos/owner/repo/commits?page=%d>; rel="next"`, r.Host, page+1))
				}
				// A commit list carries no files.
				var list []*github.RepositoryCommit
				for _, c := range d.pages[page-1] {
					list = append(list, &github.RepositoryCommit{SHA: c.SHA, Commit: c.Commit})
				}
				_ = json.NewEncoder(w).Encode(list)
			}))
			t.Cleanup(ts.Close)

			client := mustCreateTestGitHubClient(t, ts.URL+"/")

			got, err := getFileCommit(t.Context(), client, "owner", "repo", "main.yaml", "main")

			if d.wantErr && err == nil {
				t.Fatalf("expected an error but got commit %q", got.GetSHA())
			}
			if !d.wantErr && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if got.GetSHA() != d.wantSHA {
				t.Fatalf("expected commit %q but got %q", d.wantSHA, got.GetSHA())
			}
			if !slices.Equal(requests, d.wantReqs) {
				t.Fatalf("expected requests %q but got %q", d.wantReqs, requests)
			}
		})
	}
}
