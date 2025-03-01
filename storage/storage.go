package storage

import (
	"context"
	"io"
	"time"
)

type (
	Interface interface {
		Put(ctx context.Context, params PutParams) (PutResult, error)
		Del(ctx context.Context, params DelParams) error
		Url(ctx context.Context, params UrlParams) (UrlResult, error)
	}
	Config struct {
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
