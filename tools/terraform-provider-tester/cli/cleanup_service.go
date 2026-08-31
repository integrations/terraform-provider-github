package cli

import (
	"context"
	"errors"
	"reflect"
	"sort"

	"github.com/github/terraform-provider-tester/provider"
)

type cleanupSnapshot struct {
	Owner     string
	Resources []provider.Resource
}

type sweepRequest struct {
	Owner     string
	Phrase    string
	Resources []provider.Resource
}

type sweepResult struct {
	Remaining       []provider.Resource
	SnapshotChanged bool
	attempted       bool
	residualUnknown bool
}

type cleanupService struct {
	prov   provider.Provider
	mode   string
	getenv func(string) string
}

func (s cleanupService) List(ctx context.Context) (cleanupSnapshot, error) {
	owner := s.currentOwner()
	if owner == "" {
		return cleanupSnapshot{}, errors.New("GITHUB_OWNER environment variable not set")
	}

	resources, err := s.prov.Orphans(ctx, s.mode)
	if err != nil {
		return cleanupSnapshot{}, err
	}

	return cleanupSnapshot{
		Owner:     owner,
		Resources: cloneSnapshotResources(resources),
	}, nil
}

func (s cleanupService) Sweep(ctx context.Context, req sweepRequest) (sweepResult, error) {
	currentOwner := s.currentOwner()
	if req.Owner == "" || currentOwner == "" || currentOwner != req.Owner {
		return sweepResult{}, errors.New("sweep owner changed; list orphans again")
	}
	if req.Phrase != "SWEEP "+currentOwner {
		return sweepResult{}, errors.New("sweep confirmation phrase does not match")
	}
	if len(req.Resources) == 0 {
		return sweepResult{}, errors.New("no orphaned resources to sweep")
	}

	fresh, err := s.prov.Orphans(ctx, s.mode)
	if err != nil {
		return sweepResult{}, err
	}
	if !sameResources(req.Resources, fresh) {
		return sweepResult{
			Remaining:       cloneSnapshotResources(fresh),
			SnapshotChanged: true,
		}, nil
	}

	exact := cloneSnapshotResources(fresh)
	sweepErr := s.prov.Sweep(ctx, s.mode, provider.SweepOpts{
		Targets:        []string{"repositories", "teams"},
		Confirm:        true,
		ExactResources: exact,
	})
	remaining, refreshErr := s.prov.Orphans(ctx, s.mode)
	if refreshErr != nil {
		return sweepResult{
			attempted:       true,
			residualUnknown: true,
		}, errors.Join(sweepErr, refreshErr)
	}
	return sweepResult{
		Remaining: cloneSnapshotResources(remaining),
		attempted: true,
	}, errors.Join(sweepErr, refreshErr)
}

func (s cleanupService) currentOwner() string {
	if s.getenv == nil {
		return ""
	}
	return s.getenv("GITHUB_OWNER")
}

func normalizedResources(in []provider.Resource) []provider.Resource {
	out := cloneSnapshotResources(in)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Kind != out[j].Kind {
			return out[i].Kind < out[j].Kind
		}
		if out[i].Name != out[j].Name {
			return out[i].Name < out[j].Name
		}
		return out[i].URL < out[j].URL
	})
	return out
}

func sameResources(a, b []provider.Resource) bool {
	return reflect.DeepEqual(normalizedResources(a), normalizedResources(b))
}

func cloneSnapshotResources(in []provider.Resource) []provider.Resource {
	if in == nil {
		return nil
	}
	out := make([]provider.Resource, len(in))
	copy(out, in)
	return out
}
