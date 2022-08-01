package database

import (
	"order_service/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func setClauses(db *gorm.DB, filters []*entity.Filter) *gorm.DB {
	if filters == nil || len(filters) == 0 {
		return db
	}

	es := []clause.Expression{}
	for _, f := range filters {
		switch f.Cond {
		case entity.FilterCondEqual:
			es = append(es, clause.Eq{
				Column: f.Param,
				Value:  f.Values[0],
			})

		case entity.FilterCondNotEqual:
			es = append(es, clause.Neq{
				Column: f.Param,
				Value:  f.Values[0],
			})
		case entity.FilterCondGreater:
			es = append(es, clause.Gt{
				Column: f.Param,
				Value:  f.Values[0],
			})

		case entity.FilterCondGreaterEqual:
			es = append(es, clause.Gte{
				Column: f.Param,
				Value:  f.Values[0],
			})

		case entity.FilterCondLess:
			es = append(es, clause.Lt{
				Column: f.Param,
				Value:  f.Values[0],
			})

		case entity.FilterCondLessEqual:
			es = append(es, clause.Lte{
				Column: f.Param,
				Value:  f.Values[0],
			})

		case entity.FilterCondIN:

			es = append(es, clause.IN{
				Column: f.Param,
				Values: f.Values,
			})

		case entity.FilterCondNotIn:
			es = append(es, clause.Not(clause.IN{
				Column: f.Param,
				Values: f.Values,
			}))

			// simulate `between` clause
		case entity.FilterCondBetween:
			exs := []clause.Expression{}
			gte := clause.Gte{
				Column: f.Param,
				Value:  f.Values[0],
			}
			lte := clause.Lte{
				Column: f.Param,
				Value:  f.Values[1],
			}
			exs = append(exs, gte, lte)
			es = append(es, clause.And(exs...))

		default:
			continue
		}

	}
	return db.Clauses(es...)

}
