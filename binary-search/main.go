package main

import "fmt"

type Given struct {
	array  []int
	target int
}

func main() {
	testCases := []Given{
		{
			[]int{-1, 0, 3, 5, 9, 12},
			9,
		},
	}

	for _, test := range testCases {
		fmt.Printf("result:\"%v\", case %v\n", search(test.array, test.target), test)
	}
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	if len(nums) < 2 && target == nums[0] {
		return 0
	}

	for left <= right {
		ipivot := (right-left)/2 + left

		if ((right-left)/2+left)%2 == 0 {
			ipivot++
		}

		if target == nums[ipivot] {
			return ipivot
		}

		if target > nums[ipivot] {
			left = ipivot + 1
			continue
		}

		if target < nums[ipivot] {
			right = ipivot - 1
			continue
		}
	}

	return -1
}
