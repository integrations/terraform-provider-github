package github

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/go-github/v84/github"
)

// EnterpriseBillingUsageOptions specifies optional parameters for the
// enterprise billing usage report endpoint.
type EnterpriseBillingUsageOptions struct {
	Year         *int    `url:"year,omitempty"`
	Month        *int    `url:"month,omitempty"`
	Day          *int    `url:"day,omitempty"`
	CostCenterID *string `url:"cost_center_id,omitempty"`
}

// EnterprisePremiumRequestUsageOptions specifies optional parameters for the
// enterprise billing premium request usage report endpoint.
type EnterprisePremiumRequestUsageOptions struct {
	Year         *int    `url:"year,omitempty"`
	Month        *int    `url:"month,omitempty"`
	Day          *int    `url:"day,omitempty"`
	Organization *string `url:"organization,omitempty"`
	User         *string `url:"user,omitempty"`
	Model        *string `url:"model,omitempty"`
	Product      *string `url:"product,omitempty"`
	CostCenterID *string `url:"cost_center_id,omitempty"`
}

// EnterpriseUsageSummaryOptions specifies optional parameters for the
// enterprise billing usage summary endpoint.
type EnterpriseUsageSummaryOptions struct {
	Year         *int    `url:"year,omitempty"`
	Month        *int    `url:"month,omitempty"`
	Day          *int    `url:"day,omitempty"`
	Organization *string `url:"organization,omitempty"`
	Repository   *string `url:"repository,omitempty"`
	Product      *string `url:"product,omitempty"`
	SKU          *string `url:"sku,omitempty"`
	CostCenterID *string `url:"cost_center_id,omitempty"`
}

// EnterprisePremiumRequestUsageReport represents the enterprise-level
// premium request usage report response.
type EnterprisePremiumRequestUsageReport struct {
	TimePeriod github.PremiumRequestUsageTimePeriod `json:"timePeriod"`
	Enterprise *string                              `json:"enterprise,omitempty"`
	UsageItems []*github.PremiumRequestUsageItem    `json:"usageItems"`
}

// EnterpriseUsageSummaryItem represents a single usage line item in an
// enterprise billing usage summary report.
type EnterpriseUsageSummaryItem struct {
	Product          string  `json:"product"`
	SKU              string  `json:"sku"`
	UnitType         string  `json:"unitType"`
	PricePerUnit     float64 `json:"pricePerUnit"`
	GrossQuantity    float64 `json:"grossQuantity"`
	GrossAmount      float64 `json:"grossAmount"`
	DiscountQuantity float64 `json:"discountQuantity"`
	DiscountAmount   float64 `json:"discountAmount"`
	NetQuantity      float64 `json:"netQuantity"`
	NetAmount        float64 `json:"netAmount"`
}

// EnterpriseUsageSummaryReport represents the enterprise-level billing
// usage summary report response.
type EnterpriseUsageSummaryReport struct {
	TimePeriod github.PremiumRequestUsageTimePeriod `json:"timePeriod"`
	Enterprise *string                              `json:"enterprise,omitempty"`
	UsageItems []*EnterpriseUsageSummaryItem        `json:"usageItems"`
}

// buildQueryURL constructs a URL with non-empty query parameters.
func buildQueryURL(base string, params map[string]string) string {
	values := url.Values{}
	for key, val := range params {
		if val != "" {
			values.Set(key, val)
		}
	}

	if len(values) == 0 {
		return base
	}

	return base + "?" + values.Encode()
}

// intToString converts an int to its string representation.
// Returns an empty string if the value is zero.
func intToString(value int) string {
	if value == 0 {
		return ""
	}
	return strconv.Itoa(value)
}

