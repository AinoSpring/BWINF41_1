package main

func CopyMap[K comparable, T any](m map[K]T) (result map[K]T) {
	result = make(map[K]T)
	for k, v := range m {
		result[k] = v
	}
	return
}

func CompareMap[K comparable, T comparable](a map[K]T, b map[K]T) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
