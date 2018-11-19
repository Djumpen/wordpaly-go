package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) GetDictionaries() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userId := 100500
		dictionaries, err := api.storage.GetDictionaries(userId)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.NewEncoder(w).Encode(dictionaries)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (api *API) GetDictionary() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			// TODO: error response
		}
		fmt.Printf("Dict id: %s\n", id)

		userId := 100500
		dictionary, err := api.storage.GetDictionary(userId, 100)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.NewEncoder(w).Encode(dictionary)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
