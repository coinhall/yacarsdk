package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Binary struct {
	Id     string `json:"id"`
	Entity string `json:"entity"`
	Label  string `json:"label"`
}

func (b Binary) IsMinimallyPopulated() bool {
	return b.Id != "" && b.Entity != "" && b.Label != ""
}

type ByEnforcedBinaryOrder []Binary

func (bins ByEnforcedBinaryOrder) Len() int      { return len(bins) }
func (bins ByEnforcedBinaryOrder) Swap(i, j int) { bins[i], bins[j] = bins[j], bins[i] }
func (bins ByEnforcedBinaryOrder) Less(i, j int) bool {
	collator := collate.New(language.MustParse("en-US"), collate.Numeric)

	a := bins[i]
	b := bins[j]

	// Compare by ID only
	if c := collator.CompareString(a.Id, b.Id); c != 0 {
		return c < 0
	}

	return false
}
