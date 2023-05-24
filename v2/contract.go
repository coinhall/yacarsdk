package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Contract struct {
	Id     string `json:"id"`
	Entity string `json:"entity"`
	Label  string `json:"label"`
}

func (c Contract) IsMinimallyPopulated() bool {
	return c.Id != "" && c.Entity != "" && c.Label != ""
}

type ByEnforcedContractOrder []Contract

func (cons ByEnforcedContractOrder) Len() int      { return len(cons) }
func (cons ByEnforcedContractOrder) Swap(i, j int) { cons[i], cons[j] = cons[j], cons[i] }
func (cons ByEnforcedContractOrder) Less(i, j int) bool {
	collator := collate.New(language.MustParse("en-US"))

	a := cons[i]
	b := cons[j]

	// Compare by:
	// 1. Entity
	// 2. Label
	// 3. Id

	if c := collator.CompareString(a.Entity, b.Entity); c != 0 {
		return c < 0
	}

	if c := collator.CompareString(a.Label, b.Label); c != 0 {
		return c < 0
	}

	if c := collator.CompareString(a.Id, b.Id); c != 0 {
		return c < 0
	}

	return false
}
