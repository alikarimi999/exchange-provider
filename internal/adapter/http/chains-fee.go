package http

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type chainsFee struct {
	mux   *sync.Mutex
	v     *viper.Viper
	chain map[string]float64
}

func (c *chainsFee) readConfig() {
	cs := c.v.GetStringMap("chains_fee")

	chains := make(map[string]float64)
	for id, f := range cs {
		fee, ok := f.(float64)
		if !ok {
			continue
		}
		chains[id] = fee
	}
	c.update(chains)
}

func (c *chainsFee) update(chains map[string]float64) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range chains {
		c.chain[k] = v
	}
}

func (c *chainsFee) lowerEq(c1, c2 string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.chain[c1] <= c.chain[c2]
}

func (s *Server) UpdateChainsFee(ctx *gin.Context) {
	c := struct {
		Chains map[string]float64
		Msg    string
	}{}

	if err := ctx.Bind(&c); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	s.cf.update(c.Chains)

	s.v.Set("chains_fee", c.Chains)
	if err := s.cf.v.WriteConfig(); err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	c.Msg = "done"
	ctx.JSON(200, c)

}
