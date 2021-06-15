package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var arSum int32
	for _, number := range ar {
		arSum += number
	}
	return arSum
}

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}
	result := simpleArraySum(ar)
	fmt.Println(result)
}
