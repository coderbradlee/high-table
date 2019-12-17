// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package core

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"github.com/iotexproject/high-table/config"
)

const (
	maxUint = ^uint(0)
	maxInt  = int64(maxUint >> 1)
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	return sql.Open("sqlite3", cfg.DBPath)
}

//Uint64ToInt64 converts uint64 to int64
func Uint64ToInt64(u uint64) int {
	if u > uint64(maxInt) {
		zap.L().Panic("Height can't be converted to int64")
	}
	return int(u)
}
