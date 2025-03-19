package collections

func Map[T any, R any](values []T, fn func(i int, v T) R) (results []R) {
	for i, v := range values {
		results = append(results, fn(i, v))
	}
	return results
}

func MapWithErr[T any, R any](values []T, fn func(i int, v T) (R, error)) (results []R, err error) {
	for i, v := range values {
		result, err := fn(i, v)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
