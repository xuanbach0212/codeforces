//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n int

	fmt.Fscan(reader, &n)
	var x, y, z int
	sumX, sumY, sumZ := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x, &y, &z)
		sumX += x
		sumY += y
		sumZ += z
	}

	if sumX == 0 && sumY == 0 && sumZ == 0 {
		fmt.Fprintln(writer, "YES")
		return
	}

	fmt.Fprintln(writer, "NO")
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
