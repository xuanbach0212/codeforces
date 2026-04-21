//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var y int
	fmt.Fscan(reader, &y)

	for {
		y++

		t := (y / 1000) % 10
		h := y / 100 % 10
		ty := (y / 10) % 10
		d := (y / 1) % 10

		if t != h && h != ty && ty != d && d != t && t != ty && d != h {
			fmt.Fprintln(writer, y)
			return
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
