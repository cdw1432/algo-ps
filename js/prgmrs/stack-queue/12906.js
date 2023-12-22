/*
같은 숫자는 싫어
https://school.programmers.co.kr/learn/courses/30/lessons/12906
*/

const tests = [[1,1,3,3,0,1,1],[4,4,4,3,3]]
for(let test of tests) {
    console.log(solution(test));
}
//filter를 쓰면 한줄 컷, 그러나 stack의 원리를 이용해보자.
//return arr.filter((val,index) => val != arr[index+1]);
function solution(arr)
{
    var stack = [];
    stack.push(arr.pop());
    while(arr.length > 0) {
        let ele = arr.pop();
        if(ele != stack[stack.length-1])
            stack.push(ele)
    }
    return stack.reverse();
}