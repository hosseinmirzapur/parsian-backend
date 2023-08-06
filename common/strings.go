package common

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func CheckPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !HasLetter(password) {
		return false
	}

	if !HasDigits(password) {
		return false
	}

	if !HasLower(password) {
		return false
	}

	if !HasUpper(password) {
		return false
	}

	return true
}

func GenerateOtp() int {
	seed := time.Now().UnixMicro()

	genFromSeed := rand.New(rand.NewSource(seed))

	randNum := (genFromSeed.NormFloat64() * 100000) - 1

	desiredInt := int(math.Abs((randNum)))

	if len(fmt.Sprint(desiredInt)) < 6 {
		desiredInt *= 10
	}

	return desiredInt
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// To snake case : CountryId -> country_id
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
