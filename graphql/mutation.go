// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package graphql

import (
	"context"

	"github.com/iotexproject/high-table/api"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is hte resolver that handles GraphQL request
type Mutationer struct {
	Cli api.Protocol
}

//// Query returns a query resolver
func (m *Mutationer) Mutation() MutationResolver {
	return &mutationResolver{m}
}

type mutationResolver struct{ *Mutationer }

// Delegate handles delegate requests
func (r *mutationResolver) Delegate(ctx context.Context, delegate InputDelegate) (bool, error) {
	input := &api.Delegate{
		EpochNumber:    delegate.EpochNumber,
		DelegateID:     delegate.DelegateID,
		DelegateName:   delegate.DelegateName,
		DelegateNodeid: delegate.DelegateNodeid,
		GroupID:        delegate.GroupID,
		GroupName:      delegate.GroupName,
		ConsensusType:  delegate.ConsensusType,
		MaxTransNum:    delegate.MaxTransNum,
		GasLimit:       delegate.GasLimit,
	}
	return r.Cli.UpdateDelegates(input)
}
