package traininggolangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/training_db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println(db.Ping())

}
