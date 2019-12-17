// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package api

import (
	"context"
	"net"

	"github.com/iotexproject/iotex-core/pkg/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iotexproject/high-table/config"
	"github.com/iotexproject/high-table/core"
	iotexapi "github.com/iotexproject/high-table/proto/golang/api"
	iotextypes "github.com/iotexproject/high-table/proto/golang/types"
)

// Server provides api for user to query data
type Server struct {
	cfg        *config.Config
	grpcserver *grpc.Server
	protocol   *core.Delegates
}

// NewServer creates a new server
func NewServer(
	cfg *config.Config,
	protocol *core.Delegates,
) (*Server, error) {
	svr := &Server{
		cfg:      cfg,
		protocol: protocol,
	}
	svr.grpcserver = grpc.NewServer()
	iotexapi.RegisterAPIServiceServer(svr.grpcserver, svr)
	reflection.Register(svr.grpcserver)
	return svr, nil
}

// Start starts the API server
func (api *Server) Start() error {
	lis, err := net.Listen("tcp", ":"+api.cfg.Port)
	if err != nil {
		log.L().Error("API server failed to listen.", zap.Error(err))
		return errors.Wrap(err, "API server failed to listen")
	}
	log.L().Info("API server is listening.", zap.String("addr", lis.Addr().String()))

	go func() {
		if err := api.grpcserver.Serve(lis); err != nil {
			log.L().Fatal("Node failed to serve.", zap.Error(err))
		}
	}()
	return nil
}

// Stop stops the API server
func (api *Server) Stop() error {
	api.grpcserver.Stop()
	return nil
}

// GetDelegate returns the delegate
func (api *Server) GetDelegate(ctx context.Context, in *iotexapi.GetDelegateRequest) (response *iotexapi.GetDelegateResponse, err error) {
	ret, err := api.protocol.GetDelegates(core.Uint64ToInt64(in.DelegateID))
	if err != nil {
		return
	}
	delegate := &iotextypes.Delegate{
		DelegateID: in.DelegateID,
		Address:    ret,
	}
	response = &iotexapi.GetDelegateResponse{}
	response.Delegate = delegate
	return
}

// UpdateDelegate update delegate info
func (api *Server) UpdateDelegate(ctx context.Context, in *iotexapi.UpdateDelegateRequest) (response *iotexapi.UpdateDelegateResponse, err error) {
	response = &iotexapi.UpdateDelegateResponse{
		Success: true,
	}
	del := &core.Delegate{
		core.Uint64ToInt64(in.Delegate.DelegateID),
		in.Delegate.Address,
	}
	err = api.protocol.UpdateDelegates(del)
	if err != nil {
		response.Success = false
	}
	return
}
