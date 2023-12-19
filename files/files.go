package files

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// To check is file (not dir) exists, will return true if exist ant return false if not file (dir) or not exist
func IsExist(pathToFile string) bool {
	fileInfo, err := os.Stat(pathToFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return !fileInfo.IsDir()
}

// Get file extension (ex. "txt, csv, docs, json, yaml")
func GetFileExtenstion(pathOrFileName string) string {
	fileName := filepath.Base(pathOrFileName)
	splitedName := strings.Split(fileName, ".")

	if len(splitedName) <= 1 {
		return ""
	}

	fileExt := splitedName[len(splitedName)-1]
	return fileExt
}

// Will return string of current file location where this function is called.
func GetCurrentFileLocation() string {
	_, file, _, isOk := runtime.Caller(1)
	if isOk {
		return file
	}
	return ""
}

// Will return string of current method location where this function is called.
func GetCurrentMethodName() string {
	pc, _, _, isOk := runtime.Caller(1)
	if !isOk {
		return ""
	}

	f := runtime.FuncForPC(pc)
	return fmt.Sprintf("%v()", f.Name())
}
