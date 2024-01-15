/*
조이스틱
https://school.programmers.co.kr/learn/courses/30/lessons/42860
*/
package greedy

import (
	"fmt"
)

func P42860_test() {
	fmt.Println(P42860("JEROEN"))
	fmt.Println(P42860("JAN"))
}

func P42860(name string) int {
	screen := make([]rune, len(name))
	for i := range screen {
		screen[i] = 'A'
	}
	cnt := 0
	cursor := len(name) - 1
	for i, v := range name {
		if int(v) != 65 {
			if int(v)-65 >= 14 {
				cnt += 90 - int(v) + 1
			} else {
				cnt += int(v) - 65
			}
		}

		next := i + 1
		for next < len(name) && name[next] == 'A' {
			next++
		}
		if cursor > 2*i+(len(name)-next) {
			cursor = 2*i + (len(name) - next)
		}
		if cursor > i+2*(len(name)-next) {
			cursor = i + 2*(len(name)-next)
		}
	}
	return cnt + cursor
}
func P42860_fail(name string) int {

	screen := make([]rune, len(name))
	for i := range screen {
		screen[i] = 'A'
	}

	cnt := 0
	starting := 1
	for i, v := range name {
		fmt.Println(int(v), string(v))
		if int(v) != 65 {
			if (i+1)-starting < (len(name))-(i+1)+starting {
				cnt += (i + 1) - starting
			} else {
				cnt += (len(name)) - (i + 1) + starting
			}
			//			fmt.Println(cnt)
			if int(v)-65 >= 14 {
				cnt += 90 - int(v) + 1
			} else {
				cnt += int(v) - 65
			}
			starting = i + 1
		}

	}
	//printRune(screen)
	return cnt
}

/*
실패 48.1
name 문자열 길이의 배열 만들기. Screen
Screen 'A'로 초기화

name 순회
name 요소와 screen 요소가 같은 지 확인

해당 요소가 13보다 큰지 확인. (ascii코드 'A' ~ 'Z'는 65 ~ 90이다)
중간 지점인  13보다 크면 왼쪽으로 이동 작으면 오른쪽으로 이동
이동거리 계산해서 cnt에 추가

현재 인덱스와 마지막 커서 위치 starting을 이용해 오른쪽 커서이동이 빠른지 왼쪽 커서이동이 빠른지 계산 후
수가 작은 쪽을 이동 거리 cnt에 추가
*/
