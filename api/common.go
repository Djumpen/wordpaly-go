package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type commonResource struct {
	resp *Responder
}

func NewCommonResource(resp *Responder) *commonResource {
	return &commonResource{
		resp: resp,
	}
}

func (r *commonResource) Sanity(c *gin.Context) {
	r.resp.OK(c, nil)
}

func (r *commonResource) NotFound(c *gin.Context) {
	r.resp.NotFound(c, errors.New("Resource not found"))
}
