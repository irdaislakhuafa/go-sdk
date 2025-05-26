package storage

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
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
		FileName:     uuid.NewString(),
		FileBuffer:   buff,
		FileSize:     int64(buff.Len()),
		ContentType:  params.Header.Get(header.KeyContentType),
		UserTags:     map[string]string{},
		UserMetadata: map[string]string{},
	}
	return result, nil
}
