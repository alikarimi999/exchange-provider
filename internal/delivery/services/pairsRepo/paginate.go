package pairsRepo

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sort"
	"strings"
)

func (pr *pairsRepo) GetPaginated(pa *entity.Paginated, admin bool) error {
	pr.mux.RLock()
	defer pr.mux.RUnlock()

	exIds := []uint{}
	pairIds := []string{}
	var up *entity.Pair

	if pa.PerPage > 0 {
		if len(pa.Filters) > 0 {
			for _, f := range pa.Filters {
				if strings.ToLower(f.Param) == "lp" {
					switch f.Operator {
					case entity.FilterOperatorIN:
						for _, v := range f.Values {
							fv, ok := v.(float64)
							if !ok {
								return errors.Wrap(errors.ErrBadRequest,
									errors.NewMesssage(fmt.Sprintf("value %v not a number", v)))
							}
							exIds = append(exIds, uint(fv))
						}
					case entity.FilterOperatorEqual:
						v, ok := f.Values[0].(float64)
						if !ok {
							return errors.Wrap(errors.ErrBadRequest,
								errors.NewMesssage(fmt.Sprintf("value %v not a number", f.Values[0])))
						}
						exIds = append(exIds, uint(v))

					default:
						return errors.Wrap(errors.ErrBadRequest,
							errors.NewMesssage("filter operator is not supported"))
					}
				}

				if strings.ToLower(f.Param) == "pairid" {
					switch f.Operator {
					case entity.FilterOperatorIN:
						for _, v := range f.Values {
							sv, ok := v.(string)
							if !ok {
								return errors.Wrap(errors.ErrBadRequest,
									errors.NewMesssage(fmt.Sprintf("value %v not a string", v)))
							}
							id, err := fixPairId(sv)
							if err != nil {
								return err
							}
							pairIds = append(pairIds, id)
						}
					case entity.FilterOperatorEqual:
						sv, ok := f.Values[0].(string)
						if !ok {
							return errors.Wrap(errors.ErrBadRequest,
								errors.NewMesssage(fmt.Sprintf("value %v not a string", f.Values[0])))
						}
						id, err := fixPairId(sv)
						if err != nil {
							return err
						}
						pairIds = append(pairIds, id)
					default:
						return errors.Wrap(errors.ErrBadRequest,
							errors.NewMesssage("filter operator is not supported"))
					}
				} else if strings.ToLower(f.Param) == "token" {
					switch f.Operator {
					case entity.FilterOperatorEqual:
						sv, ok := f.Values[0].(string)
						if !ok {
							return errors.Wrap(errors.ErrBadRequest,
								errors.NewMesssage(fmt.Sprintf("value %v not a string", f.Values[0])))
						}
						t1, t2 := pairFromString(sv)
						up = &entity.Pair{
							T1: &entity.Token{Id: t1},
							T2: &entity.Token{Id: t2},
						}

					default:
						return errors.Wrap(errors.ErrBadRequest,
							errors.NewMesssage("filter operator is not supported"))

					}
				}

			}
		}
	}

	if len(exIds) == 0 {
		for _, ep := range pr.eps {
			exIds = append(exIds, ep.exId)
		}
	}

	ps := []*entity.Pair{}
	for _, id := range exIds {
		for _, ep := range pr.eps {
			if ep.exId == id {
				ps0 := []*entity.Pair{}
				ps0 = append(ps0, ep.getAll()...)
				if !ep.ex.IsEnable() {
					for _, p := range ps0 {
						p.Enable = false
					}
				}
				ps = append(ps, ps0...)
			}
		}
	}
	ps2 := make(map[string][]*entity.Pair)
	if len(pairIds) > 0 {
		for _, p := range ps {
			if !admin && !p.Enable {
				continue
			}
			for _, pId := range pairIds {
				if pairId(p.T1.String(), p.T2.String()) == pId {
					ps2[pId] = append(ps2[pId], p)
				}
			}
		}
	} else if up != nil {
		for _, dp := range ps {
			if !admin && !dp.Enable {
				continue
			}
			if pairEqual(up, dp) {
				ps2[pairId(dp.T1.String(), dp.T2.String())] =
					append(ps2[pairId(dp.T1.String(), dp.T2.String())], dp)
			}
		}
	} else {
		for _, p := range ps {
			if !admin && !p.Enable {
				continue
			}
			ps2[pairId(p.T1.String(), p.T2.String())] =
				append(ps2[pairId(p.T1.String(), p.T2.String())], p)
		}
	}

	ps3 := []*entity.Pair{}
	for _, ps := range ps2 {
		if admin {
			ps3 = append(ps3, ps...)
		} else {
			ps3 = append(ps3, findLowestMinAndHighestMax(ps))
		}
	}
	sortPairs(ps3)

	var start, end int64
	if pa.PerPage == 0 {
		start = 0
		end = int64(len(ps3))
	} else {
		start = (pa.Page - 1) * pa.PerPage
		end = pa.Page * pa.PerPage
	}
	if len(ps3) <= int(start) {
		pa.Total = int64(len(ps3))
		return nil
	}

	if len(ps3) < int(end) {
		end = int64(len(ps3))
	}

	pa.Pairs = ps3[start:end]
	pa.Total = int64(len(ps3))
	return nil
}

