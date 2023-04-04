package kucoin

import "fmt"

func (k *kucoinExchange) agent(fn string) string {
	return fmt.Sprintf("%s.%s", k.NID(), fn)
}
