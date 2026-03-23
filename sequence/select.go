package sequence

import "sync"

type selectFunc[TIn any, TOut any] func(source TIn) TOut

func Select[TS []TIn, TIn any, T any](source TS, selectFunc selectFunc[TIn, T]) []T {
	newSequence := make([]T, 0, len(source))
	for _, s := range source {
		newSequence = append(newSequence, selectFunc(s))
	}

	return newSequence
}

func ConcurrentSelect[TS []TIn, TIn any, T any](source TS, selectFunc selectFunc[TIn, T]) []T {
	newSequence := make([]T, len(source))
	var wg sync.WaitGroup

	for i, s := range source {
		wg.Add(1)
		go func(index int, value TIn) {
			defer wg.Done()
			newSequence[index] = selectFunc(value)
		}(i, s)
	}
	wg.Wait()

	return newSequence
}
