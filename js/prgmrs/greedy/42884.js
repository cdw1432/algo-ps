/*
단속카메라
https://school.programmers.co.kr/learn/courses/30/lessons/42884
*/

console.log(solution([[-20,-15], [-14,-5], [-18,-13], [-5,-3]]))
function solution(routes) {
    var answer = 0;
    var cam = -30001
    routes.sort((a,b) => a[1] - b[1]);
    for(let r of routes) {
        if(cam < r[0]) {
            cam = r[1]
            answer++
        }
    }
    return answer;
}