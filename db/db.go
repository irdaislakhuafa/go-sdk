package db

import (
	"database/sql"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type (
	Driver  string
	Options struct {
		SSLMode   string
		ParseTime bool
	}
	Config struct {
		Driver   Driver
		User     string
		Password string
		Port     string
		Host     string
		Name     string
		Options  Options
	}
)

const (
	DriverMySQL      = Driver("mysql")
	DriverPostgresQL = Driver("postgres")
)

func Init(cfg Config) (*sql.DB, error) {
	switch cfg.Driver {
	case DriverMySQL:
		return mysql(cfg)
	case DriverPostgresQL:
		return postgresql(cfg)
	default:
		return nil, errors.NewWithCode(codes.CodeNotImplemented, "driver name %s not implemented!", cfg.Driver)
	}
}
