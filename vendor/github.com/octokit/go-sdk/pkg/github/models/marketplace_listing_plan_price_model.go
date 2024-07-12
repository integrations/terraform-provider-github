package models
import (
    "errors"
)
type MarketplaceListingPlan_price_model int

const (
    FREE_MARKETPLACELISTINGPLAN_PRICE_MODEL MarketplaceListingPlan_price_model = iota
    FLAT_RATE_MARKETPLACELISTINGPLAN_PRICE_MODEL
    PER_UNIT_MARKETPLACELISTINGPLAN_PRICE_MODEL
)

func (i MarketplaceListingPlan_price_model) String() string {
    return []string{"FREE", "FLAT_RATE", "PER_UNIT"}[i]
}
func ParseMarketplaceListingPlan_price_model(v string) (any, error) {
    result := FREE_MARKETPLACELISTINGPLAN_PRICE_MODEL
    switch v {
        case "FREE":
            result = FREE_MARKETPLACELISTINGPLAN_PRICE_MODEL
        case "FLAT_RATE":
            result = FLAT_RATE_MARKETPLACELISTINGPLAN_PRICE_MODEL
        case "PER_UNIT":
            result = PER_UNIT_MARKETPLACELISTINGPLAN_PRICE_MODEL
        default:
            return 0, errors.New("Unknown MarketplaceListingPlan_price_model value: " + v)
    }
    return &result, nil
}
func SerializeMarketplaceListingPlan_price_model(values []MarketplaceListingPlan_price_model) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MarketplaceListingPlan_price_model) isMultiValue() bool {
    return false
}
