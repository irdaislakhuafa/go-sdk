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
		InMemory  bool // only for sqlite
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

// Driver is type of database driver. Current supported driver are:
//
// - `DriverMySQL`
// - `DriverPostgresQL`
// - `DriverSQLite`
//
// This type is string, so you can use string literal to represent driver
// name. For example: `db.Driver("mysql")` is equal with `db.DriverMySQL`.
//
// This type is useful for `db.Config` struct.
const (
	DriverMySQL      = Driver("mysql")
	DriverPostgresQL = Driver("postgres")
	DriverSQLite     = Driver("sqlite3")
)

// Init connect to database and return *sql.DB.
//
// This function will be switch to mysql, postgresql, sqlite driver.
//
// If driver not found, this function will return error with code `github.com/irdaislakhuafa/go-sdk/codes.CodeNotImplemented`.
func Init(cfg Config) (*sql.DB, error) {
	switch cfg.Driver {
	case DriverMySQL:
		return mysql(cfg)
	case DriverPostgresQL:
		return postgresql(cfg)
	case DriverSQLite:
		return sqlite(cfg)
	default:
		return nil, errors.NewWithCode(codes.CodeNotImplemented, "driver name %s not implemented!", cfg.Driver)
	}
}
