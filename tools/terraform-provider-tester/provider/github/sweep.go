package github

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/github/terraform-provider-tester/provider"
	gogithub "github.com/google/go-github/v88/github"
)

// testResourcePrefix is the hard-coded name prefix for all acceptance-test resources.
// It MUST match github/acc_test.go:28 in the provider repo.
// It is never configurable and is never derived from user input.
const testResourcePrefix = "tf-acc-test-"

// isOrgMode reports whether the given test-auth mode implies org-level API access.
// Mirrors orgTestModes from github/acc_test.go:31.
func isOrgMode(mode string) bool {
	return mode == "organization" || mode == "team" || mode == "enterprise"
}

// Orphans lists all tf-acc-test-* repositories (and, in org modes, teams) owned by
// GITHUB_OWNER.  It is strictly read-only - it never issues a DELETE.
func (p *ghProvider) Orphans(ctx context.Context, mode string) ([]provider.Resource, error) {
	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		return nil, errors.New("GITHUB_OWNER not set")
	}
	c := newClient(os.Getenv("GITHUB_BASE_URL"), os.Getenv("GITHUB_TOKEN"))
	return orphansWithClient(ctx, c, owner, isOrgMode(mode))
}

// orphansWithClient is the testable core of Orphans.
// It performs only GET (list) requests and returns all prefixed resources.
func orphansWithClient(ctx context.Context, c *gogithub.Client, owner string, org bool) ([]provider.Resource, error) {
	var resources []provider.Resource

	// List repositories (paginated).
	for page := 1; ; {
		var (
			repos []*gogithub.Repository
			resp  *gogithub.Response
			err   error
		)
		if org {
			opts := &gogithub.RepositoryListByOrgOptions{
				ListOptions: gogithub.ListOptions{Page: page, PerPage: 100},
			}
			repos, resp, err = c.Repositories.ListByOrg(ctx, owner, opts)
		} else {
			opts := &gogithub.RepositoryListByUserOptions{
				ListOptions: gogithub.ListOptions{Page: page, PerPage: 100},
			}
			repos, resp, err = c.Repositories.ListByUser(ctx, owner, opts)
		}
		if err != nil {
			return nil, err
		}
		for _, r := range repos {
			if strings.HasPrefix(r.GetName(), testResourcePrefix) {
				resources = append(resources, provider.Resource{
					Kind: "repository",
					Name: r.GetName(),
					URL:  r.GetHTMLURL(),
				})
			}
		}
		if resp.NextPage == 0 {
			break
		}
		page = resp.NextPage
	}

	// List teams - org modes only.
	if org {
		for page := 1; ; {
			opts := &gogithub.ListOptions{Page: page, PerPage: 100}
			teams, resp, err := c.Teams.ListTeams(ctx, owner, opts)
			if err != nil {
				return nil, err
			}
			for _, t := range teams {
				if strings.HasPrefix(t.GetSlug(), testResourcePrefix) {
					resources = append(resources, provider.Resource{
						Kind: "team",
						Name: t.GetSlug(),
						URL:  t.GetHTMLURL(),
					})
				}
			}
			if resp.NextPage == 0 {
				break
			}
			page = resp.NextPage
		}
	}

	return resources, nil
}

// Sweep deletes tf-acc-test-* resources in the requested target kinds.
// When opts.ExactResources is non-nil, it deletes only that validated snapshot
// and performs no discovery requests.
// It refuses to run without explicit confirmation and never mutates anything outside
// the hard-coded testResourcePrefix.
func (p *ghProvider) Sweep(ctx context.Context, mode string, opts provider.SweepOpts) error {
	if !opts.Confirm {
		return errors.New("sweep refused: opts.Confirm is false (no resources deleted)")
	}
	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		return errors.New("GITHUB_OWNER not set")
	}
	c := newClient(os.Getenv("GITHUB_BASE_URL"), os.Getenv("GITHUB_TOKEN"))
	return sweepWithClient(ctx, c, owner, isOrgMode(mode), opts)
}

func resourceKeyString(r provider.Resource) string {
	return r.Kind + "\x00" + r.Name
}

