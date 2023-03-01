// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.5-0.20230118012357-f4cf8f9a5703 DO NOT EDIT.
package models

// Defines values for TransactionType.
const (
	ContractCall     TransactionType = 1
	ContractCreation TransactionType = 2
	Transfer         TransactionType = 0
)

// Address defines model for Address.
type Address = string

// AddressFilter defines model for AddressFilter.
type AddressFilter = Address

// CountResponse defines model for CountResponse.
type CountResponse struct {
	Total Counter `json:"total"`
}

// Counter defines model for Counter.
type Counter = uint64

// Dag defines model for Dag.
type Dag struct {
	Hash             Hash      `json:"hash"`
	Level            Counter   `json:"level"`
	Sender           Address   `json:"sender"`
	Timestamp        Timestamp `json:"timestamp"`
	TransactionCount Counter   `json:"transactionCount"`
}

// DagsPaginatedResponse defines model for DagsPaginatedResponse.
type DagsPaginatedResponse = PaginatedResponse

// Hash defines model for Hash.
type Hash = string

// PaginatedResponse defines model for PaginatedResponse.
type PaginatedResponse struct {
	End     Counter `json:"end"`
	HasNext bool    `json:"hasNext"`
	Start   Counter `json:"start"`
	Total   Counter `json:"total"`
}

// PaginationFilter defines model for PaginationFilter.
type PaginationFilter struct {
	Limit uint64  `json:"limit"`
	Start *uint64 `json:"start"`
}

// Pbft defines model for Pbft.
type Pbft struct {
	Author           Address   `json:"author"`
	Hash             Hash      `json:"hash"`
	Number           Counter   `json:"number"`
	Timestamp        Timestamp `json:"timestamp"`
	TransactionCount Counter   `json:"transactionCount"`
}

// PbftsPaginatedResponse defines model for PbftsPaginatedResponse.
type PbftsPaginatedResponse = PaginatedResponse

// StatsResponse defines model for StatsResponse.
type StatsResponse struct {
	DagsCount         Counter `json:"dagsCount"`
	PbftCount         Counter `json:"pbftCount"`
	TransactionsCount Counter `json:"transactionsCount"`
}

// Timestamp defines model for Timestamp.
type Timestamp = uint64

// Transaction defines model for Transaction.
type Transaction struct {
	BlockNumber      Counter         `json:"blockNumber"`
	From             Address         `json:"from"`
	GasPrice         Counter         `json:"gasPrice"`
	GasUsed          Counter         `json:"gasUsed"`
	Hash             Hash            `json:"hash"`
	Nonce            Counter         `json:"nonce"`
	Status           bool            `json:"status"`
	Timestamp        Timestamp       `json:"timestamp"`
	To               Address         `json:"to"`
	TransactionIndex Counter         `json:"transactionIndex"`
	Type             TransactionType `json:"type"`
	Value            string          `json:"value"`
}

// TransactionType defines model for Transaction.Type.
type TransactionType uint8

// TransactionsPaginatedResponse defines model for TransactionsPaginatedResponse.
type TransactionsPaginatedResponse = PaginatedResponse

// Validator defines model for Validator.
type Validator struct {
	Address   Address `json:"address"`
	PbftCount Counter `json:"pbftCount"`
}

// ValidatorsPaginatedResponse defines model for ValidatorsPaginatedResponse.
type ValidatorsPaginatedResponse = PaginatedResponse

// WeekFilter defines model for WeekFilter.
type WeekFilter struct {
	Week int32 `json:"week"`
	Year int32 `json:"year"`
}

// GetAddressDagsParams defines parameters for GetAddressDags.
type GetAddressDagsParams struct {
	// Pagination Pagination
	Pagination PaginationFilter `form:"pagination" json:"pagination"`
}

// GetAddressPbftsParams defines parameters for GetAddressPbfts.
type GetAddressPbftsParams struct {
	// Pagination Pagination
	Pagination PaginationFilter `form:"pagination" json:"pagination"`
}

// GetAddressTransactionsParams defines parameters for GetAddressTransactions.
type GetAddressTransactionsParams struct {
	// Pagination Pagination
	Pagination PaginationFilter `form:"pagination" json:"pagination"`
}

// GetValidatorsParams defines parameters for GetValidators.
type GetValidatorsParams struct {
	// Week Week to filter by
	Week WeekFilter `form:"week" json:"week"`

	// Pagination Pagination
	Pagination PaginationFilter `form:"pagination" json:"pagination"`
}

// GetValidatorsTotalParams defines parameters for GetValidatorsTotal.
type GetValidatorsTotalParams struct {
	// Filter Week to filter by
	Filter WeekFilter `form:"filter" json:"filter"`
}
