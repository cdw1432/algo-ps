/*
https://www.acmicpc.net/problem/1463
1로 만들기
*/
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P1463() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var dp = make([]int, 1000000+1)
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + 1

		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}
	fmt.Println(dp[n])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
보텀업(for문)
먼전 dp 배열을 0으로 초기화
배열에는 각 인덱스의 숫자가 1이 되기까지의 최소 연산 횟수가 저장된다.

-1을 할때
dp[i] = dp[i-1] + 1 //i의 전, i-1가 1이 되기까지의 연산 횟수 + 1

2로 나누어질때
dp[i] = min(dp[i], dp[i/2]+1) //i/2로 나눈 숫자가 1이 되기까지의 연산횟수 + 1

3으로 나누어질때
dp[i] = min(dp[i], dp[i/3]+1) //i/3으로 나눈 숫자가 1이 되기까지의 연산횟수 + 1

다른사람들의 풀이를 보니 BFS, 탑다운으로 풀수도있다. 그러나 위와 같은 풀이가 대다수였다.
*/
