package smplsecretsvaultchain_test

import (
	"testing"

	keepertest "github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/nullify"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WalletsList: []types.Wallets{
			{
				Owner: "0",
			},
			{
				Owner: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	smplsecretsvaultchain.InitGenesis(ctx, *k, genesisState)
	got := smplsecretsvaultchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WalletsList, got.WalletsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
