// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package core

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

const (
	ProtocolID          = "delegates"
	delegateTableName   = "delegate_list"
	createDelegateTable = `CREATE TABLE IF NOT EXISTS %s (delegate_id INTEGER PRIMARY KEY, address BLOB)`
	selectDelegates     = "SELECT address from %s where delegate_id=?"
	insertDelegates     = "INSERT OR IGNORE INTO %s (delegate_id, address) VALUES (?, ?)"
)

var (
	// ErrNotExist define not exist error
	ErrNotExist = errors.New("not exist")
)

type Protocol interface {
	CreateTables(context.Context) error
	Delegate(int64) (string, error)
	UpdateDelegate(*Delegate) error
}

// Delegate defines the protocol of querying tables
type Delegate struct {
	DelegateID int64  `json:"delegate_id"`
	Address    string `json:"address"`
}

// Delegates defines the delegate protocol
type Delegates struct {
	db *sql.DB
}

// NewProtocol creates a new protocol
func NewProtocol(
	db *sql.DB,
) Protocol {
	return &Delegates{
		db: db,
	}
}

// CreateTables creates tables
func (p *Delegates) CreateTables(ctx context.Context) error {
	db := p.db
	if db == nil {
		return errors.New("db is nil")
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(fmt.Sprintf(createDelegateTable, delegateTableName))
	if err != nil {
		return err
	}
	defer tx.Rollback()
	return tx.Commit()
}

// Delegate gets delegate from db
func (p *Delegates) Delegate(delegateID int64) (ret string, err error) {
	db := p.db
	if db == nil {
		err = errors.New("db is nil")
		return
	}
	getQuery := fmt.Sprintf(selectDelegates, delegateTableName)
	err = db.QueryRow(getQuery, delegateID).Scan(&ret)
	return
}

// UpdateDelegates insert or update delegate's table
func (p *Delegates) UpdateDelegate(delegate *Delegate) (err error) {
	db := p.db
	if db == nil {
		err = errors.New("db is nil")
		return
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.Exec(fmt.Sprintf(insertDelegates, delegateTableName), delegate.DelegateID, delegate.Address); err != nil {
		return err
	}
	defer tx.Rollback()
	return tx.Commit()
}
