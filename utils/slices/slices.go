package slices

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
