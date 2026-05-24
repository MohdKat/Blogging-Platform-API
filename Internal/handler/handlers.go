package handlers

import (
	"fmt"
	"net/http"
)

func CreateBlog() http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		
	}
}

// func UpdateBlog() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }

// func DeleteBlog() http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {

// 	}
// }

// func GetBlog() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }

// func GetAllBlogs() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
		
// 	}
// }