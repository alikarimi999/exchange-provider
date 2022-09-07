package uniswapv3

import "fmt"

func (u *UniSwapV3) agent(fn string) string {
	return fmt.Sprintf("%s-%s", u.NID(), fn)
}
