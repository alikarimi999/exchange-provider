package api

import "fmt"

func (a *apiService) agent(fn string) string {
	return fmt.Sprintf("ApiService.%s", fn)
}
