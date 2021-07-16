package main

import "fmt"

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var positives int
	var negatives int
	var zeroes int

	for _, v := range arr {

		if v > 0 {
			positives++
		} else if v == 0 {
			zeroes++
		} else {
			negatives++
		}

	}

	fmt.Println(float32(positives) / float32(len(arr)))
	fmt.Println(float32(negatives) / float32(len(arr)))
	fmt.Println(float32(zeroes) / float32(len(arr)))

}

func main() {

	a := []int32{1, 1, 0, -1, -1}

	plusMinus(a)

}
