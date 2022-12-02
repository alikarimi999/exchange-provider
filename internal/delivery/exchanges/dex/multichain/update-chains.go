package multichain

import "fmt"

const (
	addMsg = "added successfully!"
)

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
	agent := "multichain.updateChains"

	res := UpdateChainsRes{
		Result: make(map[ChainId]struct {
			Msg  string
			Urls []URL
		}),
	}

	for k, v := range req.Chains {
		urls := []URL{}
		var err error
		c, ok := m.cs[k]
		if !ok {
			_, urls, err = m.newChain(k, v...)
			if err != nil {
				res.Result[k] = struct {
					Msg  string
					Urls []URL
				}{Msg: err.Error(), Urls: urls}
				continue
			}
		} else {
			for _, url := range v {
				if err := c.addProvider(url); err != nil {
					urls = append(urls, URL{Url: url, Msg: err.Error()})
				} else {
					urls = append(urls, URL{Url: url, Msg: err.Error()})
				}
			}
		}

		added := []string{}
		for _, url := range urls {
			if url.Msg == addMsg {
				added = append(added, url.Url)
			}
		}
		m.v.Set(fmt.Sprintf("%s.chains.%s", m.Id(), k), added)
		res.Result[k] = struct {
			Msg  string
			Urls []URL
		}{Msg: addMsg, Urls: urls}
	}

	if err := m.v.WriteConfig(); err != nil {
		m.l.Error(agent, err.Error())
	}

	return res, nil
}
