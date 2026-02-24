package github

import (
	"testing"

	"github.com/google/go-github/v83/github"
	"github.com/stretchr/testify/assert"
)

func TestBuildQueryURL(t *testing.T) {
	t.Run("returns base URL when params are empty", func(t *testing.T) {
		result := buildQueryURL("enterprises/test/settings/billing/usage", map[string]string{})
		assert.Equal(t, "enterprises/test/settings/billing/usage", result)
	})

	t.Run("returns base URL when all param values are empty strings", func(t *testing.T) {
		result := buildQueryURL("enterprises/test/settings/billing/usage", map[string]string{
			"year":  "",
			"month": "",
		})
		assert.Equal(t, "enterprises/test/settings/billing/usage", result)
	})

	t.Run("appends single query parameter", func(t *testing.T) {
		result := buildQueryURL("enterprises/test/settings/billing/usage", map[string]string{
			"year": "2025",
		})
		assert.Equal(t, "enterprises/test/settings/billing/usage?year=2025", result)
	})

	t.Run("appends multiple query parameters", func(t *testing.T) {
		result := buildQueryURL("enterprises/test/settings/billing/usage", map[string]string{
			"year":  "2025",
			"month": "6",
		})
		assert.Contains(t, result, "year=2025")
		assert.Contains(t, result, "month=6")
		assert.Contains(t, result, "enterprises/test/settings/billing/usage?")
	})

	t.Run("skips empty values in mixed params", func(t *testing.T) {
		result := buildQueryURL("enterprises/test/settings/billing/usage", map[string]string{
			"year":           "2025",
			"month":          "",
			"cost_center_id": "cc-123",
		})
		assert.Contains(t, result, "year=2025")
		assert.Contains(t, result, "cost_center_id=cc-123")
		assert.NotContains(t, result, "month")
	})

	t.Run("returns base URL when params map is nil-like empty", func(t *testing.T) {
		result := buildQueryURL("base/path", map[string]string{})
		assert.Equal(t, "base/path", result)
	})
}

func TestIntToString(t *testing.T) {
	t.Run("returns empty string for zero", func(t *testing.T) {
		assert.Equal(t, "", intToString(0))
	})

	t.Run("converts positive integer", func(t *testing.T) {
		assert.Equal(t, "2025", intToString(2025))
	})

	t.Run("converts single digit", func(t *testing.T) {
		assert.Equal(t, "6", intToString(6))
	})
}

func TestFlattenUsageItems(t *testing.T) {
	t.Run("returns empty slice for nil input", func(t *testing.T) {
		result := flattenUsageItems(nil)
		assert.Empty(t, result)
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		result := flattenUsageItems([]*github.UsageItem{})
		assert.Empty(t, result)
	})

	t.Run("flattens usage items correctly", func(t *testing.T) {
		items := []*github.UsageItem{
			{
				Date:             "2025-01-15",
				Product:          "Actions",
				SKU:              "Actions Linux",
				Quantity:         100,
				UnitType:         "minutes",
				PricePerUnit:     0.008,
				GrossAmount:      0.8,
				DiscountAmount:   0,
				NetAmount:        0.8,
				OrganizationName: github.Ptr("test-org"),
				RepositoryName:   github.Ptr("test-org/example"),
			},
		}

		result := flattenUsageItems(items)
		assert.Len(t, result, 1)
		assert.Equal(t, "2025-01-15", result[0]["date"])
		assert.Equal(t, "Actions", result[0]["product"])
		assert.Equal(t, "Actions Linux", result[0]["sku"])
		assert.Equal(t, 100.0, result[0]["quantity"])
		assert.Equal(t, "minutes", result[0]["unit_type"])
		assert.InDelta(t, 0.008, result[0]["price_per_unit"], 0.0001)
		assert.InDelta(t, 0.8, result[0]["gross_amount"], 0.0001)
		assert.InDelta(t, 0.0, result[0]["discount_amount"], 0.0001)
		assert.InDelta(t, 0.8, result[0]["net_amount"], 0.0001)
		assert.Equal(t, github.Ptr("test-org"), result[0]["organization_name"])
		assert.Equal(t, github.Ptr("test-org/example"), result[0]["repository_name"])
	})

	t.Run("flattens items with nil optional fields", func(t *testing.T) {
		items := []*github.UsageItem{
			{
				Date:     "2025-01-15",
				Product:  "Actions",
				SKU:      "Actions Linux",
				Quantity: 50,
				UnitType: "minutes",
			},
		}

		result := flattenUsageItems(items)
		assert.Len(t, result, 1)
		assert.Nil(t, result[0]["organization_name"])
		assert.Nil(t, result[0]["repository_name"])
	})
}

