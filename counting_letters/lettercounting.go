package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	// Write your code here
	lettertocount := "a"
	countLetter := strings.Count(s, lettertocount)
	r := int(n) % len(s)
	rep := int(n) / len(s)
	countr := strings.Count(s[:r], lettertocount)
	result := int64(rep*countLetter + countr)
	return result
}

func main() {
	n := int64(1000000000000)
	s := "aba"

	result := repeatedString(s, n)

	fmt.Print(result)
}
