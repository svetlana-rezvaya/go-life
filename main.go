package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Field ...
type Field [][]bool

// NewField ...
func NewField(width int, height int) Field {
	field := make(Field, height)
	for rowIndex := range field {
		field[rowIndex] = make([]bool, width)
	}

	return field
}

// ParseField ...
func ParseField(text string) (Field, error) {
	field := Field{}
	fieldWidth := -1
	lines := strings.Split(text, "\n")
	for lineIndex, line := range lines {
		if line != "" && line[0] == '!' {
			continue
		}

		row := []bool{}
		for _, character := range line {
			if character != 'O' && character != '.' {
				return nil, fmt.Errorf("unknown character %q", character)
			}

			cell := character == 'O'
			row = append(row, cell)
		}
		if fieldWidth == -1 {
			fieldWidth = len(row)
		} else if len(row) != fieldWidth {
			return nil, fmt.Errorf("inconsistent length of line %d", lineIndex+1)
		}

		field = append(field, row)
	}

	return field, nil
}

// Width ...
func (field Field) Width() int {
	return len(field[0])
}

// Height ...
func (field Field) Height() int {
	return len(field)
}

// Cell ...
func (field Field) Cell(column int, row int) bool {
	column = (column + field.Width()) % field.Width()
	row = (row + field.Height()) % field.Height()
	return field[row][column]
}

// SetCell ...
func (field Field) SetCell(column int, row int, cell bool) {
	field[row][column] = cell
}

// NeighborCount ...
func (field Field) NeighborCount(column int, row int) int {
	count := 0
	for columnDelta := -1; columnDelta <= 1; columnDelta = columnDelta + 1 {
		for rowDelta := -1; rowDelta <= 1; rowDelta = rowDelta + 1 {
			if columnDelta == 0 && rowDelta == 0 {
				continue
			}

			cell := field.Cell(column+columnDelta, row+rowDelta)
			if cell {
				count = count + 1
			}
		}
	}

	return count
}

// NextCell ...
func (field Field) NextCell(column int, row int) bool {
	cell := field.Cell(column, row)
	neighborCount := field.NeighborCount(column, row)
	willBeBorn := !cell && neighborCount == 3
	willSurvive := cell && (neighborCount == 2 || neighborCount == 3)
	return willBeBorn || willSurvive
}

// NextField ...
func (field Field) NextField() Field {
	nextField := NewField(field.Width(), field.Height())
	for row := 0; row < field.Height(); row = row + 1 {
		for column := 0; column < field.Width(); column = column + 1 {
			nextCell := field.NextCell(column, row)
			nextField.SetCell(column, row, nextCell)
		}
	}

	return nextField
}

// String ...
func (field Field) String() string {
	result := ""
	for row := 0; row < field.Height(); row = row + 1 {
		for column := 0; column < field.Width(); column = column + 1 {
			cell := field.Cell(column, row)
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

func main() {
	outDelay :=
		flag.Duration("outDelay", 100*time.Millisecond, "delay between frames")
	flag.Parse()

	fieldBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("unable to read the field: ", err)
	}
	fieldBytes = bytes.TrimSpace(fieldBytes)

	field, err := ParseField(string(fieldBytes))
	if err != nil {
		log.Fatal("unable to unmarshal the field: ", err)
	}

	for {
		fmt.Print(field)
		time.Sleep(*outDelay)

		field = field.NextField()
	}
}
