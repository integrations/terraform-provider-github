package github

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testUser string = os.Getenv("GITHUB_TEST_USER")
var testCollaborator string = os.Getenv("GITHUB_TEST_COLLABORATOR")
var testOrganization string = os.Getenv("GITHUB_ORGANIZATION")

var testAccProviders map[string]terraform.ResourceProvider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"github": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"github": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GITHUB_TOKEN"); v == "" {
		t.Fatal("GITHUB_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_ORGANIZATION"); v == "" {
		t.Fatal("GITHUB_ORGANIZATION must be set for acceptance tests")
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

func TestProvider_individual(t *testing.T) {

	username := "hashibot"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubMembershipDestroy,
		Steps: []resource.TestStep{
			{
				// Test individual is true.  Because GITHUB_ORGANIZATION should be set for these tests, we'll pass an
				// empty string for `org` to unset the organization
				Config: configProviderOrganization("", true) + testAccCheckGithubUserDataSourceConfig(username),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_user.test", "name"),
					resource.TestCheckResourceAttr("data.github_user.test", "name", "HashiBot"),
				),
			},
			{
				// Test individual is true, but resource requires organization.  Because GITHUB_ORGANIZATION should be
				// set for these tests, we'll pass an empty string for `org` to unset the organization
				Config:      configProviderOrganization("", true) + testAccGithubMembershipConfig(username),
				ExpectError: regexp.MustCompile("This resource requires GitHub organization to be set on the provider."),
			},
			{
				// Test conflicting `individual` and `organization`
				Config:      configProviderOrganization(testOrganization, true) + testAccCheckGithubUserDataSourceConfig(username),
				ExpectError: regexp.MustCompile("If `individual` is true, `organization` cannot be set."),
			},
			{
				// Test neither `individual` or `organization` is set.  Because GITHUB_ORGANIZATION should be
				// set for these tests, we'll pass an empty string for `org` to unset the organization
				Config:      configProviderOrganization("", false) + testAccCheckGithubUserDataSourceConfig(username),
				ExpectError: regexp.MustCompile("If `individual` is false, `organization` is required."),
			},
		},
	})
}

func TestProvider_anonymous(t *testing.T) {

	username := "hashibot"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubMembershipDestroy,
		Steps: []resource.TestStep{
			{
				// Test anonymous is true.  Because GITHUB_TOKEN should be set for these tests, we'll pass an
				// empty string for `token` to unset the token
				Config: configProviderToken("", true) + testAccCheckGithubUserDataSourceConfig(username),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_user.test", "name"),
					resource.TestCheckResourceAttr("data.github_user.test", "name", "HashiBot"),
				),
			},
			{
				// Test conflicting `anonymous` and `token`
				Config:      configProviderToken(os.Getenv("GITHUB_TOKEN"), true) + testAccCheckGithubUserDataSourceConfig(username),
				ExpectError: regexp.MustCompile("If `anonymous` is true, `token` cannot be set."),
			},
			{
				// Test neither `anonymous` or `token` is set.  Because GITHUB_TOKEN should be
				// set for these tests, we'll pass an empty string for `token` to unset the token
				Config:      configProviderToken("", false) + testAccCheckGithubUserDataSourceConfig(username),
				ExpectError: regexp.MustCompile("If `anonymous` is false, `token` is required."),
			},
		},
	})
}

func TestProvider_insecure(t *testing.T) {
	// Use ephemeral port range (49152–65535)
	port := fmt.Sprintf("%d", 49152+rand.Intn(16382))

	// Use self-signed certificate
	certFile := filepath.Join("test-fixtures", "cert.pem")
	keyFile := filepath.Join("test-fixtures", "key.pem")

	url, closeFunc := githubTLSApiMock(port, certFile, keyFile, t)
	defer func() {
		err := closeFunc()
		if err != nil {
			t.Fatal(err)
		}
	}()

	oldBaseUrl := os.Getenv("GITHUB_BASE_URL")
	defer os.Setenv("GITHUB_BASE_URL", oldBaseUrl)

	// Point provider to mock API with self-signed cert
	os.Setenv("GITHUB_BASE_URL", url)

	insecureProviderConfig := `provider "github" {
	insecure = true
}
`

	username := "hashibot"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubUserDataSourceConfig(username),
				ExpectError: regexp.MustCompile("x509: certificate is valid for untrusted, not localhost"),
			},
			{
				Config: insecureProviderConfig + testAccCheckGithubUserDataSourceConfig(username),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_user.test", "name"),
					resource.TestCheckResourceAttr("data.github_user.test", "id", "1"),
					resource.TestCheckResourceAttr("data.github_user.test", "name", "HashiBot"),
				),
			},
		},
	})
}

