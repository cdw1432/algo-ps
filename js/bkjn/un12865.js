/*
평범한 배낭
https://www.acmicpc.net/problem/12865
*/
const rl = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});


let data = []
let i = 0;
rl.on('line', (line) => {
    data.push(line.split(' ').map(Number))

    if(i >= data[0][0])
        rl.close();

    i++
}).on('close', () => {
    solution();
    process.exit();
})

function solution() {
    let [N, K] = data.shift()
    data.sort((a,b) => {
        if(a[0] == b[0])
            return a[1] - b[1]
        return a[0] - b[0]
    })

    let start = 0
    let end = 0
    let ans = 0
    
    while(data[start][0] <= K) {
        let [w,v] = sumOf(data, start, end)
        //console.log(start,end)
        if(w <= K) {
            if(ans < v)
                ans = v
            end++
        } else {
            start = end
        }
        if(end >= N || start >= N) {
            break;
        } 
    }
    console.log(ans)
}

function sumOf(arr, s, e) {
    let w = 0;
    let v = 0;
    for (let i = s; i <= e; i++) {
        w += arr[i][0]
        v += arr[i][1]
    }
    return [w,v];
}