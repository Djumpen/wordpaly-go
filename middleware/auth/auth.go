package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/djumpen/wordplay-go/api"
	"github.com/djumpen/wordplay-go/storage"

	"github.com/gin-gonic/gin"
	"github.com/raja/argon2pw"
)

type UserGetter interface {
	UserByUsername(username string) (*storage.User, error)
}

func BasicAuth(ug UserGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if creds, ok := checkAuth(c.GetHeader("Authorization")); ok {
			user, err := ug.UserByUsername(creds[0])
			if err != nil {
				// TODO: breake and response error
				fmt.Println("tWO", err.Error())
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			valid, err := argon2pw.CompareHashWithPassword(user.Hash, creds[1])
			if err != nil || valid == false {
				// TODO: breake and response error
				fmt.Println("ONE", err.Error())
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Set(api.UserParam, user)
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func checkAuth(header string) ([]string, bool) {
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
