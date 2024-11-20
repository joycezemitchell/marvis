package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Client *sql.DB

	dbUser     = os.Getenv("MYSQL_USERNAME")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbName     = os.Getenv("MYSQL_DBNAME")

	/*dbUser     = "root"
	dbPassword = "root"
	dbHost     = "mysql"
	dbPort     = 3306
	dbName     = "allychat"*/
)

func NewConnection() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		dbUser, dbPassword, dbHost, dbName,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return Client, err
	}
	if err = Client.Ping(); err != nil {
		return Client, err
	}

	return Client, nil
}
