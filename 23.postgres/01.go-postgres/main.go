package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping() // Check if the connection is responding
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

// go get github.com/lib/pq
