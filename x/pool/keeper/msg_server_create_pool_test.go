package keeper_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var _ = Describe("Create Pool", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyPoolTotalFunds()
	})

	It("Create Pool", func() {
		// ACT
		proposal := &pooltypes.CreatePoolProposal{
			Title: "Create Moonbeam Pool",
			Description: "Create Moonbeam Pool",
			Name: "Moonbeam",
			Runtime: "@kyve/evm",
			Logo: "https://arweave.net/9FJDam56yBbmvn8rlamEucATH5UcYqSBw468rlCXn8E",
			Config: "{}",
			StartKey: "0",
			UploadInterval: 60,
			OperatingCost: 2_500_000_000,
			MinStake: 100_000_000_000,
			MaxBundleSize: 100,
			Version: "0.0.0",
			Binaries: "{}",
		}
		content, _ := codectypes.NewAnyWithValue(proposal)

		// params := s.App().GovKeeper.GetVotingParams(s.Ctx())
		// fmt.Println(params)

		s.RunTxGovSuccess(&govtypes.MsgSubmitProposal{
			Content: content,
			InitialDeposit: sdk.NewCoins(sdk.NewInt64Coin(i.KYVE_DENOM, int64(100*i.KYVE))),
			Proposer: i.ALICE,
			IsExpedited: false,
		})

		// TODO: still rejected
		s.RunTxGovSuccess(&govtypes.MsgVote{
			ProposalId: 1,
			Voter: "kyve1haclwclymmfszwwyq9s2uaxl3qw4c73dej0zsv",
			Option: govtypes.VoteOption(1),
		})

		proposals := s.App().GovKeeper.GetProposals(s.Ctx())
		fmt.Println(proposals)
		
		// wait for vote end
		s.CommitAfterSeconds(uint64(govtypes.DefaultVotingParams().GetVotingPeriod(false).Seconds()))
		s.CommitAfterSeconds(1)

		proposals = s.App().GovKeeper.GetProposals(s.Ctx())
		fmt.Println(proposals)

		// ASSERT
		_, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeTrue())
	})
})
