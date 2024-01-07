/*
소수 찾기
https://school.programmers.co.kr/learn/courses/30/lessons/42839
*/
package bruteforce

import (
	"math"
	"strconv"
)

type Set map[int]bool

func (s Set) add(ele int) {
	s[ele] = true
}
func (s Set) elements() []int {
	var result []int
	for ele := range s {
		result = append(result, ele)
	}
	return result
}
func P42839(numbers string) int {
	mySet := make(Set)

	var digits []string
	for i := 0; i < len(numbers); i++ {
		digits = append(digits, string(numbers[i]))
	}
	permutation42839(&mySet, digits, "")

	nums := mySet.elements()
	ans := 0
	for _, v := range nums {
		if isPrime(v) {
			ans++
		}
	}
	return ans
}
func permutation42839(set *Set, all []string, current string) {
	num, err := strconv.Atoi(current)
	if err == nil {
		set.add(num)
	}
	if len(all) == 0 {
		return
	}
	for i := 0; i < len(all); i++ {
		next := all[0]
		all = all[1:]
		permutation42839(set, all, current+next)
		all = append(all, next)
	}
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
set: 그냥 맵이다. 쓸 함수 몇개 작성
permutation 함수: 문자열안의 숫자들로 만들수 있는 모든 숫자를 set에 저장.
isPrime 함수: set안 요소가 소수인지 판별한다.
*/
