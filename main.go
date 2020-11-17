package main

func getWidth(field [][]bool) int {
	return len(field[0])
}

func getHeight(field [][]bool) int {
	return len(field)
}

func getCell(field [][]bool, column int, row int) bool {
	column = (column + getWidth(field)) % getWidth(field)
	row = (row + getHeight(field)) % getHeight(field)
	return field[row][column]
}

func setCell(field [][]bool, column int, row int, cell bool) {
	field[row][column] = cell
}

func countNeighbors(field [][]bool, column int, row int) int {
	count := 0
	for columnDelta := -1; columnDelta <= 1; columnDelta = columnDelta + 1 {
		for rowDelta := -1; rowDelta <= 1; rowDelta = rowDelta + 1 {
			if columnDelta == 0 && rowDelta == 0 {
				continue
			}

			cell := getCell(field, column+columnDelta, row+rowDelta)
			if cell {
				count = count + 1
			}
		}
	}

	return count
}
