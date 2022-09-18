package http

import (
	"io"
	"net/http"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"sync"

	"github.com/spf13/viper"
)

type authConf struct {
	Url     string `json:"url"`
	CAP     string `json:"check_access_path"`
	CheckIP bool   `json:"check_ip"`
}

type authService struct {
	mux *sync.Mutex
	v   *viper.Viper
	c   *authConf
}

func newAuthService(v *viper.Viper) *authService {
	return &authService{
		mux: &sync.Mutex{},
		v:   v,
		c: &authConf{
			Url:     v.GetString("auth_service.url"),
			CAP:     v.GetString("auth_service.check_access_path"),
			CheckIP: v.GetBool("auth_service.check_ip"),
		},
	}
}

func (a *authService) Cofigs() interface{} {
	a.mux.Lock()
	defer a.mux.Unlock()
	return &authConf{
		Url:     a.c.Url,
		CAP:     a.c.CAP,
		CheckIP: a.c.CheckIP,
	}
}

func (a *authService) changeConfigs(c *authConf) (err error) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if c.Url != "" && c.CAP != "" {
		if a.c.Url == c.Url && a.c.CAP == c.CAP && a.c.CheckIP == c.CheckIP {
			return errors.Wrap(errors.NewMesssage("not changed"))
		}
		a.v.Set("auth_service", c)
		if err := a.v.WriteConfig(); err != nil {
			return errors.Wrap(errors.NewMesssage(err.Error()))
		}
		a.c = &authConf{
			Url:     c.Url,
			CAP:     c.CAP,
			CheckIP: c.CheckIP,
		}
		return nil
	}
	return errors.Wrap(errors.NewMesssage("auth service url or path is empty"))
}

func (a *authService) request(body io.Reader) (*http.Request, error) {
	a.mux.Lock()
	defer a.mux.Unlock()
	u := a.c.Url
	p := a.c.CAP

	if u == "" || p == "" {
		return nil, errors.Wrap(errors.NewMesssage("auth service url or path is empty"))
	}

	req, err := http.NewRequest("POST", utils.JoinSafe(u, p), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (a *authService) checkIP() bool {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.c.CheckIP
}
