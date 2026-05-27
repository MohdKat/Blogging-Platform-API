package main

import (
	"net/http"
	"log"

	"os"
	"github.com/joho/godotenv"
	"github.com/MohdKat/Blogging-Platform-API.git/Internal/database"
	"github.com/MohdKat/Blogging-Platform-API.git/Internal/handler"
)

func main() {

	//Loaded the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}


	//Extracted the db credentials
	user_name := os.Getenv("USER_NAME")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("POSTGRES_PASSWORD")

	db := database.Database(user_name, db_password, db_name)
	defer db.Close()

	mux := http.NewServeMux()
	database.CreateTable(db)

	mux.HandleFunc("POST /blogs", handler.CreateBlog(db))
	mux.HandleFunc("PUT /blogs/{id}", handler.UpdateBlog(db))
	mux.HandleFunc("DELETE /blogs/{id}", handler.DeleteBlog(db))
	mux.HandleFunc("GET /blogs/{id}", handler.GetBlog(db))
	mux.HandleFunc("GET /blogs", handler.GetAllBlogs(db))

	log.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}