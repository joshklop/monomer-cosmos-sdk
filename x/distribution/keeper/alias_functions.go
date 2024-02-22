package keeper

import (
	"context"

	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/distribution/types"
)

// get outstanding rewards
func (k Keeper) GetValidatorOutstandingRewardsCoins(ctx context.Context, val sdk.ValAddress) (sdk.DecCoins, error) {
	rewards, err := k.GetValidatorOutstandingRewards(ctx, val)
	if err != nil {
		return nil, err
	}

	return rewards.Rewards, nil
}

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx context.Context) sdk.ModuleAccountI {
	return k.authKeeper.GetModuleAccount(ctx, types.ModuleName)
}
