package main

import (
	"database/sql"
	"fmt"
)

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute a query
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows and print the results
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
