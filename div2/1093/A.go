//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n int
	fmt.Fscan(reader, &n)

	ai := make([]int, n)
	counts := make(map[int]int)
	isDup := false
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &ai[i])
		counts[ai[i]]++
		if counts[ai[i]] > 1 {
			isDup = true
		}
	}

	if isDup {
		fmt.Fprintln(writer, -1)
		return
	}

	sort.Slice(ai, func(i, j int) bool {
		return ai[i] > ai[j]
	})

	temp := make([]string, len(ai))
	for i, v := range ai {
		temp[i] = strconv.Itoa(v)
	}

	fmt.Fprintln(writer, strings.Join(temp, " "))
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
