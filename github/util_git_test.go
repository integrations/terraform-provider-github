package github

import (
	"net/http"
	"testing"

	"github.com/google/go-github/v89/github"
)

func TestGitBlobSHA(t *testing.T) {
	t.Parallel()

	for content, want := range map[string]string{
		"":        "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391",
		"hello\n": "ce013625030ba8dba906f756967f9e9ca394464a",
	} {
		if got := gitBlobSHA(content); got != want {
			t.Errorf("gitBlobSHA(%q) = %s, want %s", content, got, want)
		}
	}
}

func TestCommitRepositoryFiles(t *testing.T) {
	t.Parallel()

	const (
		branchURI  = "/repos/o/r/branches/main"
		treesURI   = "/repos/o/r/git/trees"
		commitsURI = "/repos/o/r/git/commits"
		refURI     = "/repos/o/r/git/refs/heads/main"
	)
	head := func(commit, tree string) *github.Branch {
		return &github.Branch{Commit: &github.RepositoryCommit{SHA: new(commit), Commit: &github.Commit{Tree: &github.Tree{SHA: new(tree)}}}}
	}
	commit := func(sha, tree string) *github.Commit {
		return &github.Commit{SHA: new(sha), Tree: &github.Tree{SHA: new(tree)}}
	}

	for name, tt := range map[string]struct {
		responses     []*mockResponse
		wantCommitSHA string
		wantTreeSHA   string
	}{
		"commits on the branch head": {
			responses: []*mockResponse{
				mustGetTestMockResponse(t, branchURI, http.StatusOK, head("c1", "t1")),
				mustGetTestMockResponse(t, treesURI, http.StatusCreated, &github.Tree{SHA: new("t2")}),
				mustGetTestMockResponse(t, commitsURI, http.StatusCreated, commit("c2", "t2")),
				mustGetTestMockResponse(t, refURI, http.StatusOK, &github.Reference{}),
			},
			wantCommitSHA: "c2",
			wantTreeSHA:   "t2",
		},
		"skips the commit when the tree is unchanged": {
			responses: []*mockResponse{
				mustGetTestMockResponse(t, branchURI, http.StatusOK, head("c1", "t1")),
				mustGetTestMockResponse(t, treesURI, http.StatusCreated, &github.Tree{SHA: new("t1")}),
			},
			wantCommitSHA: "c1",
			wantTreeSHA:   "t1",
		},
		"rebuilds on the new head when the ref moved": {
			responses: []*mockResponse{
				mustGetTestMockResponse(t, branchURI, http.StatusOK, head("c1", "t1")),
				mustGetTestMockResponse(t, treesURI, http.StatusCreated, &github.Tree{SHA: new("t2")}),
				mustGetTestMockResponse(t, commitsURI, http.StatusCreated, commit("c2", "t2")),
				mustGetTestMockResponse(t, refURI, http.StatusUnprocessableEntity, &github.ErrorResponse{Message: "Update is not a fast forward"}),
				mustGetTestMockResponse(t, branchURI, http.StatusOK, head("c3", "t3")),
				mustGetTestMockResponse(t, treesURI, http.StatusCreated, &github.Tree{SHA: new("t4")}),
				mustGetTestMockResponse(t, commitsURI, http.StatusCreated, commit("c4", "t4")),
				mustGetTestMockResponse(t, refURI, http.StatusOK, &github.Reference{}),
			},
			wantCommitSHA: "c4",
			wantTreeSHA:   "t4",
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock(tt.responses)
			t.Cleanup(ts.Close)
			client := mustCreateTestGitHubClient(t, ts.URL+"/")

			commitSHA, treeSHA, err := commitRepositoryFiles(t.Context(), client, "o", "r", "main", "msg", nil, map[string]string{"a.txt": "alpha"}, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if commitSHA != tt.wantCommitSHA || treeSHA != tt.wantTreeSHA {
				t.Fatalf("got commit %s tree %s, want commit %s tree %s", commitSHA, treeSHA, tt.wantCommitSHA, tt.wantTreeSHA)
			}
		})
	}
}
