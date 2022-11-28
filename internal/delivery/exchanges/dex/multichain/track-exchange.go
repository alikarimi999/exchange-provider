package multichain

import "exchange-provider/internal/entity"

func (m *Multichain) TrackExchangeOrder(o *entity.Order,
	index int, done chan<- struct{}, proccessed <-chan bool) {
	m.trackSwap(o, index)
	done <- struct{}{}
	<-proccessed
}
