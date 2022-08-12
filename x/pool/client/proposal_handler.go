package client

import (
	"github.com/KYVENetwork/chain/x/pool/client/cli"
	"github.com/KYVENetwork/chain/x/pool/client/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var CreatePoolHandler = govclient.NewProposalHandler(cli.CmdSubmitCreatePoolProposal, rest.ProposalCreatePoolRESTHandler)
