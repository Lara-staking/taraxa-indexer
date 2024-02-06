// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
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
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
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
	// Returns yield for the address
	// (GET /address/{address}/yieldForInterval)
	GetAddressYieldForInterval(ctx echo.Context, address AddressParam, params GetAddressYieldForIntervalParams) error
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressYield(ctx, address, params)
	return err
}

// GetAddressYieldForInterval converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressYieldForInterval(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressYieldForIntervalParams
	// ------------- Optional query parameter "fromBlock" -------------

	err = runtime.BindQueryParameter("form", true, false, "fromBlock", ctx.QueryParams(), &params.FromBlock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fromBlock: %s", err))
	}

	// ------------- Required query parameter "toBlock" -------------

	err = runtime.BindQueryParameter("form", true, true, "toBlock", ctx.QueryParams(), &params.ToBlock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter toBlock: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressYieldForInterval(ctx, address, params)
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetHolders(ctx, params)
	return err
}

// GetTotalSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalSupply(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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
	router.GET(baseURL+"/address/:address/yieldForInterval", wrapper.GetAddressYieldForInterval)
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
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
