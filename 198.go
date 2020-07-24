package main

import (
	"fmt"
	"sort"
)

func main() {

	numWaterBottles(9, 3)
	numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
}

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	emptyBottles := numBottles
	for {
		if emptyBottles < numExchange {
			break
		}
		res += int(emptyBottles / numExchange)

		emptyBottles = emptyBottles - int(emptyBottles/numExchange)*numExchange + (emptyBottles / numExchange)

	}

	return res
}

/**
https://leetcode-cn.com/problems/number-of-good-pairs/
*/
func numIdenticalPairs(nums []int) int {
	sort.Ints(nums[:])

	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] != nums[i] {
				res = res + (j - i - 1)
				break
			}

			if j == len(nums)-1 {
				res = res + (j - i)
			}
		}

		fmt.Println(i)
		fmt.Println(res)
	}

	return res
}
