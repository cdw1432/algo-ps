/*
카펫
https://school.programmers.co.kr/learn/courses/30/lessons/42842
*/
package bruteforce

import (
	"fmt"
	"math"
)

func P42842_test() {
	tests := [][]int{
		{10, 2},
		{8, 1},
		{24, 24},
	}
	for _, v := range tests {
		fmt.Println(P42842(v[0], v[1]))
	}
}
func P42842(brown int, yellow int) []int {
	ab := findMultiplications(yellow, brown)
	//fmt.Println(ab)
	return []int{ab[1] + 2, ab[0] + 2}
}
func findMultiplications(n int, brown int) []int {
	var ab []int

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factor := n / i
			if brown == (2 * (i + factor + 2)) {
				ab = append(ab, i)
				ab = append(ab, factor)
				break
			}
		}
	}
	return ab
}

/*
a = 세로
b = 가로

조건)
yellow == a x b
brown == 2(a+b+2)

yellow를 소인수들만의 곱으로 나타내는 경우의 수를 계산. sqrt(yellow)를 하면 테스트케이스를 줄인다.
위 두 조건을 만족하는 a x b 가 yellow의 가로 세로이다.

카펫의 넓이는 b+2, a+2

*/
