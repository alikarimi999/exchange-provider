package database

import (
	"exchange-provider/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func wrapFilter(filters []*entity.Filter) bson.D {

	if len(filters) == 0 {
		return bson.D{{}}
	}

	ds := make([]bson.D, 0)

	for _, f := range filters {
		switch f.Operator {
		case entity.FilterOperatorEqual:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$eq", f.Values[0]}}}})
		case entity.FilterOpratorNotEqual:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$ne", f.Values[0]}}}})
		case entity.FilterOperatorGreater:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$gt", f.Values[0]}}}})
		case entity.FilterOperatorGreaterEqual:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$gte", f.Values[0]}}}})
		case entity.FilterOperatorLess:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$lt", f.Values[0]}}}})
		case entity.FilterOperatorLessEqual:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$lte", f.Values[0]}}}})
		case entity.FilterOperatorIN:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$in", f.Values}}}})
		case entity.FilterOperatorNotIn:
			ds = append(ds, bson.D{{f.Param, bson.D{{"$nin", f.Values}}}})
		case entity.FilterOperatorBetween:
			gte := bson.D{{"$gte", f.Values[0]}}
			lte := bson.D{{"$lte", f.Values[1]}}
			ds = append(ds, gte, lte)

		default:
			continue
		}

	}

	return bson.D{{"$and", ds}}
}