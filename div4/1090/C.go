//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		N := 3 * n
		arr := make([]int, N)
		l := 1
		h := N
		for i := 0; i < N; i += 3 {
			arr[i] = l
			arr[i+1] = h
			arr[i+2] = h - 1

			l += 1
			h -= 2
		}

		for i, num := range arr {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, num)
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	// 	var t int
	// 	fmt.Fscan(reader, &t)
	//
	// 	for i := 0; i < t; i++ {
	// 		solve(reader, writer)
	// 	}
	// }
	solve(reader, writer)
}
