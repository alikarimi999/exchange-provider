package database

import (
	"context"
	"exchange-provider/internal/delivery/storage/database/dto"
	"exchange-provider/internal/entity"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoDb) GetPaginated(p *entity.Paginated) error {
	agent := m.agent("GetPaginated")

	filter := wrapFilter(p.Filters)
	count, err := m.orders.CountDocuments(context.Background(), filter)
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}
	p.Total = count

	osDTO := []dto.Order{}
	cur, err := m.orders.Find(context.Background(), filter, wrapOptions(p.Page, p.PerPage))
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
			continue
		}
		p.Orders = append(p.Orders, eo)
	}

	return nil
}

func wrapOptions(page, perPage int64) *options.FindOptions {
	limit := perPage
	skip := perPage * (page - 1)
	return &options.FindOptions{Limit: &limit, Skip: &skip}
}
