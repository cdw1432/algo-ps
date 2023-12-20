/* 
베스트 앨범
https://school.programmers.co.kr/learn/courses/30/lessons/42579
*/
const tests = 2;
const tests_arg1 = [
    ["classic", "pop", "classic", "classic", "pop"],
    ["classic", "pop", "korean", "japan", "china", "thai", "india"],
];
const tests_arg2 = [
    [500, 600, 150, 800, 2500],
    [100, 200, 100, 300, 300, 500, 600],
];

for(let i = 0; i < tests; i++) {
    console.log(solution(tests_arg1[i], tests_arg2[i]));
}

function solution(genres, plays) {
    let result = [];
    let map = new Map();
    for(let i in genres) {
        if(map.has(genres[i]))
            map.get(genres[i]).push([plays[i],+i]);
        else
            map.set(genres[i], [[plays[i],+i]])
    }
    map.forEach((value,key) => {
        let tot = value.reduce((tot,e) => tot+e[0],0);
        value.unshift(tot);
    });
    sortedMap = new Map([...map].sort((a,b) => b[1][0] - a[1][0]));
    console.log(sortedMap);
    
    sortedMap.forEach((value, key) => {
        value.shift();
        value.sort((a,b) => b[0] - a[0]);
        result.push(value[0][1]);
        if(value.length > 1) {
            result.push(value[1][1])
        }
    })
    return result;
}
//코드를 조금 더 축약할 수 있을 거 같다.
