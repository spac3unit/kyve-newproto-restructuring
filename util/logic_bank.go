package util

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
)

// TransferToAddress sends tokens from the given module to a specified address.
func TransferFromModuleToAddress(bankKeeper bankkeeper.Keeper, ctx sdk.Context, module string, address string, amount uint64) error {
	recipient, errAddress := sdk.AccAddressFromBech32(address)
	if errAddress != nil {
		return errAddress
	}

	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))
	err := bankKeeper.SendCoinsFromModuleToAccount(ctx, module, recipient, coins)
	return err
}

// TransferToRegistry sends tokens from a specified address to the given module.
func TransferFromAddressToModule(bankKeeper bankkeeper.Keeper, ctx sdk.Context, address string, module string, amount uint64) error {
	sender, errAddress := sdk.AccAddressFromBech32(address)
	if errAddress != nil {
		return errAddress
	}
	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, sender, module, coins)
	return err
}

// TransferInterModule ...
func TransferFromModuleToModule(bankKeeper bankkeeper.Keeper, ctx sdk.Context, fromModule string, toModule string, amount uint64) error {
	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))
	err := bankKeeper.SendCoinsFromModuleToModule(ctx, fromModule, toModule, coins)
	return err
}

// transferToTreasury sends tokens from this module to the treasury (community spend pool).
func TransferFromAddressToTreasury(distrKeeper distrKeeper.Keeper, ctx sdk.Context, address string, amount uint64) error {
	sender, errAddress := sdk.AccAddressFromBech32(address)
	if errAddress != nil {
		return errAddress
	}
	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))

	if err := distrKeeper.FundCommunityPool(ctx, coins, sender); err != nil {
		return err
	}

	return nil
}

// transferToTreasury sends tokens from this module to the treasury (community spend pool).
func TransferFromModuleToTreasury(accountKeeper authkeeper.AccountKeeper, distrKeeper distrKeeper.Keeper, ctx sdk.Context, module string, amount uint64) error {
	sender := accountKeeper.GetModuleAddress(module)
	coins := sdk.NewCoins(sdk.NewInt64Coin("tkyve", int64(amount)))

	if err := distrKeeper.FundCommunityPool(ctx, coins, sender); err != nil {
		return err
	}

	return nil
}
