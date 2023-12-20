/*
회의실 배정
https://www.acmicpc.net/problem/1931

Note
use bufio, scanner. it is faster.
*/
package greedy

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func P1931() {
	var n int
	fmt.Scan(&n)

	arr := make([][]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		ab := parseLine(line)
		arr[i] = ab
	}

	sol(n, arr)
}

func parseLine(line string) []int {
	// Parse space-separated integers from the input line
	ab := make([]int, 2)
	for i, v := range splitAndTrim(line, ' ') {
		ab[i], _ = strconv.Atoi(v)
	}
	return ab
}

func splitAndTrim(s string, sep rune) []string {
	// Split the string by the specified separator and trim leading/trailing spaces
	fields := make([]string, 0)
	for _, v := range split(s, sep) {
		fields = append(fields, trimSpace(v))
	}
	return fields
}

func split(s string, sep rune) []string {
	// Split the string by the specified separator
	var fields []string
	last := 0
	for i, c := range s {
		if c == sep {
			fields = append(fields, s[last:i])
			last = i + 1
		}
	}
	fields = append(fields, s[last:])
	return fields
}

func trimSpace(s string) string {
	// Trim leading and trailing spaces from the string
	return trimRight(trimLeft(s))
}

func trimLeft(s string) string {
	// Trim leading spaces from the string
	start := 0
	for start < len(s) && isSpace(rune(s[start])) {
		start++
	}
	return s[start:]
}

func trimRight(s string) string {
	// Trim trailing spaces from the string
	end := len(s)
	for end > 0 && isSpace(rune(s[end-1])) {
		end--
	}
	return s[:end]
}

func isSpace(r rune) bool {
	// Check if the rune is a space
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func sol(n int, arr [][]int) {
	ans := 1
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1] || (arr[i][1] == arr[j][1] && arr[i][0] < arr[j][0])
	})
	curr := arr[0][1]
	for i := 1; i < n; i++ {
		if curr <= arr[i][0] {
			curr = arr[i][1]
			ans++
		}
	}
	fmt.Println(ans)
}
