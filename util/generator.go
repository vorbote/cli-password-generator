package util

import (
	"math/rand"
	"time"
)

const (
	Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lower = "abcdefghijklmnopqrstuvwxyz"
	Digit = "0123456789"
)

func GeneratePassword(length int, useDigit, useLower, useUpper bool, special string) string {
	// Set rand seed
	rand.Seed(time.Now().UnixNano())

	passwd := make([]byte, length, length)

	// Set the characters to generate a password
	var sourceStr string

	if !useDigit && !useLower && !useUpper && len(special) == 0 {
		panic("You have no available characters to use.")
	}

	// If useDigit is true, then add Digit to its source
	if useDigit {
		sourceStr += Digit
	}

	// If useLower is true, then add Lower to its source
	if useLower {
		sourceStr += Lower
	}

	// If useUpper is true, then add Upper to its source
	if useUpper {
		sourceStr += Upper
	}

	// If special is not empty string, then add special to its source
	if len(special) != 0 {
		sourceStr += special
	}

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}
