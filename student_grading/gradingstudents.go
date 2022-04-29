package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	result := make([]int32, 0)
	gradeTop := 100
	for _, v := range grades {
		if v < 38 {
			result = append(result, v)
		} else {
			for i, found := 1, false; i < gradeTop && !found; i++ {
				nextMultipleFive := v + int32(i)
				if nextMultipleFive%5 == 0 {
					diff := nextMultipleFive - v
					if diff < 3 {
						result = append(result, nextMultipleFive)
					} else {
						result = append(result, v)
					}
					found = true
				}
			}

		}

	}

	return result
}

func main() {
	ar := []int32{73, 67, 40, 33}
	result := gradingStudents(ar)
	fmt.Println(result)
}
