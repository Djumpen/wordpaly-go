package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/djumpen/wordplay-go/cfg"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type API struct {
	version string
	storage *storage.Storage
	config  *cfg.Config
}

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	UserParam = "currUser"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewApi(config *cfg.Config, storage *storage.Storage) *API {
	return &API{
		version: "1.0",
		storage: storage,
		config:  config,
	}
}

func extractUser(c *gin.Context) (*storage.User, error) {
	userIface, ok := c.Get(UserParam)
	if ok != true {
		return nil, errors.New("Unauthorized")
	}
	user, ok := userIface.(*storage.User)
	if ok != true {
		return nil, errors.New("Unauthorized")
	}
	return user, nil
}

func responseOK(c *gin.Context, res gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"data":  res,
		"error": nil,
	})
}

func responseCreated(c *gin.Context, res gin.H) {
	c.JSON(http.StatusCreated, gin.H{
		"data":  res,
		"error": nil,
	})
}

func responseErr(c *gin.Context, code int, description string, err error) {
	errResp := gin.H{
		"code":        code,
		"description": description,
	}
	if gin.Mode() == gin.DebugMode {
		withDebug(errResp, err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"data":  nil,
		"error": errResp,
	})
}

func responseNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"data": nil,
		"error": gin.H{
			"code":        404,
			"description": "NOT_FOUND",
		},
	})
}

func withDebug(errResp gin.H, err error) {
	debugResp := gin.H{
		"details": fmt.Sprintf("%s", err),
	}
	errSt, ok := err.(stackTracer)
	if ok {
		trace := make(gin.H)
		st := errSt.StackTrace()
		for i, v := range st {
			// if i > 3 {
			// 	break
			// }
			trace[strconv.Itoa(i)] = fmt.Sprintf("%+v", v)
		}
		debugResp["trace"] = trace
	}
	errResp["debug"] = debugResp
}

// func ResponseErrWithFields(c *gin.Context)
