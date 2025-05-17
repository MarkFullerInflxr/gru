package gru

func Map[T, U any](things []T, fn func(T) U) (res []U) {
	for _, t := range things {
		res = append(res, fn(t))
	}
	return res
}
