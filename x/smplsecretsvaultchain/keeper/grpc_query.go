package keeper

import (
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
)

var _ types.QueryServer = Keeper{}
