// Code generated by goctl. DO NOT EDIT.
package types

type ParamRequest struct {
	Name string `form:"name"`
}

type PathRequest struct {
	Name string `path:"name"`
}

type PostRequest struct {
	Name string `json:"name"`
}

type Response struct {
	Message string
}
