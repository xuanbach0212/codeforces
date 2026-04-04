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

	res := 0
	for i := 0; i < n; i++ {
		var p, v, t int
		fmt.Fscan(reader, &p, &v, &t)

		if p+v+t >= 2 {
			res++
		}
	}

	fmt.Fprintln(writer, res)
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
