package auth

import (
	"encoding/base64"
	"strings"

	"github.com/djumpen/wordplay-go/api/vars"
	"github.com/djumpen/wordplay-go/apierrors"

	"github.com/djumpen/wordplay-go/storage"

	"github.com/gin-gonic/gin"
	"github.com/raja/argon2pw"
)

type UserGetter interface {
	ByUsername(username string) (*storage.User, error)
}

type ErrorHandler interface {
	HandleError(*gin.Context, error)
}

func BasicAuth(ug UserGetter, handler ErrorHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.HandleError(c, func() error {
			errUnauthorized := &apierrors.Unauthorized{}
			if creds, ok := checkBasicAuth(c.GetHeader("Authorization")); ok {
				user, err := ug.ByUsername(creds[0])
				if err != nil {
					return err
				}
				if user == nil {
					return errUnauthorized
				}
				valid, err := argon2pw.CompareHashWithPassword(user.Hash, creds[1])
				if err != nil || valid == false {
					return errUnauthorized
				}
				c.Set(vars.UserParam, user)
				return nil
			}
			return errUnauthorized
		}())
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
