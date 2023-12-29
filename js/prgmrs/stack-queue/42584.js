/*
주식가격
https://school.programmers.co.kr/learn/courses/30/lessons/42584
*/
const tests = [[1, 2, 3, 2, 3]]

for(let t of tests) {
    console.log(solution(t));
}
function solution(prices) {
    var answer = [];
    for(let i = 0; i< prices.length; i++) {
        let cnt = 0;
        for(let j = i+1; j < prices.length; j++) {
            cnt++
            if(prices[i] > prices[j]) {
                break;
            } 
        }
        answer.push(cnt);
    }
    return answer;
}