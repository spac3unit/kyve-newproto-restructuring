package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
)

var _ = Describe("Defund Pool", Ordered, func() {
	s := i.NewCleanChain()

	initialBalanceAlice := s.GetBalanceFromAddress(i.ALICE)

	BeforeAll(func() {
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
		})

		_, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 100*i.KYVE,
		})
	})

	It("Defund more than funded", func() {
		s.RunTxPoolError(&pooltypes.MsgDefundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 101*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalanceAlice - balanceAfter).To(Equal(100*i.KYVE))

		Expect(pool.Funders).To(HaveLen(1))
		Expect(pool.TotalFunds).To(Equal(100*i.KYVE))

		funder, funderFound := pool.GetFunder(i.ALICE)

		Expect(funderFound).To(BeTrue())
		Expect(funder).To(Equal(pooltypes.Funder{
			Address: i.ALICE,
			Amount: 100*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funder))
	})

	It("Defund Pool with 50 $KYVE", func() {
		s.RunTxPoolSuccess(&pooltypes.MsgDefundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 50*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalanceAlice - balanceAfter).To(Equal(50*i.KYVE))

		Expect(pool.Funders).To(HaveLen(1))
		Expect(pool.TotalFunds).To(Equal(50*i.KYVE))

		funder, funderFound := pool.GetFunder(i.ALICE)

		Expect(funderFound).To(BeTrue())
		Expect(funder).To(Equal(pooltypes.Funder{
			Address: i.ALICE,
			Amount: 50*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funder))
	})

	It("Defund everything", func() {
		s.RunTxPoolSuccess(&pooltypes.MsgDefundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 50*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalanceAlice - balanceAfter).To(BeZero())

		Expect(pool.Funders).To(BeEmpty())
		Expect(pool.TotalFunds).To(BeZero())

		funder, funderFound := pool.GetFunder(i.ALICE)

		Expect(funderFound).To(BeFalse())
		Expect(funder).To(Equal(pooltypes.Funder{}))
		Expect(pool.GetLowestFunder()).To(Equal(funder))
	})

	// TODO: test kicking out lowest funder
	// TODO: fund with more than balance
	// TODO: test for two funders with same amount
	// TODO: test modifying current funds
	// TODO: test funder payout
})
