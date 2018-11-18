package main

import (
	"log"
	"net/http"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/cfg"
	"github.com/djumpen/wordplay-go/middleware/auth"
	"github.com/gorilla/mux"
)

func main() {
	configFile := "config.json"
	config, err := cfg.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	api := api.NewApi(config)

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/dictionary", api.GetDictionaries()).Methods("GET")
	r.Use(auth.Middleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
