package yacarsdk

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Entity struct {
	Name     string `json:"name"`
	Website  string `json:"website,omitempty"`
	Telegram string `json:"telegram,omitempty"`
	Twitter  string `json:"twitter,omitempty"`
	Discord  string `json:"discord,omitempty"`
}

func (e Entity) IsMinimallyPopulated() bool {
	return e.Name != ""
}

type ByEnforcedEntityOrder []Entity

func (ents ByEnforcedEntityOrder) Len() int      { return len(ents) }
func (ents ByEnforcedEntityOrder) Swap(i, j int) { ents[i], ents[j] = ents[j], ents[i] }
func (ents ByEnforcedEntityOrder) Less(i, j int) bool {
	collator := collate.New(language.MustParse("en-US"))

	a := ents[i]
	b := ents[j]

	// Compare by name only

	if c := collator.CompareString(a.Name, b.Name); c != 0 {
		return c < 0
	}

	return false
}
