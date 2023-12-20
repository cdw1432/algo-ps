/* 
완주하지 못한 선수
https://school.programmers.co.kr/learn/courses/30/lessons/42576
*/
let e1 = ["mislav", "stanko", "mislav", "ana"]
let e1c = ["stanko", "ana", "mislav"]	
let e2 = ["leo", "kiki", "eden"]
let e2c = ["eden", "kiki"]

console.log(solution(e1,e1c));
//시간 초과, 해쉬를 써야됨
function fail(participant, completion) {
    while(completion.length > 0) {
        let c = completion.shift();
        let idx = participant.findIndex((e) => e == c);
        if(idx > -1) {
            participant.splice(idx, 1);
        }
    }
    return participant.toString();
}

function solution(participant, completion) {
    let map = {};
    for(let i = 0; i < participant.length; i++) {
        if(map[participant[i]])
            map[participant[i]]++;
        else
        map[participant[i]] = 1;
    }

    for(let i = 0; i < completion.length; i++) {
        map[completion[i]]--;
    }
    return Object.keys(map).filter(key => map[key] > 0).toString();
}