package entity

type EvmStep struct {
	*Route
	IsApprove bool
	Approved  bool
}
