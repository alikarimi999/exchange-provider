package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Multichain struct {
	cfg *Config

	cs map[ChainId]*Chain

	tt     *utils.TxTracker
	pairs  *supportedPairs
	apiUrl string

	v *viper.Viper
	l logger.Logger
}

func NewMultichain(cfg *Config, v *viper.Viper, l logger.Logger, readConfigs bool) (entity.Exchange, error) {

	m := &Multichain{
		cfg:    cfg,
		cs:     make(map[ChainId]*Chain),
		apiUrl: "https://bridgeapi.anyswap.exchange/v2/history/details?params=",
		v:      v,
		l:      l,
	}
	m.tt = utils.NewTxTracker(m.Id(), time.Duration(15)*time.Second, 1, l)
	m.pairs = newSupportedPairs()

	if readConfigs {
		i := m.v.Get(fmt.Sprintf("%s.chains", m.Id()))
		if i != nil {
			fmt.Println(i)
			csi := i.(map[string]interface{})
			for id, v := range csi {
				is := v.([]interface{})
				urls := []string{}
				for _, v := range is {
					urls = append(urls, v.(string))
				}
				m.newChain(ChainId(id), urls...)
			}
		}
	}

	return m, nil
}

func (m *Multichain) Id() string {
	return m.cfg.Name
}
