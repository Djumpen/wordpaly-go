package api

import (
	"github.com/djumpen/wordplay-go/cfg"
	"github.com/djumpen/wordplay-go/storage"
	jsoniter "github.com/json-iterator/go"
)

type API struct {
	version string
	storage *storage.Storage
	config  *cfg.Config
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewApi(config *cfg.Config, storage *storage.Storage) *API {
	return &API{
		version: "1.0",
		storage: storage,
		config:  config,
	}
}
