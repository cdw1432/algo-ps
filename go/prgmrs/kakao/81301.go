/*
숫자 문자열과 영단어
https://school.programmers.co.kr/learn/courses/30/lessons/81301
*/
package kakao

import (
	"strconv"
)

func P81301(s string) int {
	wordMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	str := ""
	ans := 0
	for i := range s {
		if s[i] >= 48 && s[i] <= 57 {
			num, _ := strconv.Atoi(string(s[i]))
			ans = (10 * ans) + num
			continue
		}
		str += string(s[i])
		v, exist := wordMap[str]
		if exist {
			str = ""
			ans = (10 * ans) + v
		}
	}
	return ans
}