func TestFlattenPremiumRequestUsageItems(t *testing.T) {
	t.Run("returns empty slice for nil input", func(t *testing.T) {
		result := flattenPremiumRequestUsageItems(nil)
		assert.Empty(t, result)
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		result := flattenPremiumRequestUsageItems([]*github.PremiumRequestUsageItem{})
		assert.Empty(t, result)
	})

	t.Run("flattens premium request usage items correctly", func(t *testing.T) {
		items := []*github.PremiumRequestUsageItem{
			{
				Product:          "Copilot",
				SKU:              "Copilot Premium Request",
				Model:            "GPT-5",
				UnitType:         "requests",
				PricePerUnit:     0.04,
				GrossQuantity:    100,
				GrossAmount:      4,
				DiscountQuantity: 0,
				DiscountAmount:   0,
				NetQuantity:      100,
				NetAmount:        4,
			},
		}

		result := flattenPremiumRequestUsageItems(items)
		assert.Len(t, result, 1)
		assert.Equal(t, "Copilot", result[0]["product"])
		assert.Equal(t, "Copilot Premium Request", result[0]["sku"])
		assert.Equal(t, "GPT-5", result[0]["model"])
		assert.Equal(t, "requests", result[0]["unit_type"])
		assert.InDelta(t, 0.04, result[0]["price_per_unit"], 0.0001)
		assert.InDelta(t, 100.0, result[0]["gross_quantity"], 0.0001)
		assert.InDelta(t, 4.0, result[0]["gross_amount"], 0.0001)
		assert.InDelta(t, 0.0, result[0]["discount_quantity"], 0.0001)
		assert.InDelta(t, 0.0, result[0]["discount_amount"], 0.0001)
		assert.InDelta(t, 100.0, result[0]["net_quantity"], 0.0001)
		assert.InDelta(t, 4.0, result[0]["net_amount"], 0.0001)
	})
}

func TestFlattenUsageSummaryItems(t *testing.T) {
	t.Run("returns empty slice for nil input", func(t *testing.T) {
		result := flattenUsageSummaryItems(nil)
		assert.Empty(t, result)
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		result := flattenUsageSummaryItems([]*EnterpriseUsageSummaryItem{})
		assert.Empty(t, result)
	})

	t.Run("flattens usage summary items correctly", func(t *testing.T) {
		items := []*EnterpriseUsageSummaryItem{
			{
				Product:          "Actions",
				SKU:              "actions_linux",
				UnitType:         "minutes",
				PricePerUnit:     0.008,
				GrossQuantity:    1000,
				GrossAmount:      8,
				DiscountQuantity: 0,
				DiscountAmount:   0,
				NetQuantity:      1000,
				NetAmount:        8,
			},
		}

		result := flattenUsageSummaryItems(items)
		assert.Len(t, result, 1)
		assert.Equal(t, "Actions", result[0]["product"])
		assert.Equal(t, "actions_linux", result[0]["sku"])
		assert.Equal(t, "minutes", result[0]["unit_type"])
		assert.InDelta(t, 0.008, result[0]["price_per_unit"], 0.0001)
		assert.InDelta(t, 1000.0, result[0]["gross_quantity"], 0.0001)
		assert.InDelta(t, 8.0, result[0]["gross_amount"], 0.0001)
		assert.InDelta(t, 0.0, result[0]["discount_quantity"], 0.0001)
		assert.InDelta(t, 0.0, result[0]["discount_amount"], 0.0001)
		assert.InDelta(t, 1000.0, result[0]["net_quantity"], 0.0001)
		assert.InDelta(t, 8.0, result[0]["net_amount"], 0.0001)
	})
}

func TestFlattenTimePeriod(t *testing.T) {
	t.Run("flattens time period with year only", func(t *testing.T) {
		timePeriod := github.PremiumRequestUsageTimePeriod{
			Year: 2025,
		}

		result := flattenTimePeriod(timePeriod)
		assert.Len(t, result, 1)
		assert.Equal(t, 2025, result[0]["year"])
		assert.Equal(t, 0, result[0]["month"])
		assert.Equal(t, 0, result[0]["day"])
	})

	t.Run("flattens time period with all fields", func(t *testing.T) {
		timePeriod := github.PremiumRequestUsageTimePeriod{
			Year:  2025,
			Month: github.Ptr(6),
			Day:   github.Ptr(15),
		}

		result := flattenTimePeriod(timePeriod)
		assert.Len(t, result, 1)
		assert.Equal(t, 2025, result[0]["year"])
		assert.Equal(t, 6, result[0]["month"])
		assert.Equal(t, 15, result[0]["day"])
	})
}
