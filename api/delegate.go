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

	s "github.com/iotexproject/high-table/sql"
)

const (
	ProtocolID          = "delegates"
	delegateTableName   = "delegate_list"
	createDelegateTable = `CREATE TABLE IF NOT EXISTS %s (
     epoch_number decimal(65, 0) NOT NULL,
     delegate_id decimal(65, 0) NOT NULL,
     delegate_name varchar(255),
     delegate_nodeid varchar(255),
     group_id decimal(65, 0) NOT NULL,
	 group_name varchar(255),
	 consensus_type varchar(255) NOT NULL,
	 max_trans_num decimal(65, 0),
	 gas_limit decimal(65, 0) NOT NULL,
	 PRIMARY KEY (epoch_number, delegate_id, group_id),
	 UNIQUE INDEX delegate_group_index(epoch_number,delegate_id, group_id))`
	selectDelegates = "SELECT epoch_number,	delegate_id,delegate_name,delegate_nodeid,group_id,group_name,consensus_type,max_trans_num,gas_limit from %s where epoch_number=? and group_id=?"
	existDelegates  = "SELECT * from %s where epoch_number=? and group_id=? and delegate_id=?"
	insertDelegates = "INSERT INTO %s (epoch_number,delegate_id,delegate_name,delegate_nodeid,group_id,group_name,consensus_type,max_trans_num,gas_limit) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

var (
	ErrNotExist = errors.New("not exist")
)

type Delegate struct {
	EpochNumber    int    `json:"epoch_number"`
	DelegateID     int    `json:"delegate_id"`
	DelegateName   string `json:"delegate_name"`
	DelegateNodeid string `json:"delegate_nodeid"`
	GroupID        int    `json:"group_id"`
	GroupName      string `json:"group_name"`
	ConsensusType  string `json:"consensus_type"`
	MaxTransNum    int    `json:"max_trans_num"`
	GasLimit       int    `json:"gas_limit"`
}

// Protocol defines the protocol of querying tables
type Delegates struct {
	Store s.Store
}

// NewProtocol creates a new protocol
func NewProtocol(
	store s.Store,
) Protocol {
	return &Delegates{
		Store: store,
	}
}

// CreateTables creates tables
func (p *Delegates) CreateTables(ctx context.Context) error {
	// create reward history table
	if _, err := p.Store.GetDB().Exec(fmt.Sprintf(createDelegateTable,
		delegateTableName)); err != nil {
		return err
	}
	return nil
}
func (p *Delegates) Initialize(context.Context, *sql.Tx) error {
	return nil
}

// GetActionsByDates gets actions by start date and end date
func (p *Delegates) GetDelegates(epochNum int, groupID int) (ret []*Delegate, err error) {
	db := p.Store.GetDB()
	if db == nil {
		return nil, errors.New("db is nil")
	}
	getQuery := fmt.Sprintf(selectDelegates, delegateTableName)
	stmt, err := db.Prepare(getQuery)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare get query")
	}
	defer stmt.Close()

	rows, err := stmt.Query(epochNum, groupID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute get query")
	}

	var delegate Delegate
	parsedRows, err := s.ParseSQLRows(rows, &delegate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse results")
	}
	if len(parsedRows) == 0 {
		err = ErrNotExist
		return
	}

	for _, parsedRow := range parsedRows {
		ret = append(ret, parsedRow.(*Delegate))
	}
	return
}
func (p *Delegates) UpdateDelegates(delegate *Delegate) (ok bool, err error) {
	db := p.Store.GetDB()
	if db == nil {
		return false, errors.New("db is nil")
	}
	getQuery := fmt.Sprintf(existDelegates, delegateTableName)
	stmt, err := db.Prepare(getQuery)
	if err != nil {
		return false, errors.Wrap(err, "failed to prepare get query")
	}
	defer stmt.Close()
	exist, err := RowExists(db, getQuery, delegate.EpochNumber, delegate.GroupID, delegate.DelegateID)
	if exist {
		// update
		fmt.Println("exist")
		return false, nil
	}

	insert := fmt.Sprintf(insertDelegates, delegateTableName)
	if _, err := db.Exec(insert, delegate.EpochNumber, delegate.DelegateID, delegate.DelegateName, delegate.DelegateNodeid, delegate.GroupID, delegate.GroupName, delegate.ConsensusType, delegate.MaxTransNum, delegate.GasLimit); err != nil {
		return false, errors.Wrapf(err, "failed to update delegates")
	}
	return true, nil
}
