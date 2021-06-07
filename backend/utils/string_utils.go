package utils

import (
	"math/rand"
	"net/mail"
	"unicode"
)

// ValidPassword validates plain password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// tot: at least eight characters long.
// No empty string or whitespace.
func ValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		upp, low, num, sym bool
		total              uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			total++
		case unicode.IsLower(char):
			low = true
			total++
		case unicode.IsNumber(char):
			num = true
			total++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			total++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || total < 8 {
		return false
	}

	return true
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
