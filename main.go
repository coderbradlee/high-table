// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/iotexproject/iotex-core/pkg/log"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/high-table/api"
	"github.com/iotexproject/high-table/config"
	"github.com/iotexproject/high-table/core"
)

func main() {
	cfg := config.DefaultCfg
	configPath := os.Getenv("CONFIG")
	if configPath == "" {
		configPath = "config.yaml"
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.L().Fatal("Failed to load config file", zap.Error(err))
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.L().Fatal("failed to unmarshal config", zap.Error(err))
	}

	db, err := core.NewDB(cfg)
	if err != nil {
		log.S().Fatal("core.NewDB", zap.Error(err))
	}
	delegates := core.NewProtocol(db)
	if err = delegates.CreateTables(context.Background()); err != nil {
		log.S().Fatal("delegates.CreateTables", zap.Error(err))
	}
	ser, err := api.NewServer(cfg, delegates)
	if err != nil {
		log.S().Fatal("api.NewServer", zap.Error(err))
	}
	if err = ser.Start(); err != nil {
		log.L().Fatal("Failed to start api server", zap.Error(err))
	}
}
