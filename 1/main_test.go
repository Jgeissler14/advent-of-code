package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	exp := 142
	res, err := Calibrate(input)
	if err != nil {
		t.Fatalf("Main func returned %v", err)
	}

	if res != exp {
		t.Fatalf(`Got %v, expected %v`, res, exp)
	}
}
