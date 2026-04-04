//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, k int
	fmt.Fscan(reader, &n, &k)

	participants := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &participants[i])
	}

	res := 0
	for _, p := range participants {
		if p >= participants[k-1] && p > 0 {
			res += 1
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
