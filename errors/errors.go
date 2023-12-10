package errors

import (
	"net/http"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/language"
)

type App struct {
	Code  codes.Code `json:"code,omitempty"`
	Title string     `json:"title,omitempty"`
	Body  string     `json:"body,omitempty"`
	sys   error
}

func (a *App) Error() string {
	return a.sys.Error()
}

// Get `codes.Code` of error and will return `codes.NoCode` if error doesn't have any code
func GetCode(err error) codes.Code {
	if err, isOk := err.(*stacktrace); isOk {
		return err.code
	}
	return codes.NoCode
}

// Compile returns an error and creates new `App` errors
func Compile(err error, lang language.Language) (int, App) {
	code := GetCode(err)
	if errMsg, isOk := codes.GetCodeMessages(code)[lang]; isOk {
		return errMsg.StatusCode, App{
			Code:  code,
			Title: errMsg.Title,
			Body:  errMsg.Body,
			sys:   err,
		}
	}

	// Default error
	return http.StatusInternalServerError, App{
		Code:  code,
		Title: "Service Error Not Defined",
		Body:  "Unknown error. Please contact admin",
		sys:   err,
	}
}
