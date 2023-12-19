package strformat

import "regexp"

// Check string value, will return `true` if string value is only number and return `false` if string value not only contains number
func IsOnlyNumber(s string) bool {
	pattern := "^[0-9]+$"
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}
	return matched
}
