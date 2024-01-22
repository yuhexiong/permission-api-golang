package util

func GetPointer[T any](value T) *T {
	return &value
}
