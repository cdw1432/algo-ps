/*
[카카오 인턴] 키패드 누르기
https://school.programmers.co.kr/learn/courses/30/lessons/67256
*/

tests = [
    [
        [1, 3, 4, 5, 8, 2, 1, 4, 5, 9, 5],"right"
    ],
    [
        [7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2],"left"
    ],
    [
        [1, 2, 3, 4, 5, 6, 7, 8, 9, 0],"right"
    ]
]

for(let t of tests) {
    console.log(solution(t[0],t[1]))
}
console.log(solution([0, 0], "right"))
function solution(numbers, hand) {
    ans = ""
    keyMap = {
        1 : [0,0], 2 : [0,1], 3 : [0,2],
        4 : [1,0], 5 : [1,1], 6 : [1,2],
        7 : [2,0], 8 : [2,1], 9 : [2,2],
                   0 : [3,1]
    }

    leftx = 3; lefty = 0;
    rightx = 3; righty = 2;
    for(let n of numbers) {
        //1,4,7 왼손 사용
        if(n == 1 || n == 4 || n == 7) {
            ans += 'L'
            leftx = keyMap[n][0]
            lefty = keyMap[n][1]
            continue;
        }
        //3,6,7 오른손 사용
        if(n == 3 || n == 6 || n == 9) {
            ans += 'R'
            rightx = keyMap[n][0]
            righty = keyMap[n][1]
            continue;
        }

        //두 엄지의 위치와 목적지 거리 계산
        let leftPath = Math.abs(leftx - keyMap[n][0]) + Math.abs(lefty - keyMap[n][1])
        let rightPath = Math.abs(rightx - keyMap[n][0]) + Math.abs(righty - keyMap[n][1])
        // let leftPath = bfs(keyMap[n][0],keyMap[n][1], leftx,lefty)
        // let rightPath = bfs(keyMap[n][0],keyMap[n][1], rightx,righty)
        //console.log(leftPath,rightPath)
        if(leftPath < rightPath) {
            ans += 'L'
            leftx = keyMap[n][0]
            lefty = keyMap[n][1]
        } else if(leftPath > rightPath) {
            ans += 'R'
            rightx = keyMap[n][0]
            righty = keyMap[n][1]
        //최단경로의 값이 같다면
        } else if(leftPath == rightPath) {
            //오른손잡이
            if(hand == "right") {
                ans += 'R'
                rightx = keyMap[n][0]
                righty = keyMap[n][1]
            //왼손잡이
            } else if(hand == "left") {
                ans += 'L'
                leftx = keyMap[n][0]
                lefty = keyMap[n][1]
            }
        }
    }
    return ans;
}

function bfs(destx,desty, startx,starty) {
    if(destx == startx && desty == starty)
        return 0;


    let pad = Array.from({length: 4}, () => {
        return new Array(3).fill(1)
    })
    let dx = [-1,1,0,0];
    let dy = [0,0,-1,1];

    let queue = [];
    queue.push([startx,starty]);

    while(queue.length > 0) {
        let [x,y] = queue.shift()

        for(let i = 0; i < 4; i++) {
            let nx = x + dx[i]
            let ny = y + dy[i]

            if(nx < 0 || ny < 0 || nx > 3 || ny > 2)
                continue;

            if (pad[nx][ny] == 1) {
                pad[nx][ny] = pad[x][y] + 1;
                queue.push([nx,ny]);
            }
            if (nx == destx && ny == desty) return pad[destx][desty];
        }
    }
    return -1
}