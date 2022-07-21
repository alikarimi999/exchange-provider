package errors

type msg struct {
	text string
}

func NewMesssage(text string) *msg {
	return &msg{text: text}
}

func (m *msg) String() string {
	return m.text
}
