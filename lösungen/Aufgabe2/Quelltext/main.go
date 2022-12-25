package main

import (
	"math/rand"
	"os"
)

func main() {
	if _, err := os.Stat("template.png"); err == nil {
		GenerateTemplate("template.png", func(x int, y int) (plusSpeed Vector, minusSpeed Vector) {
			plusSpeed.x = rand.Intn(100) + 1
			plusSpeed.y = rand.Intn(100) + 1
			minusSpeed.x = rand.Intn(100) + 1
			minusSpeed.y = rand.Intn(100) + 1
			return plusSpeed, minusSpeed
		}, 0.5)
	}
	var field = importCrystals()
	DisplayField(field).Save("1.png")
	field.Generate()
	DisplayField(field).Save("Result.png")
}
