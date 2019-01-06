package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/config"
	_ "github.com/djumpen/wordplay-go/doc"
	"github.com/djumpen/wordplay-go/middleware/auth"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/djumpen/wordplay-go/services"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	db := mysqldb.New(cfg.DB)
	responder := api.NewResponder()

	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.Use(cors.New(config.GetCorsConfig()))

	usersSvc := services.NewUsersService(storage.NewUserStorage(), db.DB)
	userResource := api.NewUsersResource(usersSvc, responder)

	commonResource := api.NewCommonResource(responder)

	authMiddleware := auth.BasicAuth(usersSvc, responder)
	authR := r.Group("/").Use(authMiddleware)

	// ---------- Routes -------------------
	r.POST("/users", userResource.Create)
	authR.GET("/me", userResource.GetCurrentUser)

	r.GET("/sanity", commonResource.Sanity)
	r.NoRoute(commonResource.NotFound)
	r.Static("/swaggerui/", "doc/swaggerui")
	// --------------------------------------------

	addr := fmt.Sprintf(":%d", cfg.Port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
