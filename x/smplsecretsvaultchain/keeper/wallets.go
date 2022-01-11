package keeper

import (
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetWallets set a specific wallets in the store from its index
func (k Keeper) SetWallets(ctx sdk.Context, wallets types.Wallets) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WalletsKeyPrefix))
	b := k.cdc.MustMarshal(&wallets)
	store.Set(types.WalletsKey(
		wallets.Owner,
	), b)
}

// GetWallets returns a wallets from its index
func (k Keeper) GetWallets(
	ctx sdk.Context,
	owner string,

) (val types.Wallets, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WalletsKeyPrefix))

	b := store.Get(types.WalletsKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWallets removes a wallets from the store
func (k Keeper) RemoveWallets(
	ctx sdk.Context,
	owner string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WalletsKeyPrefix))
	store.Delete(types.WalletsKey(
		owner,
	))
}

// GetAllWallets returns all wallets
func (k Keeper) GetAllWallets(ctx sdk.Context) (list []types.Wallets) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WalletsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Wallets
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
