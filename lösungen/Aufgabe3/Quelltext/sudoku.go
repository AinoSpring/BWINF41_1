package Quelltext

type CellShape bool
type LineShape []CellShape
type LineBlockShape []LineShape

type Sudoku struct {
	field [][]uint8
}

func (sudoku *Sudoku) Copy() Sudoku {
	var field [][]uint8
	copy(field, sudoku.field)
	return Sudoku{field: field}
}

func (sudoku *Sudoku) Get(x uint, y uint) *uint8 {
	return &sudoku.field[y][x]
}

func (sudoku *Sudoku) GetRow(index uint) (row []uint8) {
	row = make([]uint8, 0)
	for i := uint(0); i < 9; i++ {
		row = append(row, *sudoku.Get(index, i))
	}
	return
}

func (sudoku *Sudoku) GetColumn(index uint) (column []uint8) {
	column = make([]uint8, 0)
	for i := uint(0); i < 9; i++ {
		column = append(column, *sudoku.Get(i, index))
	}
	return
}

func (sudoku *Sudoku) GetRowBlock(index uint) (rowBlock [][]uint8) {
	rowBlock = make([][]uint8, 0)
	for i := index * 3; i < ((index+1)*3)-1; i++ {
		rowBlock = append(rowBlock, sudoku.GetRow(i))
	}
	return
}

func (sudoku *Sudoku) GetColumnBlock(index uint) (columnBlock [][]uint8) {
	columnBlock = make([][]uint8, 0)
	for i := index * 3; i < ((index+1)*3)-1; i++ {
		columnBlock = append(columnBlock, sudoku.GetColumn(i))
	}
	return
}

func (sudoku *Sudoku) SetRow(index uint, value []uint8) {
	for i := uint(0); i < 9; i++ {
		*sudoku.Get(i, index) = value[i]
	}
}

func (sudoku *Sudoku) SetColumn(index uint, value []uint8) {
	for i := uint(0); i < 9; i++ {
		*sudoku.Get(index, i) = value[i]
	}
}

func (sudoku *Sudoku) SetRowBlock(index uint, value [][]uint8) {
	for i := uint(0); i < 3; i++ {
		sudoku.SetRow((index*3)+i, value[i])
	}
}

func (sudoku *Sudoku) SetColumnBlock(index uint, value [][]uint8) {
	for i := uint(0); i < 3; i++ {
		sudoku.SetColumn((index*3)+i, value[i])
	}
}

func (sudoku *Sudoku) SwapRows(index1 uint, index2 uint) {
	var temp1 = sudoku.GetRow(index1)
	var temp2 = sudoku.GetRow(index2)
	sudoku.SetRow(index1, temp2)
	sudoku.SetRow(index2, temp1)
}

func (sudoku *Sudoku) SwapColumns(index1 uint, index2 uint) {
	var temp1 = sudoku.GetColumn(index1)
	var temp2 = sudoku.GetColumn(index2)
	sudoku.SetColumn(index1, temp2)
	sudoku.SetColumn(index2, temp1)
}

func (sudoku *Sudoku) SwapRowBlocks(index1 uint, index2 uint) {
	var temp1 = sudoku.GetRowBlock(index1)
	var temp2 = sudoku.GetRowBlock(index2)
	sudoku.SetRowBlock(index1, temp2)
	sudoku.SetRowBlock(index2, temp1)
}

func (sudoku *Sudoku) SwapColumnBlocks(index1 uint, index2 uint) {
	var temp1 = sudoku.GetColumnBlock(index1)
	var temp2 = sudoku.GetColumnBlock(index2)
	sudoku.SetColumnBlock(index1, temp2)
	sudoku.SetColumnBlock(index2, temp1)
}

func Shape(cell *uint8) CellShape {
	return *cell > 0
}

func RowShape(row []*uint8) (rowShape LineShape) {
	rowShape = make(LineShape, 0)
	for _, cell := range row {
		rowShape = append(rowShape, Shape(cell))
	}
	return
}

func ColumnShape(column []*uint8) (columnShape LineShape) {
	columnShape = make(LineShape, 0)
	for _, cell := range column {
		columnShape = append(columnShape, Shape(cell))
	}
	return
}

func RowBlockShape(rowBlock [][]*uint8) (rowBlockShape LineBlockShape) {
	rowBlockShape = make(LineBlockShape, len(rowBlock))
	for _, row := range rowBlock {
		rowBlockShape = append(rowBlockShape, RowShape(row))
	}
	return
}

func ColumnBlockShape(columnBlock [][]*uint8) (columnBlockShape LineBlockShape) {
	columnBlockShape = make(LineBlockShape, len(columnBlock))
	for _, row := range columnBlock {
		columnBlockShape = append(columnBlockShape, ColumnShape(row))
	}
	return
}

func (cellShape CellShape) Equal(other CellShape) bool {
	return cellShape == other
}

func (lineShape LineShape) Equal(other LineShape) (isEqual bool) {
	isEqual = true
	for idx, cellShape := range lineShape {
		isEqual = isEqual && cellShape.Equal(other[idx])
	}
	return
}

func (lineBlockShape LineBlockShape) Equal(other LineBlockShape) (isEqual bool) {
	isEqual = true
	for idx, lineShape := range lineBlockShape {
		isEqual = isEqual && lineShape.Equal(other[idx])
	}
	return
}

func (lineBlockShape LineBlockShape) Count(other LineShape) (count uint) {
	for _, lineShape := range lineBlockShape {
		if lineShape.Equal(other) {
			count++
		}
	}
	return
}

func (lineBlockShape LineBlockShape) Contains(other LineShape) bool {
	return lineBlockShape.Count(other) > 0
}

func (lineBlockShape LineBlockShape) CompareWith(other LineBlockShape) bool {
	for _, lineShape := range lineBlockShape {
		if other.Count(lineShape) == 1 {
			continue
		}
		return false
	}
	return true
}
