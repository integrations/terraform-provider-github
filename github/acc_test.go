package github

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type testMode string

const (
	anonymous    testMode = "anonymous"
	individual   testMode = "individual"
	organization testMode = "organization"
	team         testMode = "team"
	enterprise   testMode = "enterprise"
)

const testResourcePrefix = "tf-acc-test-"

var (
	orgTestModes     = []testMode{organization, team, enterprise}
	paidOrgTestModes = []testMode{team, enterprise}
)

type testAccConfig struct {
	// Target configuration
	baseURL *url.URL

	// Auth configuration
	authMode testMode
	owner    string
	username string
	token    string

	// Enterprise configuration
	enterpriseSlug  string
	enterpriseIsEMU bool

	// Global test configuration
	testPublicRepository              string
	testPublicRepositoryOwner         string
	testPublicReleaseId               int
	testPublicRelaseAssetId           string
	testPublicRelaseAssetName         string
	testPublicReleaseAssetContent     string
	testPublicTemplateRepository      string
	testPublicTemplateRepositoryOwner string
	testGHActionsAppInstallationId    int

	// User test configuration
	testUserRepository string

	// Org test configuration
	testOrgUser               string
	testOrgSecretName         string
	testOrgRepository         string
	testOrgTemplateRepository string
	testOrgAppInstallationId  int

	// External test configuration
	testExternalUser      string
	testExternalUserToken string
	testExternalUser2     string

	// Enterprise test configuration
	testEnterpriseEMUGroupId int

	// Test options
	testAdvancedSecurity bool

	// Test repository configuration
	testRepositoryVisibility string
}

var testAccConf *testAccConfig

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	//nolint:unparam
	"github": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func TestMain(m *testing.M) {
	authMode := testMode(os.Getenv("GH_TEST_AUTH_MODE"))
	if len(authMode) == 0 {
		authMode = anonymous
	}

	u, ok := os.LookupEnv("GITHUB_BASE_URL")
	if !ok {
		u = DotComAPIURL
	}

	baseURL, err := url.Parse(u)
	if err != nil {
		fmt.Printf("Error parsing base URL: %s\n", err)
		os.Exit(1)
	}

	config := testAccConfig{
		baseURL:                   baseURL,
		authMode:                  authMode,
		testPublicRepository:      "terraform-provider-github",
		testPublicRepositoryOwner: "integrations",
		testPublicReleaseId:       186531906,
		// The terraform-provider-github_6.4.0_manifest.json asset ID from
		// https://github.com/integrations/terraform-provider-github/releases/tag/v6.4.0
		testPublicRelaseAssetId:           "207956097",
		testPublicRelaseAssetName:         "terraform-provider-github_6.4.0_manifest.json",
		testPublicReleaseAssetContent:     "{\n  \"version\": 1,\n  \"metadata\": {\n    \"protocol_versions\": [\n      \"5.0\"\n    ]\n  }\n}",
		testPublicTemplateRepository:      "template-repository",
		testPublicTemplateRepositoryOwner: "template-repository",
		testGHActionsAppInstallationId:    15368,
		testUserRepository:                os.Getenv("GH_TEST_USER_REPOSITORY"),
		testOrgUser:                       os.Getenv("GH_TEST_ORG_USER"),
		testOrgSecretName:                 os.Getenv("GH_TEST_ORG_SECRET_NAME"),
		testOrgRepository:                 os.Getenv("GH_TEST_ORG_REPOSITORY"),
		testOrgTemplateRepository:         os.Getenv("GH_TEST_ORG_TEMPLATE_REPOSITORY"),
		testExternalUser:                  os.Getenv("GH_TEST_EXTERNAL_USER"),
		testExternalUserToken:             os.Getenv("GH_TEST_EXTERNAL_USER_TOKEN"),
		testExternalUser2:                 os.Getenv("GH_TEST_EXTERNAL_USER2"),
		testAdvancedSecurity:              os.Getenv("GH_TEST_ADVANCED_SECURITY") == "true",
		testRepositoryVisibility:          "public",
		enterpriseIsEMU:                   authMode == enterprise && os.Getenv("GH_TEST_ENTERPRISE_IS_EMU") == "true",
	}

	if config.authMode != anonymous {
		config.owner = os.Getenv("GITHUB_OWNER")
		config.username = os.Getenv("GITHUB_USERNAME")
		config.token = os.Getenv("GITHUB_TOKEN")

		if len(config.owner) == 0 {
			fmt.Println("GITHUB_OWNER environment variable not set")
			os.Exit(1)
		}

		if len(config.username) == 0 {
			fmt.Println("GITHUB_USERNAME environment variable not set")
			os.Exit(1)
		}

		if len(config.token) == 0 {
			fmt.Println("GITHUB_TOKEN environment variable not set")
			os.Exit(1)
		}
	}

	if config.authMode == enterprise {
		config.enterpriseSlug = os.Getenv("GITHUB_ENTERPRISE_SLUG")

		if len(config.enterpriseSlug) == 0 {
			fmt.Println("GITHUB_ENTERPRISE_SLUG environment variable not set")
			os.Exit(1)
		}

		i, err := strconv.Atoi(os.Getenv("GH_TEST_ENTERPRISE_EMU_GROUP_ID"))
		if err == nil {
			config.testEnterpriseEMUGroupId = i
		}

		if config.enterpriseIsEMU {
			config.testRepositoryVisibility = "private"
		}
	}

	i, err := strconv.Atoi(os.Getenv("GH_TEST_ORG_APP_INSTALLATION_ID"))
	if err == nil {
		config.testOrgAppInstallationId = i
	}

	testAccConf = &config

	configureSweepers()

	resource.TestMain(m)
}

