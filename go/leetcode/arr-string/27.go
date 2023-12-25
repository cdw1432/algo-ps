/*
27. Remove Element
https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

import "fmt"

func P27(nums []int, val int) int {
	k := 0
	n := len(nums)
	for i, v := range nums {
		if v == val {
			for j := n - 1; j > i; j-- {
				if nums[j] != val {
					nums[i] = nums[j]
					n = j
					break
				}
			}
			k++
		}
	}
	fmt.Println(nums)
	return len(nums) - k
}

/*
당초 요소를 지우는 것이 아닌 뒤에서 부터 val이 아닌 요소를 찾아 바꿔주는 것.
다음번에 바꿀때는 마지막에 바꾼 위치에서 부터 val이 아닌 요소를 찾는다.

*/
