/*
정수 삼각형
https://school.programmers.co.kr/learn/courses/30/lessons/43105?language=javascript
*/

test = [[7], [3, 8], [8, 1, 0], [2, 7, 4, 4], [4, 5, 2, 6, 5]]

console.log(solution(test))

function solution(triangle) {
    dp = triangle.slice();
    for(let i = 1; i < triangle.length; i++) {
        for(let j = 0; j < triangle[i].length; j++) {
            if(j == 0) {
                dp[i][j] = dp[i-1][j] + dp[i][j]
                continue;
            }

            if(j == triangle[i].length-1) {
                dp[i][j] = dp[i-1][j-1] + dp[i][j]
                continue;
            }

            dp[i][j] = Math.max(dp[i-1][j],dp[i-1][j-1]) + dp[i][j]
        }
    }
    return Math.max(...dp[triangle.length-1]);
}
/*
7
3 8
8 1 0
2 7 4 4
4 5 2 6 5

[i][j] 
j == 첫
위 선택
j == 끝
대각선 왼쪽 위 선택
j == 중간
위, 대각선 왼쪽 위 중 큰 수 선택
*/

