// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.5-0.20230118012357-f4cf8f9a5703 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all DAG blocks
	// (GET /address/{address}/dags)
	GetAddressDags(ctx echo.Context, address AddressFilter, params GetAddressDagsParams) error
	// Returns all PBFT blocks
	// (GET /address/{address}/pbfts)
	GetAddressPbfts(ctx echo.Context, address AddressFilter, params GetAddressPbftsParams) error
	// Returns stats for the address
	// (GET /address/{address}/stats)
	GetAddressStats(ctx echo.Context, address AddressFilter) error
	// Returns all transactions
	// (GET /address/{address}/transactions)
	GetAddressTransactions(ctx echo.Context, address AddressFilter, params GetAddressTransactionsParams) error
	// Returns all validators
	// (GET /validators)
	GetValidators(ctx echo.Context, params GetValidatorsParams) error
	// Returns total number of PBFT blocks
	// (GET /validators/total)
	GetValidatorsTotal(ctx echo.Context, params GetValidatorsTotalParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAddressDags converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressDags(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressFilter

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressDagsParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressDags(ctx, address, params)
	return err
}

// GetAddressPbfts converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressPbfts(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressFilter

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressPbftsParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressPbfts(ctx, address, params)
	return err
}

// GetAddressStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressStats(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressFilter

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressStats(ctx, address)
	return err
}

// GetAddressTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressFilter

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressTransactionsParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressTransactions(ctx, address, params)
	return err
}

// GetValidators converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidators(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorsParams
	// ------------- Required query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, true, "week", ctx.QueryParams(), &params.Week)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter week: %s", err))
	}

	// ------------- Optional query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, false, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidators(ctx, params)
	return err
}

// GetValidatorsTotal converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidatorsTotal(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorsTotalParams
	// ------------- Required query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, true, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidatorsTotal(ctx, params)
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

	router.GET(baseURL+"/address/:address/dags", wrapper.GetAddressDags)
	router.GET(baseURL+"/address/:address/pbfts", wrapper.GetAddressPbfts)
	router.GET(baseURL+"/address/:address/stats", wrapper.GetAddressStats)
	router.GET(baseURL+"/address/:address/transactions", wrapper.GetAddressTransactions)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
