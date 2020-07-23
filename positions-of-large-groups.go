package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//res := largeGroupPositions("abcdddeeeeaabbbcd")
	//fmt.Println(res)
	//
	//nums := []int{
	//	3, 6, 2, 3,
	//}
	//res2 := largestPerimeter(nums)
	//fmt.Println(res2)
	//
	//res3 := angleClock(3, 30)
	//fmt.Println(res3)

	fmt.Println(numWays(92))
}

func numWays(n int) int {
	res := [100]int64{}
	res[1] = 1
	res[2] = 2

	for i := 3; i <= n; i++ {
		res[i] = res[i-2] + res[i-1]
		res[i] = res[i] % 1000000007
	}
	return int(res[n])
}

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	hourFloat := float64(hour)
	minutesFloat := float64(minutes)

	angleHour := float64(360/12) * (hourFloat + minutesFloat/60)
	angleMin := float64(360/60) * (minutesFloat)

	angle1 := math.Abs(angleHour - angleMin)
	angle2 := math.Abs(360 + angleHour - angleMin)

	if angle1 > angle2 {
		if angle2 > 180.0 {
			return 360.0 - angle2
		} else {
			return angle2
		}
	}

	if angle1 > 180.0 {
		return 360.0 - angle1
	} else {
		return angle1
	}
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
