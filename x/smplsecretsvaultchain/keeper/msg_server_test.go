package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
