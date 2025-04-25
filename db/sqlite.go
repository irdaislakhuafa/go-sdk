package db

import (
	"database/sql"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	_ "github.com/mattn/go-sqlite3"
)

func sqlite(cfg Config) (*sql.DB, error) {
	dsn := ""
	if cfg.Options.InMemory {
		dsn = ":memory:"
	} else {
		dsn = "./" + cfg.Name + ".db"
	}

	db, err := sql.Open(string(DriverSQLite), dsn)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLInit, "%s", err.Error())
	}

	return db, nil
}
