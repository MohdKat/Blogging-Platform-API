package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/MohdKat/Blogging-Platform-API.git/Internal/database"
	"github.com/MohdKat/Blogging-Platform-API.git/Internal/models"
)



func CreateBlog(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		var post models.BlogPost
		//Parse http request into json byte slice
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Invalid request body is not valid", http.StatusBadRequest)
			return
		}

		//Unmarshal json byte slice into BlogPost
		err = json.Unmarshal(body, &post)
 		if err != nil {
			http.Error(w, "Could not unmarshal request", http.StatusInternalServerError)
			return
		}

		//Insert a row of the BlogPost and return BlogPostResponse struct instance
		response, err := database.CreateBpost(post, db)
		if err != nil {
			log.Printf("create post failed: %v\n", err)
			http.Error(w,"could not create post" , http.StatusInternalServerError)
			
		}
		
		//Marshall BlogPostResponse into json byte slice
		jsn_response, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "could not marshal response", http.StatusInternalServerError)
			return
		}
		
		//writing the http response
		w.Header().Set("Content_type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(jsn_response)
		if err != nil {
			log.Printf("Failed to return response: %v\n", err)
		}

	}
}

func UpdateBlog(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Getting the id and converting it into a int
		str_id := r.PathValue("id")
		id, err := strconv.Atoi(str_id)
		if err != nil {
			http.Error(w, "Error: Not a calid ID !", http.StatusBadRequest)
			return
		}


		//Unmarshalling the http request to pass it to the query
		var post models.BlogPost
		err = json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Error could not unmarshal request!", http.StatusInternalServerError)
			return 
		}


		//Updating Blog post in database and fetching the BlogPostresponse struct instance
		updated_blog, err := database.UpdateBpost(id, post, db)
		if err != nil {
			http.Error(w, "Could not update blog post!", http.StatusInternalServerError)
			return
		}



		//writing the http response
		w.Header().Set("Content_type", "application/json")
		w.WriteHeader(http.StatusContinue)

		//encoding the updated blogResponse entity into and writing it into the http response
		err = json.NewEncoder(w).Encode(updated_blog)
		if err != nil {
			http.Error(w, "Coud not marshal the response!", http.StatusInternalServerError)
		}

	}
}

func DeleteBlog(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//Getting the id for the DeleteBpost function(Query/Db operation)
		str_id := req.PathValue("id")
		id, err := strconv.Atoi(str_id)
		if err != nil {
			http.Error(w, "Not a valid ID!", http.StatusBadRequest)
		}

		err = database.DeleteBpost(id, db)
		if err != nil {
			http.Error(w, "Could not find the id! ", http.StatusBadRequest)
		}


		//writing the http response
		w.Header().Set("Content_type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}

// func GetBlog() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }

// func GetAllBlogs() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
		
// 	}
// }