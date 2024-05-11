// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	hello "github.com/jzero-io/jzero/app/internal/handler/hello"
	"github.com/jzero-io/jzero/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hello",
				Handler: hello.HelloParamHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/hello",
				Handler: hello.HelloPostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/hello/:name",
				Handler: hello.HelloPathHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
