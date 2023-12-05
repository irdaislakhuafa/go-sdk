package strformat

import "regexp"

func IsOnlyNumber(s string) bool {
	pattern := "^[0-9]+$"
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}
	return matched
}
