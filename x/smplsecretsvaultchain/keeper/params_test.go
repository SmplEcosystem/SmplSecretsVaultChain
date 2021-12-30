package keeper_test

import (
	"testing"

	testkeeper "github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SmplsecretsvaultchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
