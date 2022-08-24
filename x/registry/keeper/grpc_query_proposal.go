package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// ProposalByHeight returns the proposal which contains the requested height of the datasource.
func (k Keeper) ProposalByHeight(goCtx context.Context, req *types.QueryProposalByHeightRequest) (*types.QueryProposalByHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proposalPrefixBuilder := types.KeyPrefixBuilder{Key: types.ProposalKeyPrefixIndex2}.AInt(req.PoolId)
	proposalIndexStore := prefix.NewStore(ctx.KVStore(k.storeKey), proposalPrefixBuilder.Key)
	proposalIndexIterator := proposalIndexStore.ReverseIterator(nil, types.KeyPrefixBuilder{}.AInt(req.Height+1).Key)

	defer proposalIndexIterator.Close()

	if proposalIndexIterator.Valid() {

		storageId := string(proposalIndexIterator.Value())

		proposal, found := k.GetProposal(ctx, storageId)
		if found {
			if proposal.FromHeight <= req.Height && proposal.ToHeight > req.Height {
				return &types.QueryProposalByHeightResponse{
					Proposal: proposal,
				}, nil
			}
		}
	}

	return nil, status.Error(codes.NotFound, "no bundle found")
}

// ProposalSinceFinalizedAt returns all proposals since a given finalizedAt height.
func (k Keeper) ProposalSinceFinalizedAt(goCtx context.Context, req *types.QueryProposalSinceFinalizedAtRequest) (*types.QueryProposalSinceFinalizedAtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proposalPrefixBuilder := types.KeyPrefixBuilder{Key: types.ProposalKeyPrefixIndex3}.AInt(req.PoolId)
	proposalIndexStore := prefix.NewStore(ctx.KVStore(k.storeKey), proposalPrefixBuilder.Key)

	if req.Pagination == nil {
		req.Pagination = &query.PageRequest{}
	}

	if req.Pagination.Key == nil {
		// Find optimal key for query
		proposalIndexIterator := proposalIndexStore.Iterator(types.KeyPrefixBuilder{}.AInt(req.FinalizedAt).Key, nil)
		defer proposalIndexIterator.Close()

		if proposalIndexIterator.Valid() {
			req.Pagination.Key = proposalIndexIterator.Key()
		} else {
			return nil, status.Error(codes.NotFound, "no bundle found")
		}
	}

	var proposals []types.Proposal

	pageRes, err := query.FilteredPaginate(proposalIndexStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			storageId := string(value)
			proposal, found := k.GetProposal(ctx, storageId)
			if !found {
				return false, status.Error(codes.Internal, "storageId should exist: "+storageId)
			}
			proposals = append(proposals, proposal)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProposalSinceFinalizedAtResponse{
		Proposals:  proposals,
		Pagination: pageRes,
	}, nil
}

// ProposalSinceId returns all proposals since a given id height for a given pool
// TODO should be possible to query efficiently with page + page_size additionally of pagination key
func (k Keeper) ProposalSinceId(goCtx context.Context, req *types.QueryProposalSinceIdRequest) (*types.QueryProposalSinceIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proposalPrefixBuilder := types.KeyPrefixBuilder{Key: types.ProposalKeyPrefixIndex2}.AInt(req.PoolId)
	proposalIndexStore := prefix.NewStore(ctx.KVStore(k.storeKey), proposalPrefixBuilder.Key)

	if req.Pagination == nil {
		req.Pagination = &query.PageRequest{}
	}

	if req.Pagination.Key == nil {
		// Find optimal key for query
		proposalIndexIterator := proposalIndexStore.Iterator(types.KeyPrefixBuilder{}.AInt(req.Id).Key, nil)
		defer proposalIndexIterator.Close()

		if proposalIndexIterator.Valid() {
			req.Pagination.Key = proposalIndexIterator.Key()
		} else {
			return nil, status.Error(codes.NotFound, "no bundle found")
		}
	}

	var proposals []types.Proposal

	pageRes, err := query.FilteredPaginate(proposalIndexStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			storageId := string(value)
			proposal, found := k.GetProposal(ctx, storageId)
			if !found {
				return false, status.Error(codes.Internal, "storageId should exist: "+storageId)
			}
			proposals = append(proposals, proposal)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProposalSinceIdResponse{
		Proposals:  proposals,
		Pagination: pageRes,
	}, nil
}
