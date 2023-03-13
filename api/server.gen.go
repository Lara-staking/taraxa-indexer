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
	GetAddressDags(ctx echo.Context, address AddressParam, params GetAddressDagsParams) error
	// Returns all PBFT blocks
	// (GET /address/{address}/pbfts)
	GetAddressPbfts(ctx echo.Context, address AddressParam, params GetAddressPbftsParams) error
	// Returns stats for the address
	// (GET /address/{address}/stats)
	GetAddressStats(ctx echo.Context, address AddressParam) error
	// Returns all transactions
	// (GET /address/{address}/transactions)
	GetAddressTransactions(ctx echo.Context, address AddressParam, params GetAddressTransactionsParams) error
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
	var address AddressParam

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
	var address AddressParam

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
	var address AddressParam

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
	var address AddressParam

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

	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
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
