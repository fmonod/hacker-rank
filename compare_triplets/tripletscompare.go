package main

import (
	"fmt"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	result := []int32{0, 0}
	for i := 0; i < len(a); i++ {
		valA := a[i]
		valB := b[i]
		if valA > valB {
			result[0] += 1
		} else if valB > valA {
			result[1] += 1
		}
	}
	return result

}

func main() {
	// a := []int32{5, 6, 7}
	// b := []int32{3, 6, 10}
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}

	result := compareTriplets(a, b)

	fmt.Print(result)
}
