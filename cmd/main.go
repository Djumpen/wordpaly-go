package main

import (
	"log"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/cfg"
	"github.com/djumpen/wordplay-go/middleware/auth"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
)

// Setup docker +
// Setup mysql +
// Create middleware github.com/raja/argon2pw
// Setup configuration (viper)
// Validation (swagger)
// Research migraations https://github.com/rubenv/sql-migrate
// Error engine
// Research logs https://github.com/Sirupsen/logrus
// Monitoring
// Deployment

func main() {
	configFile := "config.json"
	config, err := cfg.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	db := mysqldb.New("wpuser", "pass", "mysqldb", "3306", "wordplay")

	storage := storage.NewStorage(db)

	api := api.NewApi(config, storage)

	r := gin.Default()

	authorized := r.Group("/")
	authorized.Use(auth.Middleware(storage))
	{
		authorized.GET("/users/me", api.GetCurrentUser())

		authorized.GET("/dictionaries", api.GetDictionaries())
		authorized.GET("/dictionary/{id:[0-9]+}", api.GetDictionary())
	}

	r.POST("/users", api.CreateUser())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run(":8000")
}
