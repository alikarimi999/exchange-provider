package database

import (
	"fmt"
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

	od := dto.UoToDto(order)
	// get last seq
	var lastSeq int64
	if err := m.db.Model(&dto.Order{}).Select("seq").Where("user_id = ?", order.UserId).Order("seq desc").Limit(1).Scan(&lastSeq).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}
	fmt.Println(lastSeq)
	od.Seq = lastSeq + 1
	err := m.db.Create(od).Error
	if err != nil {
		err = errors.Wrap(err, op, errors.ErrInternal)
	}
	order.Id = int64(od.ID)
	order.Seq = od.Seq
	order.Deposit.Id = od.Deposit.Id
	order.Deposit.OrderId = order.Id
	order.Withdrawal.Id = od.Withdrawal.Id
	order.Withdrawal.OrderId = int64(od.ID)
	order.ExchangeOrder.Id = od.ExchangeOrder.Id
	order.ExchangeOrder.OrderId = int64(od.ID)
	return err
}

func (m *MySqlDB) Get(userId, id int64) (*entity.UserOrder, error) {
	const op = errors.Op("MySqlDB.Get")

	o := &dto.Order{}
	if err := m.db.Where("id = ? and user_id = ?", id, userId).
		Preload("Deposit").Preload("Withdrawal").Preload("ExchangeOrder").
		First(o).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)

	}
	return o.ToEntity(), nil
}

func (m *MySqlDB) GetBySeq(uId, seq int64) (*entity.UserOrder, error) {
	const op = errors.Op("MySqlDB.GetBySeq")

	o := &dto.Order{}
	if err := m.db.Where("seq = ? and user_id = ?", seq, uId).
		Preload("Deposit").Preload("Withdrawal").Preload("ExchangeOrder").
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
	if err := m.db.Where("user_id = ?", UserId).Preload("Deposit").Preload("Withdrawal").Preload("ExchangeOrder").
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

	if err := m.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(dto.UoToDto(order)).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}
	return nil
}

func (m *MySqlDB) UpdateDeposit(d *entity.Deposit) error {
	dd := dto.DToDto(d)
	return m.db.Save(dd).Error
}

func (m *MySqlDB) UpdateWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("MySqlDB.UpdateWithdrawal")

	if err := m.db.Save(dto.WToDto(w)).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}
	return nil
}

// check if any deposit has this tx_id
func (m *MySqlDB) CheckTxId(txId string) (bool, error) {
	const op = errors.Op("MySqlDB.CheckTxId")

	o := &dto.Deposit{}
	if err := m.db.Where("tx_id = ?", txId).First(o).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, errors.Wrap(op, err, errors.ErrInternal)
	}
	fmt.Println(o.Id)
	return true, nil
}
