package github

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-github/v89/github"
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

	repo, _, err := testAccConf.meta.v3client.Repositories.Create(t.Context(), org, req)
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
