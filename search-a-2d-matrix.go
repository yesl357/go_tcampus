package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	x, y := searchMN(5, len(matrix[0]), len(matrix))
	fmt.Println(x, y)
	fmt.Println(searchMatrix(matrix, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	totalCount := len(matrix) * len(matrix[0])

	start := 0
	end := totalCount - 1

	for {
		if start+1 == end || start == end {
			if (matrix[start/len(matrix[0])][start%len(matrix[0])] == target) || (matrix[end/len(matrix[0])][end%len(matrix[0])] == target) {
				return true
			} else {
				return false
			}
		}
		mid := (start + end) / 2
		x, y := searchMN(mid, len(matrix[0]), len(matrix))
		if matrix[x][y] > target {
			end = mid
		} else {
			start = mid
		}
	}

}

func searchMN(mid, m, n int) (x, y int) {
	y = mid % m
	x = mid / m

	return x, y
}
