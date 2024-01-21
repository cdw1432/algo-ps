/*
타겟 넘버
https://school.programmers.co.kr/learn/courses/30/lessons/43165
*/
package dfsbfs

var cnt int = 0

func P43165(numbers []int, target int) int {
	dfsP43165(numbers, target, 0, 0)
	return cnt
}

func dfsP43165(number []int, target int, idx int, v int) {
	if idx == len(number) && v == target {
		cnt++
		return
	} else if idx == len(number) {
		return
	}

	dfsP43165(number, target, idx+1, v+number[idx])
	dfsP43165(number, target, idx+1, v-number[idx])
}
