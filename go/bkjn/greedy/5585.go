/*
거스름돈
https://www.acmicpc.net/problem/5585
*/
package greedy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P5585() {
	cType := []int{500, 100, 50, 10, 5, 1}

	result := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	n := 1000 - num

	for i := 0; i < len(cType); i++ {
		result += n / cType[i]
		n %= cType[i]
		if n <= 0 {
			break
		}
	}
	fmt.Println(result)
}

/*
큰 거스름돈 액수부터 내림차순으로 계산.
*/
