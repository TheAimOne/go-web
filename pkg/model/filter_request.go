package model

type PaginationRequest struct {
	PageNumber int64 `json:"pageNumber"`
	PageSize   int64 `json:"pageSize"`
}

type Filter struct {
	PaginationRequest
	Criterias     []*Criteria `json:"criteria"`
	SortKey       string      `json:"sortKey"`
	SortDirection string      `json:"sortDirection"`
	IsAnd         bool        `json:"isAnd"`
}

type Operator string

const (
	EQUALS       Operator = "EQUALS"
	NOT_EQUALS   Operator = "NOT_EQUALS"
	CONTAINS     Operator = "CONTAINS"
	NOT_CONTAINS Operator = "NOT_CONTAINS"
	IN           Operator = "IN"
	NOT_IN       Operator = "NOT_IN"
)

type Criteria struct {
	Key      string   `json:"key"`
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
	Values   []string `json:"values"`
}

type PaginationResponse[T any] struct {
	TotalCount int64 `json:"totalCount"`
	Data       []T   `json:"data"`
}
