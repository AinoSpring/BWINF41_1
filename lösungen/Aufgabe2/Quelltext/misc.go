package main

type Vector struct {
	x int
	y int
}

func (vector Vector) Add(other Vector) (resultVector Vector) {
	resultVector = vector
	resultVector.x += other.x
	resultVector.y += other.y
	return
}

func Unit(x float64) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

func SliceContains[T comparable](slice []T, value T) bool {
	for _, element := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func sliceEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func slice2DEqual[T comparable](a [][]T, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !sliceEqual(v, b[i]) {
			return false
		}
	}
	return true
}
