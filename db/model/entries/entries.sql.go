package entries

import (
	"context"

	"github.com/psalishol/zuri/db"
)

type Queries struct {
	*db.Queries
}

func New(q *db.Queries) * Queries {
	return &Queries{q}
}

type Entries struct {
	db.Entries
}


type CreateEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (entry Entries, err error) {
	err = q.QueryRow(ctx, createEntryQuery, arg.AccountID, arg.Amount).Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
	)
	return;
}