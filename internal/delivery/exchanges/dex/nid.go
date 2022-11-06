package dex

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func (u *dex) Name() string {
	return u.cfg.Name
}

func (u *dex) AccountId() string {
	return u.accountId
}

func (u *dex) NID() string {
	return fmt.Sprintf("%s-%s-%s", u.Name(), u.cfg.Network, u.AccountId())
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func accountId(id string) string {
	return hash(hash(id))
}
