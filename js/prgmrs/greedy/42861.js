/*
섬 연결하기
https://school.programmers.co.kr/learn/courses/30/lessons/42861
*/

console.log(solution(4,[[0,1,1],[0,2,2],[1,2,5],[1,3,1],[2,3,8]]))

function solution(n, costs) {
    let find_p = (p, x) => {
        if(p[x] != x) {
            p[x] = find_p(p, p[x])
        }
        return p[x]
    }
    let union_p = (p, a,b) => {
        let x = find_p(p, a)
        let y = find_p(p, b)

        if(x < y) {
            p[y] = x
        } else {
            p[x] = y
        }
    }
    let ans = 0;
    let parent = new Array(n).fill(0).map((_, i) => i);
    costs.sort((a,b) => a[2] - b[2])
    console.log(costs)
    for(let e of costs) {
        if(find_p(parent, e[0]) != find_p(parent,e[1])) {
            union_p(parent,e[0],e[1])
            ans += e[2]
        }
    }
    return ans;
}
/*
크루스칼 알고리즘
*/