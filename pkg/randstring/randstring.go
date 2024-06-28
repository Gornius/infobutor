package randstring

import (
	"math/rand"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()-_+=")

func WithLength(length int) string {
	rstr := make([]rune, length)
	for i := range rstr {
		rstr[i] = chars[rand.Intn(len(chars))]
	}

	return string(rstr)
}
