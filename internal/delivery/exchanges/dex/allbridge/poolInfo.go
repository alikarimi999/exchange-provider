package allbridge

import (
	"bytes"
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"io"
	"net/http"
)

type pools map[string]map[string]*types.PoolInfo
type poolInfoReq struct {
	Pools []pool `json:"pools"`
}
type pool struct {
	ChainSymbol string `json:"chainSymbol"`
	PoolAddress string `json:"poolAddress"`
}

func getPoolInfo(ts []*types.TokenInfo) (pools, error) {
	purl := "https://core.api.allbridgecoreapi.net/pool-info"
	ps := []pool{}
	for _, t := range ts {
		ps = append(ps, pool{
			ChainSymbol: t.Chain,
			PoolAddress: t.PoolAddress,
		})
	}

	pr := &poolInfoReq{Pools: ps}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(pr)
	req, _ := http.NewRequest(http.MethodPost, purl, &buf)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)
	pis := pools{}
	return pis, json.Unmarshal(b, &pis)
}
