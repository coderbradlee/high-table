// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package api

import (
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/iotexproject/iotex-core/config"
)

var (
	// ErrInternalServer indicates the internal server error
	ErrInternalServer = errors.New("internal server error")
)

// Server provides api for user to query blockchain data
type Server struct {
	cfg        config.Config
	grpcserver *grpc.Server
}

// NewServer creates a new server
func NewServer(
	cfg config.Config,
) (*Server, error) {

	svr := &Server{
		cfg: cfg,
	}
	svr.grpcserver = grpc.NewServer()
	iotexapi.RegisterAPIServiceServer(svr.grpcserver, svr)

	return svr, nil
}

// Start starts the API server
func (api *Server) Start() error {
	//portStr := ":" + strconv.Itoa(api.cfg.API.Port)
	//lis, err := net.Listen("tcp", portStr)
	//if err != nil {
	//	log.L().Error("API server failed to listen.", zap.Error(err))
	//	return errors.Wrap(err, "API server failed to listen")
	//}
	//log.L().Info("API server is listening.", zap.String("addr", lis.Addr().String()))
	//
	//go func() {
	//	if err := api.grpcserver.Serve(lis); err != nil {
	//		log.L().Fatal("Node failed to serve.", zap.Error(err))
	//	}
	//}()
	//if err := api.bc.AddSubscriber(api.chainListener); err != nil {
	//	return errors.Wrap(err, "failed to subscribe to block creations")
	//}
	//if err := api.chainListener.Start(); err != nil {
	//	return errors.Wrap(err, "failed to start blockchain listener")
	//}
	return nil
}

// Stop stops the API server
func (api *Server) Stop() error {
	api.grpcserver.Stop()
}
