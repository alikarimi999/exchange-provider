package database

import (
	"context"
	"exchange-provider/internal/delivery/database/dto"
	"exchange-provider/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoDb) AddPendingWithdrawal(oId *entity.ObjectId) error {
	return m.c.addPendingWithdrawal(oId.Id)
}
func (m *mongoDb) GetPendingWithdrawals(end time.Time) ([]*entity.ObjectId, error) {
	ids, err := m.c.getPendingWithdrawals(end)
	if err != nil {
		return nil, err
	}
	oIds := []*entity.ObjectId{}
	for _, id := range ids {
		oIds = append(oIds, &entity.ObjectId{Prefix: entity.PrefOrder, Id: id})
	}
	return oIds, nil
}

func (m *mongoDb) DelPendingWithdrawal(oId *entity.ObjectId) error {
	return m.c.delPendingWithdrawal(oId.Id)
}

func (m *mongoDb) retrivePendingWithd() error {
	cur, err := m.orders.Find(context.Background(),
		bson.D{{"order.status", entity.OWaitForWithdrawalConfirm}}, nil)
	if err != nil {
		return err
	}
	osDTO := []*dto.Order{}
	if err := cur.All(context.Background(), &osDTO); err != nil {
		return err
	}

	ids := []string{}
	for _, o := range osDTO {
		ids = append(ids, o.Id.Hex())
	}
	return m.c.addPendingWithdrawal(ids...)
}
