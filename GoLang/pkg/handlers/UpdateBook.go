package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/training/go/crud/pkg/mocks"
	"github.com/training/go/crud/pkg/models"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	// Iterate over all the mock Books
	for index, book := range mocks.Books {
		if book.Id == id {
			// Update and send response when book Id matches dynamic Id
			book.Title = updateBook.Title
			book.Author = updateBook.Author
			book.Desc = updateBook.Desc

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Context-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
		}
	}
}
