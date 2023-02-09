package uniswapV3

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	em "github.com/ethereum/go-ethereum/common/math"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
)

var windowsSize = 100

func (d *dex) Prices(ps []*entity.Pair) error {
	agent := d.agent("Prices")

	d.l.Debug(agent, fmt.Sprintf("Updating price for %d pairs", len(ps)))

	wg := &sync.WaitGroup{}
	guard := make(chan struct{}, 10)
	count := atomic.NewUint64(0)
	round := 0
	t := time.Now()
	for {
		round++
		start := (round - 1) * windowsSize
		end := round * windowsSize
		if end > len(ps) {
			end = len(ps)
		}

		guard <- struct{}{}
		wg.Add(1)
		go func(start, end, round int) {
			defer func() {
				wg.Done()
				<-guard
			}()

			con, _ := contracts.NewContracts(d.contract, d.provider())
			inputs := []contracts.IPriceAggregatorpriceIn{}
			for i, p := range ps[start:end] {
				inputs = append(inputs, contracts.IPriceAggregatorpriceIn{
					Index:           big.NewInt(int64(i)),
					T0:              common.HexToAddress(p.T1.Address),
					T1:              common.HexToAddress(p.T2.Address),
					Provider:        d.factory,
					ProviderVersion: 3,
				})
			}
			// d.l.Debug(agent, fmt.Sprintf("Downloading price for %d pairs, round %d", len(inputs), round))

			outputs, err := con.GetPrices(nil, inputs)
			if err != nil {
				d.l.Error(agent, err.Error())
				return
			}

			// inCount := 0
			for _, out := range outputs {
				if out.Price.Int64() != 0 {
					// inCount++
					count.Add(1)
					p := ps[start+int(out.Index.Int64())]
					p1f := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(out.Price), big.NewFloat(0).SetInt(em.BigPow(10, int64(p.T1.Decimals))))
					p.Price1 = p1f.Text('f', 10)
					p.Price2 = big.NewFloat(0).Quo(big.NewFloat(1), p1f).Text('f', 10)
					p.FeeTier = out.Fee.Int64()
				}
			}
			// d.l.Debug(agent, fmt.Sprintf("The price of %d pairs was downloaded, round %d", inCount, round))

		}(start, end, round)
		if end == len(ps) {
			break
		}
	}
	wg.Wait()
	d.l.Debug(agent, fmt.Sprintf("The price of %d pairs updated in %v", count.Load(), time.Since(t)))
	return nil
}

// Not Ready
func (d *dex) SaveAvailablePairs(ps []types.Pair, file string) {
	agent := d.agent("SaveAvailablePairs")

	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		// create config file if not exists
		if err := v.WriteConfigAs(file); err != nil {
			panic(err)
		}
	}

	con, err := contracts.NewContracts(d.contract, d.provider())
	if err != nil {
		d.l.Error(agent, err.Error())
		panic(err)
	}

	fmt.Println("len: ", len(ps))
	round := 0
	wg := &sync.WaitGroup{}
	gurad := make(chan struct{}, 100)
	for {
		round++
		start := (round - 1) * windowsSize
		end := round * windowsSize
		if end > len(ps) {
			end = len(ps)
		}
		gurad <- struct{}{}

		wg.Add(1)
		go func(start, end, round int) {
			defer func() {
				wg.Done()
				<-gurad
			}()

			inputs := []contracts.IPriceAggregatorexistsIn{}
			for i, p := range ps[start:end] {
				inputs = append(inputs, contracts.IPriceAggregatorexistsIn{
					Index:           big.NewInt(int64(i)),
					T0:              p.T1.Address,
					T1:              p.T2.Address,
					Provider:        d.factory,
					ProviderVersion: 3,
					Min0:            em.BigPow(10, int64(p.T1.Decimals)),
					Min1:            em.BigPow(10, int64(p.T2.Decimals)),
				})
			}
			d.l.Debug(agent, fmt.Sprintf("Checking the number of available pairs out of %d pairs,round %d", len(inputs), round))
			outs, err := con.PoolsExists(nil, inputs)
			if err != nil {
				d.l.Error(agent, err.Error())
				panic(err)
			}

			count := 0
			tps := []types.Pair{}
			for _, out := range outs {
				if out.Exists {
					p := ps[start+int(out.Index.Int64())]
					count++
					tps = append(tps, p)
					v.Set(fmt.Sprintf("Pairs.%s", p.String()), p)
				}
			}
			v.WriteConfig()
			d.l.Debug(agent, fmt.Sprintf("There are %d pairs available in round %d", count, round))

		}(start, end, round)
		if end == len(ps) {
			break
		}
	}
	wg.Wait()

}
