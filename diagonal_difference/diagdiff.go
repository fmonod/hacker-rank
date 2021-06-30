package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var primary int32
	var secondary int32
	var result int32
	for i := 0; i < len(arr); i++ {
		primary += arr[i][i]
		secondary += arr[len(arr)-i-1][i]
	}
	result += int32(math.Abs(float64(primary - secondary)))
	return result
}

func main() {
	a := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	result := diagonalDifference(a)

	fmt.Print(result)
}
