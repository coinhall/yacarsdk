package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

const (
	DexAstroport     = "Astroport"
	DexTerraswap     = "Terraswap"
	DexTfm           = "TFM"
	DexPhoenix       = "Phoenix"
	DexWhiteWhale    = "White Whale"
	DexOsmosis       = "Osmosis"
	DexFortis        = "Fortis"
	DexLoop          = "Loop"
	DexMarbleFinance = "MarbleFinance"
	DexWynd          = "Wynd"
	DexJunoSwap      = "Junoswap"
	DexFin           = "FIN"

	TypeXyk        = "xyk"
	TypeStable     = "stable"
	TypeOrderbook  = "orderbook"
	TypeBalancerV1 = "balancerV1"
)

type Pool struct {
	Id        string   `json:"id" binding:"required"`
	AssetIds  []string `json:"asset_ids" binding:"required,gte=2"`
	Dex       string   `json:"dex" binding:"required"`
	Type      string   `json:"type" binding:"required"`
	LpTokenId string   `json:"lp_token_id,omitempty"`
}

func (a Pool) IsMinimallyPopulated() bool {
	return len(a.Id) > 0 &&
		len(a.AssetIds) >= 2 &&
		len(a.AssetIds[0]) > 0 &&
		len(a.AssetIds[1]) > 0 &&
		len(a.Dex) > 0 &&
		len(a.Type) > 0
}

type ByEnforcedPoolOrder []Pool

func (pools ByEnforcedPoolOrder) Len() int      { return len(pools) }
func (pools ByEnforcedPoolOrder) Swap(i, j int) { pools[i], pools[j] = pools[j], pools[i] }
func (pools ByEnforcedPoolOrder) Less(i, j int) bool {
	collator := collate.New(language.MustParse("en-US"))

	a := pools[i]
	b := pools[j]

	// Pool sorting order
	// Dex > Type > Id
	if c := collator.CompareString(a.Dex, b.Dex); c != 0 {
		return c < 0
	}

	if c := collator.CompareString(a.Type, b.Type); c != 0 {
		return c < 0
	}

	if c := collator.CompareString(a.Id, b.Id); c != 0 {
		return c < 0
	}

	return false
}
