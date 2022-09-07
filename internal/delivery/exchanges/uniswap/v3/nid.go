package uniswapv3

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func (u *UniSwapV3) Name() string {
	return "uniswapv3"
}

func (u *UniSwapV3) AccountId() string {
	u.mux.Lock()
	defer u.mux.Unlock()
	return u.accountId
}

func (u *UniSwapV3) NID() string {
	id := u.AccountId()
	n := u.Name()
	return fmt.Sprintf("%s-%s", n, id)
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
