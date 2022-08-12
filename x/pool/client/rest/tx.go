package rest

import (
	"net/http"

	"github.com/KYVENetwork/chain/x/pool/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type CreatePoolRequest struct {
	BaseReq        rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title          string       `json:"title" yaml:"title"`
	Description    string       `json:"description" yaml:"description"`
	IsExpedited    bool         `json:"is_expedited" yaml:"is_expedited"`
	Deposit        sdk.Coins    `json:"deposit" yaml:"deposit"`
	Name           string       `json:"name" yaml:"name"`
	Runtime        string       `json:"runtime" yaml:"runtime"`
	Logo           string       `json:"logo" yaml:"logo"`
	Config         string       `json:"config" yaml:"config"`
	StartKey       string       `json:"startKey" yaml:"startKey"`
	UploadInterval uint64       `json:"uploadInterval" yaml:"uploadInterval"`
	OperatingCost  uint64       `json:"operatingCost" yaml:"operatingCost"`
	MinStake       uint64       `json:"minStake" yaml:"minStake"`
	MaxBundleSize  uint64       `json:"maxBundleSize" yaml:"maxBundleSize"`
	Version        string       `json:"version" yaml:"version"`
	Binaries       string       `json:"binaries" yaml:"binaries"`
}

func ProposalCreatePoolRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "create-pool",
		Handler:  newCreatePoolHandler(clientCtx),
	}
}

func newCreatePoolHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePoolRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewCreatePoolProposal(req.Title, req.Description, req.Name, req.Runtime, req.Logo, req.Config, req.StartKey, req.UploadInterval, req.OperatingCost, req.MinStake, req.MaxBundleSize, req.Version, req.Binaries)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr, req.IsExpedited)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}