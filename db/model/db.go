package db

import (
	"context"
	"database/sql"
)


type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	db DB
}

// Creates a new instance of the Queries struct, initializing it with the provided database connection (DB).
func New(db DB) *Queries {
    return &Queries{db}
}


// QueryContext executes query and returns multiple rows of results
func (q *Queries) Query (ctx context.Context, query string, arg ...any ) (*sql.Rows, error) {
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

