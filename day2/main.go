package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputArr := strings.Split(string(input), "\n")
	res := 0
	res2 := 0
	for _, inpLine := range inputArr {
		if isReportSafePart1(inpLine) {
			res += 1
		}
		if isReportSafePart1(inpLine) || isReportSafePart2(inpLine) {
			res2 += 1
		}
	}
	fmt.Println("Part 1: ", res)
	fmt.Println("Part 2: ", res2)
}
func isReportSafePart1(inpLine string) bool {
	nums := getLevelsFromReport(inpLine)
	return validate(nums)
}

func isReportSafePart2(inpLine string) bool {
	nums := getLevelsFromReport(inpLine)
	for i := range nums {
		removedNums := append([]int{}, nums[:i]...)
		removedNums = append(removedNums, nums[i+1:]...)
		if validate(removedNums) {
			return true
		}
	}
	return false
}

func validate(input []int) bool {
	valid := true
	decreasing := true
	for j := range input {
		if j == 0 {
			continue
		}
		if math.Abs(float64(input[j]-input[j-1])) > 3 || math.Abs(float64(input[j]-input[j-1])) < 1 {
			valid = false
			break
		}
		if j == 1 {
			if input[j] > input[j-1] {
				decreasing = false
			}
		} else {
			if (decreasing && input[j] > input[j-1]) || (!decreasing && input[j] < input[j-1]) {
				valid = false
				break
			}
		}
	}
	return valid
}

func getLevelsFromReport(report string) []int {
	numsStrs := strings.Split(report, " ")
	var nums []int
	for _, numStr := range numsStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("Invalid input")
		}
		nums = append(nums, num)
	}
	return nums
}
