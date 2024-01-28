/*
https://school.programmers.co.kr/learn/courses/30/lessons/150367
표현 가능한 이진트리
*/
package kakao

import (
	"fmt"
	"math"
	"strconv"
)

func P150367_test() {
	fmt.Println(P150367([]int64{7, 42, 5}))
}
func P150367(numbers []int64) []int {
	counter := 0
	var dfs func(node int, depth int, height int, nodeCount int, bin string) int
	dfs = func(node int, depth int, height int, nodeCount int, bin string) int {
		if depth > height {
			return 1
		}
		leftChild := 2 * node
		rightChild := 2*node + 1

		dfs(leftChild, depth+1, height, nodeCount, bin)
		if leftChild <= nodeCount && rightChild <= nodeCount {
			if bin[counter-1] == '0' {
				return 0
			}
		}
		counter++
		result := dfs(rightChild, depth+1, height, nodeCount, bin)
		return result
	}
	ans := []int{}
	for i := range numbers {
		bin := strconv.FormatInt(numbers[i], 2)
		if len(bin)%2 == 0 {
			bin = "0" + bin[:]
		}

		counter = 1
		height := int(math.Log2(float64(len(bin)) + 1))
		ans = append(ans, dfs(1, 1, height, len(bin), bin))
	}

	return ans
}
