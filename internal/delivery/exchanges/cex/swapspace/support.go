package swapspace

import "exchange-provider/internal/entity"

func (ex *exchange) Support(t1, t2 *entity.Token) bool {
	return false
}
