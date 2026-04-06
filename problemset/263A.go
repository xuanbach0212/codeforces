//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var x int
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Fscan(reader, &x)
			if x == 1 {
				fmt.Fprintln(writer, math.Abs(float64(i-3))+math.Abs(float64(j-3)))
				return
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	// var t int
	// fmt.Fscan(reader, &t)
	//
	// for i := 0; i < t; i++ {
	// 	solve(reader, writer)
	// }
	solve(reader, writer)
}
