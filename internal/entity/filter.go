package entity

type Filter struct {
	Param  string
	Cond   FilterCond
	Values []interface{}
}

type FilterCond int

const (
	FilterCondEqual FilterCond = iota
	FilterCondNotEqual
	FilterCondGreater
	FilterCondGreaterEqual
	FilterCondLess
	FilterCondLessEqual
	FilterCondIN
	FilterCondNotIn
	FilterCondBetween
)

func ParseFilterCond(cond string) FilterCond {
	switch cond {
	case "eq":
		return FilterCondEqual

	case "neq":
		return FilterCondNotEqual

	case "gt":
		return FilterCondGreater

	case "gte":
		return FilterCondGreaterEqual

	case "lt":
		return FilterCondLess

	case "lte":
		return FilterCondLessEqual
	case "in":
		return FilterCondIN
	case "notin":
		return FilterCondNotIn
	case "between":
		return FilterCondBetween

	default:
		return FilterCondEqual
	}
}
