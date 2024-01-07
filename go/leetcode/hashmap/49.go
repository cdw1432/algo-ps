/*
49. Group Anagrams
https://leetcode.com/problems/group-anagrams/?envType=study-plan-v2&envId=top-interview-150
*/
package hashmap

import (
	"fmt"
	"sort"
	"strings"
)

func P49_test() {
	arg := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}
	for _, x := range arg {
		fmt.Println(P49(x))
	}
}
func P49(strs []string) [][]string {
	result := [][]string{}
	hashMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		sli := strings.Split(strs[i], "")
		sort.Strings(sli)
		str := strings.Join(sli, "")
		hashMap[str] = append(hashMap[str], strs[i])
	}
	//fmt.Println(hashMap)
	for _, value := range hashMap {
		result = append(result, value)
	}
	return result
}

/*
알파벳순으로 정렬후 해쉬맵에 저장.
*/
