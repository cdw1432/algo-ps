/*
부녀회장이 될테야
https://www.acmicpc.net/problem/2775
*/
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P2775() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	result := make([]int, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		result[i] = solution2775(k, n)
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

func solution2775(k, n int) int {
	if n == 1 {
		return 1
	} else if k == 0 {
		return n
	}
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 0; i < k+1; i++ {
		dp[i][1] = 1
	}

	for i := 1; i < k+1; i++ {
		for j := 2; j < n+1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[k][n]
}

/*
dp 테이블을 2차원으로 만든다.
점화식만 찾으면 쉽게 풀수있다.
자신의 아래층값(i-1)과 자신의 바로 옆 집값(j-1)을 더한다.
dp[i][j] = dp[i][j-1] + dp[i-1][j]

*/
