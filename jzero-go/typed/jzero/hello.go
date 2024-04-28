
// Code generated by jzero. DO NOT EDIT.
// type: hello

package jzero

import (
    "context"

    "github.com/jaronnie/jzero-go/model/jzero/types"
    
    "github.com/jaronnie/jzero-go/rest"
)

var (
    _ = context.Background()
)

type HelloGetter interface {
	Hello() HelloInterface
}

type HelloInterface interface {
	// API /api/v1/hello 
	HelloParamHandler(ctx context.Context,param *types.ParamRequest) (*types.Response, error)
	// API /api/v1/hello 
	HelloPostHandler(ctx context.Context,param *types.PostRequest) (*types.Response, error)
	// API /api/v1/hello/:name 
	HelloPathHandler(ctx context.Context,param *types.PathRequest) (*types.Response, error)
	
	HelloExpansion
}

type helloClient struct {
	client rest.Interface
}

func newHelloClient(c *JzeroClient) *helloClient {
	return &helloClient{
		client: c.RESTClient(),
	}
}

func (x *helloClient) HelloParamHandler(ctx context.Context,param *types.ParamRequest) (*types.Response, error) {
	var resp types.Response
		err := x.client.Verb("GET").
		SubPath(
			"/api/v1/hello",
		).
		Params(
			rest.QueryParam{Name: "name", Value: param.Name},
		).
		Body(nil).
		Do(ctx).
		Into(&resp, true)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (x *helloClient) HelloPostHandler(ctx context.Context,param *types.PostRequest) (*types.Response, error) {
	var resp types.Response
		err := x.client.Verb("POST").
		SubPath(
			"/api/v1/hello",
		).
		Params(
		).
		Body(nil).
		Do(ctx).
		Into(&resp, true)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (x *helloClient) HelloPathHandler(ctx context.Context,param *types.PathRequest) (*types.Response, error) {
	var resp types.Response
		err := x.client.Verb("GET").
		SubPath(
			"/api/v1/hello/:name",
			rest.PathParam{Name: "name", Value: param.Name},
		).
		Params(
		).
		Body(nil).
		Do(ctx).
		Into(&resp, true)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

