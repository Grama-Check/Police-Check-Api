package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random id with a specific length
func RandomId(n int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Random NIC
func RandomNic() string {
	return RandomId(9)
}

// Random Name
func RandomName() string {
	return RandomString(6)
}

// Random Address
func RandomAddress() string {
	return RandomString(15)
}

// Random Clear State
func RandomClear() string {
	i := rand.Intn(2-0) + 0

	if i == 1 {
		return "clear"
	} else {
		return "not-clear"
	}
}
