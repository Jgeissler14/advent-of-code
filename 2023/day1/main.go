package main

import (
	"strings"
	"unicode"
    "strconv"
)

func main() {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	_, err := Calibrate(input)
	if err != nil {
		panic(err)
	}
}

func Calibrate(input string) (int, error) {
	lines := strings.Split(input, "\n")
	sum := 0

    // Loop though all lines
	for _, line := range lines {
		l, r := 0, len(line)-1

        // Move pointers l and r inward until they reach a number
        // should do error handling if no numbers exist but im lazy
        for l <= r && !unicode.IsNumber(rune(line[l])) {
            l++
        }
        for r >= l && !unicode.IsNumber(rune(line[r])) {
            r--
        }

        //concat the two strings to make the line number
		local_sum := string(line[l]) + string(line[r])

        // convert line string to num then add to the sum
        num, err := strconv.Atoi(local_sum)
        if err != nil {
            return 0, err
        }
		sum += num
	}
	return sum, nil
}
