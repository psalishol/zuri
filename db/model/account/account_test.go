package acc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/psalishol/zuri/db"
	"github.com/psalishol/zuri/util"
	"github.com/stretchr/testify/require"
)


var query Queries

func TestMain(m *testing.M) {
	tDb, err := util.SetMainTest("../../../")

	fmt.Printf("got here %v", tDb)

	if err != nil {
		log.Fatal("unable to set up main test", err)
	}

	query = Queries{db.New(tDb)}

	fmt.Printf("got here query %v", query)

	os.Exit(m.Run())
}

func createTestAccount(t require.TestingT) (account Account) {
	arg := CreateAccountQueryParams {
		OwnerName: util.RandomOwnerName(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
		DisplayPicture: sql.NullString{String: util.RandomDisplayPictureURL(), Valid: true},
	}

 	account, err := query.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.OwnerName, account.OwnerName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.DisplayPicture.String, account.DisplayPicture.String)

	return;
}

func getTestAccount(t *testing.T, id int64) (acc Account) {
	
	arg := GetAccountQueryParams {
		ID: id,
	}

 	acc, err := query.GetAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.ID, arg.ID)

	return;
}


func listTestAccount(t *testing.T) (acc []Account) {

	arg := ListAccountsQueryParams {
		Limit: 10,
		Offset: 0,
	}

    acc, err := query.ListAccounts(context.Background(), arg)

	fmt.Printf("==> Accounts: %v", acc)

	require.NoError(t, err)

	require.LessOrEqual(t, int32(len(acc)), arg.Limit)
	return;
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t);
}


func TestListAccounts(t *testing.T) {
    listTestAccount(t)
}


func TestGetAccount(t *testing.T) {
	accs := listTestAccount(t);

	if len(accs) > 0 {
		account := accs[rand.Intn(len(accs))]

		getTestAccount(t, account.ID);
	}

}

func TestUpdateAccount(t *testing.T) {

	accs := listTestAccount(t);

	if len(accs) > 0 {
		account := accs[rand.Intn(len(accs))]
	
		arg := UpdateAccountParams {
			OwnerName: util.RandomOwnerName(),
			Balance: account.Balance,
			ID: account.ID,
			DisplayPicture: sql.NullString{String: util.RandomDisplayPictureURL(), Valid: true},
			Currency: util.RandomCurrency(),
		};
	
		fmt.Printf("before update %v\narg: %v\n\n", account, arg);
	
		err :=	query.UpdateAccount(context.Background(), arg);
	
		require.NoError(t, err);
	
		updatedAccount := getTestAccount(t, account.ID);
	
		fmt.Printf("after update %v", updatedAccount);
	
		require.NotEmpty(t, updatedAccount);
	
		require.Equal(t, arg.OwnerName, updatedAccount.OwnerName);
		require.Equal(t, arg.Balance, updatedAccount.Balance);
		require.Equal(t, arg.Currency, updatedAccount.Currency);
		require.Equal(t, arg.DisplayPicture.String, updatedAccount.DisplayPicture.String);
	}

}



//TODO: Debug delete account, --> test sometimes fails
func TestDeleteAccount(t *testing.T) {

	accs := listTestAccount(t);

	if len(accs) > 0 {
		account := accs[rand.Intn(len(accs))]

		arg := DeleteAccountQueryParam {
        	ID: account.ID,
		}

	    err := query.DeleteAccount(context.Background(), arg )

		require.NoError(t, err);

		// check if it truly deletes, get account with the id, returns error if deleted.
		_ , err = query.GetAccount(context.Background(), GetAccountQueryParams{account.ID})
		require.Error(t, err)

	}

}