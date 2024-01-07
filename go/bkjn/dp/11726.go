/*
2xn  타일링
https://www.acmicpc.net/problem/11726
*/
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P11726() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n == 1 {
		fmt.Println(1)
	} else if n == 2 {
		fmt.Println(2)
	} else {
		dp := make([]int, n+1)
		dp[1] = 1
		dp[2] = 2

		for i := 3; i < n+1; i++ {
			dp[i] = (dp[i-2] + dp[i-1])
			dp[i] %= 10007
		}
		fmt.Println(dp[n])
	}
}

/*
점화식
dp[i] = (dp[i-2] + dp[i-1])

10007 나눈 나머지를 저장해야 최종 결과가 너무
커지지 않으며 32비트 부호 있는 정수 범위 내에 유지된다.
*/
