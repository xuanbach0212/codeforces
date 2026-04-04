//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var s string
	fmt.Fscan(reader, &s)

	if len(s) > 10 {
		fmt.Fprintf(writer, "%c%d%c\n", s[0], len(s)-2, s[len(s)-1])
	} else {
		fmt.Fprintln(writer, s)
	}
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
