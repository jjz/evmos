package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tharsis/evmos/x/fees/types"
)

// GetParams returns the total set of distribution parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the distribution parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
