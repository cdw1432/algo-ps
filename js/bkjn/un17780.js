/*
새로운 게임
https://www.acmicpc.net/problem/17780
*/
const rl = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
})

let data = []
let i = 0;
rl.on('line', (line) => {
    data.push(line.split(' ').map(Number))

    if(data[0][0] + data[0][1] <= i)
        rl.close();
    i++;
}).on('close', () => {
    solution();
    process.exit();
})

/*
0 white
    X를 그 칸으로 이동.
    이동 하려는 칸에 말이 이미 있다면 X를 올려 놓는다
1 red
    X를 그 칸으로 이동.
    이동 할때 말 순서 바뀜 reverse.

2 blue

1 2 3 4
r l u d
*/
function solution() {
    [N,K] = data.shift();
    board = data.slice(0,N)
    
    horseloc = Array.from({length: N}, () => new Array(N).fill(0))
    horse = Array.from({length: K}, () => [])
    for(i = 0; i < K; i++) {
        horseloc[data[N+i][0]-1][data[N+i][1]-1] = i+1
        horse[i].push(data[N+i])
    }
    for(let i = 0; i < K; i++) {
        if(horse[i][0][2] == -1)
            continue

        let [nx,ny] = find_next_loc(horse[i][0],N)
        
        if(board[nx][ny] == 2) {
            horse[i][0][2] = opposite(horse[i][0][2])
        }
        if(board[nx][ny] == 1) {
            horse[i].reverse()
        }

        if(horseloc[nx][ny] > 0) {
            horse[horseloc[nx][ny]-1].push(...horse[i])
            horse[i][0][2] = -1
        }
    }
    console.log(horse)
}

function opposite(x) {
    if (x % 2 == 0) {
        x--
    } else {
        x++
    }
    return x
}
function outofboard(nx,ny,N) {
    if(nx < 0 || ny < 0 || nx >= N || ny >= N) {
        return true
    }
    return false
}
function find_next_loc(b_horse, N) {
    let dx = [0, 0, -1, 1]
    let dy = [1, -1, 0, 0]

    let nx = b_horse[0]-1 + dx[(b_horse[2]-1)]
    let ny = b_horse[1]-1 + dy[(b_horse[2]-1)]

    if(outofboard(nx,ny,N)) {
        nx *= -1
        ny *= -1
    }
    return [nx,ny]
}