package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var min int64
	min = math.MaxInt64
	var max int64
	max = math.MinInt64
	var sum int64
	skip := 0
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := 0; j < len(arr); j++ {
			if j != skip {
				sum += int64(arr[j])
			}
		}
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
		skip++
	}

	fmt.Printf("%v %v", min, max)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
