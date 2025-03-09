package convert

import (
	"database/sql"
	"reflect"
	"time"
)

// Convert value to pointer of data type
func ToPointer[T any](value T) *T {
	return &value
}

// Convert any value to `T` data type and if value is not valid or `nil`, will use default value of `T` (null/nil safety) as return value
func ToSafeValue[T any](value any) T {
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Pointer {
		if !rv.IsNil() {
			rv = rv.Elem()
		} else {
			break
		}
	}

	if !rv.IsValid() {
		return ToSafeValue[T](new(T))
	}

	safeValue, isOk := rv.Interface().(T)
	if !isOk {
		return ToSafeValue[T](new(T))
	}
	return safeValue
}

// Convert sql.NullTime to *time.Time is valid and to nil is invalid
func SQLNullTimeToNil(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}
	return &value.Time
}

// Convert sql.NullInt16 to *int16 is valid and to nil is invalid
func SQLNullInt16ToNil(value sql.NullInt16) *int16 {
	if !value.Valid {
		return nil
	}
	return &value.Int16
}

// Convert sql.NullInt32 to *int32 is valid and to nil is invalid
func SQLNullInt32ToNil(value sql.NullInt32) *int32 {
	if !value.Valid {
		return nil
	}
	return &value.Int32
}

// Convert sql.NullInt64 to *int64 is valid and to nil is invalid
func SQLNullInt64ToNil(value sql.NullInt64) *int64 {
	if !value.Valid {
		return nil
	}
	return &value.Int64
}

// Convert sql.NullFloat64 to *float64 is valid and to nil is invalid
func SQLNullFloat64ToNil(value sql.NullFloat64) *float64 {
	if !value.Valid {
		return nil
	}
	return &value.Float64
}

// Convert sql.NullString to *string is valid and to nil is invalid
func SQLNullStringToNil(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}

// Convert sql.NullBool to *bool is valid and to nil is invalid
func SQLNullBoolToNil(value sql.NullBool) *bool {
	if !value.Valid {
		return nil
	}
	return &value.Bool
}
