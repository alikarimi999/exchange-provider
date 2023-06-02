package kucoin

import "fmt"

func (k *exchange) agent(fn string) string {
	return fmt.Sprintf("%s.%s", k.NID(), fn)
}
