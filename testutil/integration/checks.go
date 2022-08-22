package integration

import (
	querytypes "github.com/KYVENetwork/chain/x/query/types"
	. "github.com/onsi/gomega"

	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) PerformValidityChecks() {
	// verify pool module
	suite.VerifyPoolModuleAssetsIntegrity()
	suite.VerifyPoolTotalFunds()
	suite.VerifyPoolQueries()

	// verify stakers module
	suite.VerifyStakersModuleAssetsIntegrity()
	suite.VerifyPoolTotalStake()
	suite.VerifyStakersQueries()

	// verify bundles module
	suite.VerifyBundlesQueries()
}

// ==================
// pool module checks
// ==================

func (suite *KeeperTestSuite) VerifyPoolModuleAssetsIntegrity() {
	expectedBalance := uint64(0)
	actualBalance := uint64(0)

	for _, pool := range suite.App().PoolKeeper.GetAllPools(suite.Ctx()) {
		for _, funder := range pool.Funders {
			expectedBalance += funder.Amount
		}
	}

	moduleAcc := suite.App().AccountKeeper.GetModuleAccount(suite.Ctx(), pooltypes.ModuleName).GetAddress()
	actualBalance = suite.App().BankKeeper.GetBalance(suite.Ctx(), moduleAcc, "tkyve").Amount.Uint64()

	Expect(actualBalance).To(Equal(expectedBalance))
}

func (suite *KeeperTestSuite) VerifyPoolTotalFunds() {
	for _, pool := range suite.App().PoolKeeper.GetAllPools(suite.Ctx()) {
		expectedBalance := uint64(0)
		actualBalance := pool.TotalFunds

		for _, funder := range pool.Funders {
			expectedBalance += funder.Amount
		}

		Expect(actualBalance).To(Equal(expectedBalance))
	}
}

func (suite *KeeperTestSuite) VerifyPoolQueries() {
	poolsState := suite.App().PoolKeeper.GetAllPools(suite.Ctx())

	poolsQuery := make([]querytypes.PoolResponse, 0)

	activePoolsQuery, activePoolsQueryErr := suite.App().QueryKeeper.Pools(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryPoolsRequest{})
	pausedPoolsQuery, pausedPoolsQueryErr := suite.App().QueryKeeper.Pools(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryPoolsRequest{
		Paused: true,
	})

	poolsQuery = append(poolsQuery, activePoolsQuery.Pools...)
	poolsQuery = append(poolsQuery, pausedPoolsQuery.Pools...)

	Expect(activePoolsQueryErr).To(BeNil())
	Expect(pausedPoolsQueryErr).To(BeNil())

	Expect(poolsQuery).To(HaveLen(len(poolsState)))

	for i := range poolsState {
		bundleProposalState, _ := suite.App().BundlesKeeper.GetBundleProposal(suite.Ctx(), poolsState[i].Id)
		stakersState := suite.App().StakersKeeper.GetAllStakerAddressesOfPool(suite.Ctx(), poolsState[i].Id)
		totalStakeState := suite.App().StakersKeeper.GetTotalStake(suite.Ctx(), poolsState[i].Id)

		Expect(poolsQuery[i].Id).To(Equal(poolsState[i].Id))
		Expect(*poolsQuery[i].Data).To(Equal(poolsState[i]))
		Expect(*poolsQuery[i].BundleProposal).To(Equal(bundleProposalState))
		Expect(poolsQuery[i].Stakers).To(Equal(stakersState))
		Expect(poolsQuery[i].TotalStake).To(Equal(totalStakeState))

		poolByIdQuery, poolByIdQueryErr := suite.App().QueryKeeper.Pool(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryPoolRequest{
			Id: poolsState[i].Id,
		})

		Expect(poolByIdQueryErr).To(BeNil())
		Expect(poolByIdQuery.Pool.Id).To(Equal(poolsState[i].Id))
		Expect(*poolByIdQuery.Pool.Data).To(Equal(poolsState[i]))
		Expect(*poolByIdQuery.Pool.BundleProposal).To(Equal(bundleProposalState))
		Expect(poolsQuery[i].Stakers).To(Equal(stakersState))
		Expect(poolByIdQuery.Pool.TotalStake).To(Equal(totalStakeState))
	}
}

// =====================
// stakers module checks
// =====================

