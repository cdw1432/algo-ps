/*
기능개발
https://school.programmers.co.kr/learn/courses/30/lessons/42586
*/

const tests = 2;
const test_arg1 = [[93, 30, 55]	,[95, 90, 99, 99, 80, 99]]
const test_arg2 = [[1, 30, 5],[1, 1, 1, 1, 1, 1]]
for(let i = 0; i < tests; i++)
    console.log(solution(test_arg1[i],test_arg2[i]));

function solution(progresses, speeds) {
    let queue = [];
    let answer = [];
    for(let i = 0; i < progresses.length; i++) {
        let d = Math.ceil((100-progresses[i])/speeds[i]);
        queue.push(d);
    }
    let prev = queue.shift();
    count = 1;
    for(let d of queue) {
        if(d <= prev) {
            count++;
        } else {
            answer.push(count);
            count = 1;
            prev = d;
        }
    }
    answer.push(count);
    return answer;
}