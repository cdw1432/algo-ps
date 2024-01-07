/*
383. Ransom Note
https://leetcode.com/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150
*/
package hashmap

import "fmt"

func P393_test() {
	arg1 := []string{"a", "aa", "aa"}
	arg2 := []string{"b", "ab", "aab"}

	for i := range arg1 {
		fmt.Println(P393(arg1[i], arg2[i]))
	}
}
func P393(ransomNote string, magazine string) bool {
	hashMap := make(map[string]int)

	for _, v := range magazine {
		hashMap[string(v)]++
	}
	for _, v := range ransomNote {
		if _, exists := hashMap[string(v)]; exists {
			hashMap[string(v)]--
			if hashMap[string(v)] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

/*
basic hashmap problem
*/
