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
	"REV2mCvpZvO/AAAA//8s0TRHrz8AAA==",
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
