package github

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_getFileCommit_NoCommitContainsFile locks in the errFileCommitNotFound
// sentinel. resourceGithubRepositoryFileRead relies on it to remove the resource
// from state instead of returning a hard error, so unwrapping this error would
// silently reintroduce that failure.
func Test_getFileCommit_NoCommitContainsFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if _, err := w.Write([]byte("[]")); err != nil {
			t.Errorf("failed to write response: %s", err)
		}
	}))
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL+"/")

	_, err := getFileCommit(context.Background(), client, "owner", "repo", "some/file.txt", "main")
	if err == nil {
		t.Fatal("expected an error when no commit contains the file, got nil")
	}

	if !errors.Is(err, errFileCommitNotFound) {
		t.Errorf("errors.Is(err, errFileCommitNotFound) = false, want true; got error: %s", err)
	}
}
