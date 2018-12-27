package main

import (
	"fmt"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/config"
	_ "github.com/djumpen/wordplay-go/doc"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------
// TODO Roadmap:
// Setup docker +
// Setup mysql +
// Create middleware github.com/raja/argon2pw +
// Setup configuration (viper) +
// Documentation https://www.ribice.ba/swagger-golang/ +
// Research migrations https://github.com/rubenv/sql-migrate +
// Validation
// Error engine
// JWT, oauth https://github.com/appleboy/gin-jwt
// Access levels
// Limits
// Cors
// Research logs https://github.com/Sirupsen/logrus
// Deployment (kubernetes)
// Monitoring
// --------------------------------------------------------------

func main() {
	cfg := config.ReadConfig()

	if cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	db := mysqldb.New(cfg.DB)
	storage := storage.NewStorage(db)
	api := api.NewApi(cfg, storage)

	r := gin.Default()

	api.RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
