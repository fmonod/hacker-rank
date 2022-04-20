package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for i := int32(1); i <= n; i++ {
		spaces := n - i
		line := ""
		line += strings.Repeat(" ", int(spaces))
		line += strings.Repeat("#", int(i))
		fmt.Println(line)
	}

}

func main() {
	n := int32(4)
	staircase(n)
}
