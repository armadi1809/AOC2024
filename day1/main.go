package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputArr := strings.Split(string(input), "\n")
	var list1 []int
	var list2 []int
	for _, inpLine := range inputArr {
		nums := strings.Split(inpLine, "   ")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic("Invalid input")
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic("Invalid input")
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	res := 0
	res2 := 0
	for i := 0; i < len(list1); i++ {
		res += int(math.Abs(float64(list1[i] - list2[i])))
		res2 += list1[i] * countOccurrences(list2, list1[i])
	}
	fmt.Println("Part1: ", res)
	fmt.Println("Part2: ", res2)

}

func countOccurrences(slice []int, target int) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}
	return count
}
