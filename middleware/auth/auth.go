package auth

import (
	"encoding/base64"
	"strings"

	"github.com/djumpen/wordplay-go/api/vars"

	"github.com/djumpen/wordplay-go/storage"

	"github.com/gin-gonic/gin"
	"github.com/raja/argon2pw"
)

type UserGetter interface {
	UserByUsername(username string) (*storage.User, error)
}

type Responder interface {
	ResponseInternalError(c *gin.Context, err error)
	ResponseUnauthorized(c *gin.Context, err error)
	ResponseBadRequest(c *gin.Context, code int, description string, err error)
}

func BasicAuth(ug UserGetter, responder Responder) gin.HandlerFunc {
	return func(c *gin.Context) {
		if creds, ok := checkBasicAuth(c.GetHeader("Authorization")); ok {
			user, err := ug.UserByUsername(creds[0])
			if err != nil {
				responder.ResponseInternalError(c, err)
				return
			}
			if user == nil {
				responder.ResponseBadRequest(c, 101, "WRONG_CREDS", nil)
				return
			}
			valid, err := argon2pw.CompareHashWithPassword(user.Hash, creds[1])
			if err != nil || valid == false {
				responder.ResponseUnauthorized(c, err)
				return
			}
			c.Set(vars.UserParam, user)
			return
		}
		responder.ResponseUnauthorized(c, nil)
		return
	}
}

func checkBasicAuth(header string) ([]string, bool) {
	s := strings.SplitN(header, " ", 2)
	if len(s) != 2 {
		return nil, false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return nil, false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return nil, false
	}

	return pair, true
}
