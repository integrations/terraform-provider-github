package github

import "testing"

func TestProjectV2CompositeIDPreservesOpaqueParts(t *testing.T) {
	t.Parallel()

	id2, err := buildProjectV2ID("PVT:1", "R:1")
	if err != nil {
		t.Fatalf("building two-part ID: %v", err)
	}
	projectID, repositoryID, err := parseProjectV2ID2(id2)
	if err != nil {
		t.Fatalf("parsing two-part ID: %v", err)
	}
	if projectID != "PVT:1" || repositoryID != "R:1" {
		t.Fatalf("unexpected two-part ID values: %q, %q", projectID, repositoryID)
	}

	id3, err := buildProjectV2ID("PVT:1", "PVTI:1", "PVTF:1")
	if err != nil {
		t.Fatalf("building three-part ID: %v", err)
	}
	projectID, itemID, fieldID, err := parseProjectV2ID3(id3)
	if err != nil {
		t.Fatalf("parsing three-part ID: %v", err)
	}
	if projectID != "PVT:1" || itemID != "PVTI:1" || fieldID != "PVTF:1" {
		t.Fatalf("unexpected three-part ID values: %q, %q, %q", projectID, itemID, fieldID)
	}
}
