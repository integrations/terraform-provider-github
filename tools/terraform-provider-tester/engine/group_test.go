package engine

import (
	"errors"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/fakeprovider"
)

// TestGroupTestsBucketsByProvider verifies that GroupTests correctly routes
// test names to groups using the provider's GroupOf function and that groups
// are returned sorted by name.
func TestGroupTestsBucketsByProvider(t *testing.T) {
	names := []string{
		"TestAccGithubRepositoryCreate",
		"TestAccGithubRepositoryDelete",
		"TestAccGithubTeamCreate",
		"TestAccGithubTeamMembership",
	}

	p := &fakeprovider.Fake{
		GroupFunc: func(name string) string {
			switch {
			case strings.HasPrefix(name, "TestAccGithubRepository"):
				return "repositories"
			case strings.HasPrefix(name, "TestAccGithubTeam"):
				return "teams"
			default:
				return "misc"
			}
		},
	}

	groups, unmatched := GroupTests(names, p)

	if len(unmatched) != 0 {
		t.Fatalf("expected no unmatched names, got %v", unmatched)
	}

	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d: %v", len(groups), groups)
	}

	// Groups must be sorted by name: "repositories" < "teams".
	if groups[0].Name != "repositories" {
		t.Errorf("groups[0].Name = %q; want repositories", groups[0].Name)
	}
	if groups[1].Name != "teams" {
		t.Errorf("groups[1].Name = %q; want teams", groups[1].Name)
	}

	repoSet := make(map[string]bool)
	for _, n := range groups[0].Tests {
		repoSet[n] = true
	}
	if !repoSet["TestAccGithubRepositoryCreate"] || !repoSet["TestAccGithubRepositoryDelete"] {
		t.Errorf("repositories group missing expected tests, got %v", groups[0].Tests)
	}

	teamSet := make(map[string]bool)
	for _, n := range groups[1].Tests {
		teamSet[n] = true
	}
	if !teamSet["TestAccGithubTeamCreate"] || !teamSet["TestAccGithubTeamMembership"] {
		t.Errorf("teams group missing expected tests, got %v", groups[1].Tests)
	}
}

// TestGroupTestsReportsUnmatched verifies that names mapped to "misc" appear
// in the returned unmatched slice.
func TestGroupTestsReportsUnmatched(t *testing.T) {
	names := []string{
		"TestAccGithubRepositoryCreate",
		"TestAccGithubSomethingOdd",
	}

	p := &fakeprovider.Fake{
		GroupFunc: func(name string) string {
			if strings.HasPrefix(name, "TestAccGithubRepository") {
				return "repositories"
			}
			return "misc"
		},
	}

	_, unmatched := GroupTests(names, p)

	if len(unmatched) != 1 || unmatched[0] != "TestAccGithubSomethingOdd" {
		t.Errorf("unmatched = %v; want [TestAccGithubSomethingOdd]", unmatched)
	}
}

// TestRunPatternAnchored verifies that RunPattern builds an anchored
// alternation from a list of names.
func TestRunPatternAnchored(t *testing.T) {
	names := []string{"TestAccGithubRepository", "TestAccGithubTeam"}
	got := RunPattern(names)
	want := "^(TestAccGithubRepository|TestAccGithubTeam)$"
	if got != want {
		t.Errorf("RunPattern = %q; want %q", got, want)
	}
}

// TestRunPatternEmptyInput verifies that an empty input returns an empty
// string and not a match-all pattern.
func TestRunPatternEmptyInput(t *testing.T) {
	if got := RunPattern(nil); got != "" {
		t.Errorf("RunPattern(nil) = %q; want \"\"", got)
	}
	if got := RunPattern([]string{}); got != "" {
		t.Errorf("RunPattern([]) = %q; want \"\"", got)
	}
}

