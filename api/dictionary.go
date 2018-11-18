package api

import (
	"log"
	"net/http"
)

type Dictionary struct {
	ID          string
	Title       string
	Description string
}

var dummyRes []*Dictionary = []*Dictionary{
	&Dictionary{
		"100",
		"MyDict1",
		"Funny words",
	},
	&Dictionary{
		"101",
		"MyDict2",
		"Other words",
	},
}

func (api *API) GetDictionaries() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(dummyRes)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
