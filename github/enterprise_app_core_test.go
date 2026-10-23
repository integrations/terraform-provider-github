package github

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func TestEnterpriseAppTerraformCoreRecovery(t *testing.T) {
	terraform := os.Getenv("TF_CORE_TEST_BINARY")
	if terraform == "" {
		t.Skip("set TF_CORE_TEST_BINARY to run the offline Terraform CLI recovery test")
	}
	t.Run("successful-bootstrap", func(t *testing.T) { enterpriseAppTerraformCoreScenario(t, terraform, false) })
	t.Run("failed-bootstrap-recovery", func(t *testing.T) { enterpriseAppTerraformCoreScenario(t, terraform, true) })
}

func enterpriseAppTerraformCoreScenario(t *testing.T, terraform string, failBootstrap bool) {
	t.Helper()
	fixture := newEnterpriseLifecycleFixture(t)
	fixture.missing = !failBootstrap
	if failBootstrap {
		fixture.suspended = "2026-01-01T00:00:00Z"
	}
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{"github_enterprise_organization": resourceGithubEnterpriseOrganization()},
		ConfigureContextFunc: func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
			return fixture.owner, nil
		},
	}
	serverContext, cancel := context.WithCancel(t.Context())
	reattach, closed, err := plugin.DebugServe(serverContext, &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider { return provider }, NoLogOutputOverride: true,
	})
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	defer func() {
		cancel()
		select {
		case <-closed:
		case <-time.After(10 * time.Second):
			t.Error("local provider server did not stop")
		}
	}()
	attachment, err := json.Marshal(map[string]plugin.ReattachConfig{"registry.terraform.io/integrations/github": reattach})
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	configuration := `terraform {
  required_providers {
    github = { source = "integrations/github" }
  }
}
resource "github_enterprise_organization" "pilot" {
  enterprise_id = "E_test"
  name = "test-org"
  display_name = "Test Org"
  description = "managed"
  billing_email = "billing@example.test"
  admin_logins = ["owner-one", "owner-two"]
  lifecycle { prevent_destroy = true }
}
`
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(configuration), 0o600); err != nil {
		t.Fatal(err)
	}
	cliConfig := filepath.Join(root, "test.tfrc")
	if err := os.WriteFile(cliConfig, []byte("provider_installation { dev_overrides { \"integrations/github\" = "+strconv.Quote(filepath.ToSlash(root))+" } direct {} }"), 0o600); err != nil {
		t.Fatal(err)
	}
	environment := []string{}
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		upper := strings.ToUpper(name)
		if !strings.HasPrefix(upper, "TF_") && !strings.HasPrefix(upper, "TFC_") && !strings.Contains(upper, "GITHUB") && !strings.Contains(upper, "APP_PRIVATE_KEY") {
			environment = append(environment, entry)
		}
	}
	environment = append(environment, "TF_REATTACH_PROVIDERS="+string(attachment), "TF_CLI_CONFIG_FILE="+cliConfig, "TF_INPUT=0", "CHECKPOINT_DISABLE=1")
	run := func(wantSuccess bool, arguments ...string) string {
		t.Helper()
		commandContext, stop := context.WithTimeout(t.Context(), 30*time.Second)
		defer stop()
		command := exec.CommandContext(commandContext, terraform, append([]string{"-chdir=" + root}, arguments...)...)
		command.Env = environment
		output, commandErr := command.CombinedOutput()
		if (commandErr == nil) != wantSuccess {
			t.Fatalf("terraform %v: %v\n%s", arguments, commandErr, output)
		}
		return string(output)
	}
	output := run(!failBootstrap, "apply", "-auto-approve", "-no-color")
	if failBootstrap && !strings.Contains(output, "ID retained") {
		t.Fatalf("unexpected create failure: %s", output)
	}
	stateBytes, err := os.ReadFile(filepath.Join(root, "terraform.tfstate"))
	if err != nil {
		t.Fatal(err)
	}
	var state struct {
		Resources []struct {
			Instances []struct {
				Status     string `json:"status"`
				Attributes struct {
					ID         string `json:"id"`
					DatabaseID int64  `json:"database_id"`
				} `json:"attributes"`
			} `json:"instances"`
		} `json:"resources"`
	}
	if err := json.Unmarshal(stateBytes, &state); err != nil {
		t.Fatal(err)
	}
	if len(state.Resources) != 1 || len(state.Resources[0].Instances) != 1 {
		t.Fatalf("missing partial-create state: %s", stateBytes)
	}
	instance := state.Resources[0].Instances[0]
	wantStatus, wantInstalls, wantProfiles := "", 1, 1
	if failBootstrap {
		wantStatus, wantInstalls, wantProfiles = "tainted", 0, 0
	}
	if instance.Status != wantStatus || instance.Attributes.ID != "O_test" || instance.Attributes.DatabaseID != 42 {
		t.Fatalf("unexpected partial-create identity/status: %s", stateBytes)
	}
	if failBootstrap {
		fixture.suspended = nil
		output = run(false, "plan", "-no-color")
		if !strings.Contains(output, "prevent_destroy") {
			t.Fatalf("tainted replacement was not guarded: %s", output)
		}
		run(true, "untaint", "-no-color", "github_enterprise_organization.pilot")
	}
	output = run(true, "plan", "-no-color", "-detailed-exitcode")
	if !strings.Contains(output, "No changes.") || fixture.mutations != 1 || fixture.installPosts != wantInstalls || fixture.profileWrites != wantProfiles {
		t.Fatalf("recovery repeated a mutation or retained drift: %s; creates=%d installs=%d profile=%d", output, fixture.mutations, fixture.installPosts, fixture.profileWrites)
	}
	t.Log("PASS: Terraform Core state identity, bootstrap mutation counts and stable subsequent plan verified")
}
