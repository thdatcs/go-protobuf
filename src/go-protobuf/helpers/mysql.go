package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitMySQL initializes MySQL
func InitMySQL(host string, port int, username, password, database string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
