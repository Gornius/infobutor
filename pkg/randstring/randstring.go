package randstring

import (
	"math/rand"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()-_+=")

// creates random string with a given length
func WithLength(length int) string {
	rstr := make([]rune, length)
	for i := range rstr {
		rstr[i] = chars[rand.Intn(len(chars))]
	}

	return string(rstr)
}
