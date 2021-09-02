package main

import (
	"reflect"
	"testing"
)

func TestFieldWidth(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	width := field.Width()
	if width != 3 {
		test.Fail()
	}
}

func TestFieldHeight(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	height := field.Height()
	if height != 2 {
		test.Fail()
	}
}

func TestFieldCell_withTrue(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, true /* ! */, false},
	}
	cell := field.Cell(1, 2)
	if cell != true {
		test.Fail()
	}
}

func TestFieldCell_withFalse(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false /* ! */, false, false},
		[]bool{false, true, false},
	}
	cell := field.Cell(0, 1)
	if cell != false {
		test.Fail()
	}
}

func TestFieldCell_withCoordinatesBeyondMinimum(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, true /* ! */},
		[]bool{false, false, false},
	}
	cell := field.Cell(-1, -2)
	if cell != true {
		test.Fail()
	}
}

func TestFieldCell_withCoordinatesBeyondMaximum(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{true /* ! */, false, false},
		[]bool{false, false, false},
	}
	cell := field.Cell(3, 4)
	if cell != true {
		test.Fail()
	}
}

func TestFieldSetCell_withTrue(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false /* ! */, false, false},
		[]bool{false, false, false},
	}
	field.SetCell(0, 1, true)

	wantedField := Field{
		[]bool{false, false, false},
		[]bool{true /* ! */, false, false},
		[]bool{false, false, false},
	}
	if !reflect.DeepEqual(field, wantedField) {
		test.Fail()
	}
}

func TestFieldSetCell_withFalse(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{true /* ! */, false, false},
		[]bool{false, false, false},
	}
	field.SetCell(0, 1, false)

	wantedField := Field{
		[]bool{false, false, false},
		[]bool{false /* ! */, false, false},
		[]bool{false, false, false},
	}
	if !reflect.DeepEqual(field, wantedField) {
		test.Fail()
	}
}

func TestFieldNeighborCount_withCellInMiddle(test *testing.T) {
	field := Field{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	count := field.NeighborCount(1, 1)
	if count != 5 {
		test.Fail()
	}
}

func TestFieldNeighborCount_withCellInCorner(test *testing.T) {
	field := Field{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	count := field.NeighborCount(0, 0)
	if count != 5 {
		test.Fail()
	}
}

func Test_getNextCell_withBirth(test *testing.T) {
	field := Field{
		[]bool{false, false, false, false, false},
		[]bool{false, false, true, false, false},
		[]bool{false, false /* ! */, false, true, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}
	nextCell := getNextCell(field, 1, 2)
	if nextCell != true {
		test.Fail()
	}
}

func Test_getNextCell_withSurvival(test *testing.T) {
	field := Field{
		[]bool{false, false, false, false, false},
		[]bool{false, false, true, false, false},
		[]bool{false, false, false, true /* ! */, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}
	nextCell := getNextCell(field, 3, 2)
	if nextCell != true {
		test.Fail()
	}
}

func Test_getNextCell_withDeath(test *testing.T) {
	field := Field{
		[]bool{false, false, false, false, false},
		[]bool{false, false, true /* ! */, false, false},
		[]bool{false, false, false, true, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}
	nextCell := getNextCell(field, 2, 1)
	if nextCell != false {
		test.Fail()
	}
}

func Test_getNextField(test *testing.T) {
	field := Field{
		[]bool{false, false, false, false, false},
		[]bool{false, false, true, false, false},
		[]bool{false, false, false, true, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}
	nextField := getNextField(field)

	wantedNextField := Field{
		[]bool{false, false, false, false, false},
		[]bool{false, false, false, false, false},
		[]bool{false, true, false, true, false},
		[]bool{false, false, true, true, false},
		[]bool{false, false, true, false, false},
	}
	if !reflect.DeepEqual(nextField, wantedNextField) {
		test.Fail()
	}
}

func Test_unmarshalField_successful(test *testing.T) {
	text := ".O.\n..O\nOOO"
	field, err := unmarshalField(text)

	wantedField := Field{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	if !reflect.DeepEqual(field, wantedField) {
		test.Fail()
	}
	if err != nil {
		test.Fail()
	}
}

func Test_unmarshalField_withComments(test *testing.T) {
	text := "!comment #1\n!comment #2\n.O.\n..O\nOOO"
	field, err := unmarshalField(text)

	wantedField := Field{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	if !reflect.DeepEqual(field, wantedField) {
		test.Fail()
	}
	if err != nil {
		test.Fail()
	}
}

func Test_unmarshalField_withUnknownCharacter(test *testing.T) {
	text := ".O.\n..*\nOOO"
	field, err := unmarshalField(text)

	if field != nil {
		test.Fail()
	}
	if err == nil || err.Error() != "unknown character '*'" {
		test.Fail()
	}
}

func Test_unmarshalField_withInconsistentLength(test *testing.T) {
	text := ".O.\n..\nOOO"
	field, err := unmarshalField(text)

	if field != nil {
		test.Fail()
	}
	if err == nil || err.Error() != "inconsistent length of line 2" {
		test.Fail()
	}
}

func Test_marshalField(test *testing.T) {
	field := Field{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	result := marshalField(field)

	wantedResult := ".O.\n..O\nOOO\n"
	if result != wantedResult {
		test.Fail()
	}
}
