package main

import (
	"fmt"
	"math/rand"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	var result int32
	var currentHGSum int32
	// Result takes the value of the first hourglass
	result = arr[0][0] + arr[0][1] + arr[0][2] + arr[1][1] + arr[2][0] + arr[2][1] + arr[2][2]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if i+2 < len(arr) && j+2 < len(arr[0]) {
				currentHGSum = arr[i][j] + arr[i][j+1] + arr[i][j+2]        // First row, three positions
				currentHGSum += arr[i+1][j+1]                               // Second row, second position
				currentHGSum += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2] // Third row, three positions
			}
			if currentHGSum > result {
				result = currentHGSum
			}
		}
	}

	return result
}

func main() {
	var s []string
	arr := make([][]int32, 6)
	for a := 0; a < 6; a++ {
		arr[a] = make([]int32, 6)
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			arr[i][j] = rand.Int31n(30)
			s = append(s, fmt.Sprintf("%d", arr[i][j]))
		}
	}
	fmt.Println("Matrix: ")
	fmt.Printf("%s\n", s)

	result := hourglassSum(arr)

	fmt.Println("Result: ")
	fmt.Printf("%d\n", result)
}
