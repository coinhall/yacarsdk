package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Asset struct {
	Id             string `json:"id"`                         // All
	OriginChainId  string `json:"origin_chain_id,omitempty"`  // IBC
	OriginId       string `json:"origin_id,omitempty"`        // IBC
	Entity         string `json:"entity,omitempty"`           // All (optional)
	Name           string `json:"name,omitempty"`             // Non-IBC
	Symbol         string `json:"symbol,omitempty"`           // Non-IBC
	Decimals       string `json:"decimals,omitempty"`         // Non-IBC
	Type           string `json:"type,omitempty"`             // All
	CircSupply     string `json:"circ_supply,omitempty"`      // Non-IBC
	CircSupplyAPI  string `json:"circ_supply_api,omitempty"`  // Non-IBC
	TotalSupply    string `json:"total_supply,omitempty"`     // Non-IBC
	TotalSupplyAPI string `json:"total_supply_api,omitempty"` // Non-IBC
	Icon           string `json:"icon,omitempty"`             // Non-IBC
	CoinMarketCap  string `json:"coinmarketcap,omitempty"`    // Non-IBC
	CoinGecko      string `json:"coingecko,omitempty"`        // Non-IBC
	VerificationTx string `json:"verification_tx,omitempty"`  // Non-IBC
}

// IBC assets should maximally have the following fields:
//   - Id
//   - OriginChainId
//   - OriginId
//   - Entity (optional)
//   - Type
func (a Asset) IsMinimallyPopulatedIbc() bool {
	return a.Type == "ibc" &&
		len(a.Id) > 0 &&
		len(a.OriginChainId) > 0 &&
		len(a.OriginId) > 0
}

// Returns false if the asset is not an IBC asset or if it has extra fields
func (a Asset) HasNoExtraFieldsIbc() bool {
	return a.Type == "ibc" &&
		len(a.Name) == 0 &&
		len(a.Symbol) == 0 &&
		len(a.Decimals) == 0 &&
		len(a.CircSupply) == 0 &&
		len(a.CircSupplyAPI) == 0 &&
		len(a.TotalSupply) == 0 &&
		len(a.TotalSupplyAPI) == 0 &&
		len(a.Icon) == 0 &&
		len(a.CoinMarketCap) == 0 &&
		len(a.CoinGecko) == 0 &&
		len(a.VerificationTx) == 0
}

func (a Asset) IsMinimallyPopulated() bool {
	return len(a.Id) > 0 &&
		len(a.Name) > 0 &&
		len(a.Symbol) > 0 &&
		len(a.Decimals) > 0 &&
		len(a.Type) > 0 &&
		len(a.OriginChainId) == 0
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
