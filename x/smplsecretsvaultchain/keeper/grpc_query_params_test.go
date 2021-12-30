package keeper_test

import (
	"testing"

	testkeeper "github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.SmplsecretsvaultchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
