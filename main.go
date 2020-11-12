package main

func getCell(field [][]bool, column int, row int) bool {
	return field[row][column]
}

func setCell(field [][]bool, column int, row int, cell bool) {
	field[row][column] = cell
}
