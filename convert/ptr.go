package convert

// Convert value to pointer of data type
func ToPointer[T any](value T) *T {
	return &value
}
