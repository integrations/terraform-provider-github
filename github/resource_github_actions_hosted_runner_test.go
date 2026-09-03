package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestHostedRunnerProvisioningState(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		runner           map[string]any
		expectedUpdate   map[string]any
		requirePublicIPs bool
		wantState        string
		wantErr          bool
	}{
		"provisioning": {
			runner:    map[string]any{"status": "Provisioning"},
			wantState: "pending",
		},
		"ready without public IP requirement": {
			runner:    map[string]any{"status": "Ready"},
			wantState: "ready",
		},
		"ready before public IP allocation": {
			runner:           map[string]any{"status": "Ready", "public_ips": []any{}},
			requirePublicIPs: true,
			wantState:        "pending",
		},
		"ready before update is applied": {
			runner: map[string]any{
				"status":               "Ready",
				"machine_size_details": map[string]any{"id": "2-core"},
				"public_ip_enabled":    true,
				"public_ips":           []any{map[string]any{"prefix": "192.0.2.1"}},
			},
			expectedUpdate: map[string]any{
				"size":             "4-core",
				"enable_static_ip": true,
			},
			requirePublicIPs: true,
			wantState:        "pending",
		},
		"ready after update is applied": {
			runner: map[string]any{
				"status":               "Ready",
				"name":                 "updated",
				"machine_size_details": map[string]any{"id": "4-core"},
				"runner_group_id":      float64(2),
				"maximum_runners":      float64(5),
				"public_ip_enabled":    true,
				"public_ips":           []any{map[string]any{"prefix": "192.0.2.1"}},
				"image_details":        map[string]any{"version": "2"},
			},
			expectedUpdate: map[string]any{
				"name":             "updated",
				"size":             "4-core",
				"runner_group_id":  2,
				"maximum_runners":  5,
				"enable_static_ip": true,
				"image_version":    "2",
			},
			requirePublicIPs: true,
			wantState:        "ready",
		},
		"ready with public IP allocation": {
			runner: map[string]any{
				"status":     "Ready",
				"public_ips": []any{map[string]any{"prefix": "192.0.2.1"}},
			},
			requirePublicIPs: true,
			wantState:        "ready",
		},
		"stuck": {
			runner:  map[string]any{"status": "Stuck"},
			wantErr: true,
		},
		"missing status": {
			runner:  map[string]any{},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := hostedRunnerProvisioningState(test.runner, test.expectedUpdate, test.requirePublicIPs)
			if test.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != test.wantState {
				t.Fatalf("got state %q, want %q", got, test.wantState)
			}
		})
	}
}

func TestAccGithubActionsHostedRunner(t *testing.T) {
	t.Parallel()

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates hosted runners without error", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "image.0.id",
				"2306",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "image.0.source",
				"github",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "id",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "status",
				"Ready",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "platform",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "image.0.size_gb",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.id",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.memory_gb",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.storage_gb",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates hosted runner to enable public IPs", func(t *testing.T) {
		t.Parallel()

		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-optional-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size              = "2-core"
				runner_group_id   = github_actions_runner_group.test.id
				maximum_runners   = 5
				public_ip_enabled = false
			}
		`, randomID, randomID)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-optional-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size              = "2-core"
				runner_group_id   = github_actions_runner_group.test.id
				maximum_runners   = 5
				public_ip_enabled = true
			}
		`, randomID, randomID)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "public_ip_enabled",
				"false",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "status",
				"Ready",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-optional-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"2-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "maximum_runners",
				"5",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "public_ip_enabled",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "status",
				"Ready",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "public_ips.0.prefix",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check:  checkBefore,
				},
				{
					Config: configAfter,
					Check:  checkAfter,
				},
			},
		})
	})

	t.Run("updates hosted runner configuration", func(t *testing.T) {
		t.Parallel()

		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-update-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
				maximum_runners = 3
			}
		`, randomID, randomID)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-update-%s-updated"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
				maximum_runners = 5
			}
		`, randomID, randomID)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-update-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "maximum_runners",
				"3",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-update-%s-updated", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "maximum_runners",
				"5",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check:  checkBefore,
				},
				{
					Config: configAfter,
					Check:  checkAfter,
				},
			},
		})
	})

	t.Run("updates size field", func(t *testing.T) {
		t.Parallel()

		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-size-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-size-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "8-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
				"4",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"8-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
				"8",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check:  checkBefore,
				},
				{
					Config: configAfter,
					Check:  checkAfter,
				},
			},
		})
	})

	t.Run("imports hosted runner", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-import-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "id",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-import-%s", randomID),
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:            "github_actions_hosted_runner.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"image", "image_gen"},
				},
			},
		})
	})

	t.Run("deletes hosted runner", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-delete-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(
							"github_actions_hosted_runner.test", "id",
						),
					),
				},
				// This step should successfully delete the runner
				{
					Config: fmt.Sprintf(`
							resource "github_actions_runner_group" "test" {
								name       = "tf-acc-test-group-%s"
								visibility = "all"
							}
						`, randomID),
				},
			},
		})
	})
}
