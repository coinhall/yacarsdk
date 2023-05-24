package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Asset struct {
	Id             string `json:"id"`
	Entity         string `json:"entity,omitempty"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Decimals       string `json:"decimals"`
	Type           string `json:"type"`
	CircSupplyAPI  string `json:"circ_supply_api,omitempty"`
	TotalSupplyAPI string `json:"total_supply_api,omitempty"`
	Icon           string `json:"icon,omitempty"`
	CoinMarketCap  string `json:"coinmarketcap,omitempty"`
	CoinGecko      string `json:"coingecko,omitempty"`
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

	// Assets with entity come first
	if len(a.Entity) > 0 && len(b.Entity) == 0 {
		return true
	}
	if len(a.Entity) == 0 && len(b.Entity) > 0 {
		return false
	}

	// If both assets have entity, compare by
	// 1. Entity
	// 2. Name
	// 3. Id
	if len(a.Entity) > 0 && len(b.Entity) > 0 {
		if c := collator.CompareString(a.Entity, b.Entity); c != 0 {
			return c < 0
		}

		if c := collator.CompareString(a.Name, b.Name); c != 0 {
			return c < 0
		}

		if c := collator.CompareString(a.Id, b.Id); c != 0 {
			return c < 0
		}

		return false
	}

	// If both assets don't have entity, compare by
	// 1. Name
	// 2. Id
	if len(a.Entity) == 0 && len(b.Entity) == 0 {
		if c := collator.CompareString(a.Name, b.Name); c != 0 {
			return c < 0
		}

		if c := collator.CompareString(a.Id, b.Id); c != 0 {
			return c < 0
		}

		return false
	}

	return false
}
