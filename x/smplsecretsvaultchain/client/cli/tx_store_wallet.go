package cli

import (
	"strconv"

	"encoding/json"
	"github.com/SmplEcosystem/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStoreWallet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store-wallet [name] [wallet] [passphrase]",
		Short: "Broadcast message storeWallet",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argWallet := new(types.Wallet)
			err = json.Unmarshal([]byte(args[1]), argWallet)
			if err != nil {
				return err
			}
			argPassphrase := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStoreWallet(
				clientCtx.GetFromAddress().String(),
				argName,
				argWallet,
				argPassphrase,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
