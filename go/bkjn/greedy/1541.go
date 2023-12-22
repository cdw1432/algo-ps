/*
잃어버린 괄호
https://www.acmicpc.net/problem/1541
*/
package greedy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1541() {
	scanner := bufio.NewScanner(os.Stdin)
	var arr []string
	var isStartMinus bool
	if scanner.Scan() {
		line := scanner.Text()

		isStartMinus = line[0] == '-'
		arr = strings.Split(line, "-")
	}

	sum := 0
	for _, part := range strings.Split(arr[0], "+") {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}
	if isStartMinus {
		sum = -sum
	}
	for i := 1; i < len(arr); i++ {
		parts := strings.Split(arr[i], "+")
		add := 0
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println(err)
				return
			}
			add += num
		}
		sum -= add
	}
	fmt.Println(sum)
}
