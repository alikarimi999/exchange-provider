package dto

import "fmt"

type Number float64

func (n Number) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.*f", 8, n)
	return []byte(s), nil
}
