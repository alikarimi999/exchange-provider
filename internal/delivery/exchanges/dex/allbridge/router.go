package allbridge

import "exchange-provider/internal/delivery/exchanges/dex/allbridge/types"

func (ex *exchange) isInternal(r *types.Route) bool {
	return r.ExchangeNid == ex.NID()
}
