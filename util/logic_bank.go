package util

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// TransferToAddress sends tokens from the given module to a specified address.
func TransferToAddress(bankKeeper bankkeeper.Keeper, ctx sdk.Context, module string, address string, amount uint64) error {
	recipient, errAddress := sdk.AccAddressFromBech32(address)
	if errAddress != nil {
		return errAddress
	}

	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))
	err := bankKeeper.SendCoinsFromModuleToAccount(ctx, module, recipient, coins)
	return err
}

// TransferToRegistry sends tokens from a specified address to the given module.
func TransferToModule(bankKeeper bankkeeper.Keeper, ctx sdk.Context, module string, address string, amount uint64) error {
	sender, errAddress := sdk.AccAddressFromBech32(address)
	if errAddress != nil {
		return errAddress
	}
	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, sender, module, coins)
	return err
}

// TransferInterModule ...
func TransferInterModule(bankKeeper bankkeeper.Keeper, ctx sdk.Context, fromModule string, toModule string, amount uint64) error {
	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))
	err := bankKeeper.SendCoinsFromModuleToModule(ctx, fromModule, toModule, coins)
	return err
}

// transferToTreasury sends tokens from this module to the treasury (community spend pool).
func transferToTreasury(bankKeeper bankkeeper.Keeper, ctx sdk.Context, amount uint64) error {
	//sender := bankKeeper.accountKeeper.GetModuleAddress(types.ModuleName)
	//coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))
	//
	//err := k.distrKeeper.FundCommunityPool(ctx, coins, sender)
	//return err
	// TODO implement treasury transfer
	return nil
}
