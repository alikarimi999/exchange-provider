package entity

type Step struct {
	*Route
	NeedApprove bool
	Approved    bool
}
