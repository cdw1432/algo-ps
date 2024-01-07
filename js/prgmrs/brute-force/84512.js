/*
Words Made up of Vowels
https://school.programmers.co.kr/learn/courses/30/lessons/84512
*/
solution("UUUUU")
function solution(word) {
    let vowels = ['A','E','I','O','U']

    let stack = [];
    let cnt = 0;
    stack.push(vowels[0]);
    while(stack.length > 0) {
        cnt++;
        console.log(stack.join(''), cnt)
        if(stack.join('') === word)
            break;
        if(stack.length == 5) {
            let last = stack.pop();
            while(last == 'U' && stack.length > 0) {
                last = stack.pop();
            }
            if(last != 'U')
                stack.push(vowels[vowels.indexOf(last)+1])
        } else {
            stack.push(vowels[0])
        }
    }
    return cnt;
}
/*
스택으로 사전에 있는 단어들을 순서대로 출력한다.
word에 도달하면 반복횟수를 반환.

이 방식은 느리다. 다른 코드를 보니 더 좋은 방법들이 많다.
*/