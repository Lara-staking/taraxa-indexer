// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package models

// Defines values for TransactionType.
const (
	ContractCall             TransactionType = 1
	ContractCreation         TransactionType = 2
	InternalContractCall     TransactionType = 4
	InternalContractCreation TransactionType = 5
	InternalTransfer         TransactionType = 3
	Transfer                 TransactionType = 0
)

// Account defines model for Account.
type Account struct {
	Address Address `json:"address"`
	Balance string  `json:"balance"`
}

// Address defines model for Address.
type Address = string

// AddressFilter defines model for AddressFilter.
type AddressFilter = Address

// CallData defines model for CallData.
type CallData struct {
	Name   string `json:"name"`
	Params any    `json:"params"`
}

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
	Timestamp        Timestamp `json:"timestamp"`
	TransactionCount Counter   `json:"transactionCount"`
}

// DagsPaginatedResponse defines model for DagsPaginatedResponse.
type DagsPaginatedResponse = PaginatedResponse

// EventLog defines model for EventLog.
type EventLog struct {
	Address          Address  `json:"address"`
	Data             string   `json:"data"`
	LogIndex         Counter  `json:"logIndex"`
	Name             string   `json:"name"`
	Params           any      `json:"params"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  Hash     `json:"transactionHash"`
	TransactionIndex Counter  `json:"transactionIndex"`
}

// Hash defines model for Hash.
type Hash = string

// HoldersPaginatedResponse defines model for HoldersPaginatedResponse.
type HoldersPaginatedResponse = PaginatedResponse

// InternalTransactionsResponse defines model for InternalTransactionsResponse.
type InternalTransactionsResponse struct {
	Data []Transaction `json:"data"`
}

// LastTimestamp defines model for LastTimestamp.
type LastTimestamp = uint64

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
	PbftHash         Hash      `json:"pbftHash"`
	Timestamp        Timestamp `json:"timestamp"`
	TransactionCount Counter   `json:"transactionCount"`
}

// PbftsPaginatedResponse defines model for PbftsPaginatedResponse.
type PbftsPaginatedResponse = PaginatedResponse

// Period defines model for Period.
type Period struct {
	EndDate   Timestamp `json:"endDate"`
	HasNext   bool      `json:"hasNext"`
	StartDate Timestamp `json:"startDate"`
}

// StatsResponse defines model for StatsResponse.
type StatsResponse struct {
	DagsCount                Counter        `json:"dagsCount"`
	LastDagTimestamp         *LastTimestamp `json:"lastDagTimestamp" rlp:"nil"`
	LastPbftTimestamp        *LastTimestamp `json:"lastPbftTimestamp" rlp:"nil"`
	LastTransactionTimestamp *LastTimestamp `json:"lastTransactionTimestamp" rlp:"nil"`
	PbftCount                Counter        `json:"pbftCount"`
	TransactionsCount        Counter        `json:"transactionsCount"`
}

// Timestamp defines model for Timestamp.
type Timestamp = uint64

// Transaction defines model for Transaction.
type Transaction struct {
	BlockNumber Counter         `json:"blockNumber"`
	Calldata    *CallData       `json:"calldata,omitempty" rlp:"nil"`
	From        Address         `json:"from"`
	GasCost     Counter         `json:"gasCost"`
	Hash        Hash            `json:"hash"`
	Input       string          `json:"input"`
	Status      bool            `json:"status"`
	Timestamp   Timestamp       `json:"timestamp"`
	To          Address         `json:"to"`
	Type        TransactionType `json:"type"`
	Value       string          `json:"value"`
}

// TransactionType defines model for Transaction.Type.
type TransactionType uint8

// TransactionLogsResponse defines model for TransactionLogsResponse.
type TransactionLogsResponse struct {
	Data []EventLog `json:"data"`
}

// TransactionsPaginatedResponse defines model for TransactionsPaginatedResponse.
type TransactionsPaginatedResponse = PaginatedResponse

// Validator defines model for Validator.
type Validator struct {
	Address   Address `json:"address"`
	PbftCount Counter `json:"pbftCount"`
	Rank      uint64  `json:"rank" rlp:"-"`
}

// ValidatorsPaginatedResponse defines model for ValidatorsPaginatedResponse.
type ValidatorsPaginatedResponse = PaginatedResponse

// Week defines model for Week.
type Week struct {
	Week *int32 `json:"week"`
	Year *int32 `json:"year"`
}

// WeekResponse defines model for WeekResponse.
type WeekResponse struct {
	EndDate   Timestamp `json:"endDate"`
	HasNext   bool      `json:"hasNext"`
	StartDate Timestamp `json:"startDate"`
	Week      *int32    `json:"week"`
	Year      *int32    `json:"year"`
}

// YieldResponse defines model for YieldResponse.
type YieldResponse struct {
	FromBlock Counter `json:"fromBlock"`
	ToBlock   Counter `json:"toBlock"`
	Yield     string  `json:"yield"`
}

// AddressParam defines model for addressParam.
type AddressParam = AddressFilter

// BlockNumParam defines model for blockNumParam.
type BlockNumParam = Counter

// HashParam defines model for hashParam.
type HashParam = Hash

// PaginationParam defines model for paginationParam.
type PaginationParam = PaginationFilter

// WeekParam defines model for weekParam.
type WeekParam = Week

// GetAddressDagsParams defines parameters for GetAddressDags.
type GetAddressDagsParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetAddressPbftsParams defines parameters for GetAddressPbfts.
type GetAddressPbftsParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetAddressTransactionsParams defines parameters for GetAddressTransactions.
type GetAddressTransactionsParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetAddressYieldParams defines parameters for GetAddressYield.
type GetAddressYieldParams struct {
	// BlockNumber Block Number
	BlockNumber *BlockNumParam `form:"blockNumber,omitempty" json:"blockNumber,omitempty"`
}

// GetHoldersParams defines parameters for GetHolders.
type GetHoldersParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetTotalYieldParams defines parameters for GetTotalYield.
type GetTotalYieldParams struct {
	// BlockNumber Block Number
	BlockNumber *BlockNumParam `form:"blockNumber,omitempty" json:"blockNumber,omitempty"`
}

// GetValidatorsParams defines parameters for GetValidators.
type GetValidatorsParams struct {
	// Week Week to filter by
	Week *WeekParam `form:"week,omitempty" json:"week,omitempty"`

	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetValidatorsTotalParams defines parameters for GetValidatorsTotal.
type GetValidatorsTotalParams struct {
	// Week Week to filter by
	Week *WeekParam `form:"week,omitempty" json:"week,omitempty"`
}

// GetValidatorParams defines parameters for GetValidator.
type GetValidatorParams struct {
	// Week Week to filter by
	Week *WeekParam `form:"week,omitempty" json:"week,omitempty"`
}
