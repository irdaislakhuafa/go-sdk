package files

import (
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
		t.Run(tt.name, func(t *testing.T) {
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
