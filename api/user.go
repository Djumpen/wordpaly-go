package api

import (
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	"github.com/raja/argon2pw"
)

type UserCreateReq struct {
	Username string
	Password string
}

type UserResp struct {
	ID       int64
	Username string
	Email    string
	Name     string
}

func (api *API) CreateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var json UserCreateReq
		if err := c.ShouldBindJSON(&json); err != nil {
			responseErr(c, 100, "Some error", err)
			return
		}

		// Todo: detailed validation

		hash, err := argon2pw.GenerateSaltedHash(json.Password)
		if err != nil {
			responseErr(c, 100, "Some error", err)
			return
		}
		user := &storage.User{
			Username: json.Username,
			Hash:     hash,
		}
		id, err := api.storage.CreateUser(user)
		if err != nil {
			responseErr(c, 100, "Some error", err)
			return
		}
		responseCreated(c, gin.H{
			"User": gin.H{
				"ID": id,
			},
		})
	}
}

func (api *API) GetCurrentUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		user, err := extractUser(c)
		if err != nil {
			responseErr(c, 100, "Some error", err)
			return
		}
		resp := UserResp{
			ID:       user.ID,
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		}
		responseOK(c, gin.H{
			"User": resp,
		})
	}
}
