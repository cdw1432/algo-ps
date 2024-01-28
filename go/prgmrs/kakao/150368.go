/*
이모티콘 할인행사
https://school.programmers.co.kr/learn/courses/30/lessons/150368
*/
package kakao

import "fmt"

func P150368_test() {
	fmt.Println(P150368([][]int{
		{40, 2900},
		{23, 10000},
		{11, 5200},
		{5, 5900},
		{40, 3100},
		{27, 9200},
		{32, 6900},
	}, []int{1300, 1500, 1600, 4900}))
}
func P150368(users [][]int, emoticons []int) []int {
	ans := []int{0, 0}
	var sales = []int{10, 20, 30, 40}
	cases := [][]int{}

	//경우의 수
	var dfs func(tmp []int, depth int)
	dfs = func(tmp []int, depth int) {
		if depth == len(tmp) {
			cases = append(cases, append([]int{}, tmp...))
			return
		}
		for i := range sales {
			tmp[depth] += sales[i]
			dfs(tmp, depth+1)
			tmp[depth] -= sales[i]
		}
	}
	dfs(make([]int, len(emoticons)), 0)
	//fmt.Println(cases)
	sum := func(x []int) int {
		sum := 0
		for i := range x {
			sum += x[i]
		}
		return sum
	}
	for i := range cases {
		join := 0
		price := make([]int, len(users))
		for j := range emoticons {
			for k := range users {
				if users[k][0] <= cases[i][j] {
					price[k] += emoticons[j] * (100 - cases[i][j]) / 100
				}
			}
		}

		for l := range users {
			if price[l] >= users[l][1] {
				join++
				price[l] = 0
			}
		}

		if join >= ans[0] {
			if join == ans[0] {
				if ans[1] < sum(price) {
					ans[1] = sum(price)
				}
			} else {
				ans[0] = join
				ans[1] = sum(price)
			}
		}
	}
	return ans
}

/*
dfs 완전탐색
*/
