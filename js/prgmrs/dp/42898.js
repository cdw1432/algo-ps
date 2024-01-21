/*
등굣길
https://school.programmers.co.kr/learn/courses/30/lessons/42898
*/
console.log(solution(100,1,[[1,50]]))
function solution(m, n, puddles) {
    var dp = Array.from({length: n+1}, () => {
        return new Array(m+1).fill(0);
    })
    dp[1][1] = 1
    for(let p of puddles) {
        dp[p[1]][p[0]] = -1
    }
    
    for(let i = 1; i < n+1; i++) {
        for(let j = 1; j < m+1; j++) {
            if(dp[i][j] == -1) {
                dp[i][j] = 0
                continue
            }

            dp[i][j] += (dp[i-1][j] + dp[i][j-1])%1000000007

        }
    }
    return dp[n][m];
}