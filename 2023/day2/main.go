package main

import (
	"strings"
    "regexp"
    "strconv"
    "fmt"
)

func main() {
    input := `Game 1:  3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2:  1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3:  8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4:  1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5:  6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	_, err := CalculatePossibleGames(input)
	if err != nil {
		panic(err)
	}
}

func CalculatePossibleGames(input string) (int64, error) {
	games := strings.Split(input, "\n")
	possible_games := int64(0)

    // Loop though all lines
	for _, game := range games {
        color_map := map[string]int{"red":12, "green":13, "blue":14}

        // find line id
        game_string := strings.Split(game, ":")

        split_id := strings.Split(game_string[0], " ")
        id := split_id[len(split_id)-1]
        fmt.Println(id)

        // Regular expression to match ";", " ", and ","
        re := regexp.MustCompile(`[;,]+`)

        // Replace all occurrences with an empty string
        result := re.ReplaceAllString(game_string[1], "")
        game_array := strings.Split(result[1:], " ")
        if len(game_array) % 2 != 0 {
            panic("Invalid input")
        }
        possible := true
        for i := 0; i < len(game_array); i+=2 {
            num, err := strconv.ParseInt(game_array[i], 10, 64)
            if err != nil {
                panic("Expected string, got int")
            }
            if int(num) > color_map[game_array[i+1]] {
                possible = false
            }
        }
        if possible {
            num, _ := strconv.ParseInt(id, 10, 64)
            possible_games += int64(num)
        }

	}
	return possible_games, nil
}
