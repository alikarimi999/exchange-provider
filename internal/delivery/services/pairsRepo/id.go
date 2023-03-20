package pairsRepo

import "fmt"

func pairId(t1, t2 string) string {
	if t1 < t2 {
		return fmt.Sprintf("%s/%s", t1, t2)
	}
	return fmt.Sprintf("%s/%s", t2, t1)
}
