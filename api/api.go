package api

import (
	"errors"

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

func NewApi(config *cfg.Config, storage *storage.Storage) *API {
	return &API{
		version: "1.0",
		storage: storage,
		config:  config,
	}
}

func currUser(c *gin.Context) (*storage.User, error) {
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
