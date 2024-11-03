package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	res, err := CalculatePossibleGames(string(data))
	if err != nil {
		panic(err)
	}
    fmt.Println(res)
}

func CalculatePossibleGames(input string) (int64, error) {
	games := strings.Split(input, "\n")
	possibleGames := int64(0)

	// Loop through all lines
	for _, game := range games {
		if game == "" {
			continue // Skip empty lines
		}

		colorMap := map[string]int{"red": 12, "green": 13, "blue": 14}

		// Split the game line by ":"
		gameString := strings.Split(game, ":")
		if len(gameString) < 2 {
			return 0, fmt.Errorf("invalid line format: %s", game)
		}

		// Get the ID from the line
		splitID := strings.Split(gameString[0], " ")
		id := splitID[len(splitID)-1]

		// Regular expression to remove ";", " ", and ","
		re := regexp.MustCompile(`[;,]+`)
		result := re.ReplaceAllString(gameString[1], "")
		gameArray := strings.Split(result[1:], " ")

		if len(gameArray)%2 != 0 {
			return 0, fmt.Errorf("invalid input format in line: %s", game)
		}

		possible := true
		for i := 0; i < len(gameArray); i += 2 {
			num, err := strconv.ParseInt(gameArray[i], 10, 64)
			if err != nil {
				return 0, fmt.Errorf("expected integer, got string in line: %s", game)
			}
			if int(num) > colorMap[gameArray[i+1]] {
				possible = false
				break
			}
		}

		if possible {
			num, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse ID: %s", id)
			}
			possibleGames += num
		}
	}
	return possibleGames, nil
}
