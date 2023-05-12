package database

import (
	"exchange-provider/internal/entity"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func wrapFilter(filters []*entity.Filter) (bson.D, error) {
	if len(filters) == 0 {
		return bson.D{}, nil
	}

	ds := make([]bson.D, 0)
	for _, f := range filters {
		param := strings.ToLower(f.Param)
		switch strings.ToLower(param) {
		case "id":
			for i, v := range f.Values {
				id, ok := v.(string)
				if !ok {
					return nil, invalidErr(param, v)

				}
				ss := strings.Split(id, "-")
				if len(ss) != 2 || ss[0] != string(entity.PrefOrder) {
					return nil, invalidErr(param, id)

				}
				sid, err := primitive.ObjectIDFromHex(ss[1])
				if err != nil {
					return nil, invalidErr(param, id)
				}
				f.Values[i] = sid
			}
			param = "_id"
		case "pairid":
			sd := []primitive.D{}
			for _, v := range f.Values {
				id, ok := v.(string)
				if !ok {
					return nil, invalidErr(param, v)
				}
				in, out, err := pairFromString(id)
				if err != nil {
					return nil, err
				}
				if in != nil {
					sd = append(sd, bson.D{{"order.in", bson.D{{"$eq", in}}}})
				}
				if out != nil {
					sd = append(sd, bson.D{{"order.out", bson.D{{"$eq", out}}}})
				}
			}
			ds = append(ds, bson.D{{"$or", sd}})
			continue
		case "status":
			sd := []primitive.D{}
			for _, v := range f.Values {
				s, ok := v.(string)
				if !ok {
					return nil, invalidErr(param, v)
				}

				switch s {
				case entity.OCreated.String():
					sd = append(sd, bson.D{{"$and", bson.A{bson.D{{"status", "created"}}, bson.D{{"order.expireat", bson.D{{"$gt", time.Now().Unix()}}}}}}})
				case entity.OExpired.String():
					sd = append(sd, bson.D{{"$or", bson.A{bson.D{{"status", "expired"}}, bson.D{{"$and", bson.A{bson.D{{"status", "created"}}, bson.D{{"order.expireat", bson.D{{"$lt", time.Now().Unix()}}}}}}}}}})

				case entity.OPending.String(), entity.OFailed.String():
					sd = append(sd, bson.D{{"status", s}})
				default:
					return nil, invalidErr(param, s)
				}
			}
			ds = append(ds, bson.D{{"$or", sd}})
			continue

		case "busid":
			param = "order.busid"
		case "userid":
			param = "order.userid"
		}

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
