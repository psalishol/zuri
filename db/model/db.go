package db

import (
	"context"
	"database/sql"
)


type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) *sql.Rows
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	db DB
}


// QueryContext executes query and returns multiple rows of results
func (q *Queries) Query (ctx context.Context, query string, arg ...any ) *sql.Rows {
	return q.db.QueryContext(ctx, query, arg...)
}

// QueryRowContext executes query and return a single row of result
func (q *Queries) QueryRow (ctx context.Context, query string, arg ...any ) *sql.Row {
	return q.db.QueryRowContext(ctx, query, arg...)
}

// ExecRowContext executes query and doesnt return any row as result. nil.
func (q *Queries) Exec (ctx context.Context, query string, arg ...any ) (sql.Result, error ){
	return q.db.ExecContext(ctx, query, arg...)
}

