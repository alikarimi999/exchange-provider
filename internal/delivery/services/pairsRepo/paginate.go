package pairsRepo

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
)

func (pr *pairsRepo) GetPaginated(pa *entity.Paginated) error {
	pr.mux.RLock()
	defer pr.mux.RUnlock()

	exIds := []uint{}
	pairIds := []string{}
	if len(pa.Filters) > 0 {
		for _, f := range pa.Filters {
			if strings.ToLower(f.Param) == "lp" {
				switch f.Operator {
				case entity.FilterOperatorIN:
					for _, v := range f.Values {
						fv, ok := v.(float64)
						if !ok {
							errors.Wrap(errors.ErrBadRequest,
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
							errors.Wrap(errors.ErrBadRequest,
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
						errors.Wrap(errors.ErrBadRequest,
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
				ps = append(ps, ep.getAll()...)
			}
		}
	}

	ps2 := []*entity.Pair{}
	if len(pairIds) > 0 {
		for _, p := range ps {
			for _, pId := range pairIds {
				if pairId(p.T1.String(), p.T2.String()) == pId {
					ps2 = append(ps2, p)
				}
			}
		}
	} else {
		ps2 = ps
	}

	start := (pa.Page - 1) * pa.PerPage
	end := pa.Page * pa.PerPage

	if len(ps2) <= int(start) {
		pa.Total = int64(len(ps2))
		return nil
	}

	if len(ps2) < int(end) {
		end = int64(len(ps2))
	}
	pa.Pairs = ps2[start:end]
	pa.Total = int64(len(ps2))
	return nil
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
