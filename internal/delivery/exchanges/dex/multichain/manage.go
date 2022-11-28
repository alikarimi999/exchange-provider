package multichain

import (
	"exchange-provider/internal/entity"
	"sync"
)

func (*Multichain) Type() entity.ExType {
	return entity.DEX
}

func (*Multichain) Stop()

func (*Multichain) StartAgain() (*entity.StartAgainResult, error)

func (*Multichain) Command(entity.Command) (entity.CommandResult, error)

func (*Multichain) Run(wg *sync.WaitGroup)

func (m *Multichain) Configs() interface{} {
	return m.cfg
}
