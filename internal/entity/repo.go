package entity

type OrderRepo interface {
	Add(order *UserOrder) error
	Get(userId, id int64) (*UserOrder, error)
	GetAll(userId int64) ([]*UserOrder, error)
	Update(order *UserOrder) error
	UpdateWithdrawal(w *Withdrawal) error
}
