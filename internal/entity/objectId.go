package entity

import "fmt"

type ObjectId struct {
	Prefix ObjectPrefix
	Id     string
}

func (o *ObjectId) String() string {
	return fmt.Sprintf("%s%s%s", o.Prefix, IdDelimiter, o.Id)
}

type ObjectPrefix string

const (
	PrefOrder   ObjectPrefix = "ord"
	IdDelimiter string       = "-"
)
