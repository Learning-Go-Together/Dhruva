package passwordstrength

import (
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) > 12 || len(password) == 0 {
		return false
	}
	if unicode.IsUpper(rune(password[0])) {
		return false
	}
	return true
}
