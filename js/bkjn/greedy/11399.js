/* 
ATM
https://www.acmicpc.net/problem/11399
*/
const rl = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

let i = 1;
let data = [];
rl.on('line',(line) => {
    data.push(line.split(' ').map(Number));
    if(i >= 2) 
        rl.close();
    i++;
}).on('close', () => {
    solution();
    process.exit();
})

function solution() {
    let n = data.shift();
    let arr = data.shift();
    let ans = new Array(arr.length).fill(0);
    arr.sort((a,b) => a-b);

    for(let i =0; i < arr.length; i++) {
        for(let j = 0; j <= i; j++) {
            ans[i] += arr[j];
        }
    }
    console.log(ans.reduce((tot,e) => tot + e,0));
}