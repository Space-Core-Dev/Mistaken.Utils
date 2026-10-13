package sequence

func Any[T any](s []T, p func(T) bool) bool {
    for _, v := range s {
        if p(v) {
            return true
        }
    }
    return false
}
