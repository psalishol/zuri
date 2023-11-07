package trf

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/psalishol/zuri/db"
	"github.com/psalishol/zuri/util"
)



var query Queries

func TestMain(m *testing.M) {
	tDb, err := util.SetMainTest("../../../")

	if err != nil {
		log.Fatal("unable to set up main test", err)
	}

	query = Queries{db.New(tDb)}

	os.Exit(m.Run())
}

func TestCreateTransfer(t *testing.T) {
	arg := CreateTransferParam {
		FromAccountID: int64(2),
		ToAccountID: int64(10),
		Amount: int64(40),
	}

  	transfer, err := query.CreateTransfer(context.Background(), arg);
	fmt.Printf("==>Transfer: %v \n\ne==>Error: %v", transfer, err)
}
