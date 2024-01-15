/*
큰 수 만들기
https://school.programmers.co.kr/learn/courses/30/lessons/42883
*/
package greedy

import (
	"fmt"
	"strconv"
	"strings"
)

func P42883(number string, k int) string {
	result := []int{}

	for i := range number {
		for k > 0 && len(result) > 0 && result[len(result)-1] < int(number[i]-'0') {
			result = result[:len(result)-1]
			k--
		}
		result = append(result, int(number[i]-'0'))
	}
	result = result[:len(result)-k-1]
	str := make([]string, len(result))
	for i := range result {
		str[i] = strconv.Itoa(result[i])
	}
	return strings.Join(str, "")
}

/*
스택으로 풀이
아래 코드와 비슷한 원리.
*/
func P42883_fail(number string, k int) string {
	ans := ""
	max := 1
	maxI := 0
	lose := k
	for i := 0; i < k; i++ {
		if max < int(number[i])-48 {
			max = int(number[i]) - 48
			maxI = i
		}
	}
	lose -= maxI
	number = number[maxI:]

	for i := 0; i < len(number)-1; i++ {
		if lose > 0 && int(number[i]-'0') < int(number[i+1]-'0') {
			lose--
			continue
		}
		ans += string((number[i]))
	}
	ans += string((number[len(number)-1]))

	removeChar := func(in string, char int) string {
		result := ""
		did := 0
		for _, v := range in {
			if int(v-'0') == char && did == 0 {
				did++
				continue
			}
			result += string(v)
		}
		fmt.Println(result)
		return result
	}
	for lose > 0 {
		min := 1000001
		for i := range ans {
			if min > int(ans[i]-'0') {
				min = int(ans[i] - '0')
			}
		}
		fmt.Println(min)
		ans = removeChar(ans, min)
		lose--
	}
	fmt.Println(lose, ans)
	return ans
}

/*

실패 16.7/100
예) str = "4177252841" k = 4
먼저 문자열 첫번째 숫자부터 k번째 숫자까지 '중에' 가장 큰 수를 찾는다.

4177.. 중에 3번째 자리인 7이 가장 '먼저' 큰 수이다.

큰수 앞은 다 날려버린다. 4177252841 -> 77252841
k를 업데이트 한다. 숫자 두개가 없어졌으므로 k-2 = 2 //이제 숫자 2개만 없애면 된다

77252841 를 순회하며 i자리 수가 i+1자리보다 작다면 정답에 포함시키지 않는다.
k - 포함시키지 않는 수

마지막으로 k가 0이 될때까지 최소값을 삭제한다.
*/
