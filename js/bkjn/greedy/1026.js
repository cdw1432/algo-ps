/*
보물
https://www.acmicpc.net/problem/1026
*/
const rl = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
})

let data = [];
let i = 1;
rl.on('line', (line) => {
    data.push(line.split(' ').map(Number))

    if(i > 2)
        rl.close();
    i++;
}).on('close', () => {
    //solution();
    easySolution();
    process.exit();
})
// a[0] X b[0] + ... + a[n-1] X b [n-1]
function solution() {
    let n = data.shift();
    let a = data.shift();
    let b = data.shift();
    
    let order = b.map((_,i) => i);
    order.sort((x,y) => b[x] - b[y]);
    
    let reOrder = new Array(n);
    let num = 0
    for(let i of order) {
        reOrder[i] = num;
        num++;
    }
    a.sort((a,b) => b-a);
    a = reOrder.map((index) => a[index]);

    let result = 0;
    for(let i = 0; i < n; i++)
        result += (a[i] * b[i])
    console.log(result);
}

function easySolution() {
    let n = data.shift();
    let a = data.shift();
    let b = data.shift();

    let result = 0;
    for(let i = 0; i < n; i++) {
        const aE = Math.min(...a)
        const bE = Math.max(...b)
        result += aE * bE
        a.splice(a.indexOf(aE),1);
        b.splice(b.indexOf(bE),1);
    }
    console.log(result);
}
/* 
b 배열과 똑같은 크기의 배열을 만들고 b[i]의 오름차순 순위를 t[i]에 저장한다.
ex)
b = [2,7,8,3,1]
t = [1,3,4,2,0]
a 배열을 내림차순으로 정렬한다. t 배열의 각 요소를 인덱스로 보고 a 배열을 알맞게 다시 한번 정렬한다.
ex)
a = [1,1,1,6,0]
a = [1,1,0,1,6]

각 인덱스가 똑같은 a,b 요소 끼리 곱하고 결과를 더한다.
result += (a[i] * b[i]);

풀고 보니 더 쉬운 방법은
a 배열의 가장 작은 요소를 찾고 Min, 요소 제거
b 배열의 가장 큰 요소를 찾고 Max, 요소 제거
*/