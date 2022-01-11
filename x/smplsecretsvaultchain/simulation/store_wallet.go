package simulation

import (
	"math/rand"

	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgStoreWallet(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgStoreWallet{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the StoreWallet simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "StoreWallet simulation not implemented"), nil, nil
	}
}
