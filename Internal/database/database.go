package database

import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func Database() *sql.DB{


	//Loaded the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}


	//Extracted the db credentials
	user_name := os.Getenv("USER_NAME")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("POSTGRES_PASSWORD")


	//Connection String
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user_name, db_password, db_name)

	//Open the RDS
	db, err := sql.Open("postgres", connStr)
	

	//Closed the DB after the program ends (Best Practice)
	defer db.Close()
	
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