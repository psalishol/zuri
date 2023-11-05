package account

import (
	"context"
	"fmt"
	"testing"

	util "github.com/psalishol/zuri/helper"
	"github.com/stretchr/testify/require"
)


func TestCreateAccount(t *testing.T) {
	arg := CreateAccountQueryParams {
		OwnerName: util.RandomOwnerName(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	fmt.Printf("create account params: %#v\n", arg)

	query := Queries{TQueries}

 	account, err :=	query.CreateAccount(context.Background(), arg)

	fmt.Printf("created account: %#v\n", account)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.OwnerName, account.OwnerName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.DisplayPicture, account.DisplayPicture)
}