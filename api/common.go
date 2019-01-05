package api

import (
	"github.com/gin-gonic/gin"
)

type commonResource struct {
	resp *Responder
}

func RegisterCommonResource(r *gin.Engine) {
	resource := commonResource{
		resp: &Responder{},
	}

	r.GET("/sanity", resource.Sanity)
	r.Static("/swaggerui/", "doc/swaggerui")
	r.NoRoute(resource.NotFound)
}

func (r *commonResource) Sanity(c *gin.Context) {
	r.resp.ResponseOK(c, nil)
}

func (r *commonResource) NotFound(c *gin.Context) {
	r.resp.ResponseNotFound(c, "Route not found")
}