// sweepWithClient is the testable core of Sweep.
// PRECONDITION (defense in depth): Confirm must be true; this is checked again even
// though the public Sweep method already guards on it.
func sweepWithClient(ctx context.Context, c *gogithub.Client, owner string, org bool, opts provider.SweepOpts) error {
	if !opts.Confirm {
		return errors.New("sweep refused: opts.Confirm is false (no resources deleted)")
	}

	targetSet := sweepTargetSet(opts.Targets)
	if opts.ExactResources != nil {
		return sweepExactResources(ctx, c, owner, org, targetSet, opts.ExactResources)
	}

	allowSet := map[string]struct{}(nil)
	allowedKinds := map[string]bool(nil)
	if opts.Resources != nil {
		allowSet = make(map[string]struct{}, len(opts.Resources))
		allowedKinds = make(map[string]bool, len(opts.Resources))
		for _, resource := range opts.Resources {
			allowSet[resourceKeyString(resource)] = struct{}{}
			allowedKinds[resource.Kind] = true
		}
		if len(allowSet) == 0 {
			return nil
		}
	}

	targetEnabled := func(target, kind string) bool {
		if !targetSet[target] {
			return false
		}
		if allowSet == nil {
			return true
		}
		return allowedKinds[kind]
	}

	resourceAllowed := func(kind, name string) bool {
		if allowSet == nil {
			return true
		}
		_, ok := allowSet[resourceKeyString(provider.Resource{Kind: kind, Name: name})]
		return ok
	}

	if targetEnabled("repositories", "repository") {
		var names []string
		for page := 1; ; {
			var (
				repos []*gogithub.Repository
				resp  *gogithub.Response
				err   error
			)
			if org {
				listOpts := &gogithub.RepositoryListByOrgOptions{
					ListOptions: gogithub.ListOptions{Page: page, PerPage: 100},
				}
				repos, resp, err = c.Repositories.ListByOrg(ctx, owner, listOpts)
			} else {
				listOpts := &gogithub.RepositoryListByUserOptions{
					ListOptions: gogithub.ListOptions{Page: page, PerPage: 100},
				}
				repos, resp, err = c.Repositories.ListByUser(ctx, owner, listOpts)
			}
			if err != nil {
				return err
			}
			for _, r := range repos {
				name := r.GetName()
				if strings.HasPrefix(name, testResourcePrefix) && resourceAllowed("repository", name) {
					names = append(names, name)
				}
			}
			if resp.NextPage == 0 {
				break
			}
			page = resp.NextPage
		}
		for _, name := range names {
			if !strings.HasPrefix(name, testResourcePrefix) {
				continue
			}
			if _, delErr := c.Repositories.Delete(ctx, owner, name); delErr != nil {
				return delErr
			}
		}
	}

	if targetEnabled("teams", "team") && org {
		var slugs []string
		for page := 1; ; {
			listOpts := &gogithub.ListOptions{Page: page, PerPage: 100}
			teams, resp, err := c.Teams.ListTeams(ctx, owner, listOpts)
			if err != nil {
				return err
			}
			for _, t := range teams {
				slug := t.GetSlug()
				if strings.HasPrefix(slug, testResourcePrefix) && resourceAllowed("team", slug) {
					slugs = append(slugs, slug)
				}
			}
			if resp.NextPage == 0 {
				break
			}
			page = resp.NextPage
		}
		for _, slug := range slugs {
			if !strings.HasPrefix(slug, testResourcePrefix) {
				continue
			}
			if _, delErr := c.Teams.DeleteTeamBySlug(ctx, owner, slug); delErr != nil {
				return delErr
			}
		}
	}

	return nil
}

func sweepTargetSet(targets []string) map[string]bool {
	targetSet := make(map[string]bool, len(targets))
	for _, t := range targets {
		targetSet[t] = true
	}
	return targetSet
}

func sweepExactResources(ctx context.Context, c *gogithub.Client, owner string, org bool, targetSet map[string]bool, resources []provider.Resource) error {
	repoNames, teamSlugs, err := exactSweepNames(resources, targetSet, org)
	if err != nil {
		return err
	}
	for _, name := range repoNames {
		if _, delErr := c.Repositories.Delete(ctx, owner, name); delErr != nil {
			return delErr
		}
	}
	for _, slug := range teamSlugs {
		if _, delErr := c.Teams.DeleteTeamBySlug(ctx, owner, slug); delErr != nil {
			return delErr
		}
	}
	return nil
}

func exactSweepNames(resources []provider.Resource, targetSet map[string]bool, org bool) ([]string, []string, error) {
	repoNames := make([]string, 0, len(resources))
	teamSlugs := make([]string, 0, len(resources))
	seen := make(map[resourceKey]bool, len(resources))

	for _, r := range resources {
		if !strings.HasPrefix(r.Name, testResourcePrefix) {
			return nil, nil, fmt.Errorf("sweep exact resource %q kind %q missing required prefix %q", r.Name, r.Kind, testResourcePrefix)
		}
		key := resourceKey{kind: r.Kind, name: r.Name}
		switch r.Kind {
		case "repository":
			if !targetSet["repositories"] {
				return nil, nil, fmt.Errorf("sweep exact repository %q is not enabled by targets", r.Name)
			}
			if !seen[key] {
				repoNames = append(repoNames, r.Name)
				seen[key] = true
			}
		case "team":
			if !targetSet["teams"] {
				return nil, nil, fmt.Errorf("sweep exact team %q is not enabled by targets", r.Name)
			}
			if !org {
				return nil, nil, fmt.Errorf("sweep exact team %q requires organization auth mode", r.Name)
			}
			if !seen[key] {
				teamSlugs = append(teamSlugs, r.Name)
				seen[key] = true
			}
		default:
			return nil, nil, fmt.Errorf("sweep exact resource %q has unsupported kind %q", r.Name, r.Kind)
		}
	}

	return repoNames, teamSlugs, nil
}

type resourceKey struct {
	kind string
	name string
}
