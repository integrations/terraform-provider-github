package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

type createTestRepositoryOptionsFunc func(*github.Repository)

func withInternalVisibility() createTestRepositoryOptionsFunc {
	return func(repo *github.Repository) {
		repo.Visibility = new("internal")
	}
}

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
		_, err := retryUntilOK(context.Background(), func() (*github.Repository, bool, error) {
			_, err := testAccConf.meta.v3client.Repositories.Delete(context.Background(), testAccConf.meta.name, name)
			if err != nil {
				if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusNotFound {
					return nil, true, nil
				}
				if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusConflict {
					return nil, false, nil
				}
				return nil, false, err
			}
			return nil, true, nil
		}, nil)
		if err != nil {
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

	archived := true
	_, err := retryUntilOK(t.Context(), func() (*github.Repository, bool, error) {
		repoToArchive, _, err := testAccConf.meta.v3client.Repositories.Edit(t.Context(), testAccConf.meta.name, repo.GetName(), &github.Repository{Archived: &archived})
		if err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusUnprocessableEntity {
				return repoToArchive, false, nil
			}
			if repoToArchive != nil && repoToArchive.GetArchived() {
				return repoToArchive, true, nil
			}
			return nil, false, err
		}
		return repoToArchive, true, nil
	}, nil)
	if err != nil {
		t.Fatalf("failed to archive test repository %s: %v", repo.GetName(), err)
	}
}
