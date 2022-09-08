package uniswapv3

import (
	"fmt"
	"time"
)

func (u *UniSwapV3) Stop() {
	op := fmt.Sprintf("%s.Stop", u.NID())
	close(u.stopCh)
	u.stoppedAt = time.Now()
	u.l.Debug(string(op), "stopped")
}

func (u *UniSwapV3) Conifgs() interface{} {
	return u.cfg
}
