package trf

import (
	"context"

	"github.com/psalishol/zuri/db"
)

type Queries struct {
	*db.Queries
}

func New(q *db.Queries) *Queries {
	return &Queries{q}
}

type Transfer struct {
	db.Transfer
}

type CreateTransferParam struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParam) (trf Transfer, err error) {
	err = q.QueryRow(ctx, createTransferQuery, arg.FromAccountID, arg.ToAccountID, arg.Amount).Scan(
		&trf.ID,
		&trf.FromAccountID,
		&trf.ToAccountID,
		&trf.Amount,
		&trf.CreatedAt,
	)
	return;
}
