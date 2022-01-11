package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/nullify"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWalletsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWallets(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWalletsRequest
		response *types.QueryGetWalletsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetWalletsRequest{
				Owner: msgs[0].Owner,
			},
			response: &types.QueryGetWalletsResponse{Wallets: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetWalletsRequest{
				Owner: msgs[1].Owner,
			},
			response: &types.QueryGetWalletsResponse{Wallets: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetWalletsRequest{
				Owner: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Wallets(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestWalletsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWallets(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWalletsRequest {
		return &types.QueryAllWalletsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WalletsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Wallets), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Wallets),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WalletsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Wallets), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Wallets),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WalletsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Wallets),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WalletsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
