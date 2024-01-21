/*
2024 KAKAO WINTER INTERNSHIP
가장 많이 받은 선물
https://school.programmers.co.kr/learn/courses/30/lessons/258712?language=javascript
*/

tests = [
    [
        ["muzi", "ryan", "frodo", "neo"],
        ["muzi frodo", "muzi frodo", "ryan muzi", "ryan muzi", "ryan muzi", "frodo muzi", "frodo ryan", "neo muzi"]	
    ],
    [
        ["joy", "brad", "alessandro", "conan", "david"],
        ["alessandro brad", "alessandro joy", "alessandro conan", "david alessandro", "alessandro david"]	
    ],
    [
        ["a", "b", "c"],
        ["a b", "b a", "c a", "a c", "a c", "c a"]	
    ]
]

for(let t of tests)
    console.log(solution(t[0],t[1]))

function solution(friends, gifts) {
    let table = Array.from({length: friends.length}, () => {
        return new Array(friends.length).fill(0)
    })
    let index = new Array(friends.length)
    
    let order = {}
    let ans = {}
    for(let i = 0; i < friends.length; i++) {
        order[friends[i]] = i
        ans[friends[i]] = 0
    }
    for(let g of gifts) {
        let c = g.split(' ')

        let giver = order[c[0]]
        let receiver = order[c[1]]
        
        table[giver][receiver]++ 
        table[receiver][giver]--
    }

    for(let f of friends) {
        index[order[f]] = table[order[f]].reduce((tot,v) => tot+v,0)
    }

    for(let i = 0; i < table.length; i++) {
        for(let j = 0; j < table.length; j++) {
            if(i == j)
                continue;

            //서로 주고 받음
            if(table[i][j] != table[j][i]) {
                if(table[i][j] > table[j][i]) {
                    ans[friends[i]]++
                }
            } else {
                if(index[i] > index[j]) {
                    ans[friends[i]]++
                }
            }

        }
    }
    return Math.max(...Object.values(ans))
}