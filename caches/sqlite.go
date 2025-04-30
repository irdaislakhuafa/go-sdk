package caches

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/db"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type (
	sqlite[T any] struct {
		cfg     Config
		storage *sql.DB
		dbCfg   db.Config
		table   string
	}

	tableCaches struct {
		ID   int64           `db:"id" json:"id"`
		Key  string          `db:"key" json:"key"`
		Data json.RawMessage `db:"data" json:"data"`
	}
)

func InitSQLite[T any](cfg Config) Interface[T] {
	dbCfg := db.Config{
		Driver:   db.DriverSQLite,
		User:     "caches",
		Password: "caches",
		Name:     cfg.Dir + "caches",
		Options:  db.Options{},
	}
	table := "caches"

	db, err := db.Init(dbCfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	tableDDL := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key VARCHAR(512) NOT NULL UNIQUE,
			data JSON NOT NULL
		)
	`, table)
	_, err = db.ExecContext(ctx, tableDDL)
	if err != nil {
		panic(err)
	}

	return &sqlite[T]{
		cfg:     cfg,
		storage: db,
		dbCfg:   dbCfg,
		table:   table,
	}
}

func (s *sqlite[T]) Remember(key string, ttlS uint64, callback func() (T, error)) (T, error) {
	var empty T

	// get existing cache
	rowResult, err := s.sqlGet(tableCaches{Key: key})
	if err != nil {
		code := errors.GetCode(err)
		if code != codes.CodeCacheKeyNotFound {
			return empty, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
		}
	}

	// if data is exist then check TTL
	if rowResult.ID > 0 {
		i := item[T]{}
		if err := json.Unmarshal(rowResult.Data, &i); err != nil {
			return empty, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
		}

		// if TTL is not expired then return the data
		if i.TTL > uint64(time.Now().Unix()) {
			return i.Data, nil
		} else { // if expired then call callback
			data, err := callback()
			if err != nil {
				return empty, err
			}

			i = item[T]{
				TTL:  uint64(time.Now().Unix()) + ttlS,
				Data: data,
			}

			iData, err := json.Marshal(i)
			if err != nil {
				return empty, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
			}
			rowResult.Data = iData

			// and update data on cache
			_, err = s.sqlUpdate(rowResult)
			if err != nil {
				return empty, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
			}
			return data, nil
		}
	} else { // if not exist then add to cache
		data, err := callback()
		if err != nil {
			return empty, err
		}

		i := item[T]{
			TTL:  uint64(time.Now().Unix()) + ttlS,
			Data: data,
		}

		iData, err := json.Marshal(i)
		if err != nil {
			return empty, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
		}

		_, err = s.sqlAdd(tableCaches{
			Key:  key,
			Data: iData,
		})
		if err != nil {
			return empty, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
		}

		return data, nil
	}
}

func (s *sqlite[T]) Clear() {
	s.sqlClear()
}

func (s *sqlite[T]) Forget(key string) (T, error) {
	var empty T

	data, err := s.Get(key)
	if err != nil {
		return empty, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}

	_, err = s.sqlDel(tableCaches{Key: key})
	if err != nil {
		return empty, errors.NewWithCode(codes.CodeCacheEnd, "%s", err.Error())
	}

	return data, nil
}

func (s *sqlite[T]) ForgetFn(fn func(key string) (T, error)) (T, error) {
	var empty T
	keys, err := s.sqlKeys()
	if err != nil {
		return empty, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}

	var res T
	for _, k := range keys {
		if res, err = fn(k); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (s *sqlite[T]) Length() uint64 {
	total, _ := s.sqlCount()
	return uint64(total)
}

func (s *sqlite[T]) Get(key string) (T, error) {
	var empty T
	tc, err := s.sqlGet(tableCaches{Key: key})
	if err != nil {
		return empty, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}

	item := item[T]{}
	if err = json.Unmarshal(tc.Data, &item); err != nil {
		return empty, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
	}

	return item.Data, nil
}

func (s *sqlite[T]) Keys() []string {
	results, _ := s.sqlKeys()
	return results
}

func (s *sqlite[T]) sqlAdd(value tableCaches) (tableCaches, error) {
	add := fmt.Sprintf(`INSERT INTO %s (key, data) VALUES (?, ?)`, s.table)
	row, err := s.storage.Exec(add, value.Key, value.Data)
	if err != nil {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
	}

	id, err := row.LastInsertId()
	if err != nil {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
	}
	value.ID = id
	return value, nil
}

func (s *sqlite[T]) sqlDel(value tableCaches) (tableCaches, error) {
	del := fmt.Sprintf("DELETE FROM %s WHERE key = ?", s.table)
	row, err := s.storage.Exec(del, value.Key)
	if err != nil {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
	}

	i, err := row.RowsAffected()
	if err != nil {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheDelErr, "%s", err.Error())
	}

	if i == 0 {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheKeyNotFound, "cache key '%s' not found", value.Key)
	}

	return value, nil
}

func (s *sqlite[T]) sqlGet(value tableCaches) (tableCaches, error) {
	get := fmt.Sprintf(`SELECT id, key, data FROM %s WHERE key = ?`, s.table)
	row := s.storage.QueryRow(get, value.Key)
	if err := row.Scan(&value.ID, &value.Key, &value.Data); err != nil {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheKeyNotFound, "%s", err.Error())
	}

	if value.ID == 0 {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheKeyNotFound, "cache key '%s' not found", value.Key)
	}

	return value, nil
}

func (s *sqlite[T]) sqlUpdate(value tableCaches) (tableCaches, error) {
	update := fmt.Sprintf(`UPDATE %s SET data = ? WHERE key = ?`, s.table)
	_, err := s.storage.Exec(update, value.Data, value.Key)
	if err != nil {
		return tableCaches{}, errors.NewWithCode(codes.CodeCacheAddErr, "%s", err.Error())
	}

	return value, nil
}

func (s *sqlite[T]) sqlClear() error {
	clear := fmt.Sprintf("DELETE FROM %s", s.table)
	_, err := s.storage.Exec(clear)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDelErr, "%s", err.Error())
	}
	return nil
}

func (s *sqlite[T]) sqlList() ([]tableCaches, error) {
	list := fmt.Sprintf("SELECT id, key, data FROM %s", s.table)
	rows, err := s.storage.Query(list)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
	}
	defer rows.Close()

	results := []tableCaches{}
	for rows.Next() {
		item := tableCaches{}
		err = rows.Scan(&item.ID, &item.Key, &item.Data)
		if err != nil {
			return nil, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
		}
		results = append(results, item)
	}

	return results, nil
}

func (s *sqlite[T]) sqlCount() (int64, error) {
	var total int64
	count := fmt.Sprintf("SELECT COUNT(id) AS total FROM %s", s.table)
	row := s.storage.QueryRow(count)
	if err := row.Scan(&count); err != nil {
		return 0, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
	}

	return total, nil
}

func (s *sqlite[T]) sqlKeys() ([]string, error) {
	keys := fmt.Sprintf("SELECT key FROM %s", s.table)
	rows, err := s.storage.Query(keys)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
	}
	defer rows.Close()

	results := []string{}
	for rows.Next() {
		item := ""
		if err := rows.Scan(&item); err != nil {
			return nil, errors.NewWithCode(codes.CodeCacheReadErr, "%s", err.Error())
		}
		results = append(results, item)
	}

	return results, nil
}
