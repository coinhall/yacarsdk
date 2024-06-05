package yacarsdk

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func ValidateAccounts(accounts []Account) (int, error) {
	for i, account := range accounts {
		if !account.IsMinimallyPopulated() {
			return i, fmt.Errorf("account ID %s is not minimally populated", account.Id)
		}
	}

	idCount := make(map[string]struct{})
	for i, account := range accounts {
		if _, ok := idCount[account.Id]; ok {
			return i, fmt.Errorf("duplicate account ID: %s", account.Id)
		}

		idCount[account.Id] = struct{}{}
	}

	return -1, nil
}

func ValidateAssets(assets []Asset, entities []Entity) (int, error) {
	for i, asset := range assets {
		switch asset.Type {
		case "ibc": // IBC assets check
			if !asset.IsMinimallyPopulatedIbc() {
				return i, fmt.Errorf("IBC asset ID %s is not minimally populated", asset.Id)
			}

			if !asset.HasNoExtraFieldsIbc() {
				return i, fmt.Errorf("IBC asset ID %s contains invalid fields", asset.Id)
			}
			continue
		default:
			if !asset.IsMinimallyPopulated() {
				return i, fmt.Errorf("asset ID %s is not minimally populated", asset.Id)
			}

			if asset.Id == asset.Name {
				return i, fmt.Errorf("asset name for %s cannot be the asset ID", asset.Id)
			}

			if asset.Id == asset.Symbol {
				return i, fmt.Errorf("asset symbol for %s cannot be the asset ID", asset.Id)
			}

			if len(asset.Symbol) > 20 {
				return i, fmt.Errorf("asset symbol for %s cannot be longer than 20 characters", asset.Id)
			}
		}
	}

	idCount := make(map[string]struct{})
	for i, asset := range assets {
		if _, ok := idCount[asset.Id]; ok {
			return i, fmt.Errorf("duplicate asset ID: %s", asset.Id)
		}
		idCount[asset.Id] = struct{}{}
	}

	// Circ supply check
	for i, asset := range assets {
		if len(asset.CircSupply) > 0 && len(asset.CircSupplyAPI) > 0 {
			return i, fmt.Errorf("[%s] either 'circ_supply' or 'circ_supply_api' must be specified, but not both", asset.Id)
		}

		if len(asset.CircSupply) > 0 {
			if parsed, err := strconv.ParseFloat(asset.CircSupply, 64); err != nil && parsed > 0 {
				return i, fmt.Errorf("[%s] 'circ_supply' must be float greater than 0", asset.Id)
			}
		}
	}

	// Total supply check
	for i, asset := range assets {
		if len(asset.TotalSupply) > 0 && len(asset.TotalSupplyAPI) > 0 {
			return i, fmt.Errorf("[%s] either 'total_supply' or 'total_supply_api' must be specified, but not both", asset.Id)
		}

		if len(asset.TotalSupply) > 0 {
			if parsed, err := strconv.ParseFloat(asset.TotalSupply, 64); err != nil && parsed > 0 {
				return i, fmt.Errorf("[%s] 'total_supply' must be number greater than 0", asset.Id)
			}
		}
	}

	// Corresponding entity check
	entityNameSet := map[string]struct{}{}
	for _, entity := range entities {
		entityNameSet[entity.Name] = struct{}{}
	}
	for i, asset := range assets {
		if asset.Entity == "" {
			continue
		}
		if _, ok := entityNameSet[asset.Entity]; ok {
			continue
		}

		return i, fmt.Errorf("[%s] entity '%s' does not exists", asset.Id, asset.Entity)
	}

	// Non-permissioned DEX TxHash must be unique
	txCheck := make(map[string]struct{})
	for i, asset := range assets {
		if asset.VerificationTx == "" {
			continue
		}

		bz, err := hex.DecodeString(asset.VerificationTx)
		if err != nil {
			// Assume is permissioned DEX, no need to check
			continue
		}

		if len(bz) != 32 {
			// Tx hash should be 32 bytes long
			return i, fmt.Errorf("invalid asset tx hash: %s", asset.Id)
		}

		if _, ok := txCheck[asset.VerificationTx]; ok {
			return i, fmt.Errorf("duplicate asset tx hash: %s", asset.Id)
		}
		txCheck[asset.VerificationTx] = struct{}{}
	}

	return -1, nil
}

func ValidateBinaries(binaries []Binary) (int, error) {
	for i, binary := range binaries {
		if !binary.IsMinimallyPopulated() {
			return i, fmt.Errorf("binary ID %s is not minimally populated", binary.Id)
		}
	}

	idCount := make(map[string]struct{})
	for i, binary := range binaries {
		if _, ok := idCount[binary.Id]; ok {
			return i, fmt.Errorf("duplicate binary ID: %s", binary.Id)
		}

		idCount[binary.Id] = struct{}{}
	}

	return -1, nil
}

func ValidateContracts(contracts []Contract) (int, error) {
	for i, contract := range contracts {
		if !contract.IsMinimallyPopulated() {
			return i, fmt.Errorf("contract ID %s is not minimally populated", contract.Id)
		}
	}

	idCount := make(map[string]struct{})
	for i, contract := range contracts {
		if _, ok := idCount[contract.Id]; ok {
			return i, fmt.Errorf("duplicate contract ID: %s", contract.Id)
		}

		idCount[contract.Id] = struct{}{}
	}

	return -1, nil
}

func ValidateEntities(entities []Entity, usedEntities map[string]struct{}) (int, error) {
	for i, entity := range entities {
		if !entity.IsMinimallyPopulated() {
			return i, fmt.Errorf("entity name %s is not minimally populated", entity.Name)
		}
	}

	entityCount := make(map[string]struct{})
	for i, entity := range entities {
		if _, ok := entityCount[entity.Name]; ok {
			return i, fmt.Errorf("duplicate entity name: %s", entity.Name)
		}

		entityCount[entity.Name] = struct{}{}
	}

	for i, entity := range entities {
		if _, ok := usedEntities[entity.Name]; !ok {
			return i, fmt.Errorf("unused entity: %s", entity.Name)
		}
	}

	return -1, nil
}

func ValidatePools(pools []Pool) (int, error) {
	for i, pool := range pools {
		if !pool.IsMinimallyPopulated() {
			return i, fmt.Errorf("pool ID %s is not minimally populated", pool.Id)
		}
	}

	idCount := make(map[string]struct{})
	for i, pool := range pools {
		if _, ok := idCount[pool.Id]; ok {
			return i, fmt.Errorf("duplicate pool ID: %s", pool.Id)
		}

		idCount[pool.Id] = struct{}{}
	}

	return -1, nil
}
