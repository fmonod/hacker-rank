package main

import (
	"fmt"
)

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	var result int32
	var jumped bool
	for key, value := range c {
		switch value {
		case 0:
			if key+2 < len(c) && c[key+2] == 0 && !jumped {
				jumped = true
				result += 1
			} else if key+1 < len(c) && c[key+1] == 0 {
				if jumped {
					jumped = false
				} else {
					result += 1
				}
			}
		case 1:
			jumped = false
			break
		}
	}
	return result
}

func main() {
	// c := []int32{0, 0, 1, 0, 0, 1, 0}
	c := []int32{0, 0, 0, 0, 1, 0}

	result := jumpingOnClouds(c)

	fmt.Print(result)
}
