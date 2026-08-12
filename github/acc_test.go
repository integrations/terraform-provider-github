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

	"github.com/google/go-github/v89/github"
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
	legacyClient bool
	baseURL      *url.URL
	isGHES       bool

	// Auth configuration
	authMode          testMode
	owner             string
	token             string
	appID             string
	appInstallationID string
	appPEM            string

	// Enterprise configuration
	enterpriseSlug  string
	enterpriseIsEMU bool

	// Global test configuration
	testPublicRepository              string
	testPublicRepositoryOwner         string
	testPublicReleaseId               int
	testPublicReleaseAssetId          string
	testPublicReleaseAssetName        string
	testPublicReleaseAssetContent     string
	testPublicTemplateRepository      string
	testPublicTemplateRepositoryOwner string
	testGHActionsAppInstallationId    int

	// Org test configuration
	testOrgUser1             string
	testOrgUser2             string
	testOrgUser3             string
	testOrgAppInstallationId int

	// External test configuration
	testExternalUser1      string
	testExternalUser1Token string
	testExternalUser2      string

	// Enterprise test configuration
	testEnterpriseEMUGroupId      int
	testExternalGroup1ID          int
	testExternalGroup1DisplayName string
	testExternalGroup2ID          int

	// Test options
	testAdvancedSecurity bool

	// Test repository configuration
	testRepositoryVisibility string

	// Provider metadata
	meta *Owner
}

var testAccConf *testAccConfig

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	//nolint:unparam
	"github": func() (*schema.Provider, error) {
		return NewProvider("acctest", "none")(), nil
	},
}

func TestMain(m *testing.M) {
	if os.Getenv("TF_ACC") == "" {
		os.Exit(m.Run())
	}

	authMode := testMode(os.Getenv("GH_TEST_AUTH_MODE"))
	if len(authMode) == 0 {
		authMode = anonymous
	}

	u, ok := os.LookupEnv("GITHUB_BASE_URL")
	if !ok {
		u = DotComAPIURL
	}

	baseURL, isGHES, err := getBaseURL(u)
	if err != nil {
		fmt.Printf("Error parsing base URL: %s\n", err)
		os.Exit(1)
	}

	conf := &testAccConfig{
		legacyClient:                      os.Getenv("GITHUB_LEGACY_CLIENT") != "false",
		baseURL:                           baseURL,
		isGHES:                            isGHES,
		authMode:                          authMode,
		testPublicRepository:              "terraform-provider-github",
		testPublicRepositoryOwner:         "integrations",
		testPublicReleaseId:               186531906, // The terraform-provider-github_6.4.0_manifest.json asset ID from https://github.com/integrations/terraform-provider-github/releases/tag/v6.4.0
		testPublicReleaseAssetId:          "207956097",
		testPublicReleaseAssetName:        "terraform-provider-github_6.4.0_manifest.json",
		testPublicReleaseAssetContent:     "{\n  \"version\": 1,\n  \"metadata\": {\n    \"protocol_versions\": [\n      \"5.0\"\n    ]\n  }\n}",
		testPublicTemplateRepository:      "template-repository",
		testPublicTemplateRepositoryOwner: "template-repository",
		testGHActionsAppInstallationId:    15368,
		testOrgUser1:                      os.Getenv("GH_TEST_ORG_USER1"),
		testOrgUser2:                      os.Getenv("GH_TEST_ORG_USER2"),
		testOrgUser3:                      os.Getenv("GH_TEST_ORG_USER3"),
		testExternalUser1:                 os.Getenv("GH_TEST_EXTERNAL_USER1"),
		testExternalUser1Token:            os.Getenv("GH_TEST_EXTERNAL_USER1_TOKEN"),
		testExternalUser2:                 os.Getenv("GH_TEST_EXTERNAL_USER2"),
		testAdvancedSecurity:              os.Getenv("GH_TEST_ADVANCED_SECURITY") == "true",
		testRepositoryVisibility:          "public",
	}

	if conf.authMode != anonymous {
		conf.owner = os.Getenv("GITHUB_OWNER")
		conf.token = os.Getenv("GITHUB_TOKEN")
		conf.appID = os.Getenv("GITHUB_APP_ID")
		conf.appInstallationID = os.Getenv("GITHUB_APP_INSTALLATION_ID")
		conf.appPEM = os.Getenv("GITHUB_APP_PEM_FILE")

		if len(conf.owner) == 0 {
			fmt.Println("GITHUB_OWNER environment variable not set")
			os.Exit(1)
		}

		if conf.token == "" && conf.appID == "" {
			fmt.Println("authentication not configured")
			os.Exit(1)
		}

		if conf.token != "" && conf.appID != "" {
			fmt.Println("Both token and app auth configured")
			os.Exit(1)
		}

		if conf.appID != "" && (conf.appInstallationID == "" || conf.appPEM == "") {
			fmt.Println("App auth configured without all required parameters")
			os.Exit(1)
		}
	}

	if conf.authMode != anonymous && conf.authMode != individual {
		if i, err := strconv.Atoi(os.Getenv("GH_TEST_ORG_APP_INSTALLATION_ID")); err == nil {
			conf.testOrgAppInstallationId = i
		}
	}

	if conf.authMode == enterprise {
		conf.enterpriseSlug = os.Getenv("GITHUB_ENTERPRISE_SLUG")

		if len(conf.enterpriseSlug) == 0 {
			fmt.Println("GITHUB_ENTERPRISE_SLUG environment variable not set")
			os.Exit(1)
		}

		if os.Getenv("GH_TEST_ENTERPRISE_IS_EMU") == "true" {
			conf.enterpriseIsEMU = true
			conf.testRepositoryVisibility = "private"

			if i, err := strconv.Atoi(os.Getenv("GH_TEST_ENTERPRISE_EMU_GROUP_ID")); err == nil {
				conf.testEnterpriseEMUGroupId = i
			}
		}
	}

	meta, err := getTestMeta(conf)
	if err != nil {
		fmt.Printf("Error configuring provider meta: %s\n", err)
		os.Exit(1)
	}
	conf.meta = meta

	testAccConf = conf

	configureSweepers()

	resource.TestMain(m)
}

