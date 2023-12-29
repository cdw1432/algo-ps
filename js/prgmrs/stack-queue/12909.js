/*
올바른 괄호
https://school.programmers.co.kr/learn/courses/30/lessons/12909
*/
const tests =["()()","(())()",")()(","(()("]

for (let i of tests) {
    console.log(solution(i));
}
function solution(s){
    if(s[0] == ')')
        return false
    let stack = [];
    for(let i of s) {
        if(i == '(')
            stack.push(i);
        else
            stack.pop();
    }
    if (stack.length <= 0)
        return true;

    return false;
}
/*
스택으로 확인.
*/