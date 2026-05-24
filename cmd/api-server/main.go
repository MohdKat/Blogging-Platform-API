package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Created Blog struct and Json tags for marshalling and unmarshalling data
type Blog struct {

	ID string             `json:"id"`
	Title string          `json:"title"`
	Content string		  `json:"content"`
	Tags []string         `json:"tags"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	user_name := os.Getenv("USER_NAME")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("POSTGRES_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user_name, db_password, db_name)

	db, err := sql.Open("postgres", connStr)
	
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}


}