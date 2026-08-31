package github

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

// requirementsFixtureProviderSHA is the exact provider checkout SHA used to
// generate engine/testdata/list_github.txt. It is recorded here so that any
// future fixture refresh can be verified against a pinned revision.
const requirementsFixtureProviderSHA = "c1029539a7ed4fa076d1bc229f57d10b5e587d0f"

// TestRequirementsFor_allFixtureTestsClassified verifies that every TestAcc
// name in the pinned fixture is classified (RequirementsFor returns true).
// The fixture was generated from the provider at requirementsFixtureProviderSHA
// and must contain exactly 173 names.
func TestRequirementsFor_allFixtureTestsClassified(t *testing.T) {
	p := New()

	f, err := os.Open("../../engine/testdata/list_github.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var count int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "TestAcc") {
			continue
		}
		count++
		_, ok := p.RequirementsFor(line)
		if !ok {
			t.Errorf("RequirementsFor(%q) = _, false; want classified (true)", line)
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 173 {
		t.Errorf("fixture has %d TestAcc names (SHA %s); want 173",
			count, requirementsFixtureProviderSHA)
	}
}

// requirementCases contains exact expected requirements for representative
// tests covering each rule family. Values match the approved plan exactly.
var requirementCases = []struct {
	name string
	want provider.TestRequirements
}{
	// Rule 1: IP ranges — all five modes, no auth requirements.
	{"TestAccGithubIpRangesDataSource", provider.TestRequirements{
		Modes: []string{"anonymous", "individual", "organization", "team", "enterprise"},
	}},
	// Rule 5: User SSH key — individual only, admin:public_key, user-key side effect.
	{"TestAccGithubUserSshKey", provider.TestRequirements{
		Modes:       []string{"individual"},
		Scopes:      []string{"admin:public_key"},
		SideEffects: []string{"user-key"},
	}},
	// Rule 5: User GPG key — individual only, admin:gpg_key, user-key side effect.
	{"TestAccGithubUserGpgKey", provider.TestRequirements{
		Modes:       []string{"individual"},
		Scopes:      []string{"admin:gpg_key"},
		SideEffects: []string{"user-key"},
	}},
	// Rule 5: User Codespaces secret resource — individual, codespace scope,
	// codespaces-user-secrets capability, codespaces-secret side effect.
	{"TestAccGithubCodespacesUserSecret", provider.TestRequirements{
		Modes:        []string{"individual"},
		Scopes:       []string{"codespace"},
		Capabilities: []string{"codespaces-user-secrets"},
		SideEffects:  []string{"codespaces-secret"},
	}},
	// Rule 5: User Codespaces secrets data source — individual, codespace scope,
	// codespaces-user-secrets capability, no side effects (data source).
	{"TestAccGithubCodespacesUserSecretsDataSource", provider.TestRequirements{
		Modes:        []string{"individual"},
		Scopes:       []string{"codespace"},
		Capabilities: []string{"codespaces-user-secrets"},
	}},
	// Rule 4: Organization Actions permissions resource.
	{"TestAccGithubActionsOrganizationPermissions", provider.TestRequirements{
		Modes:        []string{"organization", "team", "enterprise"},
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization"},
		SideEffects:  []string{"repository", "organization-setting"},
	}},
	// Rule 4: Repository custom-property resource — organization scope required.
	{"TestAccGithubRepositoryCustomProperty", provider.TestRequirements{
		Modes:        []string{"organization", "team", "enterprise"},
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization"},
		SideEffects:  []string{"repository", "organization-custom-property"},
	}},
	// Rule 3: Team resource — team and enterprise only.
	{"TestAccGithubTeam", provider.TestRequirements{
		Modes:        []string{"team", "enterprise"},
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization"},
		SideEffects:  []string{"repository", "team"},
	}},
	// Rule 3: Team members data source — team and enterprise only, no side effects.
	{"TestAccDataSourceGithubTeamMembers", provider.TestRequirements{
		Modes:        []string{"team", "enterprise"},
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization"},
	}},
	// Rule 3: Team repositories data source — team and enterprise only, no side effects.
	{"TestAccDataSourceGithubTeamRepositories", provider.TestRequirements{
		Modes:        []string{"team", "enterprise"},
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization"},
	}},
	// Rule 2: Enterprise data source — enterprise only, admin:enterprise scope.
	{"TestAccGithubEnterpriseDataSource", provider.TestRequirements{
		Modes:  []string{"enterprise"},
		Scopes: []string{"repo", "delete_repo", "read:org", "admin:org", "admin:enterprise"},
	}},
	// Rule 2: External groups data source is EMU/enterprise-only despite its name.
	{"TestAccDataSourceGithubExternalGroups", provider.TestRequirements{
		Modes:  []string{"enterprise"},
		Scopes: []string{"repo", "delete_repo", "read:org", "admin:org", "admin:enterprise"},
	}},
}

// templateRepositoryCapabilityCases locks down the confirmed Task 2 regression:
// only the top-level repository tests that actually exercise configured template
// repositories should receive the template-repository capability. OIDC subject
// claim customization "Template" tests must not match rule 7 just because their
// names contain the word Template.
var templateRepositoryCapabilityCases = []struct {
	name string
	want provider.TestRequirements
}{
	{
		name: "TestAccGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate",
		want: provider.TestRequirements{
			Modes:       []string{"individual", "organization", "team", "enterprise"},
			Scopes:      []string{"repo", "delete_repo", "user"},
			SideEffects: []string{"repository"},
		},
	},
	{
		name: "TestAccGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDataSource",
		want: provider.TestRequirements{
			Modes:  []string{"individual", "organization", "team", "enterprise"},
			Scopes: []string{"repo", "delete_repo", "user"},
		},
	},
	{
		name: "TestAccGithubRepository",
		want: provider.TestRequirements{
			Modes:        []string{"individual", "organization", "team", "enterprise"},
			Scopes:       []string{"repo", "delete_repo", "user"},
			Capabilities: []string{"template-repository"},
			SideEffects:  []string{"repository"},
		},
	},
	{
		name: "TestAccDataSourceGithubRepository",
		want: provider.TestRequirements{
			Modes:        []string{"individual", "organization", "team", "enterprise"},
			Scopes:       []string{"repo", "delete_repo", "user"},
			Capabilities: []string{"template-repository"},
		},
	},
}

// TestRequirementsFor_exactFamilyCases asserts each approved test case returns
// the exact requirements specified in the brief.
func TestRequirementsFor_exactFamilyCases(t *testing.T) {
	p := New()
	for _, tc := range requirementCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, ok := p.RequirementsFor(tc.name)
			if !ok {
				t.Fatalf("RequirementsFor(%q) = _, false; want true", tc.name)
			}
			if !reflect.DeepEqual(got.Modes, tc.want.Modes) {
				t.Errorf("Modes = %v, want %v", got.Modes, tc.want.Modes)
			}
			if !reflect.DeepEqual(got.Scopes, tc.want.Scopes) {
				t.Errorf("Scopes = %v, want %v", got.Scopes, tc.want.Scopes)
			}
			if !reflect.DeepEqual(got.Capabilities, tc.want.Capabilities) {
				t.Errorf("Capabilities = %v, want %v", got.Capabilities, tc.want.Capabilities)
			}
			if !reflect.DeepEqual(got.SideEffects, tc.want.SideEffects) {
				t.Errorf("SideEffects = %v, want %v", got.SideEffects, tc.want.SideEffects)
			}
		})
	}
}

