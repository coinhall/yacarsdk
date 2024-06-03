package yacarsdk

import (
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Asset struct {
	Id             string `json:"id"`
	Entity         string `json:"entity,omitempty"`
	Name           string `json:"name,omitempty"`
	Symbol         string `json:"symbol,omitempty"`
	Decimals       string `json:"decimals,omitempty"`
	Type           string `json:"type,omitempty"`
	CircSupply     string `json:"circ_supply,omitempty"`
	CircSupplyAPI  string `json:"circ_supply_api,omitempty"`
	TotalSupply    string `json:"total_supply,omitempty"`
	TotalSupplyAPI string `json:"total_supply_api,omitempty"`
	Icon           string `json:"icon,omitempty"`
	CoinMarketCap  string `json:"coinmarketcap,omitempty"`
	CoinGecko      string `json:"coingecko,omitempty"`
	VerificationTx string `json:"verification_tx,omitempty"`
	OriginChainId  string `json:"origin_chain_id,omitempty"`
}

func (a Asset) IsMinimallyPopulatedIBC() bool {
	return strings.HasPrefix("ibc/", a.Id) && len(a.OriginChainId) > 0
}

func (a Asset) IsMinimallyPopulated() bool {
	return len(a.Id) > 0 && len(a.Name) > 0 && len(a.Symbol) > 0 && len(a.Decimals) > 0 && len(a.Type) > 0
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
