package database

import (
	"database/sql"
	"time"
	"github.com/MohdKat/Blogging-Platform-API.git/Internal/models"
	"fmt"
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
		return fmt.Sprintf("Error could not execute query: %s\n", err), err
	}

	return "Table created! Success!", nil
}


//Query for creating a Blog post that returns a pointer to a BlogPostResponse
//that can marshall our response into JSON object
func CreateBpost(post models.BlogPost, db *sql.DB) (*models.BlogPostResponse, error){
	
	query := `INSERT INTO BlogPosts (title, content, tags) VALUES (post.Title, post.Content,post.Tags) Returning id, createdAt, updatedAt`

	var pk int              
	var CrtdAt time.Time 
	var UpdtAt time.Time 

	err := db.QueryRow(query, post.Title, post.Content, post.Tags).Scan(&pk, &CrtdAt, &UpdtAt)
	if err != nil {
		return nil, err
	}


	return &models.BlogPostResponse{

		ID: pk,
		Title: post.Title,
		Content: post.Content,
		Tags: post.Tags,
		CreatedAt: CrtdAt,
		UpdatedAt: UpdtAt,
	}, nil
	
}

