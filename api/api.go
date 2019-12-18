// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package api

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

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
	cfg          *config.Config
	grpcserver   *grpc.Server
	protocol     core.Protocol
	shutdown     chan struct{}
	shuttingDown int32
	interrupted  bool
	interrupt    chan os.Signal
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
	reflection.Register(svr.grpcserver)
	return svr, nil
}

// Start starts the API server
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":"+s.cfg.Port)
	if err != nil {
		log.L().Error("API server failed to listen.", zap.Error(err))
		return errors.Wrap(err, "API server failed to listen")
	}

	wg := new(sync.WaitGroup)
	once := new(sync.Once)
	signalNotify(s.interrupt)
	go handleInterrupt(once, s)

	wg.Add(1)
	go s.handleShutdown(wg, s.grpcserver)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.grpcserver.Serve(lis); err != nil {
			log.L().Fatal("Node failed to serve.", zap.Error(err))
		}
		log.L().Info("API server is listening.", zap.String("addr", lis.Addr().String()))
	}()

	wg.Wait()
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

// Shutdown server and clean up resources
func (s *Server) Shutdown() error {
	if atomic.CompareAndSwapInt32(&s.shuttingDown, 0, 1) {
		close(s.shutdown)
	}
	return nil
}

func (s *Server) handleShutdown(wg *sync.WaitGroup, gs *grpc.Server) {
	defer wg.Done()
	<-s.shutdown
	gs.Stop()
}

func handleInterrupt(once *sync.Once, s *Server) {
	once.Do(func() {
		for _ = range s.interrupt {
			if s.interrupted {
				log.L().Info("Server already shutting down")
				continue
			}
			s.interrupted = true
			log.L().Info("Shutting down... ")
			if err := s.Shutdown(); err != nil {
				log.L().Info("HTTP server Shutdown: %v", zap.Error(err))
			}
		}
	})
}

func signalNotify(interrupt chan<- os.Signal) {
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
}
