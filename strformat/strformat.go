package strformat

import (
	"bytes"
	"regexp"
	"text/template"
)

// Check string value, will return `true` if string value is only number and return `false` if string value not only contains number
func IsOnlyNumber(s string) bool {
	pattern := "^[0-9]+$"
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}
	return matched
}

func Tmpl(textTemplate string, values any) (string, error) {
	tmpl, err := template.New("").Parse(textTemplate)
	if err != nil {
		return "", err
	}

	buff := bytes.Buffer{}
	if err := tmpl.Execute(&buff, values); err != nil {
		return "", err
	}

	return buff.String(), nil
}
