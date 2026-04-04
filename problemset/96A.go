//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var players string
	fmt.Fscan(reader, &players)

	if len(players) < 7 {
		fmt.Fprintln(writer, "NO")
		return
	}

	count := 0
	var last byte
	for i := 0; i < len(players); i++ {
		if players[i] == last {
			count++
		} else {
			count = 1
			last = players[i]
		}
		if count == 7 {
			fmt.Fprintln(writer, "YES")
			return
		}
	}
	fmt.Fprintln(writer, "NO")
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
	//

	solve(reader, writer)
}
