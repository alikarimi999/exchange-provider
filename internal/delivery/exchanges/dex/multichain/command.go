package multichain

import (
	"exchange-provider/internal/entity"
)

const (
	UpdateChains string = "update_chains"
)

func (m *Multichain) Command(cs entity.Command) (entity.CommandResult, error) {
	for k, v := range cs {
		switch k {
		case UpdateChains:
			return m.updateChains(v.(*UpdateChainsReq))
		}
	}
	return nil, nil
}
