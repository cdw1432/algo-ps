/*
가장 먼 노드
https://school.programmers.co.kr/learn/courses/30/lessons/49189
*/
package graph

func P49189(n int, edge [][]int) int {
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}
	for _, i := range edge {
		graph[i[0]][i[1]] = 1
		graph[i[1]][i[0]] = 1
	}

	distance := make([]int, n+1)
	visited := make([]bool, n+1)

	for i := range distance {
		distance[i] = int(1e9)
	}
	minNode := func() int {
		min := int(1e9)
		idx := 0

		for i := 1; i < n+1; i++ {
			if distance[i] < min && !visited[i] {
				min = distance[i]
				idx = i
			}
		}
		return idx
	}
	dijk := func(x int) {
		distance[x] = 0
		visited[x] = true

		for i, v := range graph[x] {
			if v != 0 {
				distance[i] = v
			}
		}

		for i := 0; i < n-1; i++ {
			c := minNode()
			visited[c] = true
			for j, k := range graph[c] {
				if k == 0 {
					continue
				}
				cost := distance[c] + k
				if cost < distance[j] {
					distance[j] = cost
				}
			}
		}

	}
	dijk(1)
	max := -1
	cnt := 0
	for i, v := range distance {
		if i != 0 && max < v {
			max = v
		}
	}
	for _, v := range distance {
		if max == v {
			cnt++
		}
	}
	return cnt
}
