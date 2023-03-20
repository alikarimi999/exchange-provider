package database

import (
	"context"
	"exchange-provider/internal/delivery/database/dto"
	"exchange-provider/internal/entity"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDb) GetPaginated(p *entity.Paginated) error {
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
			fmt.Println(err)
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
