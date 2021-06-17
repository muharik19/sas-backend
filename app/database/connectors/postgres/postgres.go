package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// ConnectDb ...
func ConnectDb() (*sql.DB, error) {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conString)
	if err != nil {
		fmt.Println("Failed to connect database!", err)
	}

	testPing := db.Ping()
	if testPing != nil {
		log.Fatal("Error pinging database: " + testPing.Error())
	}

	return db, err
}

// Query func
func Query(sql string) (*sql.Rows, error) {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	return db.QueryContext(ctx, sql)
}

// QueryRow func
func QueryRow(sql string) *sql.Row {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	return db.QueryRowContext(ctx, sql)
}

// Exec func
func Exec(sql string) bool {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	_, err := db.ExecContext(ctx, sql)
	if err != nil {
		return false
	}
	return true
}

// QueryCount func
func QueryCount(queryCount string) int {

	var totalRows int
	rowsCount, _ := Query(queryCount)
	for rowsCount.Next() {
		errScanCount := rowsCount.Scan(&totalRows)
		if errScanCount != nil {
			return 0
		}
	}

	return totalRows
}
