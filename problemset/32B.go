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
	res := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			res = append(res, '0')
		} else {
			if s[i+1] == '.' {
				res = append(res, '1')
				i++
			} else {
				res = append(res, '2')
				i++
			}
		}
	}

	fmt.Fprintln(writer, string(res))
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
