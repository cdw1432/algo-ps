/* 
전화번호 목록
https://school.programmers.co.kr/learn/courses/30/lessons/42577
*/

const tests = [["119", "97674223", "1195524421"],
                ["123","456","789"],
                ["12","123","1235","567","88"]]
for(let i = 0; i < tests.length; i++)
    console.log(solution(tests[i]));

function solution(phone_book) {
    let map = new Map();
    for(let i = 0; i < phone_book.length; i++) {
        map.set(phone_book[i], 1);
    }

    for(let i = 0; i < phone_book.length; i++) {
        let str = '';
        for(let c of phone_book[i]) {
            str += c;

            if(map.has(str) && str !== phone_book[i])
                return false;
        }
    }
    return true;
}
//긴 문자열의 접두사의 길이를 하나씩 늘려가며 해쉬 맵에 그 값이 있는지 확인.