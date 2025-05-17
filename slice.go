package gru

func Map[T, U any](things []T, fn func(T) U) (res []U) {
	for _, t := range things {
		res = append(res, fn(t))
	}
	return res
}

func ToMap[K comparable, V any](things []V, fn func(V) K) map[K]V {
	res := make(map[K]V)
	for _, t := range things {
		res[fn(t)] = t
	}
	return res
}
