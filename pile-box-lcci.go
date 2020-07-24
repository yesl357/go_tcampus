package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	box := [][]int{
		{9, 9, 10}, {8, 10, 9}, {9, 8, 10}, {9, 8, 10}, {10, 8, 8}, {9, 8, 9}, {9, 8, 8}, {8, 9, 10}, {10, 9, 10}, {8, 8, 10}, {10, 9, 10}, {10, 9, 8}, {8, 9, 9}, {9, 10, 8}, {8, 9, 9}, {10, 10, 9}, {8, 9, 10}, {8, 10, 10}, {8, 9, 10}, {10, 10, 8}, {10, 10, 9}, {9, 10, 10}, {10, 8, 9}, {10, 10, 9}, {8, 9, 10}, {8, 8, 9}, {8, 10, 10}, {9, 9, 10}, {10, 8, 8}, {10, 10, 8}, {8, 9, 9}, {8, 9, 8}, {10, 10, 8}, {8, 10, 8}, {10, 9, 10}, {9, 9, 10}, {9, 9, 9}, {8, 9, 8}, {9, 8, 8}, {8, 9, 10}, {10, 10, 8}, {9, 9, 9}, {10, 10, 10}, {10, 9, 8}, {9, 8, 9}, {8, 8, 10}, {8, 8, 8}, {8, 8, 8}, {8, 9, 10}, {10, 9, 8}, {8, 10, 8}, {8, 10, 10}, {9, 10, 10}, {8, 8, 9}, {9, 9, 9}, {10, 8, 8}, {8, 10, 10}, {9, 10, 9}, {9, 9, 8}, {8, 10, 9}, {9, 8, 8}, {9, 9, 10}, {9, 10, 10}, {8, 8, 10},
	}

	firstIndex := 3 //按第二列排序
	sortedBox := arraySort(box, firstIndex-1)

	fmt.Println(pileBox(sortedBox))
	fmt.Println(pileBox2(box))
}

func pileBox2(box [][]int) int {
	firstIndex := 3 //按第二列排序
	box = arraySort(box, firstIndex-1)
	length := len(box)
	maxHeight := 0

	var dp [3001]int
	for i := length - 1; i >= 0; i-- {
		dp[i] = box[i][2]

		for j := length - 1; j > i; j-- {
			if (box[j][0] < box[i][0]) && (box[j][1] < box[i][1]) && (box[j][2] < box[i][2]) {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+box[i][2])))
			}
		}
		maxHeight = int(math.Max(float64(dp[i]), float64(maxHeight)))
	}

	return maxHeight
}

func pileBox(box [][]int) int {
	length := len(box)
	maxHeight := 0
	for i := 0; i < length; i++ {

		res := dfs(box, i)

		for i, _ := range res {
			if res[i] > maxHeight {
				maxHeight = res[i]
			}
		}
	}

	return maxHeight
}

func dfs(box [][]int, i int) []int {
	res2 := make([]int, 0)

	isLastFlag := true
	for j := i + 1; j < len(box); j++ {
		if (box[j][0] < box[i][0]) && (box[j][1] < box[i][1]) && (box[j][2] < box[i][2]) {
			isLastFlag = false
			resJ := dfs(box, j)
			for k, _ := range resJ {

				res2 = append(res2, resJ[k]+box[i][2])
			}
		}
	}

	if isLastFlag {
		return []int{box[i][2]}
	}

	return res2
}

//按指定规则对nums进行排序(注：此firstIndex从0开始)
func arraySort(nums [][]int, firstIndex int) [][]int {
	//检查
	if len(nums) <= 1 {
		return nums
	}
	if firstIndex < 0 || firstIndex > len(nums[0])-1 {
		fmt.Println("Warning: Param firstIndex should between 0 and len(nums)-1. The original array is returned.")
		return nums
	}
	//排序
	mIntArray := &IntArray{nums, firstIndex}
	sort.Sort(mIntArray)
	return mIntArray.mArr
}

type IntArray struct {
	mArr       [][]int
	firstIndex int
}

//IntArray实现sort.Interface接口
func (arr *IntArray) Len() int {
	return len(arr.mArr)
}
func (arr *IntArray) Swap(i, j int) {
	arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
}
func (arr *IntArray) Less(i, j int) bool {
	arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]
	for index := arr.firstIndex; index < len(arr1); index++ {
		if arr1[index] > arr2[index] {
			return true
		} else if arr1[index] < arr2[index] {
			return false
		}
	}
	return i > j
}
