package allbridge

import "fmt"

func (ex *allBridge) agent(fn string) string {
	return fmt.Sprintf("%s.%s", ex.NID(), fn)
}
