/*
1,2,3 더하기
https://www.acmicpc.net/problem/9095
*/
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P9095() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	dp := make([]int, 11)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i < 11; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	result := make([]int, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		result[i] = dp[num]
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

/*
1,2,3 으로 나타낼수있는 경우의 수를 찾으라고 하니, 1,2,3을 1,2,3으로
나태내는 경우의 수를 먼저 찾아야 겠다고 생각했다
1 = 1
2 = 2
3 = 4

그리고 점화식을 찾으려고 했다. 규칙은 금방 찾을수있었다.
dp[i] = dp[i-3] + dp[i-2] + dp[i-1]

5를 예로 들면
'4'+1 //1로 끝날때
'3'+2 //2로 끝날때
'2'+3 //3으로 끝날때
1로 끝나면서 4를 합으로 가지는 모든 수열을 포함, dp[4]
2로 끝나면서 3을 합으로 가지는 모든 수열을 포함, dp[3]
3로 끝나면서 2을 합으로 가지는 모든 수열을 포함, dp[2]

6을 예로 들면
'5'+1
'4'+2
'3'+3
dp[5] + dp[4] + dp[3]

규칙은 찾았지만 어떤 이유에서 이 규칙이 성립되는지 이해하지 못했다.
*/
