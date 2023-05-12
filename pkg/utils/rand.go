package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

func RandString(l int) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	c := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, l)
	for i := range b {
		b[i] = c[r.Intn(len(c))]
	}

	return string(b)
}

func Hash(id string) string {
	h := sha256.New()
	h.Write([]byte(id))
	return hex.EncodeToString(h.Sum(nil))
}