func githubTLSApiMock(port, certFile, keyFile string, t *testing.T) (string, func() error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/hashibot", testRespondJson(userResponseBody))
	mux.HandleFunc("/users/hashibot/gpg_keys", testRespondJson(gpgKeysResponseBody))
	mux.HandleFunc("/users/hashibot/keys", testRespondJson(keysResponseBody))
	mux.HandleFunc("/orgs/"+testOrganization, testRespondJson(orgResponseBody(port)))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// nolint: errcheck
	go server.ListenAndServeTLS(certFile, keyFile)

	return "https://localhost:" + port + "/", server.Close
}

func testRespondJson(responseBody string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(responseBody)); err != nil {
			return
		}
	}
}

func configProviderOrganization(org string, individual bool) string {
	return fmt.Sprintf(`
provider "github" {
    organization = "%s"
    individual = %s
}
`, org, strconv.FormatBool(individual))
}

func configProviderToken(token string, anonymous bool) string {
	return fmt.Sprintf(`
provider "github" {
    token = "%s"
    anonymous = %s
}
`, token, strconv.FormatBool(anonymous))
}

const userResponseBody = `{
  "login": "hashibot",
  "id": 1,
  "node_id": "MDQ6VXNlcjE=",
  "avatar_url": "https://github.com/images/error/octocat_happy.gif",
  "gravatar_id": "",
  "url": "https://api.github.com/users/octocat",
  "html_url": "https://github.com/octocat",
  "followers_url": "https://api.github.com/users/octocat/followers",
  "following_url": "https://api.github.com/users/octocat/following{/other_user}",
  "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
  "organizations_url": "https://api.github.com/users/octocat/orgs",
  "repos_url": "https://api.github.com/users/octocat/repos",
  "events_url": "https://api.github.com/users/octocat/events{/privacy}",
  "received_events_url": "https://api.github.com/users/octocat/received_events",
  "type": "User",
  "site_admin": false,
  "name": "HashiBot",
  "company": "GitHub",
  "blog": "https://github.com/blog",
  "location": "San Francisco",
  "email": "octocat@github.com",
  "hireable": false,
  "bio": "There once was...",
  "public_repos": 2,
  "public_gists": 1,
  "followers": 20,
  "following": 0,
  "created_at": "2008-01-14T04:33:35Z",
  "updated_at": "2008-01-14T04:33:35Z"
}`

const gpgKeysResponseBody = `[
  {
    "id": 3,
    "primary_key_id": null,
    "key_id": "3262EFF25BA0D270",
    "public_key": "xsBNBFayYZ...",
    "emails": [
      {
        "email": "mastahyeti@users.noreply.github.com",
        "verified": true
      }
    ],
    "subkeys": [
      {
        "id": 4,
        "primary_key_id": 3,
        "key_id": "4A595D4C72EE49C7",
        "public_key": "zsBNBFayYZ...",
        "emails": [
        ],
        "subkeys": [
        ],
        "can_sign": false,
        "can_encrypt_comms": true,
        "can_encrypt_storage": true,
        "can_certify": false,
        "created_at": "2016-03-24T11:31:04-06:00",
        "expires_at": null
      }
    ],
    "can_sign": true,
    "can_encrypt_comms": false,
    "can_encrypt_storage": false,
    "can_certify": true,
    "created_at": "2016-03-24T11:31:04-06:00",
    "expires_at": null
  }
]`

const keysResponseBody = `[
  {
    "id": 1,
    "key": "ssh-rsa AAA..."
  }
]`

func orgResponseBody(port string) string {
	url := fmt.Sprintf(`https://localhost:%s/orgs/%s`, port, testOrganization)
	return fmt.Sprintf(`
{
	"login": "%s",
	"url" : "%s",
	"repos_url": "%s/repos"
}
`, testOrganization, url, url)
}
