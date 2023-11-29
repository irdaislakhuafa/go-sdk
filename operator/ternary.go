package operator

func Ternary[T any](isOk bool, ifOk, ifNot T) T {
	if isOk {
		return ifOk
	}
	return ifNot
}
