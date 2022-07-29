package cli

import (
	"github.com/KYVENetwork/chain/x/pool/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [name]",
		Short: "Broadcast message create-pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// TODO allow json input and unmarshalling
			name := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgCreatePool{
				Creator: clientCtx.GetFromAddress().String(),
				Name:    name,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO implement missing pool modifications

//func CmdSubmitUpdatePoolProposal() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "update-pool [flags]",
//		Args:  cobra.ExactArgs(9),
//		Short: "Submit a proposal to update a pool.",
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			id, err := strconv.ParseUint(args[0], 10, 64)
//			if err != nil {
//				return err
//			}
//
//			uploadInterval, err := strconv.ParseUint(args[5], 10, 64)
//			if err != nil {
//				return err
//			}
//
//			operatingCost, err := strconv.ParseUint(args[6], 10, 64)
//			if err != nil {
//				return err
//			}
//
//			maxBundleSize, err := strconv.ParseUint(args[7], 10, 64)
//			if err != nil {
//				return err
//			}
//
//			minStake, err := strconv.ParseUint(args[8], 10, 64)
//			if err != nil {
//				return err
//			}
//
//			title, err := cmd.Flags().GetString(cli.FlagTitle)
//			if err != nil {
//				return err
//			}
//
//			description, err := cmd.Flags().GetString(cli.FlagDescription)
//			if err != nil {
//				return err
//			}
//
//			from := clientCtx.GetFromAddress()
//
//			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
//			if err != nil {
//				return err
//			}
//			deposit, err := sdk.ParseCoinsNormalized(depositStr)
//			if err != nil {
//				return err
//			}
//
//			content := types.NewUpdatePoolProposal(title, description, id, args[1], args[2], args[3], args[4], uploadInterval, operatingCost, maxBundleSize, minStake)
//
//			isExpedited, err := cmd.Flags().GetBool(cli.FlagIsExpedited)
//			if err != nil {
//				return err
//			}
//
//			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from, isExpedited)
//			if err != nil {
//				return err
//			}
//
//			if err = msg.ValidateBasic(); err != nil {
//				return err
//			}
//
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//
//	cmd.Flags().String(cli.FlagTitle, "", "The proposal title")
//	cmd.Flags().String(cli.FlagDescription, "", "The proposal description")
//	cmd.Flags().Bool(cli.FlagIsExpedited, false, "If true, makes the proposal an expedited one")
//	cmd.Flags().String(cli.FlagDeposit, "", "The proposal deposit")
//	_ = cmd.MarkFlagRequired(cli.FlagTitle)
//	_ = cmd.MarkFlagRequired(cli.FlagDescription)
//
//	return cmd
//}
