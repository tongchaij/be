package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var MAX_POSITION int
var MIN_SUM int
var MIN_OUTPUT string

var CODE_OPS = map[byte]int{
	"L"[0]: -1,
	"R"[0]: 1,
	"="[0]: 0}

var operators []int
var digits []int

func getOperators(input string) []int {

	operators = make([]int, len(input))

	for i := 0; i < len(input); i++ {
		operators[i] = CODE_OPS[input[i]]
	}

	return operators
}

func initMinSum(input string) int {

	minSumStr := "9"
	for i := 0; i < len(input); i++ {
		minSumStr += "9"
	}

	minSum, _ := strconv.Atoi(minSumStr)

	return minSum
}

func sumDigit() int {
	sum := 0
	for i := 0; i <= MAX_POSITION; i++ {
		sum += digits[i]
	}
	return sum
}

func getMinPath(position int, digit int) {

	digits[position] = digit

	if position == MAX_POSITION { // Last Digit
		if digit <= position {
			sum := sumDigit()
			if sum < MIN_SUM {
				MIN_SUM = sum
				MIN_OUTPUT = getOutput()
			}
		}
		return
	} else { // 0 to Last Digit-1

		op := operators[position]
		for i := digit + op; i >= 0 && i <= MAX_POSITION; i += op {
			getMinPath(position+1, i)
			if op == 0 {
				break
			}
		}

	}

}

func getFinalResult(input string) string {

	operators = getOperators(input)
	MIN_SUM = initMinSum(input)
	digits = make([]int, len(input)+1)
	MIN_OUTPUT = "Not Found"

	for i := 0; i <= MAX_POSITION; i++ {
		getMinPath(0, i)
	}

	return MIN_OUTPUT
}

func getOutput() string {
	output := ""
	for i := 0; i < len(digits); i++ {
		output += fmt.Sprintf("%d", digits[i])
	}
	return output
}

func test() {
	inputs := [10]string{"LLRR=", "==RLL", "=LLRR", "RRL=R", "RRRRR", "LLLLL", "=====", "LRLRL", "RLRLRLL", "RL==LR"}
	//inputs := [1]string{"RRRRRRRRRRRRRRRRRRR"}
	//inputs := [1]string{"LL"}

	fmt.Printf("inputs:%v\n", inputs)

	for i := 0; i < len(inputs); i++ {
		MAX_POSITION = len(inputs[i])
		if MAX_POSITION > 9 {
			fmt.Printf("%d : %s => %s\n", i+1, inputs[i], "Invalid Code")
		} else {
			fmt.Printf("%d : %s => %s\n", i+1, inputs[i], getFinalResult(inputs[i]))
		}
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter encoded string (Ctrl+C to Stop): ")
		input, _ := reader.ReadString('\n')

		input = input[:len(input)-1]

		if input == "test" {
			test()
		} else {
			MAX_POSITION = len(input)
			if MAX_POSITION > 9 {
				fmt.Printf("=> %s\n", "Invalid Code")
			} else {
				fmt.Printf("=> %s\n", getFinalResult(input))
			}
		}
	}
}
