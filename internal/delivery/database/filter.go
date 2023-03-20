package database

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func wrapFilter(filters []*entity.Filter) (bson.D, error) {
	if len(filters) == 0 {
		return bson.D{}, nil
	}

	ds := make([]bson.D, 0)
	for _, f := range filters {
		param := strings.ToLower(f.Param)

		switch param {
		case "id":
			for i, v := range f.Values {
				id, ok := v.(string)
				if !ok {
					return nil, errors.Wrap(errors.ErrBadRequest,
						errors.NewMesssage(fmt.Sprintf("%v is not a string", v)))
				}
				ss := strings.Split(id, "-")
				if len(ss) != 2 || ss[0] != string(entity.PrefOrder) {
					errors.Wrap(errors.ErrBadRequest,
						errors.NewMesssage(fmt.Sprintf("%s is invalid", id)))
				}
				f.Values[i] = ss[1]
			}
			param = "objectid." + param
		}

		param = "order." + param
		switch f.Operator {
		case entity.FilterOperatorEqual:
			ds = append(ds, bson.D{{param, bson.D{{"$eq", f.Values[0]}}}})
		case entity.FilterOpratorNotEqual:
			ds = append(ds, bson.D{{param, bson.D{{"$ne", f.Values[0]}}}})
		case entity.FilterOperatorGreater:
			ds = append(ds, bson.D{{param, bson.D{{"$gt", f.Values[0]}}}})
		case entity.FilterOperatorGreaterEqual:
			ds = append(ds, bson.D{{param, bson.D{{"$gte", f.Values[0]}}}})
		case entity.FilterOperatorLess:
			ds = append(ds, bson.D{{param, bson.D{{"$lt", f.Values[0]}}}})
		case entity.FilterOperatorLessEqual:
			ds = append(ds, bson.D{{param, bson.D{{"$lte", f.Values[0]}}}})
		case entity.FilterOperatorIN:
			ds = append(ds, bson.D{{param, bson.D{{"$in", f.Values}}}})
		case entity.FilterOperatorNotIn:
			ds = append(ds, bson.D{{param, bson.D{{"$nin", f.Values}}}})
		case entity.FilterOperatorBetween:
			gte := bson.D{{"$gte", f.Values[0]}}
			lte := bson.D{{"$lte", f.Values[1]}}
			ds = append(ds, gte, lte)

		default:
			continue
		}

	}

	return bson.D{{"$and", ds}}, nil
}
