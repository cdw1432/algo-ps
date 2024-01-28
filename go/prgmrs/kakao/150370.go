/*
개인정보 수집 유효기간
https://school.programmers.co.kr/learn/courses/30/lessons/150370
*/
package kakao

import (
	"fmt"
	"strconv"
	"strings"
)

func P150370_test() {
	fmt.Println(P150370("2022.05.19", []string{"A 6", "B 12", "C 3"}, []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"}))
}
func P150370(today string, terms []string, privacies []string) []int {
	ans := []int{}
	terms_hash_map := make(map[string]int)
	today_ := strings.Split(today, ".")

	tyear, _ := strconv.Atoi(today_[0])
	tmonth, _ := strconv.Atoi(today_[1])
	tday, _ := strconv.Atoi(today_[2])

	dday_ := tyear*(12*28) + (28 * tmonth) + tday

	//약관 종류와 그에 따른 유효기간 HashMap으로 정리
	for i := range terms {
		parts := strings.Fields(terms[i])
		num, _ := strconv.Atoi(parts[1])
		terms_hash_map[parts[0]] = num * 28
	}
	//개인정보 배열 순회
	for i := range privacies {
		parts := strings.Fields(privacies[i])
		term := parts[1]

		date := strings.Split(parts[0], ".")

		year, _ := strconv.Atoi(date[0])
		month, _ := strconv.Atoi(date[1])
		day, _ := strconv.Atoi(date[2])

		cday_ := year*(12*28) + (28 * month) + day
		if dday_ >= cday_+terms_hash_map[term] {
			ans = append(ans, i+1)
		}
	}
	//fmt.Println(terms_hash_map)
	return ans
}
