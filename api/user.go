package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	"github.com/raja/argon2pw"
)

type UserCreateReq struct {
	Username string
	Password string
}

type UserLoginReq struct {
	Username string
	Password string
}

func (api *API) CreateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Content-Type", "application/json")
		var json UserCreateReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Todo: detailed validation

		hash, err := argon2pw.GenerateSaltedHash(json.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := &storage.User{
			Username: json.Username,
			Hash:     hash,
		}
		id, err := api.storage.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "Created with id " + strconv.Itoa(int(id))})
	}
}

func (api *API) GetCurrentUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		// var json UserLoginReq

		user, err := currUser(c)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"resp": "Current userID",
			"User": user,
		})
	}
}