func (suite *KeeperTestSuite) VerifyStakersModuleAssetsIntegrity() {
	expectedBalance := uint64(0)
	actualBalance := uint64(0)

	for _, staker := range suite.App().StakersKeeper.GetAllStakers(suite.Ctx()) {
		expectedBalance += staker.Amount
	}

	moduleAcc := suite.App().AccountKeeper.GetModuleAccount(suite.Ctx(), stakerstypes.ModuleName).GetAddress()
	actualBalance = suite.App().BankKeeper.GetBalance(suite.Ctx(), moduleAcc, "tkyve").Amount.Uint64()

	Expect(actualBalance).To(Equal(expectedBalance))
}

func (suite *KeeperTestSuite) VerifyPoolTotalStake() {
	for _, pool := range suite.App().PoolKeeper.GetAllPools(suite.Ctx()) {
		expectedBalance := uint64(0)
		actualBalance := suite.App().StakersKeeper.GetTotalStake(suite.Ctx(), pool.Id)

		for _, stakerAddress := range suite.App().StakersKeeper.GetAllStakerAddressesOfPool(suite.Ctx(), pool.Id) {
			staker, stakerFound := suite.App().StakersKeeper.GetStaker(suite.Ctx(), stakerAddress)

			if stakerFound {
				expectedBalance += staker.Amount
			}
		}

		Expect(actualBalance).To(Equal(expectedBalance))
	}
}

func (suite *KeeperTestSuite) VerifyStakersQueries() {
	stakersState := suite.App().StakersKeeper.GetAllStakers(suite.Ctx())
	stakersQuery, stakersQueryErr := suite.App().QueryKeeper.Stakers(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryStakersRequest{})

	Expect(stakersQueryErr).To(BeNil())
	Expect(stakersQuery.Stakers).To(HaveLen(len(stakersState)))

	for i := range stakersState {
		valaccounts := suite.App().StakersKeeper.GetValaccountsFromStaker(suite.Ctx(), stakersState[i].Address)

		Expect(*stakersQuery.Stakers[i].Staker).To(Equal(stakersState[i]))
		Expect(stakersQuery.Stakers[i].Valaccounts).To(Equal(valaccounts))

		stakerByAddressQuery, stakersByAddressQueryErr := suite.App().QueryKeeper.Staker(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryStakerRequest{
			Address: stakersState[i].Address,
		})

		Expect(stakersByAddressQueryErr).To(BeNil())
		Expect(*stakerByAddressQuery.Staker.Staker).To(Equal(stakersState[i]))
		Expect(stakerByAddressQuery.Staker.Valaccounts).To(Equal(valaccounts))
	}

	unbondingState := suite.App().StakersKeeper.GetAllUnbondingStakeEntries(suite.Ctx())
	unbondingQuery, unbondingQueryErr := suite.App().QueryKeeper.AccountStakingUnbondings(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryAccountStakingUnbondingsRequest{})

	Expect(unbondingQueryErr).To(BeNil())
	Expect(unbondingState).To(ContainElements(unbondingQuery.Unbondings))
}

// =====================
// bundles module checks
// =====================

func (suite *KeeperTestSuite) VerifyBundlesQueries() {
	pools := suite.App().PoolKeeper.GetAllPools(suite.Ctx())

	for _, pool := range pools {
		finalizedBundlesState := suite.App().BundlesKeeper.GetFinalizedBundlesByPool(suite.Ctx(), pool.Id)
		finalizedBundlesQuery, finalizedBundlesQueryErr := suite.App().QueryKeeper.FinalizedBundles(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryFinalizedBundlesRequest{
			PoolId: pool.Id,
		})

		Expect(finalizedBundlesQueryErr).To(BeNil())
		Expect(finalizedBundlesQuery.FinalizedBundles).To(HaveLen(len(finalizedBundlesState)))

		for i := range finalizedBundlesState {
			Expect(finalizedBundlesQuery.FinalizedBundles[i]).To(Equal(finalizedBundlesState[i]))

			finalizedBundleQuery, finalizedBundleQueryErr := suite.App().QueryKeeper.FinalizedBundle(sdk.WrapSDKContext(suite.Ctx()), &querytypes.QueryFinalizedBundleRequest{
				PoolId: pool.Id,
				Id:     finalizedBundlesState[i].Id,
			})

			Expect(finalizedBundleQueryErr).To(BeNil())
			Expect(finalizedBundleQuery.FinalizedBundle).To(Equal(finalizedBundlesState[i]))
		}
	}
}
