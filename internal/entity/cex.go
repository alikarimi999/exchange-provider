package entity

type Cex interface {
	Exchange
}

type AddPairsResult struct {
	Added   []string
	Existed []string
	Failed  []*PairsErr
}
type PairsErr struct {
	Pair string
	Err  error
}
