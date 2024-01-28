/*
택배 배달과 수거하기
https://school.programmers.co.kr/learn/courses/30/lessons/150369
*/
package kakao

import "fmt"

func P150369_test() {
	fmt.Println(P150369(4, 5, []int{1, 0, 3, 1, 2}, []int{0, 3, 0, 4, 0}))
}
func P150369(cap int, n int, deliveries []int, pickups []int) int64 {
	ans := 0
	deli := 0
	pick := 0
	for i := n - 1; i >= 0; i-- {
		deli += deliveries[i]
		pick += pickups[i]

		for deli > 0 || pick > 0 {
			deli -= cap
			pick -= cap
			ans += (i + 1) * 2
		}
	}
	return int64(ans)
}
