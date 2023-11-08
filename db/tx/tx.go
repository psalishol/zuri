package tx

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/psalishol/zuri/db"
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
	Transfer trf.Transfer 
	FromEntry entries.Entries
	ToEntry entries.Entries
	FromAccount acc.Account
	ToAccount acc.Account
}

// Performs internal transfer from one account (from_account) to another account (to_account) 
func (s *TxStore) TransferTx(ctx context.Context, arg TransferTxParams) ( result TransferTxResult, err  error) {

	err = s.execTx(ctx, func(q *Queries)  (err error) {

		txArg := trf.CreateTransferParam {FromAccountID: arg.FromAccountID, ToAccountID: arg.ToAccountID, Amount: arg.Amount}

		result.Transfer, err = q.transfer.CreateTransfer(ctx,txArg)

		if err != nil {
			return;
		}

		frmEntryArg := entries.CreateEntriesParams{ AccountID: arg.FromAccountID, Amount: -arg.Amount }

	  	result.FromEntry, err = q.entries.CreateEntries(ctx, frmEntryArg)

		if err != nil {
			return;
		}

		toEntryArg := entries.CreateEntriesParams{ AccountID: arg.ToAccountID, Amount: arg.Amount }

		result.ToEntry, err = q.entries.CreateEntries(ctx, toEntryArg)

		if err != nil {
			return;
		}

		if arg.FromAccountID < arg.ToAccountID {

			ac1ID, ac2ID := arg.FromAccountID, arg.ToAccountID

			amount := -arg.Amount

		    result.FromAccount, result.ToAccount, err =	updateTransactionAccounts(q, ctx, ac1ID,ac2ID, amount)

			if err != nil {
				return
			}

			return;
		} else {
	
			ac1ID, ac2ID := arg.ToAccountID, arg.FromAccountID

			amount := arg.Amount

		    result.ToAccount, result.FromAccount, err =	updateTransactionAccounts(q, ctx, ac1ID,ac2ID, amount)

			if err != nil {
				return
			}

			return;
		}
	})

	return;
}

func updateTransactionAccounts(q *Queries, ctx context.Context, acc1ID int64, acc2ID int64, amount int64) (ac1 acc.Account, ac2 acc.Account, err error) {

	ac1, err = q.account.UpdateAccountBalance(ctx, acc.UpdateAccountBalanceParams{
		AccountID: acc1ID, 
		Amount: amount})

	if err != nil {
		return;
	}

	ac2, err = q.account.UpdateAccountBalance(ctx, acc.UpdateAccountBalanceParams{
		AccountID: acc2ID, 
		Amount: -amount})

	if err != nil {
		return;
	}
	return;
}