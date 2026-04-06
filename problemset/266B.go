//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, t int
	fmt.Fscan(reader, &n, &t)
	var s string
	fmt.Fscan(reader, &s)

	row := []byte(s)

	for i := 0; i < t; i++ {
		for j := 0; j < n-1; j++ {
			if row[j] == 'B' && row[j+1] == 'G' {
				row[j], row[j+1] = row[j+1], row[j]
				j++
			}
		}
	}

	fmt.Fprintln(writer, string(row))
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
