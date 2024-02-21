package convert

import "reflect"

// Convert value to pointer of data type
func ToPointer[T any](value T) *T {
	return &value
}

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
