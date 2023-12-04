package operator

// like ternary operator in other programming language `value = (condition) ? ifTrue : ifFalse`
func Ternary[T any](isOk bool, ifOk, ifNot T) T {
	if isOk {
		return ifOk
	}
	return ifNot
}
