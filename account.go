package yacarsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Account struct {
	Id     string `json:"id"`
	Entity string `json:"entity"`
	Label  string `json:"label"`
}

func (a Account) IsMinimallyPopulated() bool {
	return a.Id != "" && a.Entity != ""
}

type ByEnforcedAccountOrder []Account

func (acc ByEnforcedAccountOrder) Len() int      { return len(acc) }
func (acc ByEnforcedAccountOrder) Swap(i, j int) { acc[i], acc[j] = acc[j], acc[i] }
func (acc ByEnforcedAccountOrder) Less(i, j int) bool {
	collator := collate.New(language.MustParse("en-US"))

	a := acc[i]
	b := acc[j]

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

func GetAccounts(ctx context.Context, httpClient *http.Client, chain string) ([]Account, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/coinhall/yacar/main/%s/account.json", chain)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var accounts []Account
	if err := json.Unmarshal(b, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
