package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
)

var _ = Describe("Fund Pool", Ordered, func() {
	s := i.NewCleanChain()

	initialBalance := s.GetBalanceFromAddress(i.ALICE)

	BeforeAll(func() {
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
		})

		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.GetLowestFunder()).To(Equal(pooltypes.Funder{}))
	})

	It("Fund Pool with 100 $KYVE", func() {
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 100*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalance - balanceAfter).To(Equal(100*i.KYVE))

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

	It("Fund additional 50 $KYVE", func() {
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 50*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalance - balanceAfter).To(Equal(150*i.KYVE))

		Expect(pool.Funders).To(HaveLen(1))
		Expect(pool.TotalFunds).To(Equal(150*i.KYVE))

		funder, funderFound := pool.GetFunder(i.ALICE)

		Expect(funderFound).To(BeTrue())
		Expect(funder).To(Equal(pooltypes.Funder{
			Address: i.ALICE,
			Amount: 150*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funder))
	})

	It("Fund more than available balance", func() {
		currentBalance := s.GetBalanceFromAddress(i.ALICE)

		s.RunTxPoolError(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: currentBalance + 1,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalance - balanceAfter).To(Equal(150*i.KYVE))

		Expect(pool.Funders).To(HaveLen(1))
		Expect(pool.TotalFunds).To(Equal(150*i.KYVE))

		funder, funderFound := pool.GetFunder(i.ALICE)

		Expect(funderFound).To(BeTrue())
		Expect(funder).To(Equal(pooltypes.Funder{
			Address: i.ALICE,
			Amount: 150*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funder))
	})

	It("Fund with new funder less than existing one", func() {
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.BOB,
			Id: 0,
			Amount: 100*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.BOB)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalance - balanceAfter).To(Equal(100*i.KYVE))

		Expect(pool.Funders).To(HaveLen(2))
		Expect(pool.TotalFunds).To(Equal(250*i.KYVE))

		funder, funderFound := pool.GetFunder(i.BOB)

		Expect(funderFound).To(BeTrue())
		Expect(funder).To(Equal(pooltypes.Funder{
			Address: i.BOB,
			Amount: 100*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funder))
	})

	It("Fund new funder more than existing one", func() {
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.CHARLIE,
			Id: 0,
			Amount: 300*i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.CHARLIE)

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(initialBalance - balanceAfter).To(Equal(300*i.KYVE))

		Expect(pool.Funders).To(HaveLen(3))
		Expect(pool.TotalFunds).To(Equal(550*i.KYVE))

		funder, funderFound := pool.GetFunder(i.CHARLIE)
		funderBob, _ := pool.GetFunder(i.BOB)

		Expect(funderFound).To(BeTrue())
		Expect(funder).To(Equal(pooltypes.Funder{
			Address: i.CHARLIE,
			Amount: 300*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funderBob))
	})

	It("Try to fund less than the lowest funder with full slots", func () {
		// clean chain again
		s = i.NewCleanChain()

		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
		})

		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.GetLowestFunder()).To(Equal(pooltypes.Funder{}))

		// begin actual test
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 100*i.KYVE,
		})

		// fill remaining slots
		for a := 0; a < 49; a++ {
			s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
				Creator: i.DUMMY[a],
				Id: 0,
				Amount: 1000*i.KYVE,
			})
		}

		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(pool.Funders).To(HaveLen(50))
		Expect(pool.TotalFunds).To(Equal(49_100*i.KYVE))

		funderAlice, _ := pool.GetFunder(i.ALICE)
		Expect(pool.GetLowestFunder()).To(Equal(funderAlice))

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		Expect(initialBalance - balanceAfter).To(Equal(100*i.KYVE))

		s.RunTxPoolError(&pooltypes.MsgFundPool{
			Creator: i.DUMMY[49],
			Id: 0,
			Amount: 50*i.KYVE,
		})

		Expect(pool.Funders).To(HaveLen(50))
		Expect(pool.TotalFunds).To(Equal(49_100*i.KYVE))

		_, funderFound := pool.GetFunder(i.DUMMY[49])
		Expect(funderFound).To(BeFalse())
	})

	It("Try to fund more than the lowest funder with full slots", func () {
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.DUMMY[49],
			Id: 0,
			Amount: 200*i.KYVE,
		})

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(pool.Funders).To(HaveLen(50))
		Expect(pool.TotalFunds).To(Equal(49_200*i.KYVE))

		funderDummy, funderFound := pool.GetFunder(i.DUMMY[49])

		Expect(funderFound).To(BeTrue())
		Expect(funderDummy).To(Equal(pooltypes.Funder{
			Address: i.DUMMY[49],
			Amount: 200*i.KYVE,
		}))
		Expect(pool.GetLowestFunder()).To(Equal(funderDummy))

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		Expect(initialBalance - balanceAfter).To(BeZero())
	})
})
