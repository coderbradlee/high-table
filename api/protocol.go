// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// Protocol defines the protocol interfaces
type Protocol interface {
	CreateTables(context.Context) error
	Initialize(context.Context, *sql.Tx) error
	GetDelegates(int, int) ([]*Delegate, error)
	UpdateDelegates(*Delegate) (bool, error)
}

// RowExists checks whether a row exists
func RowExists(db *sql.DB, query string, args ...interface{}) (bool, error) {
	exists := 0
	query = fmt.Sprintf("SELECT exists (%s)", query)
	stmt, err := db.Prepare(query)
	if err != nil {
		return false, errors.Wrap(err, "failed to prepare query")
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "failed to query the row")
	}
	if exists != 0 {
		return true, nil
	}
	return false, nil
}
