package files

import "os"

// to check is file (not dir) exists, will return true if exist ant return false if not file (dir) or not exist
func IsExist(pathToFile string) bool {
	fileInfo, err := os.Stat(pathToFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return !fileInfo.IsDir()
}
