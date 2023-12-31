/*
플로이드
https://www.acmicpc.net/problem/11404
*/
package shortestpath

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P11404() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	graph := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			if i != j {
				graph[i][j] = 1e9
			}
		}
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		c, _ := strconv.Atoi(line[2])

		graph[a][b] = min11404(c, graph[a][b])
	}

	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				graph[i][j] = min11404(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	for i := range graph {
		if i == 0 {
			continue
		}
		for j := range graph[i] {
			if j == 0 {
				continue
			}
			if graph[i][j] == 1e9 {
				fmt.Printf("0 ")

			} else {
				fmt.Printf("%d ", graph[i][j])
			}
		}
		fmt.Println()
	}
}
func min11404(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
플로이드 워셜
*/
