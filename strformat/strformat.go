package strformat

import (
	"bytes"
	"regexp"
	"text/template"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
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

func Tmpl(tmplFormat string, values any) (string, error) {
	tmpl, err := template.New("").Parse(tmplFormat)
	if err != nil {
		return "", errors.NewWithCode(codes.CodeStrTemplateInvalidFormat, "cannot parse str format, %v", err)
	}

	buff := bytes.Buffer{}
	if err := tmpl.Execute(&buff, values); err != nil {
		return "", errors.NewWithCode(codes.CodeStrTemplateExecuteErr, "cannot execute template, %v", err)
	}

	return buff.String(), nil
}
