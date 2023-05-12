package api

import "fmt"

func (a *ApiService) agent(fn string) string {
	return fmt.Sprintf("ApiService.%s", fn)
}
