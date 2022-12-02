package multichain

type UpdateChainsReq struct {
	Chains map[ChainId][]string
}

type URL struct {
	Url string
	Msg string
}

type UpdateChainsRes struct {
	Result map[ChainId]struct {
		Msg  string
		Urls []URL
	}
}

func (m *Multichain) updateChains(req *UpdateChainsReq) (interface{}, error) {
	res := UpdateChainsRes{
		Result: make(map[ChainId]struct {
			Msg  string
			Urls []URL
		}),
	}
	for k, v := range req.Chains {
		var c *Chain
		var err error
		c, ok := m.cs[k]
		if !ok {
			c, err = m.newChain(k)
			if err != nil {
				res.Result[k] = struct {
					Msg  string
					Urls []URL
				}{Msg: err.Error()}
				continue
			}
		}

		urls := []URL{}
		for _, url := range v {
			if err := c.addProvider(url); err != nil {
				urls = append(urls, URL{Url: url, Msg: err.Error()})
				continue
			}
			urls = append(urls, URL{Url: url, Msg: "added successfully"})
		}
		res.Result[k] = struct {
			Msg  string
			Urls []URL
		}{Urls: urls}

	}

	return res, nil
}
