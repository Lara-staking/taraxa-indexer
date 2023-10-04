// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
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
	// Returns yield for the address
	// (GET /address/{address}/yield)
	GetAddressYield(ctx echo.Context, address AddressParam, params GetAddressYieldParams) error
	// Returns the list of TARA token holders and their balances
	// (GET /holders)
	GetHolders(ctx echo.Context, params GetHoldersParams) error
	// Returns total supply
	// (GET /totalSupply)
	GetTotalSupply(ctx echo.Context) error
	// Returns total yield
	// (GET /totalYield)
	GetTotalYield(ctx echo.Context, params GetTotalYieldParams) error
	// Returns the decoded transaction
	// (GET /transaction/{hash})
	GetTransaction(ctx echo.Context, hash HashParam) error
	// Returns internal transactions
	// (GET /transaction/{hash}/internal_transactions)
	GetInternalTransactions(ctx echo.Context, hash HashParam) error
	// Returns event logs of transaction
	// (GET /transaction/{hash}/logs)
	GetTransactionLogs(ctx echo.Context, hash HashParam) error
	// Returns all validators
	// (GET /validators)
	GetValidators(ctx echo.Context, params GetValidatorsParams) error
	// Returns total number of PBFT blocks
	// (GET /validators/total)
	GetValidatorsTotal(ctx echo.Context, params GetValidatorsTotalParams) error
	// Returns info about the validator
	// (GET /validators/{address})
	GetValidator(ctx echo.Context, address AddressParam, params GetValidatorParams) error
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

// GetAddressYield converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressYield(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressYieldParams
	// ------------- Optional query parameter "blockNumber" -------------

	err = runtime.BindQueryParameter("form", true, false, "blockNumber", ctx.QueryParams(), &params.BlockNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter blockNumber: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressYield(ctx, address, params)
	return err
}

// GetHolders converts echo context to params.
func (w *ServerInterfaceWrapper) GetHolders(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetHoldersParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHolders(ctx, params)
	return err
}

// GetTotalSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalSupply(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTotalSupply(ctx)
	return err
}

// GetTotalYield converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalYield(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTotalYieldParams
	// ------------- Optional query parameter "blockNumber" -------------

	err = runtime.BindQueryParameter("form", true, false, "blockNumber", ctx.QueryParams(), &params.BlockNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter blockNumber: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTotalYield(ctx, params)
	return err
}

// GetTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransaction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "hash", runtime.ParamLocationPath, ctx.Param("hash"), &hash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTransaction(ctx, hash)
	return err
}

// GetInternalTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetInternalTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "hash", runtime.ParamLocationPath, ctx.Param("hash"), &hash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetInternalTransactions(ctx, hash)
	return err
}

// GetTransactionLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransactionLogs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "hash", runtime.ParamLocationPath, ctx.Param("hash"), &hash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTransactionLogs(ctx, hash)
	return err
}

// GetValidators converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidators(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorsParams
	// ------------- Optional query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, false, "week", ctx.QueryParams(), &params.Week)
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
	// ------------- Optional query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, false, "week", ctx.QueryParams(), &params.Week)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter week: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidatorsTotal(ctx, params)
	return err
}

// GetValidator converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidator(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorParams
	// ------------- Optional query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, false, "week", ctx.QueryParams(), &params.Week)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter week: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidator(ctx, address, params)
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
	router.GET(baseURL+"/address/:address/yield", wrapper.GetAddressYield)
	router.GET(baseURL+"/holders", wrapper.GetHolders)
	router.GET(baseURL+"/totalSupply", wrapper.GetTotalSupply)
	router.GET(baseURL+"/totalYield", wrapper.GetTotalYield)
	router.GET(baseURL+"/transaction/:hash", wrapper.GetTransaction)
	router.GET(baseURL+"/transaction/:hash/internal_transactions", wrapper.GetInternalTransactions)
	router.GET(baseURL+"/transaction/:hash/logs", wrapper.GetTransactionLogs)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)
	router.GET(baseURL+"/validators/:address", wrapper.GetValidator)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//LrLIqwc/AAA=",
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
