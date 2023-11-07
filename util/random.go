package helper

import (
	"math/rand"
	"strings"
	"time"
)


func init () {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer within the specified range [min, max].
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random string of the given length 'n' composed of lowercase alphabets.
func RandomString(n int) string {
	var sb strings.Builder
	letterCount := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(letterCount)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates random owner name, character is limited to 8
func RandomOwnerName() string {
	return RandomString(8)
}

// RandomMoney generates a random integer amount between 0 and 1000 (inclusive).
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random currency code from a predefined list.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "NGN", "CAD", "INR"}

	randomIndex := rand.Intn(len(currencies))
	return currencies[randomIndex]
}


func RandomDisplayPictureURL () (URL string) {
	randInt := RandomInt(2, 100)

	if randInt % 2 == 0 {
		// return empty URL
		return
	}
	return "https://"+RandomString(8) + ".com"

}