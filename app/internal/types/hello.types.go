// Code generated by goctl-types plugin. DO NOT EDIT.
package types

import (
	"time"
)

var (
	_ = time.Now()
)

type ParamRequest struct {
	Name string `form:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type PostRequest struct {
	Name string `json:"name"`
}

type PathRequest struct {
	Name string `path:"name"`
}