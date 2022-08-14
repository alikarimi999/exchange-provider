package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// create order limiter

func (r *Router) NewLimiters() {
	conf := &LimiterConfig{}

	conf.Max = r.v.GetUint64("general_limiter.max")
	conf.Period = r.v.GetDuration("general_limiter.period")

	if conf.Max <= 0 || conf.Period <= 0 {
		conf.Max = uint64(defaultLimiterMax)
		conf.Period = time.Duration(defaultLimiterPeriod)
	}

	r.gls = newGeneralLimiters(conf)

	conf.Max = r.v.GetUint64("create_order_limiter.max")
	conf.Period = r.v.GetDuration("create_order_limiter.period")

	if conf.Max <= 0 || conf.Period <= 0 {
		conf.Max = uint64(defaultLimiterMax)
		conf.Period = time.Duration(defaultLimiterPeriod)
	}

	r.col = newLimiter(conf)
}

func (r *Router) changeLimitersConf(ctx *gin.Context) {
	req := struct {
		GL struct {
			Max    uint64 `json:"max"`
			Period string `json:"period"`
			Msg    string `json:"message,omitempty"`
		} `json:"general_limiter"`
		Col struct {
			Max    uint64 `json:"max"`
			Period string `json:"period"`
			Msg    string `json:"message,omitempty"`
		} `json:"create_order_limiter"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if req.GL.Max > 0 {
		gp, err := toTime(req.GL.Period)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if req.GL.Max == r.gls.conf.Max && gp == r.gls.conf.Period {
			req.GL.Msg = "no changes"
		} else {
			conf := &LimiterConfig{
				Max:    req.GL.Max,
				Period: gp,
			}
			r.v.Set("general_limiter", conf)
			if err := r.v.WriteConfig(); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			r.gls.changeConfigs(conf)
			req.GL.Msg = "configs changed"
		}
	}

	if req.Col.Max > 0 {
		cp, err := toTime(req.Col.Period)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if req.Col.Max == r.col.conf.Max && cp == r.col.conf.Period {
			req.Col.Msg = "no changes"
		} else {

			conf := &LimiterConfig{
				Max:    req.Col.Max,
				Period: cp,
			}
			r.v.Set("create_order_limiter", conf)
			if err := r.v.WriteConfig(); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			r.col.ChangeConfigs(&LimiterConfig{
				Max:    req.Col.Max,
				Period: cp,
			})
			req.Col.Msg = "configs changed"
		}
	}

	ctx.JSON(http.StatusOK, req)
}

func toTime(t string) (time.Duration, error) {
	return time.ParseDuration(t)

}
