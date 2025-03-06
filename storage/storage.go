package storage

import (
	"context"
	"io"
	"time"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type StorageType string

const (
	TypeMinio = StorageType("minio")
	TypeDisk  = StorageType("disk")
)

type (
	Interface interface {
		// Put a file to storage.
		//
		// Errors:
		//
		// - `codes.CodeStorageNoFile`: if failed to put file to storage.
		Put(ctx context.Context, params PutParams) (PutResult, error)

		// Delete data from storage.
		//
		// Errors:
		//
		// - `codes.CodeStorageDelFailure`: if cannot delete file.
		Del(ctx context.Context, params DelParams) error

		// Generate url based on given full path of file.
		//
		// Errors:
		//
		// - `codes.CodeStorageGenerateURLFailure`: if failed on generate url.
		Url(ctx context.Context, params UrlParams) (UrlResult, error)
	}

	Config struct {
		StorageType     StorageType
		BaseDir         string
		Host            string
		Port            string
		SSL             bool
		AccessKeyID     string
		AccessKeySecret string
		Token           string
		Region          string
		MaxRetries      int
	}
	PutParams struct {
		DirName                string
		FileName               string
		FileBuffer             io.Reader
		FileSize               int64
		ContentType            string
		ContentEncoding        string
		UserTags, UserMetadata map[string]string
	}
	PutResult struct {
		DirName  string
		FileName string
		FullPath string
		Region   string
		Tag      string
		Size     int64
	}
	DelParams struct {
		FullPath string
		Force    bool
	}
	UrlParams struct {
		FullPath       string
		ExpireDuration time.Duration
	}
	UrlResult struct {
		Scheme   string
		Host     string
		FullPath string
		FullURL  string
		RawQuery string
	}
)

func (c *Config) parseDefault() {
	if c.MaxRetries <= 0 {
		c.MaxRetries = 1
	}
}

func (u *UrlParams) parseDefault() {
	if u.ExpireDuration.Milliseconds() <= 0 {
		u.ExpireDuration = time.Hour * 24
	}
}

// Initialize storage client.
//
// Current supported storage types:
//
// - minio: supported.
//
// - disk: not implemented yet.
//
// Will return error below if fail:
//
// - `codes.CodeStorageNoClient`: cannot connect to client
func Init(cfg Config) (Interface, error) {
	switch cfg.StorageType {
	case TypeMinio:
		return InitMinio(cfg)
	case TypeDisk:
		// TODO: implement file operation on disk
		fallthrough
	default:
		return nil, errors.NewWithCode(codes.CodeNotImplemented, "storage type '%s' not implemented", cfg.StorageType)
	}
}
