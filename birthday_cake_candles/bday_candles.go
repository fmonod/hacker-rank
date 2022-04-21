package main

import "fmt"

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var sumTallestCandles int32
	var tallestCandle int32
	for j := 0; j < len(candles); j++ {
		if candles[j] > tallestCandle {
			tallestCandle = candles[j]
		}
	}
	for i := 0; i < len(candles); i++ {
		if candles[i] == tallestCandle {
			sumTallestCandles++
		}
	}

	return sumTallestCandles
}

func main() {
	arr := []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))
}