func TestRequirementsFor_templateRepositoryCapabilityCases(t *testing.T) {
	p := New()
	for _, tc := range templateRepositoryCapabilityCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, ok := p.RequirementsFor(tc.name)
			if !ok {
				t.Fatalf("RequirementsFor(%q) = _, false; want true", tc.name)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("RequirementsFor(%q) = %+v, want %+v", tc.name, got, tc.want)
			}
		})
	}
}

func TestRequirementsFor_returnsFreshSlices(t *testing.T) {
	p := New()

	t.Run("classified lookup", func(t *testing.T) {
		name := "TestAccGithubCodespacesUserSecret"
		original, ok := p.RequirementsFor(name)
		if !ok {
			t.Fatalf("RequirementsFor(%q) = _, false; want true", name)
		}

		original.Modes[0] = "mutated-mode"
		original.Scopes[0] = "mutated-scope"
		original.Capabilities[0] = "mutated-capability"
		original.SideEffects[0] = "mutated-side-effect"

		got, ok := p.RequirementsFor(name)
		if !ok {
			t.Fatalf("RequirementsFor(%q) second lookup = _, false; want true", name)
		}

		want := provider.TestRequirements{
			Modes:        []string{"individual"},
			Scopes:       []string{"codespace"},
			Capabilities: []string{"codespaces-user-secrets"},
			SideEffects:  []string{"codespaces-secret"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("RequirementsFor(%q) after mutation = %+v, want %+v", name, got, want)
		}
	})

	t.Run("fallback lookup", func(t *testing.T) {
		name := "TestAccUnknownDoesNotExist"
		original, ok := p.RequirementsFor(name)
		if ok {
			t.Fatalf("RequirementsFor(%q) = _, true; want false", name)
		}

		original.Modes[0] = "mutated-mode"
		original.Scopes[0] = "mutated-scope"
		original.SideEffects[0] = "mutated-side-effect"

		got, ok := p.RequirementsFor(name)
		if ok {
			t.Fatalf("RequirementsFor(%q) second lookup = _, true; want false", name)
		}

		want := conservativeRequirements()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("RequirementsFor(%q) after mutation = %+v, want %+v", name, got, want)
		}
	})
}