// getEnterpriseBillingUsage fetches the billing usage report for an enterprise.
func getEnterpriseBillingUsage(ctx context.Context, client *github.Client, enterprise string, opts *EnterpriseBillingUsageOptions) (*github.UsageReport, error) {
	urlPath := fmt.Sprintf("enterprises/%s/settings/billing/usage", enterprise)

	params := map[string]string{}
	if opts != nil {
		if opts.Year != nil {
			params["year"] = strconv.Itoa(*opts.Year)
		}
		if opts.Month != nil {
			params["month"] = strconv.Itoa(*opts.Month)
		}
		if opts.Day != nil {
			params["day"] = strconv.Itoa(*opts.Day)
		}
		if opts.CostCenterID != nil {
			params["cost_center_id"] = *opts.CostCenterID
		}
	}

	urlPath = buildQueryURL(urlPath, params)

	req, err := client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	report := new(github.UsageReport)
	_, err = client.Do(ctx, req, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

// getEnterprisePremiumRequestUsage fetches the billing premium request usage report for an enterprise.
func getEnterprisePremiumRequestUsage(ctx context.Context, client *github.Client, enterprise string, opts *EnterprisePremiumRequestUsageOptions) (*EnterprisePremiumRequestUsageReport, error) {
	urlPath := fmt.Sprintf("enterprises/%s/settings/billing/premium_request/usage", enterprise)

	params := map[string]string{}
	if opts != nil {
		if opts.Year != nil {
			params["year"] = strconv.Itoa(*opts.Year)
		}
		if opts.Month != nil {
			params["month"] = strconv.Itoa(*opts.Month)
		}
		if opts.Day != nil {
			params["day"] = strconv.Itoa(*opts.Day)
		}
		if opts.Organization != nil {
			params["organization"] = *opts.Organization
		}
		if opts.User != nil {
			params["user"] = *opts.User
		}
		if opts.Model != nil {
			params["model"] = *opts.Model
		}
		if opts.Product != nil {
			params["product"] = *opts.Product
		}
		if opts.CostCenterID != nil {
			params["cost_center_id"] = *opts.CostCenterID
		}
	}

	urlPath = buildQueryURL(urlPath, params)

	req, err := client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	report := new(EnterprisePremiumRequestUsageReport)
	_, err = client.Do(ctx, req, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

// getEnterpriseUsageSummary fetches the billing usage summary for an enterprise.
func getEnterpriseUsageSummary(ctx context.Context, client *github.Client, enterprise string, opts *EnterpriseUsageSummaryOptions) (*EnterpriseUsageSummaryReport, error) {
	urlPath := fmt.Sprintf("enterprises/%s/settings/billing/usage/summary", enterprise)

	params := map[string]string{}
	if opts != nil {
		if opts.Year != nil {
			params["year"] = strconv.Itoa(*opts.Year)
		}
		if opts.Month != nil {
			params["month"] = strconv.Itoa(*opts.Month)
		}
		if opts.Day != nil {
			params["day"] = strconv.Itoa(*opts.Day)
		}
		if opts.Organization != nil {
			params["organization"] = *opts.Organization
		}
		if opts.Repository != nil {
			params["repository"] = *opts.Repository
		}
		if opts.Product != nil {
			params["product"] = *opts.Product
		}
		if opts.SKU != nil {
			params["sku"] = *opts.SKU
		}
		if opts.CostCenterID != nil {
			params["cost_center_id"] = *opts.CostCenterID
		}
	}

	urlPath = buildQueryURL(urlPath, params)

	req, err := client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	report := new(EnterpriseUsageSummaryReport)
	_, err = client.Do(ctx, req, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

// flattenUsageItems converts billing usage items to a Terraform state-compatible format.
func flattenUsageItems(items []*github.UsageItem) []map[string]any {
	result := make([]map[string]any, len(items))
	for idx, item := range items {
		orgName := ""
		if item.OrganizationName != nil {
			orgName = *item.OrganizationName
		}
		repoName := ""
		if item.RepositoryName != nil {
			repoName = *item.RepositoryName
		}
		result[idx] = map[string]any{
			"date":              item.Date,
			"product":           item.Product,
			"sku":               item.SKU,
			"quantity":          item.Quantity,
			"unit_type":         item.UnitType,
			"price_per_unit":    item.PricePerUnit,
			"gross_amount":      item.GrossAmount,
			"discount_amount":   item.DiscountAmount,
			"net_amount":        item.NetAmount,
			"organization_name": orgName,
			"repository_name":   repoName,
		}
	}
	return result
}

// flattenPremiumRequestUsageItems converts premium request usage items to a Terraform state-compatible format.
func flattenPremiumRequestUsageItems(items []*github.PremiumRequestUsageItem) []map[string]any {
	result := make([]map[string]any, len(items))
	for idx, item := range items {
		result[idx] = map[string]any{
			"product":           item.Product,
			"sku":               item.SKU,
			"model":             item.Model,
			"unit_type":         item.UnitType,
			"price_per_unit":    item.PricePerUnit,
			"gross_quantity":    item.GrossQuantity,
			"gross_amount":      item.GrossAmount,
			"discount_quantity": item.DiscountQuantity,
			"discount_amount":   item.DiscountAmount,
			"net_quantity":      item.NetQuantity,
			"net_amount":        item.NetAmount,
		}
	}
	return result
}

// flattenUsageSummaryItems converts usage summary items to a Terraform state-compatible format.
func flattenUsageSummaryItems(items []*EnterpriseUsageSummaryItem) []map[string]any {
	result := make([]map[string]any, len(items))
	for idx, item := range items {
		result[idx] = map[string]any{
			"product":           item.Product,
			"sku":               item.SKU,
			"unit_type":         item.UnitType,
			"price_per_unit":    item.PricePerUnit,
			"gross_quantity":    item.GrossQuantity,
			"gross_amount":      item.GrossAmount,
			"discount_quantity": item.DiscountQuantity,
			"discount_amount":   item.DiscountAmount,
			"net_quantity":      item.NetQuantity,
			"net_amount":        item.NetAmount,
		}
	}
	return result
}

// flattenTimePeriod converts a PremiumRequestUsageTimePeriod to a Terraform state-compatible format.
func flattenTimePeriod(timePeriod github.PremiumRequestUsageTimePeriod) []map[string]any {
	result := map[string]any{
		"year": timePeriod.Year,
	}
	if timePeriod.Month != nil {
		result["month"] = *timePeriod.Month
	} else {
		result["month"] = 0
	}
	if timePeriod.Day != nil {
		result["day"] = *timePeriod.Day
	} else {
		result["day"] = 0
	}
	return []map[string]any{result}
}
