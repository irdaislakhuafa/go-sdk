package files

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestIsExist(t *testing.T) {
	type args struct {
		pathToFile string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test file not exist",
			args: args{
				pathToFile: "./file_not_exist.go",
			},
			want: false,
		},
		{
			name: "test file exist",
			args: args{
				pathToFile: "./files.go",
			},
			want: true,
		},
		{
			name: "test path is directory",
			args: args{
				pathToFile: "../files",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.args.pathToFile); got != tt.want {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileExtenstion(t *testing.T) {
	tests := []struct {
		name       string
		arg        string
		wantResult string
	}{
		{
			name:       "test file with txt extension without path",
			arg:        "file.txt",
			wantResult: "txt",
		},
		{
			name:       "test file with txt extension with path",
			arg:        "/path/to/file.txt",
			wantResult: "txt",
		},
		{
			name:       "test file with multi dot extension without path",
			arg:        "file.tar.gz",
			wantResult: "gz",
		},
		{
			name:       "test file with multi dot extension with path",
			arg:        "/path/to/file.tar.gz",
			wantResult: "gz",
		},
		{
			name:       "test file without extension without path",
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
	t.Run("test get current file location", func(t *testing.T) {
		result := GetCurrentFileLocation()
		if !strings.HasSuffix(filepath.ToSlash(result), "files/files_test.go") {
			t.Fatalf("want result to end with 'files/files_test.go' but got '%v'", result)
		}
	})
}
