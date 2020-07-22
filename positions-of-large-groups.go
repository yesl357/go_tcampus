package main

import (
	"fmt"
	"sort"
)

func main() {
	res := largeGroupPositions("abcdddeeeeaabbbcd")
	fmt.Println(res)

	nums := []int{
		3, 6, 2, 3,
	}
	res2 := largestPerimeter(nums)
	fmt.Println(res2)
}

func largeGroupPositions(S string) [][]int {
	//res := make([][]int, 0)
	min := 0
	res := make([][]int, 0)
	for i, _ := range S {
		if i == len(S)-1 || S[i] != S[i+1] {
			if i-min+1 >= 3 {
				res = append(res, []int{min, i})
			}
			min = i + 1
		}
	}

	return res
}

func largestPerimeter(A []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(A[:])))

	largest := 0
	for i := 0; i <= len(A)-3; i++ {
		if A[i] < A[i+1]+A[i+2] {
			largest = A[i] + A[i+1] + A[i+2]
			break
		}
	}
	return largest
}
