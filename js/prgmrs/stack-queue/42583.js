/*
다리를 지나는 트럭
https://school.programmers.co.kr/learn/courses/30/lessons/42583
*/
const tests = [
    [2,10,[7,4,5,6],8],
     [100,100,[10],101],
     [100,100,[10,10,10,10,10,10,10,10,10,10],110]
]
for(let t of tests) {
    console.log(solution(t[0],t[1],t[2],t[3]));
}

function solution(length, weight, trucks) {
    var timer = 1;
    let bridge = [];
    let cweight = [];
    while(true) {
        
        if(timer - bridge[0] == length) {
            bridge.shift();
            cweight.shift();
        }

        let c = trucks[0];
        let w = cweight.reduce((tot,v) => tot+v,0);
        if(bridge.length < length && w + c <= weight) {
            bridge.push(timer);  
            cweight.push(c);
            trucks.shift();
        }
        if(bridge.length == 0)
            break;
        
        //console.log(timer + ': ' + bridge + ' ');
        timer++;
    }
    return timer;
}
/*
우선 다리를 상징하는 큐(배열)를 선언.

큐 안에 길이가, length보다 작고 전체 요소의 합 + 다음 차례의 트럭의 무게가 weight와 같거나 작다면 큐에 삽입.
단 timer를 삽입한다.

현재타임 - 삽입 당시 타임 이 length와 같다면 큐에서 삭제.

큐 길이가 0(다리위에 트럭이 없다면..)이면 현재타임 반환

*/