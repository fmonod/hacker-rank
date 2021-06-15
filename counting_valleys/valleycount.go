package main

import (
	"fmt"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var result int32
	pathArr := []byte(path)
	altitude := 0
	for _, step := range pathArr {
		if step == 'U' {
			altitude += 1
		} else if step == 'D' {
			altitude -= 1
		}
		if altitude == 0 && step != 'D' {
			result += 1
		}
	}
	return result
}

func main() {
	steps := int32(8)

	path := "UDDDUDUU"

	result := countingValleys(steps, path)

	fmt.Print(result)
}
