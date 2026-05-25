package structs

import (
	"time"
)

//Blog Post structure for client input adding posts
type BlogPost struct {

	Title string    
	Content string		  
	Tags []string         
}

//Blogpostresponse this should contain the json with the id and the timestamps of creation and updates
//json tags for marshalling and unmarshalling
type BlogPostResponse struct {
	
	ID int     `json:"title"`
	Title string
	Content string
	Tags []string
	CreatedAt time.Time
	UpdatedAt time.Time
}