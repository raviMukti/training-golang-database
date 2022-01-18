package traininggolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // Dont forget to close

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES ('3', 'Hartadi')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // Dont forget to close

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close() // Dont forget to close

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString // agar mampu menghandle null value
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &created_at)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date : ", birthDate.Time)
		}
		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", created_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close() // Dont forget to close

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close() // Dont forget to close

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close() // Dont forget to close

	ctx := context.Background()

	query := "INSERT INTO user(username, password) VALUES (?, ?)"

	statement, error := db.PrepareContext(ctx, query)

	if error != nil {
		panic(error)
	}

	defer statement.Close()
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	// Do Transaction

	username := "ravi"
	password := "123"

	query := "INSERT INTO user(username, password) VALUES (?, ?)"

	statement, error := db.PrepareContext(ctx, query)

	if error != nil {
		panic(error)
	}

	defer statement.Close()

	_, errStatement := statement.ExecContext(ctx, username, password)

	if errStatement != nil {
		panic(err)
	}

	txError := tx.Commit()

	if txError != nil {
		panic(err)
	}
}
