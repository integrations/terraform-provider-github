package github

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubOrganizationWebhook_basic(t *testing.T) {
	var hook github.Hook

	rn := "github_organization_webhook.foo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubOrganizationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationWebhookConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubOrganizationWebhookExists(rn, &hook),
					testAccCheckGithubOrganizationWebhookAttributes(&hook, &testAccGithubOrganizationWebhookExpectedAttributes{
						Events: []string{"pull_request"},
						Configuration: map[string]interface{}{
							"url":          "https://google.de/webhook",
							"content_type": "json",
							"insecure_ssl": "true",
						},
						Active: true,
					}),
				),
			},
			{
				Config: testAccGithubOrganizationWebhookUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubOrganizationWebhookExists(rn, &hook),
					testAccCheckGithubOrganizationWebhookAttributes(&hook, &testAccGithubOrganizationWebhookExpectedAttributes{
						Events: []string{"issues"},
						Configuration: map[string]interface{}{
							"url":          "https://google.de/webhooks",
							"content_type": "form",
							"insecure_ssl": "false",
						},
						Active: false,
					}),
				),
			},
		},
	})
}

func TestAccGithubOrganizationWebhook_secret(t *testing.T) {

	rn := "github_organization_webhook.foo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubOrganizationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationWebhookConfig_secret,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubOrganizationWebhookSecret(rn, "VerySecret"),
				),
			},
		},
	})
}

func testAccCheckGithubOrganizationWebhookExists(n string, hook *github.Hook) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		hookID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}
		if hookID == 0 {
			return fmt.Errorf("No repository name is set")
		}

		org := testAccProvider.Meta().(*Organization)
		conn := org.v3client
		getHook, _, err := conn.Organizations.GetHook(context.TODO(), org.name, hookID)
		if err != nil {
			return err
		}
		*hook = *getHook
		return nil
	}
}

type testAccGithubOrganizationWebhookExpectedAttributes struct {
	Events        []string
	Configuration map[string]interface{}
	Active        bool
}

func testAccCheckGithubOrganizationWebhookAttributes(hook *github.Hook, want *testAccGithubOrganizationWebhookExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *hook.Active != want.Active {
			return fmt.Errorf("got hook %t; want %t", *hook.Active, want.Active)
		}
		if !strings.HasPrefix(*hook.URL, "https://") {
			return fmt.Errorf("got http URL %q; want to start with 'https://'", *hook.URL)
		}
		if !reflect.DeepEqual(hook.Events, want.Events) {
			return fmt.Errorf("got hook events %q; want %q", hook.Events, want.Events)
		}
		if !reflect.DeepEqual(hook.Config, want.Configuration) {
			return fmt.Errorf("got hook configuration %q; want %q", hook.Config, want.Configuration)
		}

		return nil
	}
}

func testAccCheckGithubOrganizationWebhookSecret(r, secret string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("Not Found: %s", r)
		}

		if rs.Primary.Attributes["configuration.0.secret"] != secret {
			return fmt.Errorf("Configured secret in %s does not match secret in state.  (Expected: %s, Actual: %s)", r, secret, rs.Primary.Attributes["configuration.0.secret"])
		}

		return nil
	}
}

func testAccCheckGithubOrganizationWebhookDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client
	orgName := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_organization_webhook" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		gotHook, resp, err := conn.Organizations.GetHook(context.TODO(), orgName, id)
		if err == nil {
			if gotHook != nil && *gotHook.ID == id {
				return fmt.Errorf("Webhook still exists")
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

const testAccGithubOrganizationWebhookConfig = `
resource "github_organization_webhook" "foo" {
  configuration {
    url = "https://google.de/webhook"
    content_type = "json"
    insecure_ssl = true
  }

  events = ["pull_request"]
}
`

const testAccGithubOrganizationWebhookUpdateConfig = `
resource "github_organization_webhook" "foo" {
  configuration {
    url = "https://google.de/webhooks"
    content_type = "form"
    insecure_ssl = false
  }
  active = false

  events = ["issues"]
}
`

const testAccGithubOrganizationWebhookConfig_secret = `
resource "github_organization_webhook" "foo" {
  configuration {
    url          = "https://www.terraform.io/webhook"
    content_type = "json"
    secret       = "VerySecret"
    insecure_ssl = false
  }

  events = ["pull_request"]
}
`
