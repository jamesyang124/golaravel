package data

import (
	"database/sql"
	"os"

	upperDb "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper upperDb.Session

type Models struct {
	// any model insert here
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	if os.Getenv("DATABASE_TYPE") == "mysql" || os.Getenv("DATABASE_TYPE") == "mariadb" {
		upper, _ = mysql.New(db)
	} else {
		upper, _ = postgresql.New(db)
	}

	return Models{}
}
