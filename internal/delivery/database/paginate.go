package database

import (
	"context"
	"exchange-provider/internal/delivery/database/dto"
	"exchange-provider/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDb) GetPaginated(p *entity.Paginated, onlyCount bool) error {
	agent := m.agent("GetPaginated")

	filter, err := wrapFilter(p.Filters)
	if err != nil {
		return err
	}

	count, err := m.orders.CountDocuments(context.Background(), filter)
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}
	p.Total = count
	if onlyCount {
		return nil
	}

	osDTO := []*dto.Order{}
	cur, err := m.orders.Find(context.Background(), filter,
		wrapOptions(p.Page, p.PerPage, p.Desc))
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}

	if err := cur.All(context.Background(), &osDTO); err != nil {
		m.l.Error(agent, err.Error())
		return err
	}

	for _, o := range osDTO {
		eo, err := o.ToEntity()
		if err != nil {
			m.l.Debug(agent, err.Error())
			continue
		}

		if o.Status == entity.OCreated.String() && eo.Expire() {
			o.Status = entity.OExpired.String()
			o = dto.UoToDto(eo)
			_, err := m.orders.ReplaceOne(context.Background(), bson.D{{"_id", o.Id}}, o)
			if err != nil {
				m.l.Debug(agent, err.Error())
				continue
			}
		}
		p.Orders = append(p.Orders, eo)
	}

	return nil
}

func wrapOptions(page, perPage int64, desc bool) *options.FindOptions {
	if page == 0 || perPage == 0 {
		return nil
	}

	value := 1
	if desc {
		value = -1
	}

	limit := perPage
	skip := perPage * (page - 1)
	return &options.FindOptions{Limit: &limit, Skip: &skip,
		Sort: bson.D{{"order.createdat", value}}}
}
