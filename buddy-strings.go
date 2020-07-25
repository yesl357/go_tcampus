package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	res := buddyStrings("acc", "acc")
	fmt.Println(res)
	math.Sqrt(1)

	nums := []int{3, 5, 2, 1, 6, 4}

	wiggleSort(nums)
}

func wiggleSort(nums []int) {
	res := make([]int, 0)
	for _, v := range nums {
		res = append(res, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res[:])))

	var endIndex int
	if len(res)%2 == 0 {
		endIndex = len(res)/2 - 1
	} else {
		endIndex = len(res) / 2
	}

	for i := 0; i <= endIndex; i++ {
		nums[2*i] = res[i]
		if i < endIndex || len(res)%2 == 0 {
			nums[2*i+1] = res[len(res)-1-i]
		}
	}
}

func buddyStrings(A string, B string) bool {

	lenA, lenB := len(A), len(B)
	if lenA != lenB {
		return false
	}

	diff := 0
	diffItem := make([][]string, 0)
	for i := 0; i < lenA; i++ {
		if A[i] != B[i] {
			diff++
			diffItem = append(diffItem, []string{string(A[i]), string(B[i])})
		}
	}

	if diff == 2 {
		if diffItem[0][0] == diffItem[1][1] && diffItem[1][0] == diffItem[0][1] {
			return true
		}
	}

	if diff == 0 {
		wordMapCount := make(map[string]int)
		for i := 0; i < lenA; i++ {
			wordMapCount[string(A[i])]++
		}

		for _, value := range wordMapCount {
			if value > 1 {
				return true
			}
		}
	}

	return false
}
