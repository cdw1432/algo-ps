/*
전력망을 둘로 나누기
https://school.programmers.co.kr/learn/courses/30/lessons/86971
*/

tests = [
    [9,[[1,3],[2,3],[3,4],[4,5],[4,6],[4,7],[7,8],[7,9]]],
    [4,[[1,2],[2,3],[3,4]]],
    [7,[[1,2],[2,7],[3,7],[3,4],[4,5],[6,7]]],
    [2,[1,2]]
]
for( let t of tests) {
    console.log(solution(t[0],t[1]));
}

function solution(n, wires) {
    find = function(p, a) {
        if(p[a] != a)
            p[a] = find(p, p[a])
        return p[a]   
    }
    union = function(p, a,b) {
        x = find(p, a)
        y = find(p, b)
        if(x < y)
            p[y] = x
        else
            p[x] = y
    }

    ans = n;
    for (i = 0; i < n-1; i++) {
        let parent = new Array(n+1).fill(0)
        for(let i in parent) {
            parent[i] = i;
        }
        
        for (j = 0; j < n-1; j++) {
            if(i == j)
                continue;
            else
                union(parent, wires[j][0], wires[j][1])
        }
        //console.log(parent)
        let result = []
        if(parent) {
            for(let k = 1; k < parent.length; k++)
            result.push(find(parent,parent[k]));

        let max = Math.max(...result);
        let n1 = result.filter((ele) => ele == max).length;
        let diff = Math.abs(n1-(n - n1));
        if (diff < ans) {
            ans = diff;
        }
        }
    }
    return ans
}
/*
서로소 집합으로 풀이.
순회하며 와이어를 하나 제외 시킨다
그럼 두 그룹으로 나뉘고 
이 두 그룹의 차가 가장 작은 숫자를 반환
*/


/* 
실패 30.8/100 
왜 틀린지 모르겠다. 메모.
*/
// function find_parent(parent,x) {
//     if(parent[x] != x) {
//         parent[x] = find_parent(parent,parent[x])
//     }
//     return parent[x]
// }
// function union_parent(parent, a,b) {
//     x = find_parent(parent,a)
//     y = find_parent(parent,b)
//     if (x < y) 
//         parent[y] = x
//     else
//         parent[x] = y
// }
// function solution(n, wires) {
//     len = wires.length
//     ans = n;
//     cnt = 0;
//     for(let i = 0; i < len; i++) {
//         let parent = new Array(n+1).fill(0)
//         for(let i in parent) {
//             parent[i] = i
//         }
//         for(let j = 0; j < wires.length; j++) {
//             if(j == cnt)
//                 continue;
//             union_parent(parent,wires[j][0],wires[j][1])
//         }
//         printLine(parent)
//         grouped = grouping(parent);
//         let ax = grouped[1].length
//         let by = n - grouped[1].length 
//         if(ax >= by && ax-by < ans) {
//             ans = ax - by
//         } else if(ax < by && by-ax < ans) {
//             ans = by - ax
//         }
//         //console.log(grouped)
//         cnt++;
//     }
//     return ans;
// }
// function grouping(arr) {
//     grouped = {};
//     for(let ele of arr) {
//         if(ele == 0)
//             continue;
//         if(grouped[ele])
//             grouped[ele].push(ele);
//         else
//             grouped[ele] = [ele]
//     }
//     return grouped
// }
// function printLine(arr) {
//     for (let i of arr) {
//         process.stdout.write(i + ' ')
//     }
//     console.log();
// }