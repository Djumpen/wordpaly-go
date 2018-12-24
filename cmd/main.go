package main

import (
	"fmt"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/config"
	_ "github.com/djumpen/wordplay-go/doc"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	vp "github.com/spf13/viper"
)

// Setup docker +
// Setup mysql +
// Create middleware github.com/raja/argon2pw +
// Setup configuration (viper) +
// Documentation https://www.ribice.ba/swagger-golang/ +
// Validation
// Research migraations https://github.com/rubenv/sql-migrate
// Error engine
// JWT, oauth https://github.com/appleboy/gin-jwt
// Limits
// Cors
// Research logs https://github.com/Sirupsen/logrus
// Deployment
// Monitoring

func main() {
	var cfg config.Config
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = vp.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	db := mysqldb.New(cfg.DB)
	storage := storage.NewStorage(db)
	api := api.NewApi(&cfg, storage)

	r := gin.Default()

	api.RegisterRoutes(r)

	r.Run(":8000")
}
