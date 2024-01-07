/*
최소비용 구하기
https://www.acmicpc.net/problem/1916
*/
package shortestpath

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1916() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}
	for i := range graph {
		for j := range graph[i] {
			graph[i][j] = 1e9
		}
	}
	for i := 0; i < m; i++ {
		scanner.Scan()

		part := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(part[0])
		b, _ := strconv.Atoi(part[1])
		c, _ := strconv.Atoi(part[2])

		graph[a][b] = min1916(graph[a][b], c)
	}
	scanner.Scan()
	part := strings.Fields(scanner.Text())
	s, _ := strconv.Atoi(part[0])
	e, _ := strconv.Atoi(part[1])

	//dijk
	distance := make([]int, n+1)
	visited := make([]bool, n+1)
	for i := range visited {
		visited[i] = false
		distance[i] = 1e9
	}
	smallest_node := func() int {
		minD := int(1e9)
		index := 0
		for i := 1; i < n+1; i++ {
			if distance[i] < minD && !visited[i] {
				minD = distance[i]
				index = i
			}
		}
		return index
	}
	dijk := func(s int) {
		distance[s] = 0
		for i := 1; i < n+1; i++ {
			c := smallest_node()
			visited[c] = true

			for j := 1; j < n+1; j++ {
				if !visited[j] && graph[c][j] < 1e9 {
					cost := distance[c] + graph[c][j]
					distance[j] = min1916(distance[j], cost)
				}
			}
		}
		fmt.Println(distance[e])
	}
	dijk(s)
}

func min1916(a, b int) int {
	if a > b {
		return b
	}

	return a
}

/*
dijkstra algorithm in go
*/
