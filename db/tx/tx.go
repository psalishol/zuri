package tx

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/psalishol/zuri/db/model"
	acc "github.com/psalishol/zuri/db/model/account"
	"github.com/psalishol/zuri/db/model/entries"
	trf "github.com/psalishol/zuri/db/model/transfer"
)

type Queries struct {
  account	*acc.Queries
  transfer	*trf.Queries
  entries   *entries.Queries
}

type TxStore struct {
	*Queries
	db *sql.DB
}

func NewTxStore (db *sql.DB) *TxStore {
	return &TxStore {db: db}
}


func (s *TxStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err :=	s.db.BeginTx(ctx, nil) // TODO: implement transaction isolation level.

	if err != nil {
		return err
	}

	q := db.New(tx)

	queries := Queries{ account: acc.New(q), transfer: trf.New(q), entries: entries.New(q) }

	err = fn(&queries)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("err: %v\t\trollbackErr: %v", err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
} 

type TransferTxParams struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID int64	`json:"to_account_id"`
	Amount int64  `json:"amount"`
}

type TransferTxResult struct {

}

// Performs internal transfer from one account (from_account) to another account (to_account) 
func (s *TxStore) TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, err error) {

	err = s.execTx(ctx, func(q *Queries)  error {

		txArg := trf.CreateTransferParam {FromAccountID: arg.FromAccountID, ToAccountID: arg.ToAccountID, Amount: arg.Amount}

		transfer, err := q.transfer.CreateTransfer(ctx,txArg)

		if err != nil {
			return err
		}

		frmEntryArg := entries.CreateEntriesParams{ AccountID: arg.FromAccountID, Amount: -arg.Amount }

	  	frmEntry, err := q.entries.CreateEntries(ctx, frmEntryArg)

		if err != nil {
			return err;
		}

		toEntryArg := entries.CreateEntriesParams{ AccountID: arg.ToAccountID, Amount: arg.Amount }

		toEntry, err := q.entries.CreateEntries(ctx, toEntryArg)

		if err != nil {
			return err;
		}

		// update balance;
		
		return nil
	})

	return
}