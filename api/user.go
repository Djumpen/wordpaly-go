package api

import (
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	"github.com/raja/argon2pw"
)

type UserCreateReq struct {
	User userCreateReq
}
type userCreateReq struct {
	Username string `binding:"required,min=3"`
	Password string `binding:"required,min=4"`
}

type UserCreatedResp struct {
	User userCreated
}
type userCreated struct {
	ID int64
}

// swagger:operation POST /users Users CreateUser
// ---
// parameters:
//   - name: body
//     in: body
//     schema:
//       $ref: '#/definitions/UserCreateReq'
// responses:
//   '201':
//     schema:
//       $ref: "#/definitions/UserCreatedResp"
func (api *API) CreateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var json UserCreateReq
		if err := c.ShouldBindJSON(&json); err != nil {
			if vf, ok := getValidationFailers(err); ok {
				api.ResponseErrWithFields(c, vf)
				return
			}
			api.ResponseInternalError(c, err)
			return
		}
		hash, err := argon2pw.GenerateSaltedHash(json.User.Password)
		if err != nil {
			api.ResponseInternalError(c, err)
			return
		}
		user := &storage.User{
			Username: json.User.Username,
			Hash:     hash,
		}
		id, err := api.storage.CreateUser(user)
		if err != nil {
			if storage.CheckError(err, storage.ErrDuplicate) {
				api.ResponseConflict(c, err)
				return
			}
			api.ResponseUnpocessable(c, "TODO: Some error", err)
			return
		}

		api.ResponseCreated(c, UserCreatedResp{
			User: userCreated{
				ID: id,
			},
		})
	}
}

type UserResp struct {
	User user
}

type user struct {
	ID       int64
	Username string
	Email    string
	Name     string
}

// swagger:operation GET /me Users GetMe
// ---
// security:
//   - basicAuth: []
// responses:
//   '200':
//     schema:
//       $ref: "#/definitions/UserResp"
func (api *API) GetCurrentUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		u, err := extractUser(c)
		if err != nil {
			api.ResponseUnauthorized(c, err)
			return
		}

		api.ResponseOK(c, UserResp{
			User: user{
				ID:       u.ID,
				Username: u.Username,
				Name:     u.Name,
				Email:    u.Email,
			},
		})
	}
}
