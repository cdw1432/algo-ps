/*
DFS와 BFS
https://www.acmicpc.net/problem/1260
*/
package dfsbfs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1260() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())

	n, _ := strconv.Atoi(firstLine[0])
	m, _ := strconv.Atoi(firstLine[1])
	v, _ := strconv.Atoi(firstLine[2])

	arr := make([][]int, n+1)
	visitedbfs := make([]bool, n+1)
	visiteddfs := make([]bool, n+1)
	for i := range arr {
		arr[i] = make([]int, n+1)
		visitedbfs[i] = false
		visiteddfs[i] = false
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])

		arr[a][b] = 1
		arr[b][a] = 1
	}

	dfs := func(v int) {
		var stack []int

		stack = append(stack, v)
		for len(stack) > 0 {
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !visiteddfs[c] {
				fmt.Printf("%d ", c)
				visiteddfs[c] = true

				for i := n; i > 0; i-- {
					if arr[c][i] == 1 {
						stack = append(stack, i)
					}
				}
			}
		}

	}
	bfs := func(v int) {
		var queue []int

		queue = append(queue, v)
		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			if !visitedbfs[c] {
				visitedbfs[c] = true
				fmt.Printf("%v ", c)
				for i := 1; i < n+1; i++ {
					if arr[c][i] == 1 {
						queue = append(queue, i)
					}
				}
			}
		}
	}
	dfs(v)
	fmt.Println()
	bfs(v)
}

/*
dfs -> stack
bfs -> queue
*/
