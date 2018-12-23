package main

import (
	"fmt"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/cfg"
	_ "github.com/djumpen/wordplay-go/doc"
	"github.com/djumpen/wordplay-go/middleware/auth"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	vp "github.com/spf13/viper"
)

// Setup docker +
// Setup mysql +
// Create middleware github.com/raja/argon2pw +
// Setup configuration (viper) +
// Validation (swagger)
//		-https://www.ribice.ba/swagger-golang/
// Research migraations https://github.com/rubenv/sql-migrate
// Error engine
// JWT, oauth
// Cors
// Research logs https://github.com/Sirupsen/logrus
// Monitoring
// Deployment

func main() {
	var config cfg.Config
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	a := 1 + 100

	fmt.Print(a)

	db := mysqldb.New(config.DB)
	storage := storage.NewStorage(db)
	api := api.NewApi(&config, storage)

	r := gin.Default()

	authorized := r.Group("/")
	authorized.Use(auth.BasicAuth(storage))
	{
		authorized.GET("/me", api.GetCurrentUser())

		authorized.GET("/dictionaries", api.GetDictionaries())
		authorized.GET("/dictionaries/{id:[0-9]+}", api.GetDictionary())
	}

	r.POST("/users", api.CreateUser())

	api.SetSystemRoutes(r)

	r.Run(":8000")
}
