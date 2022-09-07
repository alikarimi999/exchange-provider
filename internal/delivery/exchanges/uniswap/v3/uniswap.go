package uniswapv3

import (
	"context"
	"math/big"
	"order_service/pkg/logger"
	"order_service/pkg/wallet/eth"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Configs struct {
	Wallet          *eth.HDWallet
	Providers       []*ethclient.Client
	DefaultProvider *ethclient.Client
	ConfirmBlocks   uint64
}

type UniSwapV3 struct {
	mux *sync.Mutex

	accountId string

	confirms  uint64
	blockTime time.Duration
	chainId   *big.Int

	dp     *ethclient.Client //default provider
	ps     []*ethclient.Client
	wallet *eth.HDWallet

	tokens *supportedTokens
	pairs  *supportedPairs

	tt *txTracker
	dt *depostiTracker

	rc *redis.Client
	v  *viper.Viper
	l  logger.Logger

	graphUrl string

	stopCh chan struct{}
}

func NewExchange(cfg *Configs, rc *redis.Client, v *viper.Viper,
	l logger.Logger, readConfig bool) (*UniSwapV3, error) {

	// const op = errors.Op("UniSwapV3.NewExchange")

	v3 := &UniSwapV3{
		mux: &sync.Mutex{},

		accountId: hash(hash(cfg.Wallet.Mnemonic())),

		confirms:  cfg.ConfirmBlocks,
		blockTime: time.Duration(20 * time.Second),
		dp:        cfg.DefaultProvider,
		ps:        cfg.Providers,
		wallet:    cfg.Wallet,

		tokens: newSupportedTokens(),
		pairs:  newSupportedPairs(),

		rc: rc,
		v:  v,
		l:  l,

		graphUrl: "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3",
		stopCh:   make(chan struct{}),
	}

	chainId, err := v3.dp.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	v3.chainId = chainId
	v3.tt = newTxTracker(v3)
	v3.dt = newDepositTracker(v3)

	return v3, nil
}

func (u *UniSwapV3) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	w := &sync.WaitGroup{}

	w.Add(1)
	go u.tt.run(w, u.stopCh)

	w.Add(1)
	go u.dt.run(w, u.stopCh)

}
