package entity

type EvmStep struct {
	*Route
	NeedApprove bool
	Approved    bool
}
