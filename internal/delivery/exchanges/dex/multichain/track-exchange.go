package multichain

import "exchange-provider/internal/entity"

func (m *Multichain) TrackSwap(o *entity.CexOrder,
	index int, done chan<- struct{}, proccessed <-chan bool) {
	m.trackSwap(o, index)
	done <- struct{}{}
	<-proccessed
}
