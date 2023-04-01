package github

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/v50/github"
)

// checkRepositoryBranchExists tests if a branch exists in a repository.
func checkRepositoryBranchExists(client *github.Client, owner, repo, branch string) error {
	ctx := context.WithValue(context.Background(), ctxId, buildTwoPartID(repo, branch))
	_, _, err := client.Repositories.GetBranch(ctx, owner, repo, branch, true)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return fmt.Errorf("branch %s not found in repository %s/%s or repository is not readable", branch, owner, repo)
			}
		}
		return err
	}

	return nil
}

// checkRepositoryFileExists tests if a file exists in a repository.
func checkRepositoryFileExists(client *github.Client, owner, repo, file, branch string) error {
	ctx := context.WithValue(context.Background(), ctxId, fmt.Sprintf("%s/%s", repo, file))
	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		return nil
	}
	if fc == nil {
		return fmt.Errorf("file %s not a file in in repository %s/%s or repository is not readable", file, owner, repo)
	}

	return nil
}

func getFileCommit(client *github.Client, owner, repo, file, branch string) (*github.RepositoryCommit, error) {
	ctx := context.WithValue(context.Background(), ctxId, fmt.Sprintf("%s/%s", repo, file))
	opts := &github.CommitsListOptions{
		SHA:  branch,
		Path: file,
	}
	allCommits := []*github.RepositoryCommit{}
	for {
		commits, resp, err := client.Repositories.ListCommits(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}

		allCommits = append(allCommits, commits...)

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	for _, c := range allCommits {
		sha := c.GetSHA()

		// Skip merge commits
		if strings.Contains(c.Commit.GetMessage(), "Merge branch") {
			continue
		}

		rc, _, err := client.Repositories.GetCommit(ctx, owner, repo, sha, nil)
		if err != nil {
			return nil, err
		}

		for _, f := range rc.Files {
			if f.GetFilename() == file && f.GetStatus() != "removed" {
				return rc, nil
			}
		}
	}

	return nil, fmt.Errorf("cannot find file %s in repo %s/%s", file, owner, repo)
}

// getAutolinkByKeyPrefix returns a single autolink reference by key prefix that was configured for the given repository.
func getAutolinkByKeyPrefix(client *github.Client, owner, repo, keyPrefix string) (*github.Autolink, error) {
	autolinks, err := listAutolinks(client, owner, repo)
	if err != nil {
		return nil, err
	}

	for _, autolink := range autolinks {
		if *autolink.KeyPrefix == keyPrefix {
			return autolink, nil
		}
	}

	return nil, nil
}

// listAutolinks returns all autolink references for the given repository.
func listAutolinks(client *github.Client, owner, repo string) ([]*github.Autolink, error) {
	ctx := context.WithValue(context.Background(), ctxId, fmt.Sprintf("%s/%s", owner, repo))
	opts := &github.ListOptions{
		PerPage: maxPerPage,
	}

	var allAutolinks []*github.Autolink
	for {
		autolinks, resp, err := client.Repositories.ListAutolinks(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		allAutolinks = append(allAutolinks, autolinks...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allAutolinks, nil
}
