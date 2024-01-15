/*
구명보트
https://school.programmers.co.kr/learn/courses/30/lessons/42885?language=javascript
*/
tests = [
    [[70, 50, 80, 50],100],
    [[70, 80, 50],100]
]

for(let t of tests) {
   console.log( solution(t[0],t[1]))
}
function solution(people, limit) {
    var ans = 0;
    var front = 0
    var back = people.length-1;
    people.sort((a,b) => (a-b))
    while (front <= back) {
        if(people[front] + people[back] <= limit) {
            front++;
        }
        back--;
        ans++
    }
    return ans
}
/*
배열 오름차순 정렬 50,50,70,80
front = 0
back = people.length-1

(people[front](현재 제일 가벼운 사람) + people[back](현재 제일 무거운 사람))이 limit을 만족하는 지 확인
만족하면 제일 가벼운 사람하고 무거운 사람하고 같이 타니까 front++, back--
만족하지 못하면 무거운 사람 혼자 타니까 back--
*/