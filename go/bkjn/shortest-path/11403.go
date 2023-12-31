/*
경로 찾기
https://www.acmicpc.net/problem/11403
*/
package shortestpath

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P11403() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		row := make([]int, n+1)
		row[0] = 0
		for j, v := range line {
			row[j+1], _ = strconv.Atoi(v)
		}
		graph[i] = row
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if graph[i][j] == 0 {
				graph[i][j] = 1e9
			}
		}
	}
	//floyd marshall
	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				graph[i][j] = min11403(graph[i][j], graph[i][k]+graph[k][j])
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
				fmt.Printf("1 ")
			}
		}
		fmt.Println()
	}
}
func min11403(a, b int) int {
	if a > b {
		return b
	}

	return a
}

/*
플로이드 워셜 알고리즘 이용.
graph[i][j] = min11403(graph[i][j], graph[i][k]+graph[k][j])
*/
