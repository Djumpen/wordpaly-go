package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/cfg"
	"github.com/djumpen/wordplay-go/middleware/auth"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	configFile := "config.json"
	config, err := cfg.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	storage := storage.NewStorage(&sqlx.DB{})

	api := api.NewApi(config, storage)

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/dictionaries", api.GetDictionaries()).Methods("GET")
	r.HandleFunc("/dictionary/{id:[0-9]+}", api.GetDictionary()).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Not found")
	})

	r.Use(auth.Middleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
