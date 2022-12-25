package main

var (
	crystalIds = 1
)

type Seed struct {
	orientation uint8
	plusSpeed   Vector
	minusSpeed  Vector
}

type Crystal struct {
	seed    Seed
	id      int
	notNull bool
}

func NewCrystal(seed Seed) (crystal Crystal) {
	crystal.seed = seed
	crystal.notNull = true
	crystal.id = crystalIds
	crystalIds++
	return
}
