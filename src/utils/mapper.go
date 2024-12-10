package utils

func SliceMapper[T any, U any](input []T, converter func(T) U) []U {
	output := make([]U, len(input))
	for i, item := range input {
		output[i] = converter(item)
	}
	return output
}
