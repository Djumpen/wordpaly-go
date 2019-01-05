package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/config"
	_ "github.com/djumpen/wordplay-go/doc"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/djumpen/wordplay-go/services"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-contrib/cors"
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
// Tests, db test tools
// Research logs https://github.com/Sirupsen/logrus
// Cache
// Deployment (kubernetes)
// Monitoring
// --------------------------------------------------------------

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	db := mysqldb.New(cfg.DB)

	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.Use(cors.New(config.GetCorsConfig()))

	// ------------------------ Register Resources ------------------------
	usersSvc := services.NewUsersService(storage.NewUserStorage(), db.DB)
	api.RegisterUsersResource(router, usersSvc)
	// --------------------------------------------------------------------

	addr := fmt.Sprintf(":%d", cfg.Port)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
