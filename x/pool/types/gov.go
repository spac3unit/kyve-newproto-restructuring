package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeCreatePool          = "CreatePool"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreatePool)
	govtypes.RegisterProposalTypeCodec(&CreatePoolProposal{}, "kyve/CreatePoolProposal")
}

var (
	_ govtypes.Content = &CreatePoolProposal{}
)

func NewCreatePoolProposal(title string, description string, name string, runtime string, logo string, config string, startKey string, uploadInterval uint64, operatingCost uint64, minStake uint64, maxBundleSize uint64, version string, binaries string) govtypes.Content {
	return &CreatePoolProposal{
		Title:          title,
		Description:    description,
		Name:           name,
		Runtime:        runtime,
		Logo:           logo,
		Config:         config,
		StartKey: startKey,
		UploadInterval: uploadInterval,
		OperatingCost:  operatingCost,
		MinStake: minStake,
		MaxBundleSize:  maxBundleSize,
		Version:        version,
		Binaries:       binaries,
	}
}

func (p *CreatePoolProposal) ProposalRoute() string { return RouterKey }

func (p *CreatePoolProposal) ProposalType() string {
	return ProposalTypeCreatePool
}

func (p *CreatePoolProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
