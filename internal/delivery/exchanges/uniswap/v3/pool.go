package uniswapv3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type pool struct {
	id string
	t0 *token
	t1 *token

	t0Price string
	t1Price string
	feeTier int
}

func (u *UniSwapV3) bestPool(t0, t1 common.Address) (*pool, error) {
	j := map[string]string{
		"query": fmt.Sprintf(`
    {
    pools
        (
            where:{token0:"%s",
                    token1:"%s"
            },
            orderBy:totalValueLockedUSD
            orderDirection:desc
            first:1
        
        )
        {
        id
        token0{id,decimals}
        token1{id,decimals}
		token0Price
        token1Price
		feeTier
    }
    }
        `, t0, t1),
	}

	b, _ := json.Marshal(j)
	req, _ := http.NewRequest("POST", u.graphUrl, bytes.NewBuffer(b))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := GetBestPoolGraph{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	if len(res.Data.Pools) == 0 {
		return nil, fmt.Errorf("pool not found")
	}

	return res.Data.Pools[0].ToPool()
}
