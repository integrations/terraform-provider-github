package github

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const commitRefConflictRetryTimeout = 1 * time.Minute

func commitRepositoryFiles(ctx context.Context, client *github.Client, owner, repo, branch, message string, author *github.CommitAuthor, upserts map[string]string, deletes []string) (commitSHA, treeSHA string, err error) {
	entries := make([]*github.TreeEntry, 0, len(upserts)+len(deletes))
	for path, content := range upserts {
		entries = append(entries, &github.TreeEntry{Path: new(path), Mode: new("100644"), Type: new("blob"), Content: new(content)})
	}
	for _, path := range deletes {
		entries = append(entries, &github.TreeEntry{Path: new(path), Mode: new("100644"), Type: new("blob")})
	}

	var commit *github.Commit
	err = retry.RetryContext(ctx, commitRefConflictRetryTimeout, func() *retry.RetryError {
		head, _, err := client.Repositories.GetBranch(ctx, owner, repo, branch, 0)
		if err != nil {
			return retry.NonRetryableError(err)
		}

		baseTreeSHA := head.GetCommit().GetCommit().GetTree().GetSHA()
		tree, _, err := client.Git.CreateTree(ctx, owner, repo, baseTreeSHA, entries)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		if tree.GetSHA() == baseTreeSHA {
			commit = &github.Commit{SHA: head.GetCommit().SHA, Tree: tree}
			return nil
		}

		commit, _, err = client.Git.CreateCommit(ctx, owner, repo, github.Commit{
			Message:   new(message),
			Tree:      tree,
			Parents:   []*github.Commit{{SHA: head.GetCommit().SHA}},
			Author:    author,
			Committer: author,
		}, nil)
		if err != nil {
			return retry.NonRetryableError(err)
		}

		if _, _, err := client.Git.UpdateRef(ctx, owner, repo, "heads/"+branch, github.UpdateRef{SHA: commit.GetSHA()}); err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && (ghErr.Response.StatusCode == http.StatusConflict || ghErr.Response.StatusCode == http.StatusUnprocessableEntity) {
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return "", "", err
	}

	return commit.GetSHA(), commit.GetTree().GetSHA(), nil
}

func getRepositoryBlobSHAs(ctx context.Context, client *github.Client, owner, repo, treeSHA string) (map[string]string, error) {
	tree, _, err := client.Git.GetTree(ctx, owner, repo, treeSHA, true)
	if err != nil {
		return nil, err
	}
	if tree.GetTruncated() {
		return nil, fmt.Errorf("tree %s has more entries than the GitHub API returns for %s/%s", treeSHA, owner, repo)
	}

	shas := make(map[string]string, len(tree.Entries))
	for _, entry := range tree.Entries {
		if entry.GetType() == "blob" {
			shas[entry.GetPath()] = entry.GetSHA()
		}
	}
	return shas, nil
}

func gitBlobSHA(content string) string {
	return fmt.Sprintf("%x", sha1.Sum(fmt.Appendf(nil, "blob %d\x00%s", len(content), content)))
}
