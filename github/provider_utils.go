package github

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var testCollaborator string = os.Getenv("GITHUB_TEST_COLLABORATOR")
var testOrganization string = testOrganizationFunc()

var isEnterprise string = os.Getenv("ENTERPRISE_ACCOUNT")
var testOwner string = os.Getenv("GITHUB_OWNER")
var testToken string = os.Getenv("GITHUB_TOKEN")
var testBaseURLGHES string = os.Getenv("GHES_BASE_URL")

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GITHUB_TOKEN"); v == "" {
		t.Fatal("GITHUB_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_OWNER"); v == "" {
		t.Fatal("GITHUB_OWNER must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_TEST_ORGANIZATION"); v == "" {
		t.Fatal("GITHUB_TEST_ORGANIZATION must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_TEST_USER"); v == "" {
		t.Fatal("GITHUB_TEST_USER must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_TEST_COLLABORATOR"); v == "" {
		t.Fatal("GITHUB_TEST_COLLABORATOR must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_TEMPLATE_REPOSITORY"); v == "" {
		t.Fatal("GITHUB_TEMPLATE_REPOSITORY must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID"); v == "" {
		t.Fatal("GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID must be set for acceptance tests")
	}
}

func skipUnlessMode(t *testing.T, providerMode string) {
	log.Printf("[DEBUG] <<<<<<< skipUnlessMode")
	log.Printf("[DEBUG] <<<<<<< user type: %s", providerMode)
	switch providerMode {
	case anonymous:
		if os.Getenv("GITHUB_BASE_URL") != "" &&
			os.Getenv("GITHUB_BASE_URL") != "https://api.github.com/" {
			t.Log("anonymous mode not supported for GHES deployments")
			break
		}

		if os.Getenv("GITHUB_TOKEN") == "" {
			log.Printf("[DEBUG] <<<<<<< configuring anonymous user without token")
			return
		} else {
			t.Log("GITHUB_TOKEN environment variable should be empty in anonymous mode")
		}
	case individual:
		if os.Getenv("GITHUB_TOKEN") != "" && os.Getenv("GITHUB_OWNER") != "" {
			log.Printf("[DEBUG] <<<<<<< configuring user type: %s", providerMode)
			return
		} else {
			t.Logf("GITHUB_TOKEN and GITHUB_OWNER environment variables should be set for tests in %v mode", providerMode)
		}
		// TODO(kfcampbell): this is a problem. how are we going to know it's an organization
		// in acceptance tests?
		// ideas:
		// - we could perform a lookup online to determine the type
		// - we could use the GITHUB_ORG variable only for acceptance tests
		// - we could switch GITHUB_ORG to be a boolean
		// - we could deprecate GITHUB_ORG entirely and create
		// 	another environment variable boolean to use instead
	case organization:
		log.Printf("[DEBUG] <<<<<<< user type: %s", providerMode)
		if os.Getenv("GITHUB_TOKEN") != "" && os.Getenv("GITHUB_TEST_ORGANIZATION") != "" {
			log.Printf("[DEBUG] <<<<<<< configuring user type: %s", providerMode)
			return
		}
	}

	log.Printf("[DEBUG] <<<<<<< skipping!")
	t.Skipf("Skipping %s which requires %s mode", t.Name(), providerMode)
}

func testAccCheckOrganization() error {

	baseURL := os.Getenv("GITHUB_BASE_URL")
	token := os.Getenv("GITHUB_TOKEN")

	owner := os.Getenv("GITHUB_TEST_ORGANIZATION")
	if owner == "" && os.Getenv("GITHUB_OWNER") == "" {
		return fmt.Errorf("neither `GITHUB_TEST_ORGANIZATION` nor `GITHUB_OWNER` are set in environment")
	} else if owner == "" {
		owner = os.Getenv("GITHUB_OWNER")
	}

	config := Config{
		BaseURL: baseURL,
		Token:   token,
		Owner:   owner,
	}

	meta, err := config.Meta()
	if err != nil {
		return err
	}
	if !meta.(*Owner).IsOrganization {
		return fmt.Errorf("GITHUB_OWNER %q is a user, not an organization", meta.(*Owner).name)
	}
	return nil
}

func testOwnerFunc() string {
	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		owner = os.Getenv("GITHUB_TEST_OWNER")
	}
	return owner
}

// testOrganizationFunc returns a test organization. IMPORTANT:
// GITHUB_ORGANIZATION is deprecated. The purpose of this variable is
// to make sure that we can still test organization cases appropriately
// in integration testing.
func testOrganizationFunc() string {
	organization := os.Getenv("GITHUB_TEST_ORGANIZATION")
	if organization == "" {
		organization = os.Getenv("GITHUB_OWNER")
	}
	return organization
}

const anonymous = "anonymous"
const individual = "individual"
const organization = "organization"
