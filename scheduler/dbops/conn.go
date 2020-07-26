package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yolomc/my-video-server/api/config"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", config.MySqlDsn)

	if err != nil {
		panic(err.Error())
	}
}
