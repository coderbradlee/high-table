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

	"github.com/iotexproject/high-table/config"
	"github.com/iotexproject/high-table/core"
	iotexapi "github.com/iotexproject/high-table/proto/golang/api"
	iotextypes "github.com/iotexproject/high-table/proto/golang/types"
)

var (
	// ErrInternalServer indicates the internal server error
	ErrInternalServer = errors.New("internal server error")
)

// Server provides api for user to query blockchain data
type Server struct {
	cfg        *config.Config
	grpcserver *grpc.Server
	protocol   core.Protocol
}

// NewServer creates a new server
func NewServer(
	cfg *config.Config,
	protocol core.Protocol,
) (*Server, error) {
	svr := &Server{
		cfg:      cfg,
		protocol: protocol,
	}
	svr.grpcserver = grpc.NewServer()
	iotexapi.RegisterAPIServiceServer(svr.grpcserver, svr)
	return svr, nil
}

// Start starts the API server
func (api *Server) Start() error {
	portStr := ":" + api.cfg.Port
	lis, err := net.Listen("tcp", portStr)
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
	response.Delegate = delegate
	return
}
func (api *Server) UpdateDelegate(ctx context.Context, in *iotexapi.UpdateDelegateRequest) (response *iotexapi.UpdateDelegateResponse, err error) {
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
