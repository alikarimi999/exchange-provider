package app

import "order_service/internal/entity"

func (o *OrderUseCase) AddCoins(coins map[string]*entity.Coin) {
	o.supportedCoins.add(coins)

	chains := []*entity.Chain{}
	for _, coin := range coins {
		chains = append(chains, coin.Chain)
	}
	o.wh.addChainTickers(chains)

}

func (o *OrderUseCase) RemoveCoin(coinId, chainId string) {
	o.supportedCoins.remove(coinId, chainId)
}

func (o *OrderUseCase) RemoveChain(chainId string) {
	o.wh.removeTicker(chainId)
}

func (o *OrderUseCase) Supported(coinId, chainId string) bool {
	return o.supportedCoins.exist(coinId, chainId) && o.wh.isTickerRunning(chainId)

}

func (o *OrderUseCase) GetAllSupportedCoins() []*entity.Coin {
	return o.supportedCoins.snapshots()
}
