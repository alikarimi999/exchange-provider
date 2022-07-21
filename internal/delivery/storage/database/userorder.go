package database

import (
	"order_service/internal/delivery/storage/database/dto"
	"order_service/internal/entity"

	"order_service/pkg/errors"

	"gorm.io/gorm"
)

type MySqlDB struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) entity.OrderRepo {

	return &MySqlDB{
		db: db,
	}
}

func (m *MySqlDB) Add(order *entity.UserOrder) error {
	const op = errors.Op("MySqlDB.Add")

	err := m.db.Create(dto.UoToDto(order)).Error
	if err != nil {
		err = errors.Wrap(err, op, errors.ErrInternal)
	}
	return err
}

func (m *MySqlDB) Get(userId, id int64) (*entity.UserOrder, error) {
	const op = errors.Op("MySqlDB.Get")

	o := &dto.Order{}
	if err := m.db.Where("id = ? and user_id = ?", id, userId).
		Preload("Deposite").Preload("Withdrawal").Preload("ExchangeOrder").
		First(o).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)

	}
	return o.ToEntity(), nil
}

func (m *MySqlDB) GetAll(UserId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("MySqlDB.GetAll")

	osDTO := []*dto.Order{}
	if err := m.db.Where("user_id = ?", UserId).Preload("Deposite").Preload("Withdrawal").Preload("ExchangeOrder").
		Find(&osDTO).Error; err != nil {
		return nil, errors.Wrap(op, err, errors.ErrInternal)
	}

	if len(osDTO) == 0 {
		return nil, errors.Wrap(op, errors.ErrNotFound)
	}

	os := []*entity.UserOrder{}
	for _, o := range osDTO {
		os = append(os, o.ToEntity())
	}
	return os, nil
}

func (m *MySqlDB) Update(order *entity.UserOrder) error {
	const op = errors.Op("MySqlDB.Update")

	if err := m.db.Save(dto.UoToDto(order)).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}
	return nil
}

func (m *MySqlDB) UpdateWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("MySqlDB.UpdateWithdrawal")

	if err := m.db.Save(dto.WToDto(w)).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}
	return nil
}
