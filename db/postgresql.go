package db

import (
	"database/sql"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/strformat"
	_ "github.com/lib/pq"
)

func postgresql(cfg Config) (*sql.DB, error) {
	dsn, err := strformat.T("host={{ .Host }} port={{ .Port }} user={{ .User }} password={{ .Password }} dbname={{ .Name }} sslmode={{ .Options.SSLMode }}", cfg)
	if err != nil {
		return nil, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}

	db, err := sql.Open(string(DriverPostgresQL), dsn)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLInit, "%s", err.Error())
	}

	return db, nil
}
