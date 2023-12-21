package yacarsdk

import (
	"fmt"
	"strconv"
)

func ValidateAccounts(accounts []Account) error {
	for _, account := range accounts {
		if !account.IsMinimallyPopulated() {
			return fmt.Errorf("account ID %s is not minimally populated", account.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, account := range accounts {
		if _, ok := idCount[account.Id]; ok {
			return fmt.Errorf("duplicate account ID: %s", account.Id)
		}

		idCount[account.Id] = struct{}{}
	}

	return nil
}

func ValidateAssets(assets []Asset, entities []Entity) error {
	for _, asset := range assets {
		if !asset.IsMinimallyPopulated() {
			return fmt.Errorf("asset ID %s is not minimally populated", asset.Id)
		}

		if asset.Id == asset.Name {
			return fmt.Errorf("asset name for %s cannot be the asset ID", asset.Id)
		}

		if asset.Id == asset.Symbol {
			return fmt.Errorf("asset symbol for %s cannot be the asset ID", asset.Id)
		}

		if len(asset.Symbol) > 20 {
			return fmt.Errorf("asset symbol for %s cannot be longer than 20 characters", asset.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, asset := range assets {
		if _, ok := idCount[asset.Id]; ok {
			return fmt.Errorf("duplicate asset ID: %s", asset.Id)
		}
		idCount[asset.Id] = struct{}{}
	}

	// Circ supply check
	for _, asset := range assets {
		if len(asset.CircSupply) > 0 && len(asset.CircSupplyAPI) > 0 {
			return fmt.Errorf("[%s] either 'circ_supply' or 'circ_supply_api' must be specified, but not both", asset.Id)
		}

		if len(asset.CircSupply) > 0 {
			if parsed, err := strconv.ParseFloat(asset.CircSupply, 64); err != nil && parsed > 0 {
				return fmt.Errorf("[%s] 'circ_supply' must be float greater than 0", asset.Id)
			}
		}
	}

	// Total supply check
	for _, asset := range assets {
		if len(asset.TotalSupply) > 0 && len(asset.TotalSupplyAPI) > 0 {
			return fmt.Errorf("[%s] either 'total_supply' or 'total_supply_api' must be specified, but not both", asset.Id)
		}

		if len(asset.TotalSupply) > 0 {
			if parsed, err := strconv.ParseFloat(asset.TotalSupply, 64); err != nil && parsed > 0 {
				return fmt.Errorf("[%s] 'total_supply' must be number greater than 0", asset.Id)
			}
		}
	}

	// Corresponding entity check
	entityNameSet := map[string]struct{}{}
	for _, entity := range entities {
		entityNameSet[entity.Name] = struct{}{}
	}
	for _, asset := range assets {
		if asset.Entity == "" {
			continue
		}
		if _, ok := entityNameSet[asset.Entity]; ok {
			continue
		}

		return fmt.Errorf("[%s] entity '%s' does not exists", asset.Id, asset.Entity)
	}

	// Non-permissioned DEX TxHash must be unique
	permissionedDex := map[string]struct{}{
		"osmosis-main": {},
		"kujira-fin":   {},
	}
	txCheck := make(map[string]struct{})
	for _, asset := range assets {
		// If asset is from a permissioned DEX or is empty, skip
		if _, ok := permissionedDex[asset.VerificationTx]; ok || asset.VerificationTx == "" {
			continue
		}

		if _, ok := txCheck[asset.VerificationTx]; ok {
			return fmt.Errorf("duplicate asset tx hash: %s", asset.Id)
		}
		txCheck[asset.VerificationTx] = struct{}{}
	}

	return nil
}

func ValidateBinaries(binaries []Binary) error {
	for _, binary := range binaries {
		if !binary.IsMinimallyPopulated() {
			return fmt.Errorf("binary ID %s is not minimally populated", binary.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, binary := range binaries {
		if _, ok := idCount[binary.Id]; ok {
			return fmt.Errorf("duplicate binary ID: %s", binary.Id)
		}

		idCount[binary.Id] = struct{}{}
	}

	return nil
}

func ValidateContracts(contracts []Contract) error {
	for _, contract := range contracts {
		if !contract.IsMinimallyPopulated() {
			return fmt.Errorf("contract ID %s is not minimally populated", contract.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, contract := range contracts {
		if _, ok := idCount[contract.Id]; ok {
			return fmt.Errorf("duplicate contract ID: %s", contract.Id)
		}

		idCount[contract.Id] = struct{}{}
	}

	return nil
}

func ValidateEntities(entities []Entity) error {
	for _, entity := range entities {
		if !entity.IsMinimallyPopulated() {
			return fmt.Errorf("entity name %s is not minimally populated", entity.Name)
		}
	}

	entityCount := make(map[string]struct{})
	for _, entity := range entities {
		if _, ok := entityCount[entity.Name]; ok {
			return fmt.Errorf("duplicate entity name: %s", entity.Name)
		}

		entityCount[entity.Name] = struct{}{}
	}

	return nil
}

func ValidatePools(pools []Pool) error {
	for _, pool := range pools {
		if !pool.IsMinimallyPopulated() {
			return fmt.Errorf("pool ID %s is not minimally populated", pool.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, pool := range pools {
		if _, ok := idCount[pool.Id]; ok {
			return fmt.Errorf("duplicate pool ID: %s", pool.Id)
		}

		idCount[pool.Id] = struct{}{}
	}

	return nil
}
