package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	var length = len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(length)]
		sb.WriteByte(c)
	}

	return sb.String()

}

func RandomName() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}

	var k = len(currencies)

	return currencies[rand.Intn(k)]
}
