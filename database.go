package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// --- Databases
func SetupDatabase(db_type string, url string) (*sql.DB, error) {
	// sql.Open() will not assure the connection to database, so we also need to use sql.Ping()
	db, err := sql.Open(db_type, url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to connect database !!!\n")
	}

	return db, err
}

func GetTableRecords() {
	// select sqlite_version();
	database_type := "sqlite3"
	database_url := "./sqlite3.db"
	db, err := SetupDatabase(database_type, database_url)

	defer db.Close()

	if err != nil {
		fmt.Printf("Failed to connect %s database\n", database_type)
		fmt.Println(err.Error())
	}

	// Run queries
	q := "SELECT * FROM gosqlite"
	rows, e := db.Query(q)
	if e != nil {
		fmt.Printf("Failed to execute query: [ %s ]\n", q)
		fmt.Println(e)
	}

	// fmt.Println(rows)
	fmt.Println("result ---")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("Failed to get record\n")
		}
		fmt.Printf("---> ")
		fmt.Println(id, name)
	}
}
