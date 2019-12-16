// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// usage: go build -o ./bin/server -v .
// ./bin/server

package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/iotexproject/iotex-core/pkg/log"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/high-table/api"
	"github.com/iotexproject/high-table/graphql"
	"github.com/iotexproject/high-table/sql"
)

// Config is config
type Config struct {
	Port       string `yaml:"port"`
	Connection string `yaml:"connection"`
	DBName     string `yaml:"DBName"`
}

const defaultPort = "8089"

func main() {
	configPath := os.Getenv("CONFIG")
	if configPath == "" {
		configPath = "config.yaml"
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.L().Fatal("Failed to load config file", zap.Error(err))
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.L().Fatal("failed to unmarshal config", zap.Error(err))
	}

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = defaultPort
	}
	cfg.Connection = os.Getenv("CONNECTION_STRING")
	if cfg.Connection == "" {
		cfg.Connection = "root:rootuser@tcp(127.0.0.1:3306)/"
	}

	cfg.DBName = os.Getenv("DB_NAME")
	if cfg.DBName == "" {
		cfg.DBName = "high_table"
	}

	store := sql.NewMySQL(cfg.Connection, cfg.DBName)
	err = store.Start(context.Background())
	if err != nil {
		log.S().Error("store.Start", zap.Error(err))
		return
	}

	delegates := api.NewProtocol(store)
	err = delegates.CreateTables(context.Background())
	if err != nil {
		log.S().Error("delegates.CreateTables", zap.Error(err))
		return
	}
	http.Handle("/", graphqlHandler(handler.Playground("GraphQL playground", "/query")))
	http.Handle("/query", graphqlHandler(handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{delegates}}))))

	log.S().Infof("connect to http://0.0.0.0:%s/ for GraphQL playground", cfg.Port)

	// Start GraphQL query service
	go func() {
		if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
			log.L().Fatal("Failed to serve index query service", zap.Error(err))
		}
	}()

	select {}
}

func graphqlHandler(playgroundHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		playgroundHandler.ServeHTTP(w, r)
	})
}
