package keeper

import (
	"context"

	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StoreWallet(goCtx context.Context, msg *types.MsgStoreWallet) (*types.MsgStoreWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	//var wallets types.Wallet
	wallets, found := k.GetWallets(ctx, ownerAddress.String())
	walletAddressString := ownerAddress.String()
	if !found {
		wallets = types.Wallets{
			Owner:               walletAddressString,
			WalletStorageMap:    map[string]*types.WalletStorage{},
			TestPhrase:          "the test phrase",
			EncryptedTestPhrase: "testPhraseEnc",
		}
	}

	walletStorage := &types.WalletStorage{
		Wallet: msg.Wallet,
	}
	wallets.WalletStorageMap[msg.Name] = walletStorage

	newWallets := types.Wallets{
		Owner:            walletAddressString,
		WalletStorageMap: wallets.WalletStorageMap,
	}

	k.SetWallets(ctx, newWallets)

	return &types.MsgStoreWalletResponse{}, nil
}
