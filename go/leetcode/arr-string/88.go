/*
88. Merge Sorted Array
https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

import (
	"fmt"
	"sort"
)

func P88_test() {
	test_arg1 := [][]int{{1, 2, 3, 0, 0, 0}, {1}, {0}}
	test_arg2 := [][]int{{2, 5, 6}, {}, {1}}
	test_arg3 := []int{3, 1, 0}
	test_arg4 := []int{3, 0, 1}

	for i := 0; i < 3; i++ {
		P88_O(test_arg1[i], test_arg3[i], test_arg2[i], test_arg4[i])
		fmt.Println(test_arg1[i])
	}
}
func P88(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
func P88_O(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

/*
nums1 index m-1에 num2를 merge.
그 다음 정렬.

정석.
배열들의 끝부터 시작하여 요소들을 비교하고 nums1에 배치한다.
*/
