package account

import (
	"context"
	"database/sql"

	db "github.com/psalishol/zuri/db/model"
)

type Queries struct { 
	*db.Queries
}

type Account struct {
	db.Account
}

type AccountQueryParams struct {
	OwnerName string `json:"owner_name"`
	Balance int64 `json:"balance"`
	DisplayPicture sql.NullString `json:"display_picture"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount (ctx context.Context, arg AccountQueryParams ) (i Account, err error) {
	err = q.QueryRow(ctx, createAccountQuery, 
		arg.OwnerName, 
		arg.Balance,
		arg.DisplayPicture,
		arg.Currency).Scan(
			&i.ID, 
			&i.OwnerName, 
			&i.Balance, 
			&i.DisplayPicture, 
			&i.Currency, 
			&i.CreatedAt,
		);

	return;
}



type UpdateAccountParams struct {
	ID int64 `json:"id"`
	OwnerName string `json:"owner_name"`
	Balance int64 `json:"balance"`
	DisplayPicture sql.NullString `json:"display_picture"`
	Currency string `json:"currency"`
}


func (q *Queries) UpdateAccount (ctx context.Context, arg UpdateAccountParams) (err error) {

	_ , err = q.Exec(ctx, updateAccountQuery, 
		arg.ID, 
		arg.OwnerName, 
		arg.Balance, 
		arg.DisplayPicture, 
		arg.Currency)

	return;

}


type GetAccountQueryParams struct {
	ID int64 `json:"id"`
}

func (q *Queries) GetAccount (ctx context.Context, arg GetAccountQueryParams ) (i Account, err error) {
  err = q.QueryRow(ctx, getAccountQuery, arg.ID).Scan(	
	&i.ID, 
	&i.OwnerName, 
	&i.Balance, 
	&i.DisplayPicture, 
	&i.Currency, 
	&i.CreatedAt,
  );

  return;
}


type ListAccountsQueryParams struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts (ctx context.Context, arg ListAccountsQueryParams) ([]Account, error) {

	var accounts []Account

    rows :=	q.Query(ctx, listAccountsQuery, arg.Limit, arg.Offset)

	defer rows.Close()

	for rows.Next() {
		var i Account;

		if  err := rows.Scan(
			&i.ID, 
			&i.OwnerName, 
			&i.Balance, 
			&i.DisplayPicture, 
			&i.Currency, 
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}

	 accounts = append(accounts, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

type DeleteAccountQueryParam struct {
	ID  int64  `json:"id"`
}

func (q *Queries) DeleteAccount (ctx context.Context, arg DeleteAccountQueryParam) (err error) {
    _, err = q.Exec(ctx, deleteAccountQuery, arg.ID);
	return;
}