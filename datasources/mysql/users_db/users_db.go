package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_user_name = "mysql_users_user_name"
	mysql_users_password  = "mysql_users_password"
	mysql_user_host       = "mysql_user_host"
	mysql_users_schema    = "mysql_users_schema"
)

var (
	Client   *sql.DB
	userName = os.Getenv(mysql_users_user_name)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_user_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		userName, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database configure successfully")
}
