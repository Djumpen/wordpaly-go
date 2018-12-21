package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type DictionaryResp struct {
	ID          int
	Title       string
	Description string
}

func (api *API) GetDictionaries() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		userId := 100500
		dictionaries, err := api.storage.GetDictionaries(userId)
		if err != nil {
			log.Fatalln(err)
		}

		err = json2.NewEncoder(c.Writer).Encode(dictionaries)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (api *API) GetDictionary() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(c.Request)
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

		err = json2.NewEncoder(c.Writer).Encode(dictionary)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
