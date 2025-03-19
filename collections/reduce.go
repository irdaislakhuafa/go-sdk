package collections

func Reduce[T any, R any](
	values []T,
	initial R,
	fn func(
		c R,
		i int,
		v T,
	) R,
) R {
	for i, v := range values {
		initial = fn(initial, i, v)
	}
	return initial
}

func ReduceWithErr[T any, R any](
	values []T,
	initial R,
	fn func(
		c R,
		i int,
		v T,
	) (R, error),
) (R, error) {
	var err error
	for i, v := range values {
		if initial, err = fn(initial, i, v); err != nil {
			return initial, err
		}
	}
	return initial, nil
}
