// Package public provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List hosts involved in Playbook runs
	// (GET /api/playbook-dispatcher/v1/run_hosts)
	ApiRunHostsList(ctx echo.Context, params ApiRunHostsListParams) error
	// List Playbook runs
	// (GET /api/playbook-dispatcher/v1/runs)
	ApiRunsList(ctx echo.Context, params ApiRunsListParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ApiRunHostsList converts echo context to params.
func (w *ServerInterfaceWrapper) ApiRunHostsList(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ApiRunHostsListParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Optional query parameter "fields" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "fields", ctx.QueryParams(), &params.Fields)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fields: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiRunHostsList(ctx, params)
	return err
}

// ApiRunsList converts echo context to params.
func (w *ServerInterfaceWrapper) ApiRunsList(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ApiRunsListParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Optional query parameter "fields" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "fields", ctx.QueryParams(), &params.Fields)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fields: %s", err))
	}

	// ------------- Optional query parameter "sort_by" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort_by", ctx.QueryParams(), &params.SortBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort_by: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiRunsList(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/playbook-dispatcher/v1/run_hosts", wrapper.ApiRunHostsList)
	router.GET(baseURL+"/api/playbook-dispatcher/v1/runs", wrapper.ApiRunsList)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"HnXzmykNRrG/3/8vAAD//9PzLD5FHgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
