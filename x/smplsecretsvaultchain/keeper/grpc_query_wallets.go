package keeper

import (
	"context"

	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WalletsAll(c context.Context, req *types.QueryAllWalletsRequest) (*types.QueryAllWalletsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var walletss []types.Wallets
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	walletsStore := prefix.NewStore(store, types.KeyPrefix(types.WalletsKeyPrefix))

	pageRes, err := query.Paginate(walletsStore, req.Pagination, func(key []byte, value []byte) error {
		var wallets types.Wallets
		if err := k.cdc.Unmarshal(value, &wallets); err != nil {
			return err
		}

		walletss = append(walletss, wallets)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWalletsResponse{Wallets: walletss, Pagination: pageRes}, nil
}

func (k Keeper) Wallets(c context.Context, req *types.QueryGetWalletsRequest) (*types.QueryGetWalletsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWallets(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetWalletsResponse{Wallets: val}, nil
}
