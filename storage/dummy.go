package storage

import (
	"context"
	"fmt"
)

type (
	dummyImpl struct {
		cfg Config
	}
)

// InitDummy initializes a dummy storage client.
// It is intended for testing purposes and will always return a successful result.
func InitDummy(cfg Config) (Interface, error) {
	result := &dummyImpl{
		cfg: cfg,
	}
	return result, nil
}

// Del implements the storage Interface by simulating a successful file deletion.
// It always returns nil.
func (d *dummyImpl) Del(ctx context.Context, params DelParams) error {
	return nil
}

// Put implements the storage Interface by simulating a successful file upload.
// It returns a PutResult populated with the provided parameters and dummy metadata.
func (d *dummyImpl) Put(ctx context.Context, params PutParams) (PutResult, error) {
	params.DirName = NormalizeDir(params.DirName)
	result := PutResult{
		DirName:  params.DirName,
		FileName: params.FileName,
		FullPath: params.DirName + params.FileName,
		Region:   string(TypeDummy),
		Tag:      string(TypeDummy),
		Size:     params.FileSize,
	}

	return result, nil
}

// Url implements the storage Interface by generating a mock URL.
// It returns a dummy URL string based on the provided path.
func (d *dummyImpl) Url(ctx context.Context, params UrlParams) (UrlResult, error) {
	result := UrlResult{
		Scheme:   "https",
		Host:     string(TypeDummy) + ".com",
		FullPath: params.FullPath,
		FullURL:  fmt.Sprintf("%v.com/%v", TypeDummy, params.FullPath),
		RawQuery: "",
	}
	return result, nil
}
