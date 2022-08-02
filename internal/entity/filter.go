package entity

type Filter struct {
	Param    string
	Operator FilterOperator
	Values   []interface{}
}

type FilterOperator int

const (
	FilterOperatorEqual FilterOperator = iota
	FilterOpratorNotEqual
	FilterOperatorGreater
	FilterOperatorGreaterEqual
	FilterOperatorLess
	FilterOperatorLessEqual
	FilterOperatorIN
	FilterOperatorNotIn
	FilterOperatorBetween
)

func ParseFilterOperator(cond string) FilterOperator {
	switch cond {
	case "eq":
		return FilterOperatorEqual

	case "neq":
		return FilterOpratorNotEqual

	case "gt":
		return FilterOperatorGreater

	case "gte":
		return FilterOperatorGreaterEqual

	case "lt":
		return FilterOperatorLess

	case "lte":
		return FilterOperatorLessEqual
	case "in":
		return FilterOperatorIN
	case "notin":
		return FilterOperatorNotIn
	case "between":
		return FilterOperatorBetween

	default:
		return FilterOperatorEqual
	}
}
