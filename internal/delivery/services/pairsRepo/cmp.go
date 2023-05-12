package pairsRepo

import (
	"exchange-provider/internal/entity"
)

func pairEqual(up, dp *entity.Pair) bool {
	if tokenEqual(up.T1.Id, dp.T1.Id) && tokenEqual(up.T2.Id, dp.T2.Id) {
		return true
	} else if tokenEqual(up.T1.Id, dp.T2.Id) && tokenEqual(up.T2.Id, dp.T1.Id) {
		return true
	}
	return false
}

func tokenEqual(ut, dt entity.TokenId) bool {
	if emptyStr(ut.Symbol) && emptyStr(ut.Standard) && emptyStr(ut.Network) {
		return true
	} else if !emptyStr(ut.Symbol) && emptyStr(ut.Standard) && emptyStr(ut.Network) &&
		len(dt.Symbol) >= len(ut.Symbol) {
		return ut.Symbol == dt.Symbol[:len(ut.Symbol)]
	} else if !emptyStr(ut.Symbol) && !emptyStr(ut.Standard) && emptyStr(ut.Network) &&
		len(dt.Symbol) >= len(ut.Symbol) && len(dt.Standard) >= len(ut.Standard) {
		return ut.Symbol == dt.Symbol[:len(ut.Symbol)] && ut.Standard == dt.Standard[:len(ut.Standard)]
	} else if !emptyStr(ut.Symbol) && !emptyStr(ut.Standard) && !emptyStr(ut.Network) &&
		len(dt.Symbol) >= len(ut.Symbol) && len(dt.Standard) >= len(ut.Standard) && len(dt.Network) >= len(ut.Network) {
		return ut.Symbol == dt.Symbol[:len(ut.Symbol)] && ut.Standard == dt.Standard[:len(ut.Standard)] && ut.Network == dt.Network[:len(ut.Network)]
	}
	return false
}

func emptyStr(s string) bool { return s == "" }
