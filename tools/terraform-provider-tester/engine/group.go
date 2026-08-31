package engine

import (
	"regexp"
	"sort"
	"strings"

	"github.com/github/terraform-provider-tester/provider"
)

// Group holds a named set of tests that share the same logical bucket.
type Group struct {
	Name  string
	Tests []string
}

// GroupTests buckets names via p.GroupOf and returns groups sorted by name.
// Names that map to "misc" are also returned in the unmatched slice so callers
// can enforce an --unmatched guard.
func GroupTests(names []string, p provider.Provider) (groups []Group, unmatched []string) {
	byName := make(map[string][]string)
	for _, n := range names {
		g := p.GroupOf(n)
		byName[g] = append(byName[g], n)
		if g == "misc" {
			unmatched = append(unmatched, n)
		}
	}

	keys := make([]string, 0, len(byName))
	for k := range byName {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		groups = append(groups, Group{Name: k, Tests: byName[k]})
	}
	return groups, unmatched
}

// RunPattern builds an anchored regexp alternation over top-level test names:
//
//	["TestAccA","TestAccB"] -> "^(TestAccA|TestAccB)$"
//
// An empty input returns "" (empty string). Callers must treat an empty pattern
// as "run nothing", never as "run all".
func RunPattern(topLevel []string) string {
	if len(topLevel) == 0 {
		return ""
	}
	escaped := make([]string, len(topLevel))
	for i, name := range topLevel {
		escaped[i] = regexp.QuoteMeta(name)
	}
	return "^(" + strings.Join(escaped, "|") + ")$"
}
