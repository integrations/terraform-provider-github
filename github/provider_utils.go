package github

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

var testUser string = os.Getenv("GITHUB_TEST_USER")
var testCollaborator string = os.Getenv("GITHUB_TEST_COLLABORATOR")
var isEnterprise string = os.Getenv("ENTERPRISE_ACCOUNT")
var testOrganization string = os.Getenv("GITHUB_ORGANIZATION")
var testOwner string = os.Getenv("GITHUB_OWNER")
var testToken string = os.Getenv("GITHUB_TOKEN")

var testTokenGHES string = os.Getenv("GHES_TOKEN")
var testBaseURLGHES string = os.Getenv("GHES_BASE_URL")

func testAccPreCheckEnvironment(t *testing.T, requiredEnvironmentVariables []string) {
	for _, variable := range requiredEnvironmentVariables {
		if v := os.Getenv(variable); v == "" {
			t.Fatal(variable + " must be set for acceptance tests")
		}
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GITHUB_TOKEN"); v == "" {
		t.Fatal("GITHUB_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("GITHUB_ORGANIZATION"); v == "" && os.Getenv("GITHUB_OWNER") == "" {
		t.Fatal("GITHUB_ORGANIZATION or GITHUB_OWNER must be set for acceptance tests")
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

func testAccCheckOrganization() error {

	baseURL := os.Getenv("GITHUB_BASE_URL")
	token := os.Getenv("GITHUB_TOKEN")

	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		organization := os.Getenv("GITHUB_ORGANIZATION")
		if organization == "" {
			return fmt.Errorf("Neither `GITHUB_OWNER` or `GITHUB_ORGANIZATION` set in environment")
		}
		owner = organization
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

func githubTLSApiMock(port, certFile, keyFile string, t *testing.T) (string, func() error) {
	mux := http.NewServeMux()

	userPattern := "/v3/users/hashibot"
	orgPattern := "/v3/orgs/" + testOwner

	mux.HandleFunc(userPattern, testRespondJson(userResponseBody))
	mux.HandleFunc(userPattern+"/gpg_keys", testRespondJson(gpgKeysResponseBody))
	mux.HandleFunc(userPattern+"/keys", testRespondJson(keysResponseBody))
	mux.HandleFunc(orgPattern, testRespondJson(orgResponseBody(port)))

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
	url := fmt.Sprintf(`https://localhost:%s/v3/orgs/%s`, port, testOwner)
	return fmt.Sprintf(`
{
	"login": "%s",
	"url" : "%s",
	"repos_url": "%s/repos"
}
`, testOwner, url, url)
}

func OwnerOrOrgEnvDefaultFunc() (interface{}, error) {
	if organization := os.Getenv("GITHUB_ORGANIZATION"); organization != "" {
		return os.Getenv("GITHUB_ORGANIZATION"), nil
	}
	return os.Getenv("GITHUB_OWNER"), nil
}
