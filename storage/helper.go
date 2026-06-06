package storage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/files"
	"github.com/irdaislakhuafa/go-sdk/header"
)

func NewBufferFromMultipart(params multipart.FileHeader) (*bytes.Buffer, error) {
	f, err := params.Open()
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeStorageNoFile, "%s", err.Error())
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeStorageNoFile, "%s", err.Error())
	}

	return bytes.NewBuffer(b), nil
}

func NewPutParamsFromMultipart(dir string, params multipart.FileHeader) (PutParams, error) {
	buff, err := NewBufferFromMultipart(params)
	if err != nil {
		return PutParams{}, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}

	result := PutParams{
		DirName:      dir,
		FileName:     uuid.NewString() + "." + files.GetFileExtenstion(params.Filename),
		FileBuffer:   buff,
		FileSize:     int64(buff.Len()),
		ContentType:  params.Header.Get(header.KeyContentType),
		UserTags:     map[string]string{},
		UserMetadata: map[string]string{},
	}
	return result, nil
}

func NormalizeDir(dir string) string {
	if dir[len(dir):] == "" {
		dir = dir + "/"
	}
	return dir
}

func FormatBytes(bytes uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	if bytes < 1024 {
		return fmt.Sprintf("%.2f B", float64(bytes))
	}

	value := float64(bytes)
	unitIdx := 0

	for value >= 1024 && unitIdx < len(units)-1 {
		value /= 1024
		unitIdx++
	}

	return fmt.Sprintf("%.2f %v", value, units[unitIdx])
}

func PercentBytes(used, total uint64) float64 {
	if total <= 0 {
		return 100
	}

	return (float64(used) / float64(total)) * 100
}
