package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandInt64(l int) int64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	c := []rune("0123456789")
	b := make([]rune, l)
	for i := range b {
		b[i] = c[r.Intn(len(c))]
	}

	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n

}
