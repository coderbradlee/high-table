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
	 UNIQUE INDEX delegate_group_index(epoch_number,delegate_id, group_id))`
	selectActionHistoryByTimestamp = "SELECT action_hash, block_hash, timestamp, action_type, `from`, `to`, amount, t1.gas_price*t1.gas_consumed " +
		"FROM %s AS t1 LEFT JOIN %s AS t2 ON t1.block_height=t2.block_height " +
		"WHERE timestamp >= ? AND timestamp <= ? ORDER BY `timestamp` desc limit ?,?"
)

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
//func (p *Protocol) GetActionsByDates(startDate, endDate uint64, offset, size int) ([]*ActionInfo, error) {
//	if _, ok := p.indexer.Registry.Find(actions.ProtocolID); !ok {
//		return nil, errors.New("actions protocol is unregistered")
//	}
//
//	db := p.indexer.Store.GetDB()
//
//	getQuery := fmt.Sprintf(selectActionHistoryByTimestamp, actions.ActionHistoryTableName, blocks.BlockHistoryTableName)
//	stmt, err := db.Prepare(getQuery)
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to prepare get query")
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(startDate, endDate, offset, size)
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to execute get query")
//	}
//
//	var actInfo ActionInfo
//	parsedRows, err := s.ParseSQLRows(rows, &actInfo)
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to parse results")
//	}
//	if len(parsedRows) == 0 {
//		err = indexprotocol.ErrNotExist
//		return nil, err
//	}
//
//	actionInfoList := make([]*ActionInfo, 0)
//	for _, parsedRow := range parsedRows {
//		actionInfoList = append(actionInfoList, parsedRow.(*ActionInfo))
//	}
//	return actionInfoList, nil
//}
