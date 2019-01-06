package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/djumpen/wordplay-go/apierrors"
	validator "gopkg.in/go-playground/validator.v8"

	"github.com/djumpen/wordplay-go/api/vars"
	"github.com/djumpen/wordplay-go/storage"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

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

type objJson map[string]interface{}

func (j objJson) extractStringValue(key string) (string, error) {
	valIface, ok := j[key]
	if ok == false {
		// TODO: return validation error
		return "", errors.Errorf("No key %s", key)
	}
	val, ok := valIface.(string)
	if ok == false {
		// TODO: return validation error
		return "", errors.Errorf("%s is not a string", key)
	}
	return val, nil
}

type SimpleResponder interface {
	OK(c *gin.Context, res interface{})
	Created(c *gin.Context, res interface{})
	NotFound(c *gin.Context, err error)
	HandleError(c *gin.Context, err error)
}

type Response struct {
	Success bool             `json:"success"`
	Data    interface{}      `json:"data"`
	Error   *errResponsePart `json:"error"`
}

type errResponsePart struct {
	Code        int                          `json:"code"`
	Description string                       `json:"description"`
	Fields      map[string]fieldResponsePart `json:"fields,omitempty"`
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

type Responder struct{}

func NewResponder() *Responder {
	return &Responder{}
}

func (r *Responder) OK(c *gin.Context, res interface{}) {
	response(c, http.StatusOK, res)
}

func (r *Responder) Created(c *gin.Context, res interface{}) {
	response(c, http.StatusCreated, res)
}

func (r *Responder) BadRequest(c *gin.Context, code int, description string, err error) {
	reponseErr(c, http.StatusBadRequest, code, description, err, nil)
}

func (r *Responder) Unauthorized(c *gin.Context, err error) {
	reponseErr(c, http.StatusUnauthorized, 401, "UNAUTHORIZED", err, nil)
}

func (r *Responder) Forbidden(c *gin.Context) {
	reponseErr(c, http.StatusForbidden, 403, "FORBIDDEN", nil, nil)
}

func (r *Responder) NotFound(c *gin.Context, err error) {
	reponseErr(c, http.StatusNotFound, 404, "NOT_FOUND", err, nil)
}

func (r *Responder) Conflict(c *gin.Context, err error) {
	reponseErr(c, http.StatusConflict, 409, "CONFLICT", err, nil)
}

func (r *Responder) Unprocessable(c *gin.Context, description string, err error) {
	reponseErr(c, http.StatusUnprocessableEntity, 422, description, err, nil)
}

func (r *Responder) InternalError(c *gin.Context, err error) {
	reponseErr(c, http.StatusInternalServerError, 500, "SERVER_ERROR", err, nil)
}

func (r *Responder) ResponseErrWithFields(c *gin.Context, vf ValidationFailers) {
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
		Success: true,
		Data:    res,
	})
	c.Abort()
}

func reponseErr(c *gin.Context, httpCode, code int, description string, err error, fields map[string]fieldResponsePart) {
	errResp := &errResponsePart{
		Code:        code,
		Description: description,
		Fields:      fields,
	}
	if gin.Mode() == gin.DebugMode {
		withDebug(errResp, err)
	}
	c.JSON(httpCode, Response{
		Success: false,
		Error:   errResp,
	})
	c.Abort()
}

func withDebug(errResp *errResponsePart, err error) {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
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

type ValidationFailers []validationFailer

type validationFailer struct {
	NameSpace string
	Field     string
	Rule      string
	RuleValue string
	Value     interface{}
}

func getValidationFailers(err error) (vf ValidationFailers, ok bool) {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, v := range ve {
			vf = append(vf, validationFailer{
				NameSpace: v.NameNamespace,
				Field:     v.Field,
				Rule:      v.Tag,
				RuleValue: v.Param,
				Value:     v.Value,
			})
		}
		return vf, true
	}
	return nil, false
}

func (r *Responder) HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	if vf, ok := getValidationFailers(err); ok {
		r.ResponseErrWithFields(c, vf)
		return
	}
	switch err.(type) {
	case *apierrors.NoRows:
		r.NotFound(c, err)
		return
	case *apierrors.Unauthorized:
		r.Unauthorized(c, err)
		return
	case *apierrors.DuplicateEntry:
		r.Conflict(c, err)
		return
	}
	r.InternalError(c, err)
}
