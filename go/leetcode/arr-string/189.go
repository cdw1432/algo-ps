/*
189. Rotate Array
https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

import "fmt"

func P189(nums []int, k int) {
	v := k
	if k > len(nums) {
		v = k % len(nums)
	}
	slice := nums[len(nums)-v:]
	slice = append(slice, nums[:len(nums)-v]...)

	copy(nums, slice)
	fmt.Println(nums, slice)
}
func P189_O(nums []int, k int) {
	k %= len(nums)
	var slice = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		idx := (i + k) % len(nums)
		slice[idx] = nums[i]
	}
	copy(nums, slice)
}

/*
slice를 이어 붙이는 방식.

모듈러를 이용해 원형순회.
*/
