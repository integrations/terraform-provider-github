package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func mustCreateTestEnterpriseCostCenter(t *testing.T) *github.CostCenter {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	cc, _, err := testAccConf.meta.v3client.Enterprise.CreateCostCenter(t.Context(), testAccConf.enterpriseSlug, github.CostCenterRequest{Name: name})
	if err != nil {
		t.Fatalf("failed to create test enterprise cost center: %v", err)
	}

	t.Cleanup(func() {
		if _, _, err := testAccConf.meta.v3client.Enterprise.DeleteCostCenter(context.Background(), testAccConf.enterpriseSlug, cc.ID); err != nil {
			if errIs404(err) {
				return
			}
			t.Logf("failed to delete test enterprise cost center %s: %v", cc.ID, err)
		}
	})

	return cc
}

func mustAddTestEnterpriseCostCenterResources(t *testing.T, cc *github.CostCenter, resources github.CostCenterResourceRequest) {
	t.Helper()

	if _, _, err := testAccConf.meta.v3client.Enterprise.AddResourcesToCostCenter(t.Context(), testAccConf.enterpriseSlug, cc.ID, resources); err != nil {
		t.Fatalf("failed to add resources to test enterprise cost center %s: %v", cc.ID, err)
	}

	t.Cleanup(func() {
		if _, _, err := testAccConf.meta.v3client.Enterprise.RemoveResourcesFromCostCenter(context.Background(), testAccConf.enterpriseSlug, cc.ID, resources); err != nil {
			if errIs404(err) {
				return
			}
			t.Logf("failed to remove resources from test enterprise cost center %s: %v", cc.ID, err)
		}
	})
}
