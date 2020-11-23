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

func getNextCell(field [][]bool, column int, row int) bool {
	cell := getCell(field, column, row)
	neighborCount := countNeighbors(field, column, row)
	willBeBorn := !cell && neighborCount == 3
	willSurvive := cell && (neighborCount == 2 || neighborCount == 3)
	return willBeBorn || willSurvive
}

func getNextField(field [][]bool) [][]bool {
	nextField := [][]bool{}
	for row := 0; row < getHeight(field); row = row + 1 {
		nextRow := []bool{}
		for column := 0; column < getWidth(field); column = column + 1 {
			nextCell := getNextCell(field, column, row)
			nextRow = append(nextRow, nextCell)
		}

		nextField = append(nextField, nextRow)
	}

	return nextField
}

func marshalField(field [][]bool) string {
	result := ""
	for row := 0; row < getHeight(field); row = row + 1 {
		for column := 0; column < getWidth(field); column = column + 1 {
			cell := getCell(field, column, row)
			if cell {
				result = result + "O"
			} else {
				result = result + "."
			}
		}

		result = result + "\n"
	}

	return result
}
