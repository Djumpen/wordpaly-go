package api

import (
	"github.com/djumpen/wordplay-go/middleware/auth"
	"github.com/gin-gonic/gin"
)

func (api *API) RegisterRoutes(r *gin.Engine) {
	authorized := r.Group("/")
	authorized.Use(auth.BasicAuth(api.storage, api))
	{
		authorized.GET("/me", api.GetCurrentUser())

		authorized.GET("/dictionaries", api.GetDictionaries())
		authorized.GET("/dictionaries/{id:[0-9]+}", api.GetDictionary())
	}

	r.POST("/users", api.CreateUser())

	r.Static("/swaggerui/", "doc/swaggerui")

	r.NoRoute(func(c *gin.Context) {
		api.ResponseNotFound(c)
	})
}
