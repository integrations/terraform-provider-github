package github

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

type createTestRepositoryOptionsFunc func(*github.Repository)

func mustCreateTestRepository(t *testing.T, f ...createTestRepositoryOptionsFunc) *github.Repository {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	req := &github.Repository{
		Name:     &name,
		AutoInit: new(true),
	}

	for _, fn := range f {
		if fn != nil {
			fn(req)
		}
	}

	var org string
	if testAccConf.meta.IsOrganization {
		org = testAccConf.meta.name
	}

	repo, _, err := testAccConf.meta.v3client.Repositories.Create(t.Context(), org, repositoryCreateRequest(req))
	if err != nil {
		t.Fatalf("failed to create test repository: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Repositories.Delete(context.Background(), testAccConf.meta.name, name); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test repository %s: %v", name, err)
		}
	})

	return repo
}

func mustRenameTestRepository(t *testing.T, repo *github.Repository, newName string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Repositories.Edit(t.Context(), testAccConf.meta.name, repo.GetName(), &github.Repository{Name: &newName})
	if err != nil {
		t.Fatalf("failed to rename test repository %s to %s: %v", repo.GetName(), newName, err)
	}
}

func mustArchiveTestRepository(t *testing.T, repo *github.Repository) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Repositories.Edit(t.Context(), testAccConf.meta.name, repo.GetName(), &github.Repository{Archived: new(true)})
	if err != nil {
		t.Fatalf("failed to archive test repository %s: %v", repo.GetName(), err)
	}
}

func mustCreateTestRepositoryDeployKey(t *testing.T, repo *github.Repository) *github.Key {
	randomID := acctest.RandString(testRandomIDLength)
	title := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	pub, _ := mustNewSSHKey(t)

	req := github.CreateDeployKeyRequest{
		Key:      pub,
		Title:    new(title),
		ReadOnly: new(true),
	}

	k, _, err := testAccConf.meta.v3client.Repositories.CreateKey(t.Context(), testAccConf.meta.name, repo.GetName(), req)
	if err != nil {
		t.Fatalf("failed to create test repository deploy key: %v", err)
	}

	return k
}

type createTestRepositoryFileOptionsFunc func(*github.RepositoryContentFileOptions)

func mustCreateTestRepositoryFile(t *testing.T, repo *github.Repository, path, content string, f ...createTestRepositoryFileOptionsFunc) *github.RepositoryContentResponse {
	t.Helper()

	req := &github.RepositoryContentFileOptions{
		Message: new(fmt.Sprintf("Add %s", path)),
		Content: []byte(content),
		Branch:  new(repo.GetDefaultBranch()),
	}

	for _, fn := range f {
		if fn != nil {
			fn(req)
		}
	}

	resp, _, err := testAccConf.meta.v3client.Repositories.CreateFile(t.Context(), testAccConf.meta.name, repo.GetName(), path, req)
	if err != nil {
		t.Fatalf("failed to create %s in test repository %s: %v", path, repo.GetName(), err)
	}

	return resp
}

func mustDeleteRepositoryFile(t *testing.T, repo *github.Repository, path string) {
	t.Helper()

	file, _, _, err := testAccConf.meta.v3client.Repositories.GetContents(t.Context(), testAccConf.meta.name, repo.GetName(), path, nil)
	if err != nil {
		t.Fatalf("failed to read %s of test repository %s: %v", path, repo.GetName(), err)
	}

	_, _, err = testAccConf.meta.v3client.Repositories.DeleteFile(t.Context(), testAccConf.meta.name, repo.GetName(), path, &github.RepositoryContentFileOptions{
		Message: new(fmt.Sprintf("Remove %s", path)),
		SHA:     file.SHA,
	})
	if err != nil {
		t.Fatalf("failed to delete %s of test repository %s: %v", path, repo.GetName(), err)
	}
}

func mustCreateTestRepositoryBranch(t *testing.T, repo *github.Repository) string {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	// Get the SHA of the default branch
	defaultBranch, _, err := testAccConf.meta.v3client.Repositories.GetBranch(t.Context(), testAccConf.meta.name, repo.GetName(), repo.GetDefaultBranch(), 0)
	if err != nil {
		t.Fatalf("failed to get default branch for test repository %s: %v", repo.GetName(), err)
	}

	req := github.CreateRef{
		Ref: fmt.Sprintf("refs/heads/%s", name),
		SHA: defaultBranch.GetCommit().GetSHA(),
	}
	if _, _, err := testAccConf.meta.v3client.Git.CreateRef(t.Context(), testAccConf.meta.name, repo.GetName(), req); err != nil {
		t.Fatalf("failed to create test branch %s for repository %s: %v", name, repo.GetName(), err)
	}

	return name
}

func mustCreateTestRepositoryBranchProtection(t *testing.T, repo *github.Repository, branch string) *github.Protection {
	t.Helper()

	protection, _, err := testAccConf.meta.v3client.Repositories.UpdateBranchProtection(t.Context(), testAccConf.meta.name, repo.GetName(), branch, &github.ProtectionRequest{})
	if err != nil {
		t.Fatalf("failed to create branch protection for branch %s of test repository %s: %v", branch, repo.GetName(), err)
	}

	return protection
}

func mustCreateTestRepositoryPullRequest(t *testing.T, repo *github.Repository, head string) *github.PullRequest {
	t.Helper()

	pr, _, err := testAccConf.meta.v3client.PullRequests.Create(t.Context(), testAccConf.meta.name, repo.GetName(), github.CreatePullRequest{
		Title: new("test title"),
		Body:  new("test body"),
		Head:  head,
		Base:  repo.GetDefaultBranch(),
	})
	if err != nil {
		t.Fatalf("failed to create pull request from %s in test repository %s: %v", head, repo.GetName(), err)
	}

	return pr
}
