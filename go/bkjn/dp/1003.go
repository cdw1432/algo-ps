/*
피보나치 함수
https://www.acmicpc.net/problem/1003
*/
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P1003() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	tests := make([]int, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		tests[i] = n
	}

	for i := range tests {
		if tests[i] == 0 {
			fmt.Println("1 0")
			continue
		} else if tests[i] == 1 {
			fmt.Println("0 1")
			continue
		}
		solution1003(tests[i])
	}
}
func solution1003(n int) {
	dp := make([][]int, n+1)

	dp[0] = []int{1, 0}
	dp[1] = []int{0, 1}
	for i := 2; i <= n; i++ {
		dp[i] = []int{dp[i-1][0] + dp[i-2][0], dp[i-1][1] + dp[i-2][1]}
	}
	fmt.Printf("%d %d\n", dp[n][0], dp[n][1])
}

/*
보텀업 방식(반복문)
각 0과 1이 몇번 호출 되는 지를 저장하고 사용한다.
*/
