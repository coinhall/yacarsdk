package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Asset struct {
	Id             string `json:"id"`                         // All
	OriginChainId  string `json:"origin_chain_id,omitempty"`  // IBC only
	OriginId       string `json:"origin_id,omitempty"`        // IBC only
	Entity         string `json:"entity,omitempty"`           // All (Optional)
	Name           string `json:"name,omitempty"`             // All
	Symbol         string `json:"symbol,omitempty"`           // All
	Decimals       string `json:"decimals,omitempty"`         // All
	Type           string `json:"type,omitempty"`             // All
	CircSupply     string `json:"circ_supply,omitempty"`      // All
	CircSupplyAPI  string `json:"circ_supply_api,omitempty"`  // All
	TotalSupply    string `json:"total_supply,omitempty"`     // All
	TotalSupplyAPI string `json:"total_supply_api,omitempty"` // All
	Icon           string `json:"icon,omitempty"`             // All
	CoinMarketCap  string `json:"coinmarketcap,omitempty"`    // All
	CoinGecko      string `json:"coingecko,omitempty"`        // All
	VerificationTx string `json:"verification_tx,omitempty"`  // All
}

func (a Asset) IsMinimallyPopulated() bool {
	hasBasicInfo := func(a Asset) bool {
		return len(a.Id) > 0 &&
			len(a.Name) > 0 &&
			len(a.Symbol) > 0 &&
			len(a.Decimals) > 0 &&
			len(a.Type) > 0
	}

	switch a.Type {
	case "ibc":
		return hasBasicInfo(a) && len(a.OriginChainId) >= 0 && len(a.OriginId) >= 0
	default:
		return hasBasicInfo(a) && len(a.OriginChainId) == 0 && len(a.OriginId) == 0
	}
}

type ByEnforcedAssetOrder []Asset

func (assets ByEnforcedAssetOrder) Len() int      { return len(assets) }
func (assets ByEnforcedAssetOrder) Swap(i, j int) { assets[i], assets[j] = assets[j], assets[i] }
func (assets ByEnforcedAssetOrder) Less(i, j int) bool {
	collator := collate.New(language.MustParse("en-US"))

	a := assets[i]
	b := assets[j]

	// Compare by ID only
	if c := collator.CompareString(a.Id, b.Id); c != 0 {
		return c < 0
	}

	return false
}
