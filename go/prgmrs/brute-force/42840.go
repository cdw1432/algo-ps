/*
모의고사
https://school.programmers.co.kr/learn/courses/30/lessons/42840?language=go
*/
package bruteforce

func P42840(answers []int) []int {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 1, 2, 3, 2, 4, 2, 5}
	c := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
	cnt := make([]int, 3)
	for i := range answers {
		if answers[i] == a[i%len(a)] {
			cnt[0]++
		}
		if answers[i] == b[i%len(b)] {
			cnt[1]++
		}
		if answers[i] == c[i%len(c)] {
			cnt[2]++
		}
	}
	max := 0
	for _, v := range cnt {
		if max < v {
			max = v
		}
	}
	answer := []int{}
	for k, v := range cnt {
		if max == v {
			answer = append(answer, k+1)
		}
	}
	return answer
}

/*
모듈러를 이용하면 쉽게 풀수있음.

최대값을 찾고 최대값과 같은 인덱스만 배열에 포함시켜 반환.
*/
