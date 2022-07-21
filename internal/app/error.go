package app

type ErrMsg struct {
	msg string
}

func (e *ErrMsg) String() string {
	return e.msg
}
