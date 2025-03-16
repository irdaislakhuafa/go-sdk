package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/strformat"
)

func mysql(cfg Config) (*sql.DB, error) {
	dsn, err := strformat.T("{{ .User }}:{{ .Password }}@({{ .Host }}:{{ .Port }})/{{ .Name }}?parseTime={{ .Options.ParseTime }}", cfg)
	if err != nil {
		return nil, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}

	db, err := sql.Open(string(DriverMySQL), dsn)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLInit, "%s", err.Error())
	}

	return db, nil
}
