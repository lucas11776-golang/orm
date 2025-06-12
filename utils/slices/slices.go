package slices

// Comment
func Map[T any, R any](items []T, callback func(item T) R) []R {
	mapped := []R{}

	for _, item := range items {
		mapped = append(mapped, callback(item))
	}

	return mapped
}

// Comment
func Filter[T any](items []T, callback func(item T) bool) []T {
	list := make([]T, 0)

	for i := range items {
		if !callback(items[i]) {
			list = append(list, items[i])
		}
	}

	return list
}
