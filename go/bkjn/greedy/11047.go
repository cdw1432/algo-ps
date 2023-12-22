/*
동전 0
https://www.acmicpc.net/problem/11047
*/
package greedy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P11047() {
	var k []int
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	_, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < n; i++ {
		cInput, _ := reader.ReadString('\n')
		c, _ := strconv.Atoi(strings.TrimSpace(cInput))

		k = append(k, c)
	}
	solution(n, m, k)
}

func solution(n int, m int, k []int) {
	ans := 0
	for i := len(k) - 1; i >= 0; i-- {
		if m >= k[i] {
			ans += m / k[i]
			m %= k[i]
		}
	}
	fmt.Println(ans)
}
