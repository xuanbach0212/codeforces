//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	ai := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &ai[i])
	}

	fmt.Fprintln(writer, "YES")
	fmt.Fprintln(writer, "NO")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		solve(reader, writer)
	}
}