func getTestMeta() (*Owner, error) {
	config := Config{
		Token:   testAccConf.token,
		Owner:   testAccConf.owner,
		BaseURL: testAccConf.baseURL,
	}

	meta, err := config.Meta()
	if err != nil {
		return nil, fmt.Errorf("error getting GitHub meta parameter")
	}

	return meta.(*Owner), nil
}

func configureSweepers() {
	resource.AddTestSweepers("repositories", &resource.Sweeper{
		Name: "repositories",
		F:    sweepRepositories,
	})

	resource.AddTestSweepers("teams", &resource.Sweeper{
		Name: "teams",
		F:    sweepTeams,
	})
}

func sweepTeams(_ string) error {
	if !slices.Contains(orgTestModes, testMode(os.Getenv("GH_TEST_AUTH_MODE"))) {
		return nil
	}

	fmt.Println("sweeping teams")

	meta, err := getTestMeta()
	if err != nil {
		return fmt.Errorf("could not get test meta for sweeper: %w", err)
	}

	client := meta.v3client
	owner := meta.name
	ctx := context.Background()

	teams, _, err := client.Teams.ListTeams(ctx, owner, nil)
	if err != nil {
		return err
	}

	for _, t := range teams {
		if slug := t.GetSlug(); strings.HasPrefix(slug, testResourcePrefix) {
			fmt.Printf("destroying team %s\n", slug)

			if _, err := client.Teams.DeleteTeamBySlug(ctx, owner, slug); err != nil {
				return err
			}
		}
	}

	return nil
}

func sweepRepositories(_ string) error {
	fmt.Println("sweeping repositories")

	meta, err := getTestMeta()
	if err != nil {
		return fmt.Errorf("could not get test meta for sweeper: %w", err)
	}

	client := meta.v3client
	owner := meta.name
	ctx := context.Background()

	var repos []*github.Repository
	var err2 error
	if slices.Contains(orgTestModes, testMode(os.Getenv("GH_TEST_AUTH_MODE"))) {
		repos, _, err2 = client.Repositories.ListByOrg(ctx, owner, nil)
	} else {
		repos, _, err2 = client.Repositories.ListByUser(ctx, owner, nil)
	}
	if err2 != nil {
		return err2
	}

	for _, r := range repos {
		if name := r.GetName(); strings.HasPrefix(name, testResourcePrefix) {
			fmt.Printf("destroying repository %s\n", name)

			if _, err := client.Repositories.Delete(ctx, owner, name); err != nil {
				return err
			}
		}
	}

	return nil
}

func skipUnauthenticated(t *testing.T) {
	if testAccConf.authMode == anonymous {
		t.Skip("Skipping as test mode not authenticated")
	}
}

func skipUnlessHasOrgs(t *testing.T) {
	if !slices.Contains(orgTestModes, testAccConf.authMode) {
		t.Skip("Skipping as test mode doesn't have orgs")
	}
}

func skipUnlessHasPaidOrgs(t *testing.T) {
	if !slices.Contains(paidOrgTestModes, testAccConf.authMode) {
		t.Skip("Skipping as test mode doesn't have paid orgs")
	}
}

func skipUnlessEnterprise(t *testing.T) {
	if testAccConf.authMode != enterprise {
		t.Skip("Skipping as test mode is not enterprise")
	}
}

func skipUnlessHasAppInstallations(t *testing.T) {
	t.Helper()

	meta, err := getTestMeta()
	if err != nil {
		t.Fatalf("failed to get test meta: %s", err)
	}

	installations, _, err := meta.v3client.Organizations.ListInstallations(context.Background(), meta.name, nil)
	if err != nil {
		t.Fatalf("failed to list app installations: %s", err)
	}

	if len(installations.Installations) == 0 {
		t.Skip("Skipping because no GitHub App installations found in the test organization")
	}
}

func skipUnlessEMUEnterprise(t *testing.T) {
	if !testAccConf.enterpriseIsEMU {
		t.Skip("Skipping as test mode is not EMU enterprise")
	}
}

func skipIfEMUEnterprise(t *testing.T) {
	if testAccConf.enterpriseIsEMU {
		t.Skip("Skipping as this test is not supported for EMU enterprise")
	}
}

func skipUnlessMode(t *testing.T, testModes ...testMode) {
	if !slices.Contains(testModes, testAccConf.authMode) {
		t.Skip("Skipping as not supported test mode")
	}
}
