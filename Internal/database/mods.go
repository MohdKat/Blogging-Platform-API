package database

import (
	"database/sql"
	"log"
	"time"
)

//Simpler not optimal design,
//I can create a table for tags using the same ID and mapping it to many tags on different rows
//Prioritizing handler logic and http understanding
func CreateTable(db *sql.DB) (string, error){

	query := `CREATE TABLE IF NOT EXISTS BlogPosts (
			id SERIAL PRIMARY KEY,
			title VARCHAR NOT NUll,
			content VARCHAR NOT NULL,
			tags TEXT[],
			createdAt timestamp DEFAULT NOW(),
			updatedAt timestamp DEFAULT NOW()		
	)`

	_, err := db.Exec(query);if err != nil {
		log.Fatal(err)
	}
	Id = 1
}


//Query for creating a Blog post
func CreateBpost(post BlogPost, db *sql.DB) {
	
	query := `INSERT INTO BlogPosts (title, content, tags) VALUES (post.Title, post.Content,post.Tags) Returning id, createdAt, updatedAt`

	var pk int              
	var CreatedAT time.Time 
	var UpdatedAt time.Time 

	err := db.QueryRow(query, post.Title, post.Content, post.Tags).Scan(&pk, &CreatedAT, &UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	PostResponse:= BlogPostResponse{}
}

