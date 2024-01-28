/*
Watering the Fields
https://www.acmicpc.net/problem/10021
*/
package bkjn

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func P10021() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(line[0])
	C, _ := strconv.Atoi(line[1])

	fields := [][]int{}
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		fields = append(fields, []int{x, y})
	}

	edges := [][]int{}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			xi := float64(fields[i][0])
			yi := float64(fields[i][1])
			xj := float64(fields[j][0])
			yj := float64(fields[j][1])
			cost := math.Pow(math.Abs(xi-xj), 2) + math.Pow(math.Abs(yi-yj), 2)
			if int(cost) >= C {
				edges = append(edges, []int{int(cost), i, j})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	ans := 0
	parents := make([]int, N+1)
	edges_cnt := 0
	for i := range parents {
		parents[i] = i
	}
	for i := range edges {
		cost := edges[i][0]
		a := edges[i][1]
		b := edges[i][2]
		if findP10021(parents, a) != findP10021(parents, b) {
			unionP10021(parents, a, b)
			ans += cost
			edges_cnt++
		}

	}
	if edges_cnt != N-1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func findP10021(p []int, x int) int {
	if p[x] != x {
		p[x] = findP10021(p, p[x])
	}
	return p[x]
}
func unionP10021(p []int, a, b int) {
	x := findP10021(p, a)
	y := findP10021(p, b)

	if x < y {
		p[y] = x
	} else {
		p[x] = y
	}
}
