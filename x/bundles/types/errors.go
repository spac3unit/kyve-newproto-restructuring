package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bundles module sentinel errors
var (
	ErrUploaderAlreadyClaimed = sdkerrors.Register(ModuleName, 1100, "uploader role already claimed")
	ErrInvalidArgs            = sdkerrors.Register(ModuleName, 1107, "invalid args")
	ErrFromHeight             = sdkerrors.Register(ModuleName, 1118, "invalid from height")
	ErrToHeight               = sdkerrors.Register(ModuleName, 1123, "invalid to height")
	ErrNotDesignatedUploader  = sdkerrors.Register(ModuleName, 1113, "not designated uploader")
	ErrUploadInterval         = sdkerrors.Register(ModuleName, 1108, "upload interval not surpassed")
	ErrMaxBundleSize          = sdkerrors.Register(ModuleName, 1109, "max bundle size was surpassed")
	ErrFromKey                = sdkerrors.Register(ModuleName, 1124, "invalid from key")
	ErrVoterIsUploader        = sdkerrors.Register(ModuleName, 1112, "voter is uploader")
	ErrInvalidVote            = sdkerrors.Register(ModuleName, 1119, "invalid vote %v")
	ErrAlreadyVoted           = sdkerrors.Register(ModuleName, 1110, "already voted on proposal %v")
	ErrInvalidStorageId       = sdkerrors.Register(ModuleName, 1209, "current storageId %v does not match provided storageId")
)
