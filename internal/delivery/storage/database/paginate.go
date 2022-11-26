package database

import (
	"exchange-provider/internal/delivery/storage/database/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"gorm.io/gorm"
)

func (m *MySqlDB) GetPaginated(p *entity.PaginatedOrders) error {
	const op = errors.Op("MySqlDB.GetAllPaging")

	osDTO := []*dto.Order{}
	if err := setClauses(m.db, p.Filters).Scopes(paginate(p.Page, p.PerPage)).
		Preload("Deposit").Preload("Withdrawal").Preload("Swaps").Find(&osDTO).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}

	for _, o := range osDTO {
		p.Orders = append(p.Orders, o.ToEntity())
	}

	var count int64
	if err := setClauses(m.db, p.Filters).Model(&dto.Order{}).Count(&count).Error; err != nil {
		return errors.Wrap(op, err, errors.ErrInternal)
	}

	p.Total = count

	return nil
}

func (m *MySqlDB) GetPaginatedByParams(page, perPage int, params map[string]string) ([]*entity.Order, error) {
	const op = errors.Op("MySqlDB.GetPaging")

	osDTO := []*dto.Order{}
	if err := m.db.Where(params).Preload("Deposit").Preload("Withdrawal").Preload("Swaps").
		Offset(page * perPage).Limit(perPage).Find(&osDTO).Error; err != nil {
		return nil, errors.Wrap(op, err, errors.ErrInternal)
	}
	res := []*entity.Order{}
	for _, o := range osDTO {
		res = append(res, o.ToEntity())
	}

	return res, nil
}

func paginate(page, perPage int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		if perPage <= 0 {
			perPage = 10
		}
		if perPage >= 100 {
			perPage = 100
		}

		return db.Offset(int((page - 1) * perPage)).Limit(int(perPage))
	}
}
