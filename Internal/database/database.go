package database

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/lib/pq"
)



func Database(user_name string, db_password string, db_name string) *sql.DB{

	//Connection String
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user_name, db_password, db_name)

	//Open the RDS
	db, err := sql.Open("postgres", connStr)
	

	
	//Does db open or does it not?
	if err != nil {
		log.Fatal(err)
	}


	//Can I connect to the db?
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connection Established.")

	//return db for imports
	return db
}