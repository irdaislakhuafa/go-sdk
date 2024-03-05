package errors

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/language"
	"github.com/irdaislakhuafa/go-sdk/operator"
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

func shortFuncName(f *runtime.Func) string {
	// f.Name() is like one of these:
	// - "github.com/irdaislakhuafa/go-sdk/<package>.<FuncName>"
	// - "github.com/irdaislakhuafa/go-sdk/<package>.<Receiver>.<MethodName>"
	// - "github.com/irdaislakhuafa/go-sdk/<package>.<*PtrReceiver>.<MethodName>"
	longName := f.Name()
	withoutPath := longName[strings.LastIndex(longName, "/")+1:]
	withoutPackage := withoutPath[strings.LastIndex(withoutPath, ".")+1:]

	shortName := withoutPackage
	shortName = strings.Replace(shortName, "(", "", 1)
	shortName = strings.Replace(shortName, "*", "", 1)
	shortName = strings.Replace(shortName, ")", "", 1)

	return shortName
}

func create(cause error, code codes.Code, msg string, val ...any) error {
	if code == codes.NoCode {
		code = GetCode(cause)
	}

	err := &stacktrace{
		message:  fmt.Sprintf(msg, val...),
		cause:    cause,
		code:     code,
		file:     "",
		function: "",
		line:     0,
	}

	pc, file, line, isOk := runtime.Caller(1)
	if !isOk {
		return err
	}

	err.file, err.line = file, line

	function := runtime.FuncForPC(pc)
	if function == nil {
		return err
	}

	err.function = shortFuncName(function)

	return err
}

func GetCaller(cause error) (file string, line int, message string, err error) {
	if st, isOk := cause.(*stacktrace); isOk {
		return st.file, st.line, st.message, nil
	} else {
		return "", 0, "", create(nil, codes.NoCode, operator.Ternary(cause == nil, "failed to cast error to stacktrace", cause.Error()))
	}
}

func NewWithCode(code codes.Code, msg string, val ...any) error {
	return create(nil, code, msg, val...)
}
