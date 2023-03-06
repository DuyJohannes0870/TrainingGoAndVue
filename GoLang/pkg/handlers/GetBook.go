package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/training/go/crud/pkg/mocks"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	//Read to request body
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Iterate over all the mock books
	for _, book := range mocks.Books {
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Context-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
	//Send a 201 created response
}
