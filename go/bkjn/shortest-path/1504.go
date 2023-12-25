/*
특정한 최단 경로
https://www.acmicpc.net/problem/1504
*/
package shortestpath

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1504() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)

	//vertex and edges
	v, _ := strconv.Atoi(parts[0])
	e, _ := strconv.Atoi(parts[1])

	var table = make([][]int, v+1)
	for i := 1; i < v+1; i++ {
		table[i] = make([]int, v+1)
	}
	for i := range table {
		for j := range table[i] {
			if i == j {
				table[i][j] = 0
			} else {
				table[i][j] = int(1e9)
			}
		}
	}
	for i := 0; i < e; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		table[a][b] = min(table[a][b], c)
		table[b][a] = min(table[b][a], c)
	}

	scanner.Scan()
	lastLine := scanner.Text()
	parts = strings.Fields(lastLine)

	v1, _ := strconv.Atoi(parts[0])
	v2, _ := strconv.Atoi(parts[1])

	//pathfinding
	for k := 1; k < v+1; k++ {
		for i := 1; i < v+1; i++ {
			for j := 1; j < v+1; j++ {
				table[i][j] = min(table[i][j], table[i][k]+table[k][j])
			}
		}
	}
	result := min(table[1][v1]+table[v1][v2]+table[v2][v],
		table[1][v2]+table[v2][v1]+table[v1][v])
	if result < int(1e9) {
		fmt.Println(result)
	} else {
		fmt.Println(-1)
	}
}

/*
플로이드 워셜 알고리즘으로 풀이.
경우의 수
시작 -> v1 -> v2 -> 끝
시작 -> v2 -> v1 -> 끝

두가지 중 작은 하나가 정답이다.

나는 플로이드워셜을 사용했다. 이유는 구현이 쉽기때문,
성능면에서 다익스트라 알고리즘을 사용하는게 좋다.
*/
