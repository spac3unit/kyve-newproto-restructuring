package cli

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/KYVENetwork/chain/x/pool/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/spf13/cobra"
)

func CmdSubmitCreatePoolProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool-proposal [flags]",
		Args:  cobra.ExactArgs(10),
		Short: "Submit a proposal to create a pool.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			argName := args[0]

			argRuntime := args[1]

			argLogo := args[2]

			argConfig := args[3]

			argStartKey := args[4]

			argUploadInterval, err := strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				return err
			}

			argOperatingCost, err := strconv.ParseUint(args[6], 10, 64)
			if err != nil {
				return err
			}

			argMinStake, err := strconv.ParseUint(args[7], 10, 64)
			if err != nil {
				return err
			}

			argMaxBundleSize, err := strconv.ParseUint(args[8], 10, 64)
			if err != nil {
				return err
			}

			argVersion := args[9]

			argBinaries := args[10]

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewCreatePoolProposal(title, description, argName, argRuntime, argLogo, argConfig, argStartKey, argUploadInterval, argOperatingCost, argMinStake, argMaxBundleSize, argVersion, argBinaries)

			isExpedited, err := cmd.Flags().GetBool(cli.FlagIsExpedited)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from, isExpedited)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}