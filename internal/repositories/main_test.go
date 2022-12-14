package repositories

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	util "github.com/vtv-us/kahoot-backend/internal/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("can't load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
