/*
80. Remove Duplicates from Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

import "fmt"

func P80(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	var arr []int

	arr = append(arr, nums[0])
	arr = append(arr, nums[1])
	for i := 2; i < len(nums); i++ {
		if arr[len(arr)-2] != nums[i] {
			arr = append(arr, nums[i])
		}
	}
	copy(nums, arr)
	fmt.Println(nums)
	return len(arr)
}
func P80_O(nums []int) int {
	var k int = 0
	for _, v := range nums {
		if k == 0 || k == 1 || nums[k-2] != v {
			nums[k] = v
			k++
		}
	}
	fmt.Println(nums)
	return k
}

/*
런타임은 훌륭. 메모리가 아쉽다.

아래 함수와 위 함수의 원리는 비슷하다.
*/
