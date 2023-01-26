package database

import (
	"context"
	"exchange-provider/internal/delivery/database/dto"
	"exchange-provider/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoDb) AddPendingWithdrawal(orderId string) error {
	return m.c.addPendingWithdrawal(orderId)
}
func (m *mongoDb) GetPendingWithdrawals(end time.Time) ([]string, error) {
	return m.c.getPendingWithdrawals(end)
}
func (m *mongoDb) DelPendingWithdrawal(orderId string) error {
	return m.c.delPendingWithdrawal(orderId)
}

func (m *mongoDb) retrivePendingWithd() error {
	cur, err := m.orders.Find(context.Background(),
		bson.D{{"status", entity.OWaitForWithdrawalConfirm}}, nil)
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
