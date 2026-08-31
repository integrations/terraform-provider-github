package engine

import (
	"sort"

	"github.com/github/terraform-provider-tester/provider"
)

const (
	CleanupNotApplicable = "not-applicable"
	CleanupBaselineOnly  = "baseline-only"
	CleanupComplete      = "complete"
	CleanupUnknown       = "unknown"
)

type OrphanAccounting struct {
	Mode             string              `json:"mode"`
	Baseline         []provider.Resource `json:"baseline"`
	Final            []provider.Resource `json:"final"`
	PreExisting      []provider.Resource `json:"pre_existing"`
	New              []provider.Resource `json:"new"`
	BaselineCaptured bool                `json:"baseline_captured"`
	FinalCaptured    bool                `json:"final_captured"`
	CleanupStatus    string              `json:"cleanup_status"`
}

func orphanResourceKey(r provider.Resource) string {
	return r.Kind + "\x00" + r.Name
}

func cloneAndSortResources(resources []provider.Resource) []provider.Resource {
	if resources == nil {
		return nil
	}

	byKey := make(map[string]provider.Resource, len(resources))
	for _, resource := range resources {
		byKey[orphanResourceKey(resource)] = resource
	}

	sorted := make([]provider.Resource, 0, len(byKey))
	for _, resource := range byKey {
		sorted = append(sorted, resource)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Kind != sorted[j].Kind {
			return sorted[i].Kind < sorted[j].Kind
		}
		return sorted[i].Name < sorted[j].Name
	})
	return sorted
}

// ComputeOrphanDelta compares baseline membership by kind+name only; duplicate
// baseline entries and URL changes are ignored, and final is deduplicated by
// kind+name, keeping the last observed URL and sorting by kind then name before
// partitioning.
func ComputeOrphanDelta(
	baseline []provider.Resource,
	final []provider.Resource,
) ([]provider.Resource, []provider.Resource) {
	baselineSet := make(map[string]struct{}, len(baseline))
	for _, resource := range baseline {
		baselineSet[orphanResourceKey(resource)] = struct{}{}
	}

	finalResources := cloneAndSortResources(final)
	preExisting := make([]provider.Resource, 0, len(finalResources))
	newlyLeaked := make([]provider.Resource, 0, len(finalResources))
	for _, resource := range finalResources {
		if _, ok := baselineSet[orphanResourceKey(resource)]; ok {
			preExisting = append(preExisting, resource)
			continue
		}
		newlyLeaked = append(newlyLeaked, resource)
	}
	return preExisting, newlyLeaked
}
