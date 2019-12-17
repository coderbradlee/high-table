########################################################################################################################
# Copyright (c) 2018 IoTeX
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: gogen
gogen:
	@protoc --go_out=plugins=grpc:${GOPATH}/src ./proto/proto/types/*
	@protoc -I. -I./proto/proto/types --go_out=plugins=grpc:${GOPATH}/src ./proto/proto/api/*
	@protoc -I. --grpc-gateway_out=logtostderr=true:${GOPATH}/src ./proto/proto/api/*
