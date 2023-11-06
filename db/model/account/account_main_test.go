package account

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/psalishol/zuri/db/model"
	"github.com/psalishol/zuri/helper"
)

var TQueries *db.Queries

func TestMain(m *testing.M) {
	conf, err := helper.ReadConfig(".")

	if err != nil {
		log.Fatal("unable to read env config ", err)
	}

   conn, err :=	sql.Open(conf.DbDriver, conf.Dsn);

   if err != nil {
	log.Fatal("error opening connection to db: ", err)
   }

   if err := conn.Ping(); err != nil {
	log.Fatal("connection to db no longer alive", err)
   }


   TQueries = db.New(conn)

   os.Exit(m.Run())
}

