package storage

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/strformat"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type (
	minioimpl struct {
		client *minio.Client
		cfg    Config
	}
)

// Initialize minio object storage.
func InitMinio(cfg Config) (Interface, error) {
	cfg.parseDefault()

	endpoint := ""
	if port := strings.Trim(cfg.Port, " "); port != "" {
		if !strings.HasSuffix(port, ":") {
			port = ":" + port
		}

		endpoint = cfg.Host + port
	} else {
		endpoint = cfg.Host
	}

	m, err := minio.New(endpoint, &minio.Options{
		Creds:      credentials.NewStaticV4(cfg.AccessKeyID, cfg.AccessKeySecret, cfg.Token),
		Secure:     cfg.SSL,
		Region:     cfg.Region,
		MaxRetries: cfg.MaxRetries,
	})
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeStorageNoClient, "%s", err.Error())
	}

	if isExist, err := m.BucketExists(context.Background(), cfg.BaseDir); err != nil {
		return nil, errors.NewWithCode(codes.CodeStorageNoClient, "%s", err.Error())
	} else {
		if !isExist {
			err := m.MakeBucket(context.Background(), cfg.BaseDir, minio.MakeBucketOptions{
				Region: cfg.Region,
			})
			if err != nil {
				return nil, errors.NewWithCode(codes.CodeStorageNoClient, "%s", err.Error())
			}
		}
	}

	return &minioimpl{
		client: m,
		cfg:    cfg,
	}, nil
}

func (m *minioimpl) Put(ctx context.Context, params PutParams) (PutResult, error) {
	if params.DirName[len(params.DirName):] == "" {
		params.DirName = params.DirName + "/"
	}

	clientRes, err := m.client.PutObject(ctx, m.cfg.BaseDir, fmt.Sprintf("%s%s", params.DirName, params.FileName), params.FileBuffer, params.FileSize, minio.PutObjectOptions{
		ContentType:     params.ContentType,
		ContentEncoding: params.ContentEncoding,
	})
	if err != nil {
		return PutResult{}, errors.NewWithCode(codes.CodeStorageNoFile, "%s", err.Error())
	}

	result := PutResult{
		DirName:  params.DirName,
		FileName: params.FileName,
		FullPath: clientRes.Key,
		Region:   clientRes.Location,
		Tag:      clientRes.ETag,
		Size:     clientRes.Size,
	}
	return result, nil
}

func (m *minioimpl) Del(ctx context.Context, params DelParams) error {
	err := m.client.RemoveObject(ctx, m.cfg.BaseDir, params.FullPath, minio.RemoveObjectOptions{
		ForceDelete: params.Force,
	})
	if err != nil {
		return errors.NewWithCode(codes.CodeStorageDelFailure, "%s", err.Error())
	}

	return nil
}

func (m *minioimpl) Url(ctx context.Context, params UrlParams) (UrlResult, error) {
	params.parseDefault()
	clientRes, err := m.client.PresignedGetObject(ctx, m.cfg.BaseDir, params.FullPath, params.ExpireDuration, url.Values{})
	if err != nil {
		return UrlResult{}, errors.NewWithCode(codes.CodeStorageGenerateURLFailure, "%s", err.Error())
	}

	clientRes.RawQuery = url.QueryEscape(clientRes.RawQuery)

	url, err := strformat.T("{{ .Scheme }}://{{ .Host }}{{ .Path }}?{{ .RawQuery }}", clientRes)
	if err != nil {
		return UrlResult{}, errors.NewWithCode(codes.CodeStorageGenerateURLFailure, "%s", err.Error())
	}

	result := UrlResult{
		Scheme:   clientRes.Scheme,
		Host:     clientRes.Host,
		FullPath: clientRes.Path,
		FullURL:  url,
		RawQuery: clientRes.RawQuery,
	}

	return result, nil
}
