// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package graphql

import (
	"context"

	"github.com/pkg/errors"

	"github.com/iotexproject/high-table/api"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is the resolver that handles GraphQL request
type Resolver struct {
	Cli api.Protocol
}

// Query returns a query resolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// Delegates handles delegate requests
func (r *queryResolver) Delegates(ctx context.Context, epochNum int, groupID int) ([]*Delegate, error) {
	return r.getDelegates(ctx, epochNum, groupID)
}

func (r *queryResolver) getDelegates(ctx context.Context, epochNum int, groupID int) (ret []*Delegate, err error) {
	delegates, err := r.Cli.GetDelegates(epochNum, groupID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get delegates information")
	}
	for _, d := range delegates {
		ret = append(ret, &Delegate{
			EpochNumber:    d.EpochNumber,
			DelegateID:     d.DelegateID,
			DelegateName:   d.DelegateName,
			DelegateNodeid: d.DelegateNodeid,
			GroupID:        d.GroupID,
			GroupName:      d.GroupName,
			ConsensusType:  d.ConsensusType,
			MaxTransNum:    d.MaxTransNum,
			GasLimit:       d.GasLimit,
		})
	}
	return
}

// UpdateDelegate handles delegate requests
func (r *queryResolver) UpdateDelegate(ctx context.Context, delegate InputDelegate) (bool, error) {
	input := &api.Delegate{
		EpochNumber: delegate.EpochNumber,
		DelegateID:  delegate.DelegateID,
		GroupID:     delegate.GroupID,
	}
	if delegate.DelegateName != nil {
		input.DelegateName = *delegate.DelegateName
	}
	if delegate.DelegateNodeid != nil {
		input.DelegateNodeid = *delegate.DelegateNodeid
	}
	if delegate.GroupName != nil {
		input.GroupName = *delegate.GroupName
	}
	if delegate.ConsensusType != nil {
		input.ConsensusType = *delegate.ConsensusType
	}
	if delegate.MaxTransNum != nil {
		input.MaxTransNum = *delegate.MaxTransNum
	}
	if delegate.GasLimit != nil {
		input.GasLimit = *delegate.GasLimit
	}
	return r.Cli.UpdateDelegates(input)
}
