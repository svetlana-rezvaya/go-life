package main

func getWidth(field [][]bool) int {
	return len(field[0])
}

func getHeight(field [][]bool) int {
	return len(field)
}

func getCell(field [][]bool, column int, row int) bool {
	return field[row][column]
}

func setCell(field [][]bool, column int, row int, cell bool) {
	field[row][column] = cell
}
