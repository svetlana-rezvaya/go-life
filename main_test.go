package main

import "testing"

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
