package github

import "testing"

func TestGithubBranchIsUpdatedWhenBranchChanges(t *testing.T) {
	resource := resourceGithubBranch()

	branchSchema := resource.Schema["branch"]
	if branchSchema == nil {
		t.Fatal("branch field should exist in schema")
		return
	}
	if branchSchema.ForceNew {
		t.Error("branch field should not be ForceNew so renames are handled via update")
	}
}

func TestGithubBranchIsRecreatedWhenRepositoryChanges(t *testing.T) {
	resource := resourceGithubBranch()

	repositorySchema := resource.Schema["repository"]
	if repositorySchema == nil {
		t.Fatal("repository field should exist in schema")
		return
	}
	if !repositorySchema.ForceNew {
		t.Error("repository field should be ForceNew so changes recreate the resource")
	}
}

func TestGithubBranchIsRecreatedWhenSourceBranchChanges(t *testing.T) {
	resource := resourceGithubBranch()

	sourceBranchSchema := resource.Schema["source_branch"]
	if sourceBranchSchema == nil {
		t.Fatal("source_branch field should exist in schema")
		return
	}
	if !sourceBranchSchema.ForceNew {
		t.Error("source_branch field should be ForceNew so changes recreate the resource")
	}
}

func TestGithubBranchIsRecreatedWhenSourceSHAChanges(t *testing.T) {
	resource := resourceGithubBranch()

	sourceSHASchema := resource.Schema["source_sha"]
	if sourceSHASchema == nil {
		t.Fatal("source_sha field should exist in schema")
		return
	}
	if !sourceSHASchema.ForceNew {
		t.Error("source_sha field should be ForceNew so changes recreate the resource")
	}
}
