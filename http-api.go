package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string = []string{}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/test", test)
	router.HandleFunc("/add/{item}", addItem)

	http.ListenAndServe(":5000", router)

}

func test(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		ID string
	}{"777"})

}

func addItem(w http.ResponseWriter, r *http.Request) {

	item := mux.Vars(r)["item"]
	data = append(data, item)

}
