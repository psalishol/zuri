package entries

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

	fmt.Printf("got here %v", tDb)

	if err != nil {
		log.Fatal("unable to set up main test", err)
	}

	query = Queries{db.New(tDb)}

	fmt.Printf("got here query %v", query)

	os.Exit(m.Run())
}

func TestCreateEntries(t *testing.T) {
	fmt.Println("got here now====>>>>>>>>")
	arg := CreateEntriesParams {
		AccountID: int64(10),
		Amount: int64(40),
	}

  	entry, err := query.CreateEntries(context.Background(), arg);
	fmt.Printf("==>Entry: %v \n\ne==>Error: %v", entry, err)
}