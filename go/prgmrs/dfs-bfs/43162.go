/*
네트워크
https://school.programmers.co.kr/learn/courses/30/lessons/43162
*/
package dfsbfs

func P43162(n int, computers [][]int) int {
	cnt := 0
	visited := make([]bool, n)
	for !allVisited(visited) {
		for i, v := range visited {
			if !v {
				dfsP43162(computers, visited, i)
				break
			}
		}

		cnt++
	}
	return cnt
}

func allVisited(vi []bool) bool {
	for _, v := range vi {
		if v == false {
			return false
		}
	}
	return true
}
func dfsP43162(computers [][]int, visited []bool, start int) {
	stack := []int{}
	stack = append(stack, start)
	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited[c] = true
		for i := range computers[c] {
			if !visited[i] && computers[c][i] == 1 {
				stack = append(stack, i)
			}
		}
	}
}
