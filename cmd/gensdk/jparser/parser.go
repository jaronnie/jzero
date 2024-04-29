package jparser

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/jzero-io/jzero/cmd/gensdk/config"

	"github.com/jzero-io/jzero/daemon/pkg/stringx"

	"github.com/jzero-io/jzero/cmd/gensdk/jparser/api"
	"github.com/jzero-io/jzero/cmd/gensdk/jparser/gateway"
	"github.com/jzero-io/jzero/cmd/gensdk/vars"

	"github.com/jhump/protoreflect/desc"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
)

func Parse(config *config.Config, fds []*desc.FileDescriptor, apiSpecs []*spec.ApiSpec) (vars.ScopeResourceHTTPInterfaceMap, error) {
	interfaces, err := genHTTPInterfaces(config, fds, apiSpecs)
	if err != nil {
		return nil, err
	}
	return convertToMap(interfaces), nil
}

func genHTTPInterfaces(config *config.Config, fds []*desc.FileDescriptor, apiSpecs []*spec.ApiSpec) ([]*vars.HTTPInterface, error) {
	var httpInterfaces []*vars.HTTPInterface

	for _, fd := range fds {
		services := fd.GetServices()
		for _, service := range services {
			methods := service.GetMethods()
			for _, method := range methods {
				ext := proto.GetExtension(method.GetMethodOptions(), annotations.E_Http)
				var httpInterface vars.HTTPInterface
				switch rule := ext.(type) {
				case *annotations.HttpRule:
					if rule == nil {
						continue
					}

					switch httpRule := rule.GetPattern().(type) {
					case *annotations.HttpRule_Get:
						httpInterface.Method = http.MethodGet
						httpInterface.URL = httpRule.Get
					case *annotations.HttpRule_Post:
						httpInterface.Method = http.MethodPost
						httpInterface.URL = httpRule.Post
					}

					var requestBodyName string
					if (httpInterface.Method == http.MethodPost) && rule.Body != "*" {
						for _, v := range method.GetInputType().GetFields() {
							if rule.Body == v.GetName() {
								requestBodyName = v.GetName()
							}
						}
					} else {
						requestBodyName = method.GetInputType().GetName()
					}
					httpInterface.RequestBody = &vars.RequestBody{
						RealBodyName: requestBodyName,
						Name:         stringx.FirstUpper(method.GetInputType().GetName()),
						Body:         rule.Body,
						Type:         "proto",
						Package:      strings.TrimPrefix(*service.GetFile().GetFileOptions().GoPackage, "./"),
					}
					httpInterface.ResponseBody = &vars.ResponseBody{
						Name:    stringx.FirstUpper(method.GetOutputType().GetName()),
						Package: strings.TrimPrefix(*service.GetFile().GetFileOptions().GoPackage, "./"),
					}
				}
				httpInterface.MethodName = method.GetName()
				// TODO if multiple. Use config to get scope
				httpInterface.Scope = vars.Scope(config.APP)
				httpInterface.Resource = vars.Resource(service.GetName())

				pathParams, err := gateway.PathParam(httpInterface.URL)
				if err != nil {
					return nil, nil
				}
				httpInterface.PathParams = pathParams
				queryParams := gateway.CreateQueryParams(method)
				httpInterface.QueryParams = queryParams

				httpInterfaces = append(httpInterfaces, &httpInterface)
			}
		}
	}

	for _, apiSpec := range apiSpecs {
		for _, group := range apiSpec.Service.Groups {
			for _, route := range group.Routes {
				path, _ := url.JoinPath(group.Annotation.Properties["prefix"], route.Path)

				httpInterface := vars.HTTPInterface{
					Scope:      vars.Scope(config.APP),
					Resource:   vars.Resource(group.Annotation.Properties["group"]),
					Method:     strings.ToUpper(route.Method),
					URL:        path,
					MethodName: route.Handler,
					Comments:   route.AtDoc.Text,
				}
				if route.RequestType != nil {
					httpInterface.RequestBody = &vars.RequestBody{
						Name:         stringx.FirstUpper(route.RequestType.Name()),
						RealBodyName: stringx.FirstUpper(route.RequestType.Name()),
						Package:      "types",
						Type:         "api",
					}
				} else {
					continue
				}

				if route.ResponseType != nil {
					httpInterface.ResponseBody = &vars.ResponseBody{
						Name:    stringx.FirstUpper(route.ResponseType.Name()),
						Package: "types",
					}
				} else {
					continue
				}

				pathParams, err := api.CreatePathParam(httpInterface.URL, &route)
				if err != nil {
					return nil, nil
				}
				httpInterface.PathParams = pathParams
				queryParams := api.CreateQueryParams(&route)
				httpInterface.QueryParams = queryParams

				httpInterfaces = append(httpInterfaces, &httpInterface)
			}
		}
	}
	return httpInterfaces, nil
}

func convertToMap(interfaces []*vars.HTTPInterface) vars.ScopeResourceHTTPInterfaceMap {
	scopeResourceHTTPInterfaceMap := make(vars.ScopeResourceHTTPInterfaceMap)

	for _, inf := range interfaces {
		scope := inf.Scope
		resource := inf.Resource

		if _, ok := scopeResourceHTTPInterfaceMap[scope]; !ok {
			scopeResourceHTTPInterfaceMap[scope] = make(map[vars.Resource][]*vars.HTTPInterface)
		}

		scopeResourceHTTPInterfaceMap[scope][resource] = append(scopeResourceHTTPInterfaceMap[scope][resource], inf)
	}

	return scopeResourceHTTPInterfaceMap
}
