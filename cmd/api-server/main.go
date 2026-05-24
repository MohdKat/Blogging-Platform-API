package main

import (
	"fmt"
	"net/http"
	"time"
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
	
}