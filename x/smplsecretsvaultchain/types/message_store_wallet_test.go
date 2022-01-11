package types

import (
	"testing"

	"github.com/SmplEcosystem/SmplSecretsVaultChain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgStoreWallet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgStoreWallet
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgStoreWallet{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgStoreWallet{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
