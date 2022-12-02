package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"time"
)

type Multichain struct {
	cfg *Config

	cs map[ChainId]*Chain

	tt     *utils.TxTracker
	pairs  *supportedPairs
	apiUrl string
	l      logger.Logger
}

func NewMultichain(cfg *Config, l logger.Logger) (entity.Exchange, error) {

	m := &Multichain{
		cfg:    cfg,
		cs:     make(map[ChainId]*Chain),
		apiUrl: "https://bridgeapi.anyswap.exchange/v2/history/details?params=",
		l:      l,
	}
	m.tt = utils.NewTxTracker(m.NID(), time.Duration(15)*time.Second, 1, l)
	m.pairs = newSupportedPairs()

	return m, nil
}
