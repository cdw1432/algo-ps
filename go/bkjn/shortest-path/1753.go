/*
최단경로
https://www.acmicpc.net/problem/1753
*/
package shortestpath

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INF = 1e9

func P1753() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	first := scanner.Text()

	parts := strings.Fields(first)

	//vertex and edges
	v, _ := strconv.Atoi(parts[0])
	e, _ := strconv.Atoi(parts[1])
	//make 2d array, sized [v+1][v+1]
	arr := make([][][]int, v+1)
	for i := 1; i < v+1; i++ {
		arr[i] = make([][]int, 0)
	}

	//start node
	scanner.Scan()
	second := scanner.Text()

	startNode, _ := strconv.Atoi(second)

	//edge info
	for i := 0; i < e; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		tmp := []int{b, c}
		arr[a] = append(arr[a], tmp)
	}
	for i, v := range dijkstra(v, e, startNode, arr) {
		if i == 0 {
			continue
		}
		if v == INF {
			fmt.Println("INF")
		} else {
			fmt.Println(v)
		}
	}
}

func get_smallest_node(arr []int, visited []bool) int {
	min_v := int(INF)
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < int(min_v) && !visited[i] {
			min_v = arr[i]
			index = i
		}
	}
	return index
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func dijkstra(v, e, n int, arr [][][]int) []int {
	visited := make([]bool, v+1)
	for i := range visited {
		visited[i] = false
	}
	distance := make([]int, v+1)
	for i := range distance {
		distance[i] = INF
	}

	distance[n] = 0
	visited[n] = true
	// should take the minimum of the existing weight and the new weight.
	for _, v := range arr[n] {
		distance[v[0]] = min(distance[v[0]], v[1])
	}

	for i := 0; i < v-1; i++ {
		now := get_smallest_node(distance, visited)
		visited[now] = true
		for _, v := range arr[now] {
			cost := distance[now] + v[1]
			if cost < distance[v[0]] {
				distance[v[0]] = cost
			}
		}
	}
	return distance
}

//just a Dijkstra algorithm.
