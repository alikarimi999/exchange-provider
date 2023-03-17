package swapspace

import "fmt"

func (ex *exchange) agent(fn string) string {
	return fmt.Sprintf("%s.%s", ex.Id(), fn)
}
