package model

type PaginationRequest struct {
	StartsWith int64 `json:"startWith"`
	PageSize   int64 `json:"pageSize"`
}

type Filter struct {
	PaginationRequest
	Criterias     []*Criteria `json:"criteria"`
	SortKey       string      `json:"sortKey"`
	SortDirection string      `json:"sortDirection"`
}

type Operator string

const (
	EQUALS     Operator = "EQUALS"
	NOT_EQUALS Operator = "NOT_EQUALS"
	LIKE       Operator = "LIKE"
	IN         Operator = "IN"
	NOT_IN     Operator = "NOT_IN"
)

type Criteria struct {
	Key      string   `json:"key"`
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
	Values   []string `json:"values"`
}