func sortPairs(ps []*entity.Pair) {
	sort.Slice(ps, func(i, j int) bool {
		return pairId(ps[i].T1.String(), ps[i].T2.String()) <
			pairId(ps[j].T1.String(), ps[j].T2.String())
	})
}

func fixPairId(id string) (string, error) {
	ss := strings.Split(id, "/")
	if len(ss) != 2 {
		return "", errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid pairId"))
	}
	if ss[0] < ss[1] {
		return ss[0] + "/" + ss[1], nil
	} else {
		return ss[1] + "/" + ss[0], nil
	}
}

func pairFromString(id string) (entity.TokenId, entity.TokenId) {
	ss := strings.Split(strings.ToUpper(id), "/")
	if len(ss) == 1 {
		return string2TokenId(ss[0]), entity.TokenId{}
	} else if len(ss) > 1 {
		return string2TokenId(ss[0]), string2TokenId(ss[1])
	}
	return entity.TokenId{}, entity.TokenId{}
}

func string2TokenId(id string) entity.TokenId {
	t := entity.TokenId{}
	ts := strings.Split(id, "-")
	t.Symbol = ts[0]
	if len(ts) >= 2 {
		t.Standard = ts[1]
	}
	if len(ts) == 3 {
		t.Network = ts[2]
	}
	return t
}
func toPointer(n float64) *float64 { return &n }

func findLowestMinAndHighestMax(ps []*entity.Pair) *entity.Pair {

	var min1, min2, max1, max2 *float64

	for _, p := range ps {
		if min1 == nil || p.T2.Min < *min1 {
			min1 = toPointer(p.T1.Min)
		}

		if min2 == nil || p.T2.Min < *min2 {
			min2 = toPointer(p.T2.Min)
		}

		if max1 == nil || (p.T1.Max == 0 || p.T1.Max > *max1) {
			max1 = toPointer(p.T1.Max)
		}

		if max2 == nil || (p.T2.Max == 0 || p.T2.Max > *max2) {
			max2 = toPointer(p.T2.Max)
		}
	}

	return &entity.Pair{
		T1: &entity.Token{
			Id:     ps[0].T1.Id,
			Native: ps[0].T1.Native,
			Min:    *min1,
			Max:    *max1,
		},
		T2: &entity.Token{
			Id:     ps[0].T2.Id,
			Native: ps[0].T2.Native,
			Min:    *min2,
			Max:    *max2,
		},

		FeeRate1:    ps[0].FeeRate1,
		FeeRate2:    ps[0].FeeRate2,
		ExchangeFee: ps[0].ExchangeFee,
	}

}
