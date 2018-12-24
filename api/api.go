package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/djumpen/wordplay-go/api/vars"
	"github.com/djumpen/wordplay-go/config"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type API struct {
	version string
	storage *storage.Storage
	config  *config.Config
}

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewApi(cfg *config.Config, storage *storage.Storage) *API {
	return &API{
		storage: storage,
		config:  cfg,
	}
}

func extractUser(c *gin.Context) (*storage.User, error) {
	userIface, ok := c.Get(vars.UserParam)
	if ok != true {
		return nil, errors.New("Unauthorized")
	}
	user, ok := userIface.(*storage.User)
	if ok != true {
		return nil, errors.New("Unauthorized")
	}
	return user, nil
}

type Response struct {
	Data  interface{}      `json:"data"`
	Error *errResponsePart `json:"error"`
}

type errResponsePart struct {
	Code        int                          `json:"code"`
	Description string                       `json:"description"`
	Fileds      map[string]fieldResponsePart `json:"fields,omitempty"`
	Debug       debugResponsePart            `json:"debug,omitempty"`
}

type fieldResponsePart struct {
	Rule      string `json:"rule"`
	RuleValue string `json:"ruleValue"`
}

type debugResponsePart struct {
	Details string `json:"details"`
	Trace   gin.H  `json:"trace,omimtempty"`
}

func (api *API) ResponseOK(c *gin.Context, res interface{}) {
	response(c, http.StatusOK, res)
}

func (api *API) ResponseCreated(c *gin.Context, res interface{}) {
	response(c, http.StatusCreated, res)
}

func (api *API) ResponseBadRequest(c *gin.Context, code int, description string, err error) {
	reponseErr(c, http.StatusBadRequest, code, description, err, nil)
}

func (api *API) ResponseUnauthorized(c *gin.Context, err error) {
	reponseErr(c, http.StatusUnauthorized, 401, "UNAUTHORIZED", err, nil)
}

func (api *API) ResponseForbidden(c *gin.Context) {
	reponseErr(c, http.StatusForbidden, 403, "FORBIDDEN", nil, nil)
}

func (api *API) ResponseNotFound(c *gin.Context) {
	reponseErr(c, http.StatusNotFound, 404, "NOT_FOUND", nil, nil)
}

func (api *API) ResponseConflict(c *gin.Context, err error) {
	reponseErr(c, http.StatusConflict, 409, "CONFLICT", err, nil)
}

func (api *API) ResponseUnpocessable(c *gin.Context, description string, err error) {
	reponseErr(c, http.StatusUnprocessableEntity, 422, description, err, nil)
}

func (api *API) ResponseInternalError(c *gin.Context, err error) {
	reponseErr(c, http.StatusInternalServerError, 500, "SERVER_ERROR", err, nil)
}

func (api *API) ResponseErrWithFields(c *gin.Context, vf ValidationFailers) {
	fields := make(map[string]fieldResponsePart)
	for _, f := range vf {
		fields[f.NameSpace] = fieldResponsePart{
			Rule:      f.Rule,
			RuleValue: f.RuleValue,
		}
	}
	reponseErr(c, http.StatusUnprocessableEntity, 422, "VALIDATION", nil, fields)
}

func response(c *gin.Context, httpCode int, res interface{}) {
	c.JSON(httpCode, Response{
		Data: res,
	})
	c.Abort()
}

func reponseErr(c *gin.Context, httpCode, code int, description string, err error, fields map[string]fieldResponsePart) {
	errResp := &errResponsePart{
		Code:        code,
		Description: description,
		Fileds:      fields,
	}
	if gin.Mode() == gin.DebugMode {
		withDebug(errResp, err)
	}
	c.JSON(http.StatusBadRequest, Response{
		Error: errResp,
	})
	c.Abort()
}

func withDebug(errResp *errResponsePart, err error) {
	if err == nil {
		return
	}
	debugResp := debugResponsePart{
		Details: fmt.Sprintf("%s", err),
	}
	errSt, ok := err.(stackTracer)
	if ok {
		trace := make(gin.H)
		st := errSt.StackTrace()
		for i, v := range st {
			if i > 3 {
				break
			}
			trace[strconv.Itoa(i)] = fmt.Sprintf("%+v", v)
		}
		debugResp.Trace = trace
	}
	errResp.Debug = debugResp
}
