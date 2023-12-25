/*
26. Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

import "fmt"

func P26(nums []int) int {
	var arr []int
	arr = append(arr, nums[0])
	for i := 1; i < len(nums); i++ {
		if arr[len(arr)-1] != nums[i] {
			arr = append(arr, nums[i])
		}

	}
	copy(nums, arr)
	fmt.Println(nums)
	return len(arr)
}
func P26_O(nums []int) int {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}
	return l
}

/*
새로운 배열에 담고 그것을 마지막에 nums에 copy한다.
런타임은 빠르지만 메모리 측면에서는 별로다.

정석.
nums를 순회하며 반복된 숫자가 끝나는 지점을 찾는다.
찾았다면 그 요소를 알맞은 순서에 넣는다. 알맞은 순서는 찾을때마다 한칸씩 늘린다.
반복되는 숫자를 업데이트해준다.
*/
