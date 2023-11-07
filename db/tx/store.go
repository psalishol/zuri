package tx

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/psalishol/zuri/db/model"
	account "github.com/psalishol/zuri/db/model/account"
)

type Queries struct {
	*account.AQueries
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore (db *sql.DB) *Store {
	return &Store {db: db}
}


func (s *Store) execTx (ctx context.Context, fn func(*Queries) error) error {
	tx, err :=	s.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(&Queries{&account.AQueries{Queries: q}})

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("err: %v\t\trollbackErr: %v", err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
} 