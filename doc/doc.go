// Package classification Wordplay-go API.
// version: 1.0
//
//     Security:
//     - basicAuth:
//
//     SecurityDefinitions:
//       basicAuth:
//         type: basic
//         in: header
//
// swagger:meta
package doc

import "github.com/djumpen/wordplay-go/api"

// swagger:parameters UserCreateReq
type UserCreateReq struct {
	Body api.UserCreateReq
}

// swagger:response
type UserCreatedResp struct {
	Body api.UserCreatedResp
}

// swagger:response
type UserResp struct {
	Body api.UserResp
}
