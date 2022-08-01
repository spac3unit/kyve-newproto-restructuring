package cli

import (
	"github.com/KYVENetwork/chain/x/registry/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query params",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			_ = queryClient

			//params := &types.QueryParamsRequest{}

			//res, err := queryClient.Params(cmd.Context(), params)
			//if err != nil {
			//	return err
			//}

			//return clientCtx.PrintProto(res)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
