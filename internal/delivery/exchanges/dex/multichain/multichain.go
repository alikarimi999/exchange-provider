package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"time"
)

type Multichain struct {
	cfg *Config

	cs map[chainId]*Chain

	tt     *utils.TxTracker
	pairs  *supportedPairs
	apiUrl string
	l      logger.Logger
}

func NewMultichain(cfg *Config, l logger.Logger) (entity.Exchange, error) {

	m := &Multichain{
		cfg:    cfg,
		cs:     make(map[chainId]*Chain),
		apiUrl: "https://bridgeapi.anyswap.exchange/v2/history/details?params=",
		l:      l,
	}
	m.tt = utils.NewTxTracker(m.NID(), time.Duration(15)*time.Second, 1, l)
	m.pairs = newSupportedPairs()

	return m, nil
}

// func readFile() (map[chainId]*tokens, error) {

// 	var b []byte
// 	f, err := os.Open("./tokenlistv4.json")
// 	if err != nil {
// 		res, err := http.Get("https://bridgeapi.anyswap.exchange/v4/tokenlistv4/all")
// 		if err != nil {
// 			return nil, err
// 		}

// 		b, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			return nil, err
// 		}

// 		file, err := os.Create("./tokenlistv4.json")
// 		if err == nil {
// 			file.Write(b)
// 			file.Close()
// 		}
// 	} else {
// 		b, err = io.ReadAll(f)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	cs := make(map[chainId]*tokens)

// 	json.Unmarshal(b, &cs)
// 	return cs, nil

// }
