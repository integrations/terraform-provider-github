package engine

import (
	"reflect"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

func baselineOrphanResources() []provider.Resource {
	return []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-baseline"},
		{Kind: "repository", Name: "tf-acc-test-removed", URL: "https://example.test/repos/tf-acc-test-removed"},
	}
}

func finalOrphanResources() []provider.Resource {
	return []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
		{Kind: "team", Name: "tf-acc-test-a", URL: "https://example.test/teams/tf-acc-test-a"},
		{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new"},
	}
}

func TestComputeOrphanDeltaSeparatesPreExistingAndNew(t *testing.T) {
	preExisting, newlyLeaked := ComputeOrphanDelta(baselineOrphanResources(), finalOrphanResources())

	wantPreExisting := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
	}
	wantNew := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new"},
		{Kind: "team", Name: "tf-acc-test-a", URL: "https://example.test/teams/tf-acc-test-a"},
	}
	if !reflect.DeepEqual(preExisting, wantPreExisting) {
		t.Fatalf("preExisting = %#v, want %#v", preExisting, wantPreExisting)
	}
	if !reflect.DeepEqual(newlyLeaked, wantNew) {
		t.Fatalf("newlyLeaked = %#v, want %#v", newlyLeaked, wantNew)
	}
}

func TestComputeOrphanDeltaIgnoresDeletedBaselineResources(t *testing.T) {
	preExisting, newlyLeaked := ComputeOrphanDelta(
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-baseline"},
			{Kind: "repository", Name: "tf-acc-test-removed", URL: "https://example.test/repos/tf-acc-test-removed"},
		},
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
		},
	)

	wantPreExisting := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
	}
	if !reflect.DeepEqual(preExisting, wantPreExisting) {
		t.Fatalf("preExisting = %#v, want %#v", preExisting, wantPreExisting)
	}
	if len(newlyLeaked) != 0 {
		t.Fatalf("newlyLeaked = %#v, want empty", newlyLeaked)
	}
}

func TestComputeOrphanDeltaIncludesKindInIdentity(t *testing.T) {
	preExisting, newlyLeaked := ComputeOrphanDelta(
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-baseline"},
		},
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
			{Kind: "team", Name: "tf-acc-test-a", URL: "https://example.test/teams/tf-acc-test-a"},
		},
	)

	wantPreExisting := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
	}
	wantNew := []provider.Resource{
		{Kind: "team", Name: "tf-acc-test-a", URL: "https://example.test/teams/tf-acc-test-a"},
	}
	if !reflect.DeepEqual(preExisting, wantPreExisting) {
		t.Fatalf("preExisting = %#v, want %#v", preExisting, wantPreExisting)
	}
	if !reflect.DeepEqual(newlyLeaked, wantNew) {
		t.Fatalf("newlyLeaked = %#v, want %#v", newlyLeaked, wantNew)
	}
}

func TestComputeOrphanDeltaIgnoresURLChanges(t *testing.T) {
	preExisting, newlyLeaked := ComputeOrphanDelta(
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-baseline"},
		},
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
		},
	)

	wantPreExisting := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
	}
	if !reflect.DeepEqual(preExisting, wantPreExisting) {
		t.Fatalf("preExisting = %#v, want %#v", preExisting, wantPreExisting)
	}
	if len(newlyLeaked) != 0 {
		t.Fatalf("newlyLeaked = %#v, want empty", newlyLeaked)
	}
}

func TestComputeOrphanDeltaDeduplicatesResources(t *testing.T) {
	preExisting, newlyLeaked := ComputeOrphanDelta(
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-first"},
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-last-baseline"},
		},
		[]provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-first-final"},
			{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-last-final"},
			{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new-first"},
			{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new-last"},
		},
	)

	wantPreExisting := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-last-final"},
	}
	wantNew := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new-last"},
	}
	if !reflect.DeepEqual(preExisting, wantPreExisting) {
		t.Fatalf("preExisting = %#v, want %#v", preExisting, wantPreExisting)
	}
	if !reflect.DeepEqual(newlyLeaked, wantNew) {
		t.Fatalf("newlyLeaked = %#v, want %#v", newlyLeaked, wantNew)
	}
}

func TestComputeOrphanDeltaSortsByKindThenName(t *testing.T) {
	preExisting, newlyLeaked := ComputeOrphanDelta(nil, []provider.Resource{
		{Kind: "team", Name: "tf-acc-test-b", URL: "https://example.test/teams/tf-acc-test-b"},
		{Kind: "repository", Name: "tf-acc-test-c", URL: "https://example.test/repos/tf-acc-test-c"},
		{Kind: "team", Name: "tf-acc-test-a", URL: "https://example.test/teams/tf-acc-test-a"},
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a"},
	})

	if len(preExisting) != 0 {
		t.Fatalf("preExisting = %#v, want empty", preExisting)
	}
	wantNew := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a"},
		{Kind: "repository", Name: "tf-acc-test-c", URL: "https://example.test/repos/tf-acc-test-c"},
		{Kind: "team", Name: "tf-acc-test-a", URL: "https://example.test/teams/tf-acc-test-a"},
		{Kind: "team", Name: "tf-acc-test-b", URL: "https://example.test/teams/tf-acc-test-b"},
	}
	if !reflect.DeepEqual(newlyLeaked, wantNew) {
		t.Fatalf("newlyLeaked = %#v, want %#v", newlyLeaked, wantNew)
	}
}
