package main

import (
	"reflect"
	"testing"
)

func Test_getWidth(test *testing.T) {
	field := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	width := getWidth(field)
	if width != 3 {
		test.Fail()
	}
}

func Test_getHeight(test *testing.T) {
	field := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	height := getHeight(field)
	if height != 2 {
		test.Fail()
	}
}

func Test_getCell_withTrue(test *testing.T) {
	field := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, true /* ! */, false},
	}
	cell := getCell(field, 1, 2)
	if cell != true {
		test.Fail()
	}
}

func Test_getCell_withFalse(test *testing.T) {
	field := [][]bool{
		[]bool{false, false, false},
		[]bool{false /* ! */, false, false},
		[]bool{false, true, false},
	}
	cell := getCell(field, 0, 1)
	if cell != false {
		test.Fail()
	}
}

func Test_setCell_withTrue(test *testing.T) {
	field := [][]bool{
		[]bool{false, false, false},
		[]bool{false /* ! */, false, false},
		[]bool{false, false, false},
	}
	setCell(field, 0, 1, true)

	wantedField := [][]bool{
		[]bool{false, false, false},
		[]bool{true /* ! */, false, false},
		[]bool{false, false, false},
	}
	if !reflect.DeepEqual(field, wantedField) {
		test.Fail()
	}
}

func Test_setCell_withFalse(test *testing.T) {
	field := [][]bool{
		[]bool{false, false, false},
		[]bool{true /* ! */, false, false},
		[]bool{false, false, false},
	}
	setCell(field, 0, 1, false)

	wantedField := [][]bool{
		[]bool{false, false, false},
		[]bool{false /* ! */, false, false},
		[]bool{false, false, false},
	}
	if !reflect.DeepEqual(field, wantedField) {
		test.Fail()
	}
}
