//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n int
	fmt.Fscanln(reader, &n)

	x := 0
	for i := 0; i < n; i++ {
		var op string
		fmt.Fscan(reader, &op)
		if op[1] == '+' {
			x += 1
		} else {
			x -= 1
		}
	}

	fmt.Fprintln(writer, x)
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
