package crisis

import (
	"context"
	"time"

	"github.com/joshklop/monomer-cosmos-sdk/telemetry"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/crisis/keeper"
	"github.com/joshklop/monomer-cosmos-sdk/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx context.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.InvCheckPeriod() == 0 || sdkCtx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(sdkCtx)
}
