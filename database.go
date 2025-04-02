package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_TYPE string = "sqlite3"
const DB_URL string = "./sqlite3.db"

func InitializeDatabase() bool {

	fmt.Println(">>> Checking SQLite3 database file exists ...")
	// os.Stat() returns nil when the file is there
	_, err := os.Stat(DB_URL)
	if err != nil {
		fmt.Println("No database file exists, will create new.")
		_, err := os.Create(DB_URL)
		if err != nil {
			fmt.Println("Failed to create file ...")
			return false
		}
	} else {
		fmt.Println("Found, will recreate with initial data.")
		err := os.Remove(DB_URL)
		if err != nil {
			fmt.Printf("Failed to clean up %s\n", DB_URL)
			return false
		}
	}
	fmt.Println("OK")
	return true
}

// --- Databases
func GetDatabaseInstance() (*sql.DB, error) {
	// sql.Open() will not assure the connection to database, so we also need to use sql.Ping()
	db, err := sql.Open(DB_TYPE, DB_URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to connect database !!!\n")
	}

	return db, err
}

func Seed() bool {
	db, err := GetDatabaseInstance()

	defer db.Close()

	if err != nil {
		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
		fmt.Println(err.Error())
		return false
	}

	// Initialize tables
	table_name := "companies"
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id PRIMARY KEY, name, founder, year);", table_name)
	_, e := db.Exec(sql)
	if e != nil {
		fmt.Printf("Failed to create table\n")
		fmt.Println(e)
		return false
	}

	// Check data duplications
	var n int
	er := db.QueryRow(fmt.Sprintf("SELECT count(*) FROM %s;", table_name)).Scan(&n)
	if er != nil {
		fmt.Println("Failed to get number of records")
		fmt.Println(er)
		return false
	}

	if n != 0 {
		fmt.Printf("Some data already exists in [ %s ] table, skipped\n", table_name)
	} else {
		// Get seed data and insert after the check
		seeds := GetCompanyList()

		q := fmt.Sprintf("INSERT INTO `%s` VALUES (?, ?, ?, ?);\n", table_name)
		for _, s := range seeds {
			_, err := db.Exec(q, s.Id, s.Name, s.Founders[0], s.Year)
			if err != nil {
				fmt.Println("Failed to exec INSERT.")
				fmt.Println(err)
			}
		}
	}

	return true
}

// func GetRecordsFromTable() {
// 	// select sqlite_version();
// 	db, err := GetDatabaseInstance()
// 	if err != nil {
// 		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
// 		fmt.Println(err.Error())
// 	}

// 	// Run queries
// 	q := "SELECT * FROM gosqlite"
// 	rows, e := db.Query(q)
// 	if e != nil {
// 		fmt.Printf("Failed to execute query: [ %s ]\n", q)
// 		fmt.Println(e)
// 	}

// 	// fmt.Println(rows)
// 	fmt.Println("result ---")
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			fmt.Printf("Failed to get record\n")
// 		}
// 		fmt.Printf("---> ")
// 		fmt.Println(id, name)
// 	}
// }
