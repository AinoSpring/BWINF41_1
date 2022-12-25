package main

import (
	"fmt"
)

type Field struct {
	field       [][]Crystal
	size        Vector
	longestTime int
}

func NewField(size Vector) (field Field) {
	field.size = size
	field.longestTime = -1
	field.field = make([][]Crystal, 0)
	for y := 0; y < size.y; y++ {
		field.field = append(field.field, make([]Crystal, size.x))
	}
	return
}

func (field *Field) Position(position Vector) *Crystal {
	return &field.field[position.y][position.x]
}

func (field *Field) SetCell(position Vector, crystal Crystal) bool {
	if !(0 <= position.x && position.x < field.size.x && 0 <= position.y && position.y < field.size.y) {
		return false
	}
	if positionPointer := field.Position(position); !positionPointer.notNull {
		*positionPointer = crystal
		return true
	}
	return false
}

func (field *Field) UpdateCell(position Vector, frame int) ([]int, bool) {
	var currentlyMade = make([]int, 0)
	var changed = false
	if positionPointer := field.Position(position); positionPointer.notNull {
		if (frame % positionPointer.seed.plusSpeed.x) == 0 {
			var currentPosition = position.Add(Vector{x: 1})
			var crystal = NewCrystal(positionPointer.seed)
			currentlyMade = append(currentlyMade, crystal.id)
			if field.SetCell(currentPosition, crystal) {
				changed = true
			}
		}
		if (frame % positionPointer.seed.plusSpeed.y) == 0 {
			var currentPosition = position.Add(Vector{y: 1})
			var crystal = NewCrystal(positionPointer.seed)
			currentlyMade = append(currentlyMade, crystal.id)
			if field.SetCell(currentPosition, crystal) {
				changed = true
			}
		}
		if (frame % positionPointer.seed.minusSpeed.x) == 0 {
			var currentPosition = position.Add(Vector{x: -1})
			var crystal = NewCrystal(positionPointer.seed)
			currentlyMade = append(currentlyMade, crystal.id)
			if field.SetCell(currentPosition, crystal) {
				changed = true
			}
		}
		if (frame % positionPointer.seed.minusSpeed.y) == 0 {
			var currentPosition = position.Add(Vector{y: -1})
			var crystal = NewCrystal(positionPointer.seed)
			currentlyMade = append(currentlyMade, crystal.id)
			if field.SetCell(currentPosition, crystal) {
				changed = true
			}
		}
	}
	return currentlyMade, changed
}

func (field *Field) AddCrystal(crystal Crystal, position Vector) {
	field.SetCell(position, crystal)
	for _, time := range []int{crystal.seed.plusSpeed.x, crystal.seed.plusSpeed.y, crystal.seed.minusSpeed.x, crystal.seed.minusSpeed.y} {
		if field.longestTime < time {
			field.longestTime = time
		}
	}
}

func (field *Field) Update(frame int) bool {
	var changed = false
	var finishedIds = make([]int, 0)
	for x := 0; x < field.size.x; x++ {
		for y := 0; y < field.size.y; y++ {
			var position = Vector{x, y}
			if positionPointer := field.Position(position); !SliceContains(finishedIds, positionPointer.id) {
				changedIds, fieldChanged := field.UpdateCell(position, frame)
				for _, id := range changedIds {
					finishedIds = append(finishedIds, id)
				}
				changed = changed || fieldChanged
			}
		}
	}
	return changed
}

func (field *Field) Generate() {
	var noDeltaFrames = 0
	for i := 0; true; i++ {
		fmt.Printf("%v\n", i)
		var changed = field.Update(i)
		if !changed {
			noDeltaFrames++
		} else {
			noDeltaFrames = 0
		}
		if noDeltaFrames >= field.longestTime {
			break
		}
	}
}

func (field Field) String() (resultString string) {
	for x := 0; x < field.size.x; x++ {
		for y := 0; y < field.size.y; y++ {
			resultString += fmt.Sprintf("%v ", field.Position(Vector{x, y}).seed.orientation)
		}
		resultString += "\n"
	}
	return
}
