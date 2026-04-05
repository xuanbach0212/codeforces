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

	arr := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	for _, a := range arr {
		fmt.Fprintln(writer, a)
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
