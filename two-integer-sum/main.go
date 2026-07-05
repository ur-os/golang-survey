package main

import "fmt"

type Given struct {
	integers []int
	target   int
}

func main() {
	testCases := []Given{
		Given{
			[]int{2, 7, 11, 15},
			9,
		},
	}

	for _, payload := range testCases {
		fmt.Printf("\"%s\", result: %v\n", payload, twoSum(payload.integers, payload.target))
	}
}

func twoSum(numbers []int, target int) []int {
	for i, value := range numbers {

	}

	return []int{}
}
