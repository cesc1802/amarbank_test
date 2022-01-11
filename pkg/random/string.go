package random

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var alphaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")

func Random(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumeric[rand.Intn(99999999)%len(alphaNumeric)]
	}
	return string(b)
}
