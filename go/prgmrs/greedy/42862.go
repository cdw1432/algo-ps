/*
체육복
https://school.programmers.co.kr/learn/courses/30/lessons/42862?language=go
*/
package greedy

import (
	"fmt"
	"sort"
)

func P42862_test() {
	test := [][]interface{}{
		{5, []int{2, 4}, []int{1, 3, 5}},
		{5, []int{2, 4}, []int{3}},
		{3, []int{3}, []int{1}},
		{6, []int{2, 4, 5, 6}, []int{1, 3, 6}},
	}
	for _, t := range test {
		// if i == 2 || i == 1 {
		// 	continue
		// }
		fmt.Println(P42862(t[0].(int), t[1].([]int), t[2].([]int)))
	}
}
func P42862(n int, lost []int, reserve []int) int {
	sort.Ints(lost)
	sort.Ints(reserve)

	myMap := make(map[int]int)
	for _, v := range reserve {
		myMap[v]++
	}

	cnt := 0
	for _, v := range lost {
		if myMap[v] > 0 {
			myMap[v] = -1
			cnt++
		}
	}
	for _, v := range lost {
		if myMap[v] == -1 {
			continue
		}
		if v%2 == 0 {
			if myMap[v-1] > 0 {
				myMap[v-1] = 0
				myMap[v] = -1
				cnt++
			}
		} else {
			if myMap[v+1] > 0 {
				myMap[v+1] = 0
				myMap[v] = -1
				cnt++
			}
		}
	}
	for _, v := range lost {
		if myMap[v] == -1 {
			continue
		}
		if v%2 == 0 {
			if myMap[v+1] > 0 {
				myMap[v+1] = 0
				cnt++
			}
		} else {
			if myMap[v-1] > 0 {
				myMap[v-1] = 0
				cnt++
			}
		}
	}
	return cnt + (n - len(lost))
}

/*
reserve의 요소를 맵의 키로 설정. reserve 요소는 단 한 사람에게 빌려줄수있으므로 초기 값은 1

myMap에 lost 요소가 들어있다면 빌려줄수없으므로 값을 -1로 설정

lost를 순회하며 myMap확인. -1이면 빌려줄수도 없고, 빌릴 필요도 없으므로 스킵.
요소가 먼저 짝수이면 앞사람에게 빌리고
		홀수이면 뒷사람에게서 빌린다.

		lost를 다시한번 순회하며
이번엔 요소가 짝수면 뒷사람한테 빌리고
		홀수면 앞사람한테 빌린다.

잃어버린 요소가 자기자신, 앞사람, 혹은 뒷사람에게 빌리면 cnt를 증가시킨다.

(잃어버리지 않은 사람 + cnt)를 반환.
*/
