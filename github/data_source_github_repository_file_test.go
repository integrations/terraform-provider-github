package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestAccGithubRepositoryFileDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("read files", func(t *testing.T) {
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}

			data "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = github_repository_file.test.file
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "content",
				"bar",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "sha",
				"ba0e162e1c47469e3fe4b393a8bf8c569f302116",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_author",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_email",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_message",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_sha",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}

func TestDataSourceGithubRepositoryFileRead(t *testing.T) {
	// helper function to simplify marshalling.
	marshal := func(t *testing.T, msg interface{}) string {
		data, err := json.MarshalIndent(msg, "", "    ")
		if err != nil {
			t.Fatalf("cant encode to json: %v", err)
		}

		return string(data)
	}

	sha := "some-test-sha"
	committerName := "some-test-user"
	committerEmail := "some-test-user@github.com"
	commitMessage := "test commit message"

	enc := "base64"
	fileContent := "here-goes-content-of-our-glorious-config.json"
	b64FileContent := base64.StdEncoding.EncodeToString([]byte(fileContent))

	fileName := "test-file.json"
	branch := "main"

	// setting up some org/owner info
	owner := "test-owner"
	org := "test-org"
	repo := "test-repo"

	// preparing mashalled objects
	branchRespBody := marshal(t, &github.Branch{Name: &branch})
	repoContentRespBody := marshal(t, &github.RepositoryContent{
		Encoding: &enc,
		Content:  &b64FileContent,
		SHA:      &sha,
	})
	repoCommitRespBody := marshal(t, &github.RepositoryCommit{
		SHA: &sha,
		Committer: &github.User{
			Name:  &committerName,
			Email: &committerEmail,
		},
		Commit: &github.Commit{
			Message: &commitMessage,
		},
	})

	t.Run("extracting org and repo if full_name is passed", func(t *testing.T) {
		// test setup
		repositoryFullName := fmt.Sprintf("%s/%s", org, repo)
		expectedID := fmt.Sprintf("%s/%s", repo, fileName)
		expectedRepo := "test-repo"

		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/branches/%s", org, repo, branch),
				ResponseBody: branchRespBody,
				StatusCode:   http.StatusOK,
			},
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/contents/%s?ref=%s", org, repo, fileName, branch),
				ResponseBody: repoContentRespBody,
				StatusCode:   http.StatusOK,
			},
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/commits/%s", org, repo, sha),
				ResponseBody: repoCommitRespBody,
				StatusCode:   http.StatusOK,
			},
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/commits", org, repo),
				ResponseBody: repoCommitRespBody,
				StatusCode:   http.StatusOK,
			},
		})
		defer ts.Close()

		httpCl := http.DefaultClient
		httpCl.Transport = http.DefaultTransport

		client := github.NewClient(httpCl)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		meta := &Owner{
			name:     owner,
			v3client: client,
		}

		testSchema := map[string]*schema.Schema{
			"repository": {Type: schema.TypeString},
			"file":       {Type: schema.TypeString},
			"branch":     {Type: schema.TypeString},
			"commit_sha": {Type: schema.TypeString},
			"content":    {Type: schema.TypeString},
			"id":         {Type: schema.TypeString},
		}

		schema := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
			"repository": repositoryFullName,
			"file":       fileName,
			"branch":     branch,
			"commit_sha": sha,
			"content":    "",
			"id":         "",
		})

		// actual call
		err := dataSourceGithubRepositoryFileRead(schema, meta)

		// assertions
		assert.Nil(t, err)
		assert.Equal(t, expectedRepo, schema.Get("repository"))
		assert.Equal(t, fileContent, schema.Get("content"))
		assert.Equal(t, expectedID, schema.Get("id"))
	})
	t.Run("using user as owner if just name is passed", func(t *testing.T) {
		// test setup
		repositoryFullName := repo
		expectedID := fmt.Sprintf("%s/%s", repo, fileName)
		expectedRepo := "test-repo"

		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/branches/%s", owner, repo, branch),
				ResponseBody: branchRespBody,
				StatusCode:   http.StatusOK,
			},
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/contents/%s?ref=%s", owner, repo, fileName, branch),
				ResponseBody: repoContentRespBody,
				StatusCode:   http.StatusOK,
			},
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/commits/%s", owner, repo, sha),
				ResponseBody: repoCommitRespBody,
				StatusCode:   http.StatusOK,
			},
			{
				ExpectedUri:  fmt.Sprintf("/repos/%s/%s/commits", owner, repo),
				ResponseBody: repoCommitRespBody,
				StatusCode:   http.StatusOK,
			},
		})
		defer ts.Close()

		httpCl := http.DefaultClient
		httpCl.Transport = http.DefaultTransport

		client := github.NewClient(httpCl)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		meta := &Owner{
			name:     owner,
			v3client: client,
		}

		testSchema := map[string]*schema.Schema{
			"repository": {Type: schema.TypeString},
			"file":       {Type: schema.TypeString},
			"branch":     {Type: schema.TypeString},
			"commit_sha": {Type: schema.TypeString},
			"content":    {Type: schema.TypeString},
			"id":         {Type: schema.TypeString},
		}

		schema := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
			"repository": repositoryFullName,
			"file":       fileName,
			"branch":     branch,
			"commit_sha": sha,
			"content":    "",
			"id":         "",
		})

		// actual call
		err := dataSourceGithubRepositoryFileRead(schema, meta)

		// assertions
		assert.Nil(t, err)
		assert.Equal(t, expectedRepo, schema.Get("repository"))
		assert.Equal(t, fileContent, schema.Get("content"))
		assert.Equal(t, expectedID, schema.Get("id"))
	})
}
