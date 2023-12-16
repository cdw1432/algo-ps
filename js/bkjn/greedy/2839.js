/*
설탕 배달
https://www.acmicpc.net/problem/2839
*/
const rl = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

let data = [];
rl.on('line', (line) => {
    data.push(parseInt(line));
    rl.close();
}).on('close', () => {
    solution();
    process.exit();
})

function solution() {
    let n = data.shift();
    let ans = 0;
    while(n >= 0) {
        if(n % 5 == 0) {
            ans += Math.floor(n/5);
            n %= 5;
            break;
        }
        n -= 3;
        ans++;
    }
    if(n != 0)
        console.log(-1);
    else
        console.log(ans);
}