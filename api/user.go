package api

import (
	"github.com/djumpen/wordplay-go/apierrors"
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

type UserResp struct {
	User userRespPart
}

type userRespPart struct {
	ID       int64
	Username string
	Email    string
	Name     string
}

// ----------------------------------

type usersSerivce interface {
	Create(*storage.User) (int64, error)
}

type usersResource struct {
	svc  usersSerivce
	resp SimpleResponder
}

func NewUsersResource(svc usersSerivce, resp SimpleResponder) *usersResource {
	return &usersResource{
		svc:  svc,
		resp: resp,
	}
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
func (r *usersResource) Create(c *gin.Context) {
	r.resp.HandleError(c, func() error {
		var json UserCreateReq
		if err := c.ShouldBindJSON(&json); err != nil {
			return err
		}
		hash, err := argon2pw.GenerateSaltedHash(json.User.Password)
		if err != nil {
			return err
		}
		user := &storage.User{
			Username: json.User.Username,
			Hash:     hash,
		}
		id, err := r.svc.Create(user)
		if err != nil {
			return err
		}

		r.resp.Created(c, UserCreatedResp{
			User: userCreated{
				ID: id,
			},
		})
		return nil
	}())
}

// swagger:operation GET /me Users GetMe
// ---
// security:
//   - basicAuth: []
// responses:
//   '200':
//     schema:
//       $ref: "#/definitions/UserResp"
func (r *usersResource) GetCurrentUser(c *gin.Context) {
	r.resp.HandleError(c, func() error {
		u, err := extractUser(c)
		if err != nil {
			return apierrors.NewUnauthorized("Can't get user")
		}

		r.resp.OK(c, UserResp{
			User: userRespPart{
				ID:       u.ID,
				Username: u.Username,
				Name:     u.Name,
				Email:    u.Email,
			},
		})
		return nil
	}())
}
