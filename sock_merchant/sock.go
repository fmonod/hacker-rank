package main

import (
	"fmt"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var pairs int32
	pairsMap := make(map[int32]int32)
	for _, val := range ar {
		if _, hasKey := pairsMap[val]; hasKey {
			pairsMap[val] = pairsMap[val] + 1
		} else {
			pairsMap[val] = 1
		}
	}
	for _, value := range pairsMap {
		if value > 1 {
			pair := value / 2
			pairs += pair
		}
	}
	return pairs
}

func main() {

	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	n := int32(len(ar))

	result := sockMerchant(n, ar)

	fmt.Print(result)
}
