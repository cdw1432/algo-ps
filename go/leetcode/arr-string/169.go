/*
169. Majority Element
https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

import (
	"fmt"
	"sort"
	"strconv"
)

func P169(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)

	result := 0
	maxN := 0
	cnt := 0
	for i := range nums {
		if i == 0 || nums[i] == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if maxN < cnt {
			result = i - 1
			maxN = cnt
		}
	}
	fmt.Println(nums)
	return nums[result]
}
func P169_O(nums []int) int {
	m := make(map[string]int)
	for i := range nums {
		m[strconv.Itoa(nums[i])]++
	}

	for k, v := range m {
		if v > len(nums)/2 {
			toInt, _ := strconv.Atoi(k)
			return toInt
		}
	}
	return -1
}

/*
배열을 정렬 후 연속된 수 횟수를 세어서 가장 많이 연속되는 수의 인덱스를 반환

더 좋은 방법은
- Boyer-Moore Majority Voting Algo.
- hashmap
O(n)
*/