func getTestAppToken(conf *testAccConfig) (string, error) {
	if conf.appID == "" || conf.appInstallationID == "" || conf.appPEM == "" {
		return "", fmt.Errorf("app auth not configured")
	}

	restAPIPath := ""
	if conf.isGHES {
		restAPIPath = GHESRESTAPIPath
	}

	appToken, err := GenerateOAuthTokenFromApp(conf.baseURL.JoinPath(restAPIPath), conf.appID, conf.appInstallationID, conf.appPEM)
	if err != nil {
		return "", err
	}

	return appToken, nil
}

func getTestToken(conf *testAccConfig) (string, error) {
	if conf.token != "" {
		return conf.token, nil
	}

	return getTestAppToken(conf)
}

func getTestMeta(conf *testAccConfig) (*Owner, error) {
	config := &Config{
		LegacyClient: conf.legacyClient,
		BaseURL:      conf.baseURL,
		IsGHES:       conf.isGHES,
		Owner:        conf.owner,
	}

	if config.LegacyClient || conf.appID == "" {
		token, err := getTestToken(conf)
		if err != nil {
			return nil, fmt.Errorf("error getting test token: %w", err)
		}
		config.Token = token
	} else {
		config.AppID = &conf.appID
		config.AppInstallationID = &conf.appInstallationID
		config.AppPEM = []byte(conf.appPEM)
	}

	return configureProviderMeta(context.Background(), "test", config)
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

	client := testAccConf.meta.v3client
	owner := testAccConf.meta.name
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

	client := testAccConf.meta.v3client
	owner := testAccConf.meta.name
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

func skipApp(t *testing.T) {
	if testAccConf.appID != "" {
		t.Skip("Skipping as app not configured")
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

	installations, _, err := testAccConf.meta.v3client.Organizations.ListInstallations(t.Context(), testAccConf.meta.name, nil)
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

func skipUnlessHasOrgUser1(t *testing.T) {
	if testAccConf.testOrgUser1 == "" {
		t.Skip("Skipping as no test org user is configured")
	}
}

func skipUnlessHasOrgUser2(t *testing.T) {
	if testAccConf.testOrgUser2 == "" {
		t.Skip("Skipping as no test org user 2 is configured")
	}
}

// func skipUnlessHasOrgUser3(t *testing.T) {
// 	if testAccConf.testOrgUser3 == "" {
// 		t.Skip("Skipping as no test org user 3 is configured")
// 	}
// }
