/*
입국심사
https://school.programmers.co.kr/learn/courses/30/lessons/43238?language=go
*/
package binsearch

func P43238(n int, times []int) int64 {
	ans := 0

	left := 1
	right := maxP43238(times) * n

	for left <= right {
		mid := (left + right) / 2
		ppl := 0
		for _, v := range times {
			ppl += mid / v
			if ppl >= n {
				break
			}
		}

		if ppl >= n {
			ans = mid
			right = mid - 1
		} else if ppl < n {
			left = mid + 1
		}
	}

	return int64(ans)
}

func maxP43238(a []int) int {
	max := -1

	for _, v := range a {
		if max < v {
			max = v
		}
	}
	return max
}
