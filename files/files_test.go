package files

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func Test_IsExist(t *testing.T) {
	type test struct {
		name          string
		argPathToFile string
		runBeforeExec func(pathIntoFile string)
		runAfterExec  func(pathIntoFile string)
		wantResult    bool
	}
	tests := []test{
		{
			name:          "test file not exist",
			argPathToFile: "/tmp/file.txt",
			wantResult:    false,
		},
		{
			name:          "test file is exist",
			argPathToFile: "/tmp/file.txt",
			wantResult:    true,
			runBeforeExec: func(pathIntoFile string) {
				dirs := filepath.Dir(pathIntoFile)
				_ = os.MkdirAll(dirs, fs.FileMode(0777))
				_, _ = os.Create(pathIntoFile)
			},
			runAfterExec: func(pathIntoFile string) {
				_ = os.RemoveAll(pathIntoFile)
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v: %v", GetCurrentMethodName(), tt.name), func(t *testing.T) {
			if tt.runBeforeExec != nil {
				tt.runBeforeExec(tt.argPathToFile)
			}

			result := IsExist(tt.argPathToFile)
			if result != tt.wantResult {
				t.Errorf("error: want result is '%v' but got '%v'\n", tt.wantResult, result)
			}

			if tt.runAfterExec != nil {
				tt.runAfterExec(tt.argPathToFile)
			}
		})
	}
}

func Test_GetFileExtenstion(t *testing.T) {
	type test struct {
		name       string
		arg        string
		wantResult string
	}
	tests := []test{
		{
			name:       "test file csv",
			arg:        "file.csv",
			wantResult: "csv",
		},
		{
			name:       "test file csv with path",
			arg:        "/path/to/file.csv",
			wantResult: "csv",
		},
		{
			name:       "test file without extension",
			arg:        "file",
			wantResult: "",
		},
		{
			name:       "test file without extension with path",
			arg:        "/path/to/file",
			wantResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFileExtenstion(tt.arg)
			if result != tt.wantResult {
				t.Fatalf("error: want result is '%v' but got '%v'", tt.wantResult, result)
			}
		})
	}
}

func Test_GetCurrentFileLocation(t *testing.T) {
	type test struct {
		name       string
		wantResult string
	}

	tests := []test{
		{
			name:       "test get current file location",
			wantResult: "/media/Projects/Golang/go-sdk/files/files_test.go", // NOTE: or replace with this test file location on your local PC
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetCurrentFileLocation()
			if result != tt.wantResult {
				t.Fatalf("want result is '%v' but got '%v'", tt.wantResult, result)
			}
		})
	}
}
