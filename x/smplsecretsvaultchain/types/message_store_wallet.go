package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStoreWallet = "store_wallet"

var _ sdk.Msg = &MsgStoreWallet{}

func NewMsgStoreWallet(creator string, name string, wallet *Wallet, passphrase string) *MsgStoreWallet {
	return &MsgStoreWallet{
		Creator:    creator,
		Name:       name,
		Wallet:     wallet,
		Passphrase: passphrase,
	}
}

func (msg *MsgStoreWallet) Route() string {
	return RouterKey
}

func (msg *MsgStoreWallet) Type() string {
	return TypeMsgStoreWallet
}

func (msg *MsgStoreWallet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStoreWallet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStoreWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
