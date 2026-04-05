//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var t int
	fmt.Fscan(reader, &t)

	mulNums := make([][]int, t)
	for i := 0; i < t; i++ {
		for j := 0; j < 7; j++ {
			var a int
			fmt.Fscan(reader, &a)
			mulNums[i] = append(mulNums[i], a)
		}
	}

	for _, nums := range mulNums {
		sort.Ints(nums)
		sum := 0
		for i, num := range nums {
			if i < 6 {
				sum -= num
			} else {
				sum += num
			}
		}

		fmt.Fprintln(writer, sum)
		sum = 0
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
