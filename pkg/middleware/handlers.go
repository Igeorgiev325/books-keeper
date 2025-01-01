package middleware

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToPsql() {
	connStr := "postgres://ivan:785532@localhost:5432/mydb"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
}
