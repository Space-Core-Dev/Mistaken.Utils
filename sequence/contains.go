package sequence

import "sort"

type comparator[T any] func(a, b T) bool

func Contains[T comparable](source *[]T, obj T) bool {
	for _, item := range *source {
		if item == obj {
			return true
		}
	}

	return false
}

func ContainsBinarySearch[T any](source *[]T, obj T, less comparator[T]) bool {
	index := sort.Search(len(*source), func(i int) bool {
		return !less((*source)[i], obj)
	})
	return index < len(*source) && !less(obj, (*source)[index])
}
