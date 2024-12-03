package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)
	i := 0
	res := 0
	res2 := 0
	enabled := true
	for i < len(inputStr) {
		if inputStr[i] == 'd' {
			if i < len(inputStr)-4 && inputStr[i:i+4] == "do()" {
				enabled = true
				i += 2
			} else if i < len(inputStr)-7 && inputStr[i:i+7] == "don't()" {
				enabled = false
				i += 5
			} else {
				i += 1
			}
			continue
		}
		if inputStr[i] != 'm' {
			i += 1
			continue
		}
		if i < len(inputStr)-4 && inputStr[i:i+4] != "mul(" {
			i += 1
			continue
		}

		i += 4
		j := i
		for j < len(inputStr) && inputStr[j] != ',' && j < i+3 && unicode.IsDigit([]rune(inputStr)[j]) {
			j += 1
		}
		if inputStr[j] != ',' {
			i = j
			continue
		}
		num1, err := strconv.Atoi(inputStr[i:j])
		if err != nil {
			panic("incorrect processing, should never happen")
		}
		j = j + 1
		i = j
		for j < len(inputStr) && inputStr[j] != ')' && j < i+3 && unicode.IsDigit([]rune(inputStr)[j]) {
			j += 1
		}
		if inputStr[j] != ')' {
			i = j
			continue
		}
		num2, err := strconv.Atoi(inputStr[i:j])
		if err != nil {
			panic("incorrect processing, should never happen")

		}
		res += num1 * num2
		if enabled {
			res2 += num1 * num2
		}

		i = j
	}
	fmt.Println("Part1: ", res)
	fmt.Println("Part2: ", res2)

}
