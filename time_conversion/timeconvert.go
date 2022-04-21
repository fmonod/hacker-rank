package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	var timeResult string
	time12, err := time.Parse("15:04:05PM", s)
	if err != nil {
		return err.Error()
	}
	timeResult = time12.Format("15:04:05")
	return timeResult
}

func main() {
	time := "07:05:45PM"
	fmt.Println(timeConversion(time))
}