// TestRunPatternEscapesMetacharacters verifies that metacharacters in test
// names are quoted, and that the $ anchor prevents prefix matches.
func TestRunPatternEscapesMetacharacters(t *testing.T) {
	// A plain name (no metacharacters) stays literal and is anchored.
	plain := RunPattern([]string{"TestAccGithubRepository"})
	if plain != "^(TestAccGithubRepository)$" {
		t.Errorf("plain pattern = %q; want ^(TestAccGithubRepository)$", plain)
	}

	// The $ anchor must prevent TestAccGithubRepository from matching
	// TestAccGithubRepositoryFile.
	re := regexp.MustCompile(plain)
	if re.MatchString("TestAccGithubRepositoryFile") {
		t.Errorf("pattern %q should NOT match TestAccGithubRepositoryFile", plain)
	}
	if !re.MatchString("TestAccGithubRepository") {
		t.Errorf("pattern %q should match TestAccGithubRepository", plain)
	}

	// A name with regexp metacharacters must be escaped.
	withMeta := RunPattern([]string{"TestAcc.With+Meta"})
	// QuoteMeta turns '.' into '\.' and '+' into '\+' etc.
	// The resulting pattern must NOT match other strings via unescaped '.'.
	reMeta := regexp.MustCompile(withMeta)
	if reMeta.MatchString("TestAccXWith+Meta") {
		t.Errorf("unescaped dot: pattern %q should NOT match TestAccXWith+Meta", withMeta)
	}
	if !reMeta.MatchString("TestAcc.With+Meta") {
		t.Errorf("pattern %q should match the literal name TestAcc.With+Meta", withMeta)
	}
}

// TestInterpretListResultParsesNames feeds real captured -list output and
// verifies all names are returned in declaration order with the ok line stripped.
func TestInterpretListResultParsesNames(t *testing.T) {
	t.Run("single_package", func(t *testing.T) {
		data, err := os.ReadFile("testdata/list_github.txt")
		if err != nil {
			t.Fatal(err)
		}
		names, err := interpretListResult(string(data), nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Independently derive the expected ordered list from the fixture:
		// keep only lines beginning with the literal "TestAcc" prefix. This is
		// a different predicate than the parser's general Test/Benchmark/Example
		// /Fuzz rule, so it asserts full membership AND order without merely
		// re-implementing the function under test.
		var expected []string
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "TestAcc") {
				expected = append(expected, line)
			}
		}
		if len(expected) != 173 {
			t.Fatalf("fixture changed: expected 173 TestAcc lines from the pinned-SHA refresh, got %d", len(expected))
		}
		if !reflect.DeepEqual(names, expected) {
			limit := len(names)
			if len(expected) < limit {
				limit = len(expected)
			}
			for i := 0; i < limit; i++ {
				if names[i] != expected[i] {
					t.Fatalf("names differ at index %d: got %q, want %q", i, names[i], expected[i])
				}
			}
			t.Fatalf("names length %d != expected length %d", len(names), len(expected))
		}
	})

	t.Run("multi_package", func(t *testing.T) {
		data, err := os.ReadFile("testdata/list_multipkg.txt")
		if err != nil {
			t.Fatal(err)
		}
		names, err := interpretListResult(string(data), nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(names) != 2 {
			t.Errorf("expected 2 names, got %d: %v", len(names), names)
		}
		if names[0] != "TestAccA1" {
			t.Errorf("names[0] = %q; want TestAccA1", names[0])
		}
		if names[1] != "TestAccB1" {
			t.Errorf("names[1] = %q; want TestAccB1", names[1])
		}
	})
}

// TestInterpretListResultEmptyOnNoMatch verifies that a no-match -list run
// (ok line only, exit 0) returns an empty slice without error.
func TestInterpretListResultEmptyOnNoMatch(t *testing.T) {
	data, err := os.ReadFile("testdata/list_nomatch.txt")
	if err != nil {
		t.Fatal(err)
	}
	names, err := interpretListResult(string(data), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected empty slice, got %v", names)
	}
}

// TestInterpretListResultErrorsOnBuildFailure verifies that a non-zero exit
// (e.g. build failure) returns an error that includes the compiler output and
// returns no names.
func TestInterpretListResultErrorsOnBuildFailure(t *testing.T) {
	data, err := os.ReadFile("testdata/list_buildfail.txt")
	if err != nil {
		t.Fatal(err)
	}
	fakeRunErr := errors.New("exit status 1")
	names, err := interpretListResult(string(data), fakeRunErr)
	if err == nil {
		t.Fatal("expected an error for build failure, got nil")
	}
	if len(names) != 0 {
		t.Errorf("expected no names on build failure, got %v", names)
	}
	// The error message must surface the compiler output so the operator
	// can see what went wrong.
	if !strings.Contains(err.Error(), "[build failed]") {
		t.Errorf("error %q does not contain '[build failed]'", err.Error())
	}
}
