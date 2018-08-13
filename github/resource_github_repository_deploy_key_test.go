package github

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestSuppressDeployKeyDiff(t *testing.T) {
	oldV := "ssh-rsa AAAABB...cd+=="
	newV := "ssh-rsa AAAABB...cd+== terraform-acctest@hashicorp.com\n"
	if !suppressDeployKeyDiff("test", oldV, newV, nil) {
		t.Fatalf("Expected %q and %q to be suppressed", oldV, newV)
	}

	oldV = "ssh-rsa AAAABB...cd+=="
	newV = "ssh-rsa AAAABB...cd+=="
	if !suppressDeployKeyDiff("test", oldV, newV, nil) {
		t.Fatalf("Expected %q and %q to be suppressed", oldV, newV)
	}

	oldV = "ssh-rsa AAAABV...cd+=="
	newV = "ssh-rsa DIFFERENT...cd+=="
	if suppressDeployKeyDiff("test", oldV, newV, nil) {
		t.Fatalf("Expected %q and %q NOT to be suppressed", oldV, newV)
	}
}

func TestAccGithubRepositoryDeployKey_basic(t *testing.T) {
	rs := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	repositoryName := fmt.Sprintf("acctest-%s", rs)
	keyPath := filepath.Join("test-fixtures", "id_rsa.pub")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDeployKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryDeployKeyConfig(repositoryName, keyPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryDeployKeyExists("github_repository_deploy_key.test_repo_deploy_key"),
					resource.TestCheckResourceAttr("github_repository_deploy_key.test_repo_deploy_key", "read_only", "false"),
					resource.TestCheckResourceAttr("github_repository_deploy_key.test_repo_deploy_key", "repository", repositoryName),
					resource.TestMatchResourceAttr("github_repository_deploy_key.test_repo_deploy_key", "key", regexp.MustCompile(`^ssh-rsa [^\s]+$`)),
					resource.TestCheckResourceAttr("github_repository_deploy_key.test_repo_deploy_key", "title", "title"),
				),
			},
		},
	})
}

func TestAccGithubRepositoryDeployKey_importBasic(t *testing.T) {
	rs := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	repositoryName := fmt.Sprintf("acctest-%s", rs)
	keyPath := filepath.Join("test-fixtures", "id_rsa.pub")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDeployKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryDeployKeyConfig(repositoryName, keyPath),
			},
			{
				ResourceName:      "github_repository_deploy_key.test_repo_deploy_key",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGithubRepositoryDeployKeyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository_deploy_key" {
			continue
		}

		orgName := testAccProvider.Meta().(*Organization).name
		repoName, idString, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(idString, err)
		}

		_, resp, err := conn.Repositories.GetKey(context.TODO(), orgName, repoName, id)

		if err != nil && resp.Response.StatusCode != 404 {
			return err
		}
		return nil
	}

	return nil
}

func testAccCheckGithubRepositoryDeployKeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, idString, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(idString, err)
		}

		_, _, err = conn.Repositories.GetKey(context.TODO(), orgName, repoName, id)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccGithubRepositoryDeployKeyConfig(name, keyPath string) string {
	return fmt.Sprintf(`
  resource "github_repository" "test_repo" {
		name = "%s"
	}

	resource "github_repository_deploy_key" "test_repo_deploy_key" {
    key = "${file("%s")}"
    read_only = "false"
    repository = "${github_repository.test_repo.name}"
    title = "title"
  }
`, name, keyPath)
}
