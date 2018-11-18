package api

import (
	"github.com/djumpen/wordplay-go/cfg"
	jsoniter "github.com/json-iterator/go"
)

type API struct {
	version string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewApi(config *cfg.Config) *API {
	return &API{
		version: "1.0",
	}
}
