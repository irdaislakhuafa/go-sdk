package collections

func Map[T any, R any](values []T, fn func(value T) R) (results []R) {
	for _, v := range values {
		results = append(results, fn(v))
	}
	return results
}

func MapWithErr[T any, R any](values []T, fn func(value T) (R, error)) (results []R, err error) {
	for _, v := range values {
		result, err := fn(v)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
