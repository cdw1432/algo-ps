/*
피로도
https://school.programmers.co.kr/learn/courses/30/lessons/87946
*/

console.log(solution(80,[[80,20],[50,40],[30,10]]))

function solution(k, dungeons) {
    dungeons = dungeons.filter((v) => v[0] <= k)
    
    let pmts = permutation(dungeons)
    //console.log(pmts)
    let ans = 0;
    for(let i = 0; i < pmts.length; i++) {
        let t = k;
        let cnt = 0;
        for(let j = 0; j < pmts[i].length; j++) {
            if(pmts[i][j][0] <= t) {
                t -= pmts[i][j][1];
                cnt++;
            } else {
                break;
            }
        }
        ans = Math.max(ans,cnt);
    }
    return ans
}
function permutation(arr) {
    let result = [];
    
    function permute(curr, rem) {
        if(rem.length === 0) {
            result.push(curr);
            return;
        }
        for(let i = 0; i < rem.length; i++) {
            let ncurr = [...curr,rem[i]];
            let nrem = [...rem.slice(0,i), ...rem.slice(i+1)]
            permute(ncurr,nrem);
        }
    }

    permute([], arr);
    return result;
}
/*
permutation 함수
경우의 수 계산.
abc
acb
bac
bca
cab
cba
*/