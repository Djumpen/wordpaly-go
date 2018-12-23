package api

import (
	"github.com/gin-gonic/gin"
)

func (api *API) SetSystemRoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		responseNotFound(c)
	})
}
