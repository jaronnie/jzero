// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "github.com/jzero-io/jzero/app/internal/handler/file"
	hello "github.com/jzero-io/jzero/app/internal/handler/hello"
	"github.com/jzero-io/jzero/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/static/:file",
				Handler: file.DownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: file.UploadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

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
