package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var sliceWanted []int
	var wantedInt, wantedInt1 int
	result := 0
	result1 := 0

	if (args[1] != "+") && (args[1] != "-") && (args[1] != "*") && (args[1] != "/") && (args[1] != "%") {
		return
	}
	if len(args[0]) == 1 {
		wantedInt = int(args[0][0]) - '0'
		sliceWanted = append(sliceWanted, wantedInt)
	} else {
		for _, ch := range args[0] {
			result = result*10 + int(ch) - '0'
		}
		if (result > -922337203685477580) && (result < 922337203685477580) {
			sliceWanted = append(sliceWanted, result)
		} else {
			return
		}

	}
	if len(args[2]) == 1 {
		wantedInt1 = int(args[2][0]) - '0'
		sliceWanted = append(sliceWanted, wantedInt1)
	} else {
		for _, ch := range args[2] {
			result1 = result1*10 + int(ch) - '0'
		}
		if (result1 > -922337203685477580) && (result1 < 922337203685477580) {
			sliceWanted = append(sliceWanted, result1)
		} else {
			return
		}

	}

	fmt.Println(sliceWanted)

	switch args[1] {
	case "+":
		fmt.Println(sliceWanted[0] + sliceWanted[1])
	case "-":
		fmt.Println(sliceWanted[0] - sliceWanted[1])
	case "*":
		fmt.Println(sliceWanted[0] * sliceWanted[1])
	case "/":
		fmt.Println(sliceWanted[0] / sliceWanted[1])
	case "%":
		fmt.Println(sliceWanted[0] % sliceWanted[1])
	}
}
