package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bundles module sentinel errors
var (
	ErrUploaderAlreadyClaimed = sdkerrors.Register(ModuleName, 1100, "uploader role already claimed")
)
